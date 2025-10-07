import type { VxeGridProps } from '#/adapter/vxe-table';
import type { DictionaryDetailInfo } from '#/api/sys/model/dictionaryDetailModel';

import { h } from 'vue';

import { type VbenFormProps } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { InputNumber, Switch } from 'ant-design-vue';

import { z } from '#/adapter/form';
import { updateDictionary } from '#/api/sys/dictionary';
import { updateDictionaryDetail } from '#/api/sys/dictionaryDetail';

export const tableColumns: VxeGridProps = {
  columns: [
    {
      title: $t('common.displayName'),
      field: 'trans',
    },
    {
      title: $t('sys.dictionary.name'),
      field: 'name',
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
              updateDictionary({ id: e.row.id, status: newStatus }).then(() => {
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
      label: $t('sys.dictionary.name'),
      component: 'Input',
      rules: z.string().max(50).optional(),
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
      defaultValue: 0,
    },
    {
      fieldName: 'title',
      label: $t('common.displayName'),
      component: 'Input',
      rules: z.string().min(1).max(50),
    },
    {
      fieldName: 'name',
      label: $t('sys.dictionary.name'),
      component: 'Input',
      rules: z.string().min(1).max(50),
    },
    {
      fieldName: 'desc',
      label: $t('common.description'),
      component: 'Input',
      rules: z.string().max(200).optional(),
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
    },
  ],
};

export const detailTableColumns: VxeGridProps<DictionaryDetailInfo> = {
  columns: [
    {
      type: 'checkbox',
      width: 60,
    },
    {
      editRender: { name: 'input' },
      title: $t('common.displayName'),
      field: 'title',
    },
    {
      editRender: { name: 'input' },
      title: $t('sys.dictionary.key'),
      field: 'key',
    },
    {
      editRender: { name: 'input' },
      title: $t('sys.dictionary.value'),
      field: 'value',
    },
    {
      editRender: {},
      title: $t('common.sort'),
      field: 'sort',
      slots: {
        edit: (e) => {
          return h(InputNumber, {
            modelValue: e.row.sort,
            onInput: (n) => {
              e.row.sort = Number(n);
            },
          });
        },
      },
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
              updateDictionaryDetail({ id: e.row.id, status: newStatus }).then(
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
