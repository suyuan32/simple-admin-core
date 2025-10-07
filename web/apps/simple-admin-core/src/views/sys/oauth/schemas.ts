import type { VxeGridProps } from '#/adapter/vxe-table';

import { type VbenFormProps } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { z } from '#/adapter/form';

export const tableColumns: VxeGridProps = {
  columns: [
    {
      type: 'checkbox',
      width: 60,
    },

    {
      title: $t('common.name'),
      field: 'name',
    },
    {
      title: $t('sys.oauth.clientId'),
      field: 'clientId',
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
      label: $t('common.name'),
      component: 'Input',
      rules: z.string().max(30).optional(),
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
      label: $t('common.name'),
      component: 'Input',
      rules: z.string().max(30),
    },
    {
      fieldName: 'clientId',
      label: $t('sys.oauth.clientId'),
      component: 'Input',
      rules: z.string().max(80),
    },
    {
      fieldName: 'clientSecret',
      label: $t('sys.oauth.clientSecret'),
      component: 'Input',
      rules: z.string().max(100),
    },
    {
      fieldName: 'redirectUrl',
      label: $t('sys.oauth.redirectURL'),
      component: 'Input',
      rules: z.string().max(300),
    },
    {
      fieldName: 'scopes',
      label: $t('sys.oauth.scope'),
      component: 'Input',
      rules: z.string().max(100),
    },
    {
      fieldName: 'authUrl',
      label: $t('sys.oauth.authURL'),
      component: 'Input',
      rules: z.string().max(300),
    },
    {
      fieldName: 'tokenUrl',
      label: $t('sys.oauth.tokenURL'),
      component: 'Input',
      rules: z.string().max(300),
    },
    {
      fieldName: 'infoUrl',
      label: $t('sys.oauth.infoURL'),
      component: 'Input',
      rules: z.string().max(300),
    },
    {
      fieldName: 'authStyle',
      label: $t('sys.oauth.authStyle'),
      component: 'Select',
      componentProps: {
        options: [
          { label: $t('sys.oauth.params'), value: 1 },
          { label: $t('sys.oauth.header'), value: 2 },
        ],
        class: 'w-full',
      },
    },
  ],
};
