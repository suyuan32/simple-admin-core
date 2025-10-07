<script lang="ts" setup>
import type { VxeGridListeners, VxeGridProps } from '#/adapter/vxe-table';

import { h, ref } from 'vue';

import { Page, useVbenModal, type VbenFormProps } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { useClipboard } from '@vueuse/core';
import { Button, Image, message, Modal } from 'ant-design-vue';
import { isPlainObject } from 'remeda';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { deleteFile, downloadFile, getFileList } from '#/api/fms/file';
import { UploadDragger } from '#/components/form';
import { type ActionItem, TableAction } from '#/components/table/table-action';

import FileForm from './form.vue';
import { searchFormSchemas, tableColumns } from './schemas';

defineOptions({
  name: 'FileManagement',
});

// ---------------- form -----------------

const [FormModal, formModalApi] = useVbenModal({
  connectedComponent: FileForm,
});

const showDeleteButton = ref<boolean>(false);

const gridEvents: VxeGridListeners<any> = {
  checkboxChange(e) {
    showDeleteButton.value = e.$table.getCheckboxRecords().length > 0;
  },
  checkboxAll(e) {
    showDeleteButton.value = e.$table.getCheckboxRecords().length > 0;
  },
};

const formOptions: VbenFormProps = {
  // 默认展开
  collapsed: false,
  schema: [...(searchFormSchemas.schema as any)],
  // 控制表单是否显示折叠按钮
  showCollapseButton: true,
  // 按下回车时是否提交表单
  submitOnEnter: false,
};

// ------------- table --------------------

const gridOptions: VxeGridProps = {
  checkboxConfig: {
    highlight: true,
  },
  toolbarConfig: {
    slots: {
      buttons: 'toolbar-buttons',
    },
  },
  columns: [
    ...(tableColumns.columns as any),
    {
      title: $t('common.action'),
      fixed: 'right',
      field: 'action',
      slots: {
        default: ({ row }) =>
          h(TableAction, {
            actions: [
              {
                type: 'link',
                icon: 'ant-design:cloud-download-outlined',
                tooltip: $t('fms.file.download'),
                onClick: handleDownload.bind(null, row),
              },
              {
                type: 'link',
                icon: 'ant-design:copy-outlined',
                tooltip: $t('fms.file.copyURL'),
                onClick: handleCopyPath.bind(null, row),
              },
              {
                type: 'link',
                icon: 'clarity:note-edit-line',
                tooltip: $t('common.edit'),
                onClick: openFormModal.bind(null, row),
              },
              {
                icon: 'ant-design:delete-outlined',
                type: 'link',
                color: 'error',
                tooltip: $t('common.delete'),
                popConfirm: {
                  title: $t('common.deleteConfirm'),
                  placement: 'left',
                  confirm: batchDelete.bind(null, [row.id]),
                },
              },
            ] as ActionItem[],
          }),
      },
    },
  ],
  height: 'auto',
  keepSource: true,
  pagerConfig: {},
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        const res = await getFileList({
          page: page.currentPage,
          pageSize: page.pageSize,
          ...formValues,
        });
        return res.data;
      },
    },
  },
};

const [Grid, gridApi] = useVbenVxeGrid({
  formOptions,
  gridOptions,
  gridEvents,
});

function openFormModal(record: any) {
  if (isPlainObject(record)) {
    formModalApi.setData({
      record,
      isUpdate: true,
      gridApi,
    });
  } else {
    formModalApi.setData({
      record: null,
      isUpdate: false,
      gridApi,
    });
  }
  formModalApi.open();
}

function handleBatchDelete() {
  Modal.confirm({
    title: $t('common.deleteConfirm'),
    async onOk() {
      const ids = gridApi.grid.getCheckboxRecords().map((item: any) => item.id);

      batchDelete(ids);
    },
  });
}

async function batchDelete(ids: any[]) {
  const result = await deleteFile({
    ids,
  });
  if (result.code === 0) {
    await gridApi.reload();
    showDeleteButton.value = false;
  }
}

