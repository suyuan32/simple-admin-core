import type { VxeGridProps } from '#/adapter/vxe-table';

import { h } from 'vue';

import { type VbenFormProps } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { Switch } from 'ant-design-vue';

import { z } from '#/adapter/form';
import { updateConfiguration } from '#/api/sys/configuration';

export const tableColumns: VxeGridProps = {
  columns: [
    {
      type: 'checkbox',
      width: 60,
    },
    {
      title: $t('sys.configuration.name'),
      field: 'name',
    },
    {
      title: $t('sys.configuration.key'),
      field: 'key',
    },
    {
      title: $t('sys.configuration.category'),
      field: 'category',
    },
    {
      title: $t('sys.configuration.remark'),
      field: 'remark',
    },
    {
      title: $t('common.status'),
      field: 'state',
      slots: {
        default: (e) =>
          h(Switch, {
            checked: e.row.state,
            onClick: () => {
              const newStatus = !e.row.state;
              updateConfiguration({ id: e.row.id, state: newStatus }).then(
                () => {
                  e.row.state = newStatus;
                },
              );
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
      label: $t('sys.configuration.name'),
      component: 'Input',
      rules: z.string().optional(),
    },
    {
      fieldName: 'key',
      label: $t('sys.configuration.key'),
      component: 'Input',
      rules: z.string().optional(),
    },
    {
      fieldName: 'category',
      label: $t('sys.configuration.category'),
      component: 'Input',
      rules: z.string().optional(),
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
      label: $t('sys.configuration.name'),
      component: 'Input',
      rules: 'required',
    },
    {
      fieldName: 'key',
      label: $t('sys.configuration.key'),
      component: 'Input',
      rules: 'required',
    },
    {
      fieldName: 'value',
      label: $t('sys.configuration.value'),
      component: 'Input',
      rules: 'required',
    },
    {
      fieldName: 'category',
      label: $t('sys.configuration.category'),
      component: 'Input',
      rules: 'required',
    },
    {
      fieldName: 'remark',
      label: $t('common.remark'),
      component: 'Input',
      rules: z.string().optional(),
    },
    {
      fieldName: 'sort',
      label: $t('common.sort'),
      component: 'InputNumber',
      rules: 'required',
      defaultValue: 1,
    },
    {
      fieldName: 'state',
      label: $t('common.status'),
      component: 'RadioButtonGroup',
      defaultValue: true,
      componentProps: {
        options: [
          { label: $t('common.on'), value: true },
          { label: $t('common.off'), value: false },
        ],
      },
    },
  ],
};
