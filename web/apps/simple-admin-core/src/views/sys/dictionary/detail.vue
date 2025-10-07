<script lang="ts" setup>
import type { VxeGridListeners, VxeGridProps } from '#/adapter/vxe-table';
import type { DictionaryDetailInfo } from '#/api/sys/model/dictionaryDetailModel';

import { h, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { Button, message, Modal } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import {
  createDictionaryDetail,
  deleteDictionaryDetail,
  getDictionaryDetailList,
  updateDictionaryDetail,
} from '#/api/sys/dictionaryDetail';
import { type ActionItem, TableAction } from '#/components/table/table-action';

import { detailTableColumns } from './schemas';

defineOptions({
  name: 'DictionaryDetailModal',
});
const dictionaryId = ref();

// ---------------- form -----------------

const showDeleteButton = ref<boolean>(false);

const gridEvents: VxeGridListeners<any> = {
  checkboxChange(e) {
    showDeleteButton.value = e.$table.getCheckboxRecords().length > 0;
  },
  checkboxAll(e) {
    showDeleteButton.value = e.$table.getCheckboxRecords().length > 0;
  },
  // insertEvent(e) {
  //   // e.$table.
  // },
};

// ------------- table --------------------

const gridOptions: VxeGridProps<DictionaryDetailInfo> = {
  checkboxConfig: {
    highlight: true,
  },
  toolbarConfig: {
    slots: {
      buttons: 'toolbar-buttons',
    },
  },
  editConfig: {
    mode: 'row',
    trigger: 'click',
  },
  showOverflow: true,
  columns: [
    ...(detailTableColumns.columns as any),
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
                icon: 'clarity:note-edit-line',
                tooltip: $t('common.edit'),
                onClick: editRowEvent.bind(null, row),
                // eslint-disable-next-line no-use-before-define
                ifShow: !gridApi.grid?.isEditByRow(row),
              },
              {
                type: 'link',
                icon: 'lucide-lab:save',
                tooltip: $t('common.saveText'),
                onClick: saveRowEvent.bind(null, row),
                // eslint-disable-next-line no-use-before-define
                ifShow: gridApi.grid?.isEditByRow(row),
              },
              {
                type: 'link',
                icon: 'material-symbols:cancel-outline',
                tooltip: $t('common.cancelText'),
                onClick: editRowEvent.bind(null, row),
                // eslint-disable-next-line no-use-before-define
                ifShow: gridApi.grid?.isEditByRow(row),
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
                // eslint-disable-next-line no-use-before-define
                ifShow: !gridApi.grid?.isEditByRow(row),
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
      query: async ({ page }) => {
        const res = await getDictionaryDetailList({
          page: page.currentPage,
          pageSize: page.pageSize,
          dictionaryId: dictionaryId.value,
        } as any);
        return res.data;
      },
    },
  },
};

const [Grid, gridApi] = useVbenVxeGrid({
  gridOptions,
  gridEvents,
});

function editRowEvent(row: DictionaryDetailInfo) {
  gridApi.grid?.setEditRow(row);
}

async function saveRowEvent(row: DictionaryDetailInfo) {
  await gridApi.grid?.clearEdit();

  const result =
    row.id === 0
      ? await createDictionaryDetail(row)
      : await updateDictionaryDetail(row);
  if (result.code === 0) {
    message.success($t('common.successful'));
    await gridApi.reload();
  }
}

function insertRowEvent() {
  const newRow: DictionaryDetailInfo = {
    id: 0,
    createdAt: Date.now(),
    status: 1,
    dictionaryId: dictionaryId.value,
  };
  gridApi.grid?.insertAt(newRow, -1);
  gridApi.grid?.setEditRow(newRow);
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
  const result = await deleteDictionaryDetail({
    ids,
  });
  if (result.code === 0) {
    showDeleteButton.value = false;
  }
  await gridApi.reload();
}

const [TableModal, modalApi] = useVbenModal({
  fullscreenButton: false,
  onCancel() {
    modalApi.close();
  },
  onConfirm: async () => {
    modalApi.close();
  },
  onOpenChange(isOpen: boolean) {
    if (isOpen) {
      dictionaryId.value = modalApi.getData()?.id;
    }
  },
  title: $t('sys.dictionary.editDictionaryDetail'),
});

defineExpose(modalApi);
</script>

<template>
  <TableModal class="h-1/2 w-1/2">
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
        <Button type="primary" @click="insertRowEvent">
          {{ $t('sys.dictionary.addDictionaryDetail') }}
        </Button>
      </template>
    </Grid>
  </TableModal>
</template>