// ---------------- upload modal ------------------
const [UploadModal, uploadModalApi] = useVbenModal({
  fullscreenButton: false,
  onCancel() {
    uploadModalApi.close();
  },
  onConfirm: async () => {
    uploadModalApi.close();
  },
  onOpenChange() {},
  title: $t('component.upload.upload'),
});

const imagePath = ref<string>('');
const videoPath = ref<string>('');
const videoTitle = ref<string>('');
const imageTitle = ref<string>('');
const currentFileName = ref<string>('');

// ------------- preview video modal --------------------
const [PreviewVideoModal, previewVideoModalApi] = useVbenModal({
  fullscreenButton: false,
  onCancel() {
    previewVideoModalApi.close();
  },
  onConfirm: async () => {
    previewVideoModalApi.close();
  },
  onOpenChange() {
    previewVideoModalApi.setState({ title: videoTitle.value });
  },
  title: videoTitle.value,
});

function handleDownloadVideo() {
  const link = document.createElement('a');
  link.href = videoPath.value;
  link.download = currentFileName.value;
  link.click();
  link.remove();
  URL.revokeObjectURL(link.href);
  previewVideoModalApi.close();
}

// -------------- preview image modal ----------------------
const [PreviewImageModal, previewImageModalApi] = useVbenModal({
  fullscreenButton: false,
  onCancel() {
    previewImageModalApi.close();
  },
  onConfirm: async () => {
    previewImageModalApi.close();
  },
  onOpenChange() {
    previewImageModalApi.setState({ title: imageTitle.value });
  },
});

function handleDownloadImage() {
  const link = document.createElement('a');
  link.href = imagePath.value;
  link.download = currentFileName.value;
  link.click();
  link.remove();
  URL.revokeObjectURL(link.href);
  previewImageModalApi.close();
}

const { copy } = useClipboard();

function handleCopyPath(record: any) {
  copy(record.publicPath);
  message.success($t('common.successful'));
}

async function handleDownload(record: any) {
  const file: any = await downloadFile(record.id);
  const fileType = file.type.split('/')[0];

  currentFileName.value = `${record.name}.${record.path.split('.')[1]}`;
  if (fileType === 'image') {
    imageTitle.value = currentFileName.value;
    imagePath.value = URL.createObjectURL(file);
    previewImageModalApi.open();
  } else if (fileType === 'video') {
    videoTitle.value = currentFileName.value;
    videoPath.value = URL.createObjectURL(file);
    previewVideoModalApi.open();
  } else {
    const link = document.createElement('a');
    link.href = URL.createObjectURL(file);
    link.download = currentFileName.value;
    link.click();
    link.remove();
    URL.revokeObjectURL(link.href);
  }
}
</script>

<template>
  <Page auto-content-height>
    <PreviewImageModal>
      <Image :src="imagePath" style="" width="100%" />
      <template #footer>
        <Button key="download" type="primary" @click="handleDownloadImage">
          {{ $t('fms.file.download') }}
        </Button>
      </template>
    </PreviewImageModal>
    <PreviewVideoModal>
      <template #footer>
        <Button key="download" type="primary" @click="handleDownloadVideo">
          {{ $t('fms.file.download') }}
        </Button>
      </template>
      <video controls height="720" width="1280">
        <source :src="videoPath" type="video/mp4" />
      </video>
    </PreviewVideoModal>
    <FormModal />
    <UploadModal>
      <UploadDragger />
    </UploadModal>
    <Grid>
      <template #toolbar-buttons>
        <Button
          v-show="showDeleteButton"
          danger
          type="primary"
          @click="handleBatchDelete"
        >
          {{ $t('common.delete') }}
        </Button>
      </template>

      <template #toolbar-tools>
        <Button type="primary" @click="uploadModalApi.open">
          {{ $t('component.upload.upload') }}
        </Button>
      </template>
    </Grid>
  </Page>
</template>
