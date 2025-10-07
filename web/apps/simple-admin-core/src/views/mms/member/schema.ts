import type { VxeGridProps } from '#/adapter/vxe-table';

import { h } from 'vue';

import { type VbenFormProps } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { Switch } from 'ant-design-vue';

import { z } from '#/adapter/form';
import { updateMember } from '#/api/member/member';
import { getMemberRankList } from '#/api/member/memberRank';

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
      title: $t('sys.user.nickname'),
      field: 'nickname',
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
              updateMember({ id: e.row.id, status: newStatus }).then(() => {
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
      fieldName: 'username',
      label: $t('sys.login.username'),
      component: 'Input',
    },
    {
      fieldName: 'nickname',
      label: $t('sys.user.nickname'),
      component: 'Input',
    },
    {
      fieldName: 'mobile',
      label: $t('sys.login.mobile'),
      component: 'Input',
    },
    {
      fieldName: 'email',
      label: $t('sys.login.email'),
      component: 'Input',
    },
  ],
};

export const dataFormSchemas: VbenFormProps = {
  schema: [
    {
      fieldName: 'avatar',
      label: $t('sys.user.avatar'),
      component: 'ImageUpload',
      componentProps: {
        accept: ['png', 'jpeg', 'jpg'],
        maxSize: 2,
        maxNumber: 1,
        multiple: false,
        provider: 'cloud-default',
      },
    },
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
      fieldName: 'username',
      label: $t('sys.login.username'),
      component: 'Input',
      rules: z.string().max(50),
    },
    {
      fieldName: 'nickname',
      label: $t('sys.user.nickname'),
      component: 'Input',
      rules: z.string().max(40),
    },
    {
      fieldName: 'rankId',
      label: $t('sys.member.rankId'),
      component: 'ApiSelect',
      defaultValue: 1,
      componentProps: {
        api: getMemberRankList,
        params: {
          page: 1,
          pageSize: 1000,
        },
        resultField: 'data.data',
        labelField: 'trans',
        valueField: 'id',
      },
    },
    {
      fieldName: 'expiredAt',
      label: $t('sys.member.expiredAt'),
      component: 'SimpleTimePicker',
      componentProps: {
        timeMode: 'date',
      },
    },
    {
      fieldName: 'mobile',
      label: $t('sys.login.mobile'),
      component: 'Input',
      rules: z.string().max(20).optional(),
    },
    {
      fieldName: 'email',
      label: $t('sys.login.email'),
      component: 'Input',
      rules: z.string().email(),
    },
    {
      fieldName: 'password',
      label: $t('sys.login.password'),
      component: 'Input',
      rules: z.string().min(6).max(100),
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
      rules: 'required',
    },
  ],
};
