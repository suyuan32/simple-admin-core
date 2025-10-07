<script lang="ts" setup>
import type { UploadProps } from 'ant-design-vue';

import { uploadCloudFile } from '#/api/fms/cloudFile';
import { uploadFile } from '#/api/fms/file';
import { InboxOutlined } from '@ant-design/icons-vue';
import { $t } from '@vben/locales';
import { useClipboard, useVModel } from '@vueuse/core';
import { message, UploadDragger } from 'ant-design-vue';
import { values } from 'remeda';
import { ref, watch } from 'vue';

defineOptions({
  name: 'UploadDragger',
});

const props = defineProps({
  multiple: {
    type: Boolean,
    default: true,
  },
  provider: {
    type: String,
    default: 'local',
  },
  value: {
    type: [Array, Object, String, Number],
    default: undefined,
  },
});

const emits = defineEmits(['update:value']);

const { copy } = useClipboard();

const fileUrlDict: { [key: string]: string } = {};

const fileList = ref<UploadProps['fileList']>();

const state = useVModel(props, 'value', emits, {
  defaultValue: props.value,
  passive: true,
});

async function handleUpload(file: any) {
  if (props.provider === 'local') {
    const result = (await uploadFile(file.file)) as any;
    if (fileList.value !== undefined) {
      fileList.value.forEach((item) => {
        if (item.uid === file.file.uid) {
          const resultStatus = result.code !== undefined && result.code === 0;
          item.status = resultStatus ? 'done' : 'error';
          if (resultStatus) {
            fileUrlDict[item.uid] = result.data.url;
            message.success($t('component.upload.uploadSuccess'));
          }
        }
      });
    }
  } else {
    const result = (await uploadCloudFile(file.file, props.provider)) as any;
    if (fileList.value !== undefined) {
      fileList.value.forEach((item) => {
        if (item.uid === file.file.uid) {
          const resultStatus = result.code !== undefined && result.code === 0;
          item.status = resultStatus ? 'done' : 'error';
          if (resultStatus) {
            fileUrlDict[item.uid] = result.data.url;
            message.success($t('component.upload.uploadSuccess'));
          }
        }
      });
    }
  }
}

watch(
  () => fileUrlDict.value,
  () => {
    state.value = props.multiple ? values(fileUrlDict) : values(fileUrlDict)[0];
  },
  { deep: true },
);

function handleCopyPath(file: any) {
  copy(fileUrlDict[file.uid] as any);
  message.success($t('fms.file.copyPathSuccess'));
}
</script>

<template>
  <UploadDragger
    v-model:file-list="fileList"
    :custom-request="handleUpload"
    :multiple="props.multiple"
    :show-upload-list="{ showDownloadIcon: true, showRemoveIcon: false }"
    @download="handleCopyPath"
    v-bind="$attrs"
  >
    <p class="ant-upload-drag-icon">
      <InboxOutlined />
    </p>
    <p class="ant-upload-text">
      {{ $t('component.upload.uploadHelpMessage') }}
    </p>
    <p class="ant-upload-hint">
      {{ $t('component.upload.supportAnyFormat') }}
    </p>
    <template #download-icon>
      {{ $t('fms.file.copyURL') }}
    </template>
  </UploadDragger>
</template>
