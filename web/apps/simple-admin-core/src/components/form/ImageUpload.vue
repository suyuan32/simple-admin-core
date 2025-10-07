<script lang="ts" setup>
import type { UploadFile, UploadProps } from 'ant-design-vue';
import type { UploadRequestOption } from 'ant-design-vue/lib/vc-upload/interface';
import type { PropType, Ref } from 'vue';

import { uploadCloudFile } from '#/api/fms/cloudFile';
import { uploadFile } from '#/api/fms/file';
import { PlusOutlined } from '@ant-design/icons-vue';
import { $t } from '@vben/locales';
import { useVModel } from '@vueuse/core';
import { message, Modal, Upload } from 'ant-design-vue';
import { isArray, isObjectType, isString } from 'remeda';
import { computed, ref, toRefs, unref, watch } from 'vue';

defineOptions({ name: 'ImageUpload' });

const props = defineProps({
  value: {
    type: [Array<string>, String],
    default: undefined,
  },
  listType: {
    // eslint-disable-next-line no-use-before-define
    type: String as PropType<ListType>,
    default: 'picture-card',
  },
  helpText: {
    type: String as PropType<string>,
    default: '',
  },
  // 文件最大多少MB
  maxSize: {
    type: Number as PropType<number>,
    default: 2,
  },
  // 最大数量的文件，Infinity不限制
  maxNumber: {
    type: Number as PropType<number>,
    default: 1,
  },
  // 根据后缀，或者其他
  accept: {
    type: Array as PropType<string[]>,
    default: () => [],
  },
  multiple: {
    type: Boolean as PropType<boolean>,
    default: false,
  },
  fileListOpenDrag: {
    type: Boolean,
    default: true,
  },
  fileListDragOptions: {
    type: Object,
    default: () => ({}),
  },
  resultField: {
    type: String as PropType<string>,
    default: 'data.url',
  },
  calculateMd5: {
    type: Boolean as PropType<boolean>,
    default: false,
  },
  showPreviewNumber: {
    type: Boolean as PropType<boolean>,
    default: true,
  },
  emptyHidePreview: {
    type: Boolean as PropType<boolean>,
    default: false,
  },
  provider: {
    type: String,
    default: 'local',
  },
});

const emits = defineEmits(['change', 'update:value', 'delete']);
const state = useVModel(props, 'value', emits, {
  defaultValue: props.value,
  passive: true,
});

enum UploadResultStatus {
  // eslint-disable-next-line no-unused-vars
  DONE = 'done',
  // eslint-disable-next-line no-unused-vars
  ERROR = 'error',
  // eslint-disable-next-line no-unused-vars
  SUCCESS = 'success',
  // eslint-disable-next-line no-unused-vars
  UPLOADING = 'uploading',
}

type ListType = 'picture' | 'picture-card' | 'text';

function useUploadType({
  acceptRef,
  helpTextRef,
  maxNumberRef,
  maxSizeRef,
}: {
  acceptRef: Ref<string[]>;
  helpTextRef: Ref<string>;
  maxNumberRef: Ref<number>;
  maxSizeRef: Ref<number>;
}) {
  // 文件类型限制
  const getAccept = computed(() => {
    const accept = unref(acceptRef);
    if (accept && accept.length > 0) {
      return accept;
    }
    return [];
  });
  const getStringAccept = computed(() => {
    return unref(getAccept)
      .map((item) => {
        return item.indexOf('/') > 0 || item.startsWith('.')
          ? item
          : `.${item}`;
      })
      .join(',');
  });

  // 支持jpg、jpeg、png格式，不超过2M，最多可选择10张图片，。
  const getHelpText = computed(() => {
    const helpText = unref(helpTextRef);
    if (helpText) {
      return helpText;
    }
    const helpTexts: string[] = [];

    const accept = unref(acceptRef);
    if (accept.length > 0) {
      helpTexts.push($t('component.upload.accept', [accept.join(',')]));
    }

    const maxSize = unref(maxSizeRef);
    if (maxSize) {
      helpTexts.push($t('component.upload.maxSize', [maxSize]));
    }

    const maxNumber = unref(maxNumberRef);
    if (maxNumber && maxNumber !== Infinity) {
      helpTexts.push($t('component.upload.maxNumber', [maxNumber]));
    }
    return helpTexts.join('，');
  });
  return { getAccept, getStringAccept, getHelpText };
}

const { accept, helpText, maxNumber, maxSize } = toRefs(props);
const isInnerOperate = ref<boolean>(false);
const { getStringAccept } = useUploadType({
  acceptRef: accept,
  helpTextRef: helpText,
  maxNumberRef: maxNumber,
  maxSizeRef: maxSize,
});
const previewOpen = ref<boolean>(false);
const previewImage = ref<string>('');
const previewTitle = ref<string>('');

