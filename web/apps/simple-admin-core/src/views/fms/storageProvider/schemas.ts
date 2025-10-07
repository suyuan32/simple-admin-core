import type { VxeGridProps } from '#/adapter/vxe-table';

import { h } from 'vue';

import { type VbenFormProps } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { Switch, Tag } from 'ant-design-vue';

import { updateStorageProvider } from '#/api/fms/storageProvider';

export const tableColumns: VxeGridProps = {
  columns: [
    {
      type: 'checkbox',
      width: 60,
    },

    {
      title: $t('fms.storageProvider.name'),
      field: 'name',
    },
    {
      title: $t('fms.storageProvider.isDefault'),
      field: 'isDefault',
      slots: {
        default: (record) => {
          let resultText = '';
          resultText = record.row.isDefault
            ? $t('common.yes')
            : $t('common.no');
          return h(
            Tag,
            {
              color: record.row.isDefault ? 'green' : 'red',
            },
            () => resultText,
          );
        },
      },
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
              updateStorageProvider({ id: e.row.id, state: newStatus }).then(
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
      label: $t('fms.storageProvider.name'),
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
      label: $t('fms.storageProvider.name'),
      component: 'Input',
      help: $t('fms.storageProvider.nameHelpMessage'),
      rules: 'required',
    },
    {
      fieldName: 'bucket',
      label: $t('fms.storageProvider.bucket'),
      component: 'Input',
      rules: 'required',
    },
    {
      fieldName: 'secretId',
      label: $t('fms.storageProvider.secretId'),
      component: 'Input',
      rules: 'required',
    },
    {
      fieldName: 'secretKey',
      label: $t('fms.storageProvider.secretKey'),
      component: 'Input',
      rules: 'required',
    },
    {
      fieldName: 'endpoint',
      label: $t('fms.storageProvider.endpoint'),
      component: 'Input',
      rules: 'required',
    },
    {
      fieldName: 'folder',
      label: $t('fms.storageProvider.folder'),
      component: 'Input',
    },
    {
      fieldName: 'region',
      label: $t('fms.storageProvider.region'),
      component: 'Input',
    },
    {
      fieldName: 'isDefault',
      label: $t('fms.storageProvider.isDefault'),
      component: 'RadioButtonGroup',
      defaultValue: false,
      componentProps: {
        options: [
          { label: $t('common.on'), value: true },
          { label: $t('common.off'), value: false },
        ],
      },
    },
    {
      fieldName: 'useCdn',
      label: $t('fms.storageProvider.useCdn'),
      component: 'RadioButtonGroup',
      defaultValue: true,
      componentProps: {
        options: [
          { label: $t('common.on'), value: true },
          { label: $t('common.off'), value: false },
        ],
      },
    },
    {
      fieldName: 'cdnUrl',
      label: $t('fms.storageProvider.cdnUrl'),
      component: 'Input',
    },
    {
      fieldName: 'state',
      label: $t('fms.storageProvider.state'),
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
