<script lang="ts" setup>
import type { VxeGridListeners, VxeGridProps } from '#/adapter/vxe-table';
import type { ApiInfo } from '#/api/sys/model/apiModel';

import { h, ref } from 'vue';

import { Page, useVbenModal, type VbenFormProps } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { Button, Modal } from 'ant-design-vue';
import { isPlainObject } from 'remeda';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { deleteApi, getApiList } from '#/api/sys/api';
import { type ActionItem, TableAction } from '#/components/table/table-action';

import ApiForm from './form.vue';
import { searchFormSchemas, tableColumns } from './schemas';

defineOptions({
  name: 'ApiManagement',
});

// ---------------- form -----------------

const [FormModal, formModalApi] = useVbenModal({
  connectedComponent: ApiForm,
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

const gridOptions: VxeGridProps<ApiInfo> = {
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
                tooltip: $t('common.edit'),
                icon: 'clarity:note-edit-line',
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
        const res = await getApiList({
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
  const result = await deleteApi({
    ids,
  });
  if (result.code === 0) {
    await gridApi.reload();
    showDeleteButton.value = false;
  }
}
</script>

<template>
  <Page auto-content-height>
    <FormModal />
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
        <Button type="primary" @click="openFormModal">
          {{ $t('sys.apis.addApi') }}
        </Button>
      </template>
    </Grid>
  </Page>
</template>