const fileList = ref<UploadProps['fileList']>([]);
const isLtMsg = ref<boolean>(true);
const isActMsg = ref<boolean>(true);

function isImgTypeByName(name: string) {
  return /\.(?:jpg|jpeg|png|gif|webp)$/i.test(name);
}

watch(
  () => state.value,
  (v) => {
    if (isInnerOperate.value) {
      isInnerOperate.value = false;
      return;
    }
    if (v) {
      let value: string[] = [];
      if (isArray(v)) {
        value = v as any;
      } else {
        value.push(v);
      }
      // eslint-disable-next-line array-callback-return
      fileList.value = value.map((item, i) => {
        if (item && isString(item)) {
          return {
            uid: `${-i}`,
            name: item.slice(Math.max(0, item.lastIndexOf('/') + 1)),
            status: 'done',
            url: item,
          };
        } else if (item && isObjectType(item)) {
          return item;
        }
      }) as UploadProps['fileList'];
    }
  },
);

function getValue() {
  const list = (fileList.value || [])
    .filter((item) => item?.status === UploadResultStatus.DONE)
    .map((item: any) => {
      return item?.url || item?.response?.url;
    });

  if (props.multiple) {
    return list;
  } else {
    return list.length > 0 ? list[0] : '';
  }
}

function getBase64<T extends ArrayBuffer | null | string>(file: File) {
  return new Promise<T>((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.addEventListener('load', () => {
      resolve(reader.result as T);
    });
    // eslint-disable-next-line unicorn/prefer-add-event-listener
    reader.onerror = (error) => reject(error);
  });
}

const handlePreview = async (file: UploadFile) => {
  if (!file.url && !file.preview) {
    file.preview = await getBase64<string>(file.originFileObj!);
  }
  previewImage.value = file.url || file.preview || '';
  previewOpen.value = true;
  previewTitle.value =
    file.name ||
    previewImage.value.slice(
      Math.max(0, previewImage.value.lastIndexOf('/') + 1),
    );
};

const handleRemove = async (file: UploadFile) => {
  if (fileList.value) {
    const index = fileList.value.findIndex((item) => item.uid === file.uid);
    index !== -1 && fileList.value.splice(index, 1);
    const value = getValue();
    isInnerOperate.value = true;
    emits('update:value', value);
  }
};

function handleCancel() {
  previewOpen.value = false;
  previewTitle.value = '';
}

const beforeUpload = (file: File) => {
  const { maxSize, accept } = props;
  const { name } = file;
  const isAct = isImgTypeByName(name);
  if (!isAct) {
    message.error($t('component.upload.acceptUpload', [accept]));
    isActMsg.value = false;
    // 防止弹出多个错误提示
    setTimeout(() => (isActMsg.value = true), 1000);
  }
  const isLt = file.size / 1024 / 1024 > maxSize;
  if (isLt) {
    message.error($t('component.upload.maxSizeMultiple', [maxSize]));
    isLtMsg.value = false;
    // 防止弹出多个错误提示
    setTimeout(() => (isLtMsg.value = true), 1000);
  }
  return (isAct && !isLt) || Upload.LIST_IGNORE;
};

async function customRequest(info: UploadRequestOption<any>) {
  try {
    const result =
      props.provider === 'local'
        ? await uploadFile(info.file as any)
        : await uploadCloudFile(
            info.file as any,
            props.provider === 'cloud-default' ? '' : props.provider,
          );

    info.onSuccess!(result.data);
    const value = getValue();
    isInnerOperate.value = true;
    emits('update:value', value);
  } catch (error: any) {
    info.onError!(error);
  }
}
</script>

<template>
  <div>
    <Upload
      v-model:file-list="fileList"
      :accept="getStringAccept"
      :before-upload="beforeUpload"
      :custom-request="customRequest"
      :list-type="listType"
      :max-count="maxNumber"
      :multiple="multiple"
      v-bind="$attrs"
      @preview="handlePreview"
      @remove="handleRemove"
    >
      <div v-if="fileList && fileList.length < maxNumber">
        <PlusOutlined />
        <div style="margin-top: 8px">{{ $t('component.upload.upload') }}</div>
      </div>
    </Upload>
    <Modal
      :footer="null"
      :open="previewOpen"
      :title="previewTitle"
      @cancel="handleCancel"
    >
      <img :src="previewImage" alt="" style="width: 100%" />
    </Modal>
  </div>
</template>
