<script lang="ts" setup>
import type { VxeGridListeners, VxeGridProps } from '#/adapter/vxe-table';
import type { DictionaryDetailInfo } from '#/api/sys/model/dictionaryDetailModel';

import { h, ref } from 'vue';

import { useVbenModal, type VbenFormProps } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { Button, Modal } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { deleteTaskLog, getTaskLogList } from '#/api/sys/taskLog';
import { type ActionItem, TableAction } from '#/components/table/table-action';

import { taskLogSearchFormSchemas, taskLogTableColumns } from './schemas';

defineOptions({
  name: 'TaskLogModal',
});
const taskId = ref();

// ---------------- form -----------------

const formOptions: VbenFormProps = {
  // 默认展开
  collapsed: false,
  schema: [...(taskLogSearchFormSchemas.schema as any)],
  // 控制表单是否显示折叠按钮
  showCollapseButton: false,
  // 按下回车时是否提交表单
  submitOnEnter: false,
};

const showDeleteButton = ref<boolean>(false);

const gridEvents: VxeGridListeners<any> = {
  checkboxChange(e) {
    showDeleteButton.value = e.$table.getCheckboxRecords().length > 0;
  },
  checkboxAll(e) {
    showDeleteButton.value = e.$table.getCheckboxRecords().length > 0;
  },
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
    ...(taskLogTableColumns.columns as any),
    {
      title: $t('common.action'),
      fixed: 'right',
      field: 'action',
      slots: {
        default: ({ row }) =>
          h(TableAction, {
            actions: [
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
            ] as unknown as ActionItem[],
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
        const res = await getTaskLogList({
          page: page.currentPage,
          pageSize: page.pageSize,
          taskId: taskId.value,
          result: formValues.result ?? 0,
        } as any);
        return res.data;
      },
    },
  },
};

const [Grid, gridApi] = useVbenVxeGrid({
  gridOptions,
  formOptions,
  gridEvents,
});

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
  const result = await deleteTaskLog({
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
      taskId.value = modalApi.getData()?.taskId;
    }
  },
  title: $t('sys.taskLog.taskLogList'),
});

defineExpose(modalApi);
</script>

<template>
  <TableModal class="h-2/3 w-1/2">
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
    </Grid>
  </TableModal>
</template>
