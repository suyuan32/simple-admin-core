<script lang="ts" setup>
import type { VxeGridListeners, VxeGridProps } from '#/adapter/vxe-table';
import type { RoleInfo } from '#/api/sys/model/roleModel';

import { h, ref } from 'vue';

import { Page, useVbenModal, type VbenFormProps } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { Button, Modal } from 'ant-design-vue';
import { isPlainObject } from 'remeda';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { deleteRole, getRoleList } from '#/api/sys/role';
import { type ActionItem, TableAction } from '#/components/table/table-action';
import ApiAuthorityModal from '#/views/sys/role/api-authority.vue';
import MenuAuthorityModal from '#/views/sys/role/menu-authority.vue';

import RoleForm from './form.vue';
import { searchFormSchemas, tableColumns } from './schemas';

defineOptions({
  name: 'RoleManagement',
});

// ------------ api authority --------------
const [ApiModal, apiModalApi] = useVbenModal({
  connectedComponent: ApiAuthorityModal,
});

// ---------------- menu authority --------------
const [MenuModal, menuModalApi] = useVbenModal({
  connectedComponent: MenuAuthorityModal,
});

// ---------------- form -----------------

const [FormModal, formModalApi] = useVbenModal({
  connectedComponent: RoleForm,
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

const gridOptions: VxeGridProps<RoleInfo> = {
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
                size: 'small',
                icon: 'clarity:note-edit-line',
                tooltip: $t('common.edit'),
                onClick: openFormModal.bind(null, row),
              },
              {
                type: 'link',
                size: 'small',
                icon: 'ep:menu',
                tooltip: $t('sys.authority.menuAuthority'),
                onClick: openMenuModal.bind(null, row),
              },
              {
                type: 'link',
                size: 'small',
                icon: 'carbon:network-interface',
                tooltip: $t('sys.authority.apiAuthority'),
                onClick: openApiModal.bind(null, row),
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
        const res = await getRoleList({
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

function openMenuModal(record: any) {
  if (isPlainObject(record)) {
    menuModalApi.setData({ roleId: record.id });
  }
  menuModalApi.open();
}

function openApiModal(record: any) {
  if (isPlainObject(record)) {
    apiModalApi.setData({ roleId: record.id });
  }
  apiModalApi.open();
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
  const result = await deleteRole({
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
    <MenuModal />
    <FormModal />
    <ApiModal />
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
          {{ $t('sys.role.addRole') }}
        </Button>
      </template>
    </Grid>
  </Page>
</template>
