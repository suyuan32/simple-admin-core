import type { VxeGridProps } from '#/adapter/vxe-table';

import { h } from 'vue';

import { type VbenFormProps } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { Tag } from 'ant-design-vue';

import { z } from '#/adapter/form';

export const tableColumns: VxeGridProps = {
  columns: [
    {
      type: 'checkbox',
      width: 60,
    },
    {
      title: $t('sys.apis.path'),
      field: 'path',
    },
    {
      title: $t('sys.apis.group'),
      field: 'group',
    },
    {
      title: $t('sys.apis.serviceName'),
      field: 'serviceName',
      width: 120,
    },
    {
      title: $t('sys.apis.description'),
      field: 'trans',
    },
    {
      title: $t('sys.apis.method'),
      field: 'method',
    },
    {
      title: $t('common.required'),
      field: 'isRequired',
      width: 80,
      slots: {
        default: (record) => {
          let resultText = '';
          resultText = record.row.isRequired
            ? $t('common.yes')
            : $t('common.no');
          return h(
            Tag,
            {
              color: record.row.isRequired ? 'green' : 'red',
            },
            () => resultText,
          );
        },
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
      fieldName: 'path',
      label: $t('sys.apis.path'),
      component: 'Input',
      rules: z.string().max(200).optional(),
    },
    {
      fieldName: 'serviceName',
      label: $t('sys.apis.serviceName'),
      component: 'Input',
      rules: z.string().max(20).optional(),
    },
    {
      fieldName: 'group',
      label: $t('sys.apis.group'),
      component: 'Input',
      rules: z.string().max(80).optional(),
    },
    {
      fieldName: 'description',
      label: $t('sys.apis.description'),
      component: 'Input',
      rules: z.string().max(200).optional(),
    },
    {
      fieldName: 'method',
      label: $t('sys.apis.method'),
      component: 'Select',
      componentProps: {
        options: [
          { label: 'GET', value: 'GET' },
          { label: 'POST', value: 'POST' },
          { label: 'DELETE', value: 'DELETE' },
          { label: 'PUT', value: 'PUT' },
          { label: 'PATCH', value: 'PATCH' },
        ],
      },
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
      fieldName: 'path',
      label: $t('sys.apis.path'),
      component: 'Input',
      rules: z.string().max(200),
    },
    {
      fieldName: 'serviceName',
      label: $t('sys.apis.serviceName'),
      component: 'Input',
      help: $t('sys.apis.serviceNameHelpMessage'),
      rules: z.string().max(20),
    },
    {
      fieldName: 'group',
      label: $t('sys.apis.group'),
      component: 'Input',
      rules: z.string().max(80),
    },
    {
      fieldName: 'description',
      label: $t('sys.apis.description'),
      component: 'Input',
      rules: z.string().max(100),
    },
    {
      fieldName: 'method',
      label: $t('sys.apis.method'),
      component: 'Select',
      componentProps: {
        options: [
          { label: 'GET', value: 'GET' },
          { label: 'POST', value: 'POST' },
          { label: 'DELETE', value: 'DELETE' },
          { label: 'PUT', value: 'PUT' },
          { label: 'PATCH', value: 'PATCH' },
        ],
      },
    },
    {
      fieldName: 'isRequired',
      label: $t('common.required'),
      component: 'RadioButtonGroup',
      defaultValue: false,
      componentProps: {
        options: [
          { label: $t('common.yes'), value: true },
          { label: $t('common.no'), value: false },
        ],
      },
      help: $t('sys.apis.isRequiredHelpMessage'),
    },
  ],
};
