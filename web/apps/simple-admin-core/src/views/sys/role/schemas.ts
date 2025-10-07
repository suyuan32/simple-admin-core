import type { VxeGridProps } from '#/adapter/vxe-table';
import type { VbenFormProps } from '@vben/common-ui';

import { z } from '#/adapter/form';
import { updateRole } from '#/api/sys/role';
import { $t } from '@vben/locales';
import { Switch } from 'ant-design-vue';
import { h } from 'vue';

export const tableColumns: VxeGridProps = {
  columns: [
    {
      type: 'checkbox',
      width: 60,
    },

    {
      title: $t('sys.role.roleName'),
      field: 'trans',
    },
    {
      title: $t('sys.role.roleValue'),
      field: 'code',
    },
    {
      title: $t('common.status'),
      field: 'status',
      slots: {
        default: (e) =>
          h(Switch, {
            checked: e.row.status === 1,
            onClick: () => {
              const newStatus = e.row.status === 1 ? 2 : 1;
              updateRole({ id: e.row.id, status: newStatus }).then(() => {
                e.row.status = newStatus;
              });
            },
          }),
      },
    },

    {
      title: $t('common.createTime'),
      field: 'createdAt',
      formatter: 'formatDateTime',
    },
  ],
};

export const searchFormSchemas: VbenFormProps = {
  schema: [
    {
      fieldName: 'name',
      label: $t('sys.role.roleName'),
      component: 'Input',
    },
  ],
};

export const dataFormSchemas: VbenFormProps = {
  schema: [
    {
      fieldName: 'id',
      label: 'ID',
      component: 'Input',
      dependencies: {
        show: false,
        triggerFields: ['id'],
      },
    },
    {
      fieldName: 'name',
      label: $t('sys.role.roleName'),
      component: 'Input',
      rules: 'required',
    },
    {
      fieldName: 'code',
      label: $t('sys.role.roleValue'),
      component: 'Input',
      rules: 'required',
      help: $t('sys.role.roleValueHelpMessage'),
    },
    {
      fieldName: 'remark',
      label: $t('common.remark'),
      component: 'Input',
    },
    {
      fieldName: 'sort',
      label: $t('sys.menu.order'),
      component: 'InputNumber',
      rules: z.number().max(10_000).min(1),
    },
    {
      fieldName: 'status',
      label: $t('common.status'),
      component: 'RadioButtonGroup',
      defaultValue: 1,
      componentProps: {
        options: [
          { label: $t('common.on'), value: 1 },
          { label: $t('common.off'), value: 2 },
        ],
      },
    },
  ],
};
