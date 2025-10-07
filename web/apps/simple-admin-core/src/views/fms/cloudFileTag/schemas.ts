import type { VxeGridProps } from '#/adapter/vxe-table';

import { h } from 'vue';

import { type VbenFormProps } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { Switch } from 'ant-design-vue';

import { updateCloudFileTag } from '#/api/fms/cloudFileTag';

export const tableColumns: VxeGridProps = {
  columns: [
    {
      type: 'checkbox',
      width: 60,
    },

    {
      title: $t('fms.cloudFileTag.name'),
      field: 'name',
    },
    {
      title: $t('fms.cloudFileTag.remark'),
      field: 'remark',
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
              updateCloudFileTag({ id: e.row.id, status: newStatus }).then(
                () => {
                  e.row.status = newStatus;
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
      label: $t('fms.cloudFileTag.name'),
      component: 'Input',
    },
    {
      fieldName: 'remark',
      label: $t('fms.cloudFileTag.remark'),
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
      label: $t('fms.cloudFileTag.name'),
      component: 'Input',
    },
    {
      fieldName: 'remark',
      label: $t('fms.cloudFileTag.remark'),
      component: 'Input',
    },
    {
      fieldName: 'status',
      label: $t('fms.cloudFileTag.status'),
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
