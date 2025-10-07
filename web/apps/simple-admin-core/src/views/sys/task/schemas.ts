import type { VxeGridProps } from '#/adapter/vxe-table';

import { h } from 'vue';

import { type VbenFormProps } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { Switch, Tag } from 'ant-design-vue';

import { z } from '#/adapter/form';
import { updateTask } from '#/api/sys/task';

export const tableColumns: VxeGridProps = {
  columns: [
    {
      type: 'checkbox',
      width: 60,
    },

    {
      title: $t('sys.task.name'),
      field: 'name',
    },
    {
      title: $t('sys.task.taskGroup'),
      field: 'taskGroup',
    },
    {
      title: $t('sys.task.cronExpression'),
      field: 'cronExpression',
    },
    {
      title: $t('sys.task.pattern'),
      field: 'pattern',
    },
    {
      title: $t('sys.task.payload'),
      field: 'payload',
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
              updateTask({ id: e.row.id, status: newStatus }).then(() => {
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
      fieldName: 'name',
      label: $t('sys.task.name'),
      component: 'Input',
    },
    {
      fieldName: 'taskGroup',
      label: $t('sys.task.taskGroup'),
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
      label: $t('sys.task.name'),
      component: 'Input',
      rules: z.string().max(50),
    },
    {
      fieldName: 'taskGroup',
      label: $t('sys.task.taskGroup'),
      component: 'Input',
      rules: z.string().max(40),
    },
    {
      fieldName: 'cronExpression',
      label: $t('sys.task.cronExpression'),
      component: 'Input',
      rules: z.string().max(80),
    },
    {
      fieldName: 'pattern',
      label: $t('sys.task.pattern'),
      component: 'Input',
      rules: z.string().max(100),
    },
    {
      fieldName: 'payload',
      label: $t('sys.task.payload'),
      component: 'Input',
    },
    {
      fieldName: 'status',
      label: $t('sys.task.status'),
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

export const taskLogTableColumns: VxeGridProps = {
  columns: [
    {
      type: 'checkbox',
      width: 60,
    },
    {
      title: $t('sys.taskLog.startedAt'),
      field: 'startedAt',
      formatter: 'formatDateTime',
    },
    {
      title: $t('sys.taskLog.finishedAt'),
      field: 'finishedAt',
      formatter: 'formatDateTime',
    },
    {
      title: $t('sys.taskLog.result'),
      field: 'result',
      width: 70,
      slots: {
        default: (record) => {
          let resultText = '';
          resultText =
            record.row.result === 1
              ? $t('common.successful')
              : $t('common.failed');
          return h(
            Tag,
            {
              color: record.row.result === 1 ? 'green' : 'red',
            },
            () => resultText,
          );
        },
      },
    },
  ],
};

export const taskLogSearchFormSchemas: VbenFormProps = {
  schema: [
    {
      fieldName: 'result',
      label: $t('sys.taskLog.result'),
      component: 'Select',
      defaultValue: 0,
      componentProps: {
        options: [
          { label: $t('common.all'), value: 0 },
          { label: $t('common.successful'), value: 1 },
          { label: $t('common.failed'), value: 2 },
        ],
      },
    },
  ],
};
