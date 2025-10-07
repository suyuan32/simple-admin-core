import type { VxeGridProps } from '#/adapter/vxe-table';
import type { VbenFormProps } from '@vben/common-ui';

import { getEmailProviderList } from '#/api/mcms/emailProvider';
import { $t } from '@vben/locales';
import { Tag } from 'ant-design-vue';
import { h } from 'vue';

export const tableColumns: VxeGridProps = {
  columns: [
    {
      type: 'checkbox',
      width: 60,
    },
    {
      title: $t('mcms.emailProvider.name'),
      field: 'name',
    },

    {
      title: $t('mcms.emailProvider.isDefault'),
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
      label: $t('mcms.emailProvider.name'),
      component: 'Input',
    },
    {
      fieldName: 'emailAddr',
      label: $t('mcms.emailProvider.emailAddr'),
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
      label: $t('mcms.emailProvider.name'),
      component: 'Input',
      rules: 'required',
    },
    {
      fieldName: 'authType',
      label: $t('mcms.emailProvider.authType'),
      component: 'Select',
      componentProps: {
        options: [
          { label: 'plain', value: 1 },
          { label: 'CRAMMD5', value: 2 },
        ],
        class: 'w-full',
      },
      help: $t('mcms.emailProvider.authTypeHelp'),
      rules: 'required',
    },
    {
      fieldName: 'emailAddr',
      label: $t('mcms.emailProvider.emailAddr'),
      component: 'Input',
      rules: 'required',
      dependencies: {
        triggerFields: ['authType'],
        if(values) {
          return values.authType === 1;
        },
      },
    },
    {
      fieldName: 'password',
      label: $t('mcms.emailProvider.password'),
      component: 'Input',
      rules: 'required',
      dependencies: {
        triggerFields: ['authType'],
        if(values) {
          return values.authType === 1;
        },
      },
      help: $t('mcms.emailProvider.passwordHelp'),
    },
    {
      fieldName: 'hostName',
      label: $t('mcms.emailProvider.hostName'),
      component: 'Input',
      rules: 'required',
      dependencies: {
        triggerFields: ['authType'],
        if(values) {
          return values.authType === 1;
        },
      },
    },
    {
      fieldName: 'identify',
      label: $t('mcms.emailProvider.identify'),
      component: 'Input',
      dependencies: {
        triggerFields: ['authType'],
        if(values) {
          return values.authType === 2;
        },
      },
    },
    {
      fieldName: 'secret',
      label: $t('mcms.emailProvider.secret'),
      component: 'Input',
      rules: 'required',
      dependencies: {
        triggerFields: ['authType'],
        if(values) {
          return values.authType === 2;
        },
      },
    },
    {
      fieldName: 'port',
      label: $t('mcms.emailProvider.port'),
      component: 'InputNumber',
      componentProps: {
        class: 'w-full',
      },
      dependencies: {
        triggerFields: ['authType'],
        if(values) {
          return values.authType === 1;
        },
      },
    },
    {
      fieldName: 'tls',
      label: $t('mcms.emailProvider.tls'),
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
      fieldName: 'isDefault',
      label: $t('mcms.emailProvider.isDefault'),
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

// / ------------------- email log -------------

export const emailLogTableColumns: VxeGridProps = {
  columns: [
    {
      type: 'checkbox',
      width: 60,
    },

    {
      title: $t('mcms.emailLog.target'),
      field: 'target',
    },
    {
      title: $t('mcms.emailLog.subject'),
      field: 'subject',
    },
    {
      title: $t('mcms.emailLog.content'),
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

export const emailLogSearchFormSchemas: VbenFormProps = {
  schema: [
    {
      fieldName: 'target',
      label: $t('mcms.emailLog.target'),
      component: 'Input',
    },
    {
      fieldName: 'subject',
      label: $t('mcms.emailLog.subject'),
      component: 'Input',
    },
    {
      fieldName: 'provider',
      label: $t('mcms.emailLog.provider'),
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

// -------------- email sender ----------

export const emailSenderFormSchemas: VbenFormProps = {
  schema: [
    {
      fieldName: 'target',
      label: $t('mcms.email.targetAddress'),
      component: 'Input',
      rules: 'required',
    },
    {
      fieldName: 'subject',
      label: $t('mcms.email.subject'),
      component: 'Input',
      rules: 'required',
    },
    {
      fieldName: 'content',
      label: $t('mcms.email.content'),
      component: 'Textarea',
      componentProps: {
        rows: 4,
      },
      rules: 'required',
    },
    {
      fieldName: 'provider',
      label: $t('mcms.emailLog.provider'),
      component: 'ApiSelect',
      rules: 'required',
      componentProps: {
        api: getEmailProviderList,
        params: {
          page: 1,
          pageSize: 1000,
        },
        resultField: 'data.data',
        labelField: 'name',
        valueField: 'name',
      },
    },
  ],
};
