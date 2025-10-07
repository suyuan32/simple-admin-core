import type { VxeGridProps } from '#/adapter/vxe-table';

import { type VbenFormProps } from '@vben/common-ui';
import { $t } from '@vben/locales';

export const tableColumns: VxeGridProps = {
  columns: [
    {
      type: 'checkbox',
      width: 60,
    },

    {
      title: $t('sys.memberRank.name'),
      field: 'trans',
    },
    {
      title: $t('sys.memberRank.description'),
      field: 'description',
    },
    {
      title: $t('sys.memberRank.remark'),
      field: 'remark',
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
      label: $t('sys.memberRank.name'),
      component: 'Input',
    },
    {
      fieldName: 'description',
      label: $t('sys.memberRank.description'),
      component: 'Input',
    },
    {
      fieldName: 'remark',
      label: $t('sys.memberRank.remark'),
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
      label: $t('sys.memberRank.name'),
      component: 'Input',
      rules: 'required',
    },
    {
      fieldName: 'code',
      label: $t('sys.memberRank.code'),
      component: 'Input',
      rules: 'required',
    },
    {
      fieldName: 'description',
      label: $t('sys.memberRank.description'),
      component: 'Input',
    },
    {
      fieldName: 'remark',
      label: $t('sys.memberRank.remark'),
      component: 'Input',
    },
  ],
};
