<script lang="ts" setup>
import type { VxeGridProps } from '#/adapter/vxe-table';

import { h } from 'vue';

import { Page, useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { Button } from 'ant-design-vue';
import { isPlainObject } from 'remeda';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { deleteMenu, getMenuList } from '#/api/sys/menu';
import { type ActionItem, TableAction } from '#/components/table/table-action';

import MenuForm from './form.vue';
import { tableColumns } from './schemas';

defineOptions({
  name: 'MenuManagement',
});

// ---------------- form -----------------

const [FormModal, formModalApi] = useVbenModal({
  connectedComponent: MenuForm,
});

// ------------- table --------------------

const gridOptions: VxeGridProps = {
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
                size: 'small',
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
  pagerConfig: {
    enabled: false,
  },
  proxyConfig: {
    ajax: {
      query: async (_formValues) => {
        const res = await getMenuList();
        return res.data;
      },
    },
  },
  treeConfig: {
    transform: true,
    parentField: 'parentId',
    rowField: 'id',
  },
};

const [Grid, gridApi] = useVbenVxeGrid({
  gridOptions,
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

async function batchDelete(ids: any) {
  const result = await deleteMenu({
    id: ids[0],
  });
  if (result.code === 0) {
    await gridApi.reload();
  }
}
</script>

<template>
  <Page auto-content-height>
    <FormModal />
    <Grid>
      <template #toolbar-tools>
        <Button type="primary" @click="openFormModal">
          {{ $t('sys.menu.addMenu') }}
        </Button>
      </template>
    </Grid>
  </Page>
</template>
