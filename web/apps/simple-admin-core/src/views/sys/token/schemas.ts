import type { VxeGridProps } from '#/adapter/vxe-table';

import { h } from 'vue';

import { type VbenFormProps } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { Switch } from 'ant-design-vue';

import { z } from '#/adapter/form';
import { updateToken } from '#/api/sys/token';

export const tableColumns: VxeGridProps = {
  columns: [
    {
      type: 'checkbox',
      width: 60,
    },

    {
      title: $t('sys.login.username'),
      field: 'username',
    },
    {
      title: 'Token',
      field: 'token',
    },
    {
      title: $t('common.source'),
      field: 'source',
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
              updateToken({ id: e.row.id, status: newStatus }).then(() => {
                e.row.status = newStatus;
              });
            },
          }),
      },
    },
    {
      title: $t('common.expiredAt'),
      field: 'expiredAt',
      formatter: 'formatDateTime',
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
      fieldName: 'username',
      label: $t('sys.login.username'),
      component: 'Input',
      rules: z.string().max(30).optional(),
    },
    {
      fieldName: 'nickname',
      label: $t('sys.user.nickname'),
      component: 'Input',
      rules: z.string().max(30).optional(),
    },
    {
      fieldName: 'email',
      label: $t('sys.login.email'),
      component: 'Input',
      rules: z.string().max(50).optional(),
    },
    {
      fieldName: 'uuid',
      label: 'UUID',
      component: 'Input',
      rules: z.string().max(30).optional(),
    },
  ],
};
