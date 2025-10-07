import type { VxeGridProps } from '#/adapter/vxe-table';

import { h } from 'vue';

import { type VbenFormProps } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { Tag } from 'ant-design-vue';

export const tableColumns: VxeGridProps = {
  columns: [
    {
      type: 'checkbox',
      width: 60,
    },

    {
      title: $t('mcms.smsProvider.name'),
      field: 'name',
    },
    {
      title: $t('mcms.smsProvider.isDefault'),
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
      label: $t('mcms.smsProvider.name'),
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
      label: $t('mcms.smsProvider.name'),
      component: 'Select',
      rules: 'required',
      componentProps: {
        options: [
          { label: $t('mcms.smsProvider.tencent'), value: 'tencent' },
          { label: $t('mcms.smsProvider.aliyun'), value: 'aliyun' },
          { label: $t('mcms.smsProvider.uni'), value: 'uni' },
          { label: $t('mcms.smsProvider.smsbao'), value: 'smsbao' },
        ],
        class: 'w-full',
      },
    },
    {
      fieldName: 'secretId',
      label: $t('mcms.smsProvider.secretId'),
      component: 'Input',
      rules: 'required',
    },
    {
      fieldName: 'secretKey',
      label: $t('mcms.smsProvider.secretKey'),
      component: 'Input',
      rules: 'required',
    },
    {
      fieldName: 'region',
      label: $t('mcms.smsProvider.region'),
      component: 'Input',
    },
    {
      fieldName: 'isDefault',
      label: $t('mcms.smsProvider.isDefault'),
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

// ------------- sms log ------------------------

export const smsLogTableColumns: VxeGridProps = {
  columns: [
    {
      type: 'checkbox',
      width: 60,
    },

    {
      title: $t('mcms.smsLog.phoneNumber'),
      field: 'phoneNumber',
    },
    {
      title: $t('mcms.smsLog.content'),
      field: 'content',
    },
    {
      title: $t('mcms.emailLog.sendStatus'),
      field: 'sendStatus',
      slots: {
        default: (record) => {
          let resultText = '';
          resultText =
            record.row.sendStatus === 1 ? $t('common.yes') : $t('common.no');
          return h(
            Tag,
            {
              color: record.row.sendStatus === 1 ? 'green' : 'red',
            },
            () => resultText,
          );
        },
      },
    },
    {
      title: $t('mcms.smsLog.provider'),
      field: 'provider',
    },
    {
      title: $t('common.createTime'),
      field: 'createdAt',
      formatter: 'formatDateTime',
    },
  ],
};

export const smsLogSearchFormSchemas: VbenFormProps = {
  schema: [
    {
      fieldName: 'phoneNumber',
      label: $t('mcms.smsLog.phoneNumber'),
      component: 'Input',
    },
    {
      fieldName: 'content',
      label: $t('mcms.smsLog.content'),
      component: 'Input',
    },
    {
      fieldName: 'provider',
      label: $t('mcms.smsLog.provider'),
      component: 'Input',
    },
    {
      fieldName: 'sendStatus',
      label: $t('mcms.emailLog.sendStatus'),
      component: 'Select',
      defaultValue: 0,
      componentProps: {
        options: [
          { label: $t('common.all'), value: 0 },
          { label: $t('common.successful'), value: 1 },
          { label: $t('common.failed'), value: 2 },
        ],
        class: 'w-full',
      },
    },
  ],
};

// -------------- sms sender --------------

export const smsSenderFormSchemas: VbenFormProps = {
  schema: [
    {
      fieldName: 'phoneNumber',
      label: $t('mcms.smsLog.phoneNumber'),
      component: 'Input',
      rules: 'required',
    },
    {
      fieldName: 'params',
      label: $t('mcms.sms.params'),
      component: 'Input',
      help: $t('mcms.sms.paramsHelp'),
      rules: 'required',
    },
    {
      fieldName: 'templateId',
      label: $t('mcms.sms.templateId'),
      component: 'Input',
      help: $t('mcms.sms.templateIdHelp'),
    },
    {
      fieldName: 'signName',
      label: $t('mcms.sms.signName'),
      component: 'Input',
    },
    {
      fieldName: 'provider',
      label: $t('mcms.smsLog.provider'),
      component: 'Select',
      rules: 'required',
      componentProps: {
        options: [
          { label: $t('mcms.smsProvider.tencent'), value: 'tencent' },
          { label: $t('mcms.smsProvider.aliyun'), value: 'aliyun' },
          { label: $t('mcms.smsProvider.uni'), value: 'uni' },
          { label: $t('mcms.smsProvider.smsbao'), value: 'smsbao' },
        ],
        class: 'w-full',
      },
    },
  ],
};
