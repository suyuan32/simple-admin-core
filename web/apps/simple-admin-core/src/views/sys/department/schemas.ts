import type { VxeGridProps } from '#/adapter/vxe-table';
import type { VbenFormProps } from '@vben/common-ui';

import { z } from '#/adapter/form';
import { getDepartmentList, updateDepartment } from '#/api/sys/department';
import { ParentIdEnum } from '#/enums/common';
import { $t } from '@vben/locales';
import { Switch } from 'ant-design-vue';
import { h } from 'vue';

export const tableColumns: VxeGridProps = {
  columns: [
    {
      type: 'checkbox',
      width: 60,
    },
    {
      title: $t('sys.department.name'),
      field: 'trans',
      treeNode: true,
    },
    {
      title: $t('sys.department.leader'),
      field: 'leader',
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
              updateDepartment({
                id: e.row.id,
                status: newStatus,
                parentId: e.row.parentId,
              }).then(() => {
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
      label: $t('sys.department.name'),
      component: 'Input',
      rules: z.string().max(50).optional(),
    },
    {
      fieldName: 'leader',
      label: $t('sys.department.leader'),
      component: 'Input',
      rules: z.string().max(20).optional(),
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
      label: $t('sys.department.name'),
      component: 'Input',
      rules: z.string().min(1).max(200),
    },
    {
      fieldName: 'parentId',
      label: $t('sys.department.parentId'),
      component: 'ApiTreeSelect',
      rules: 'required',
      componentProps: {
        api: getDepartmentList,
        params: {
          page: 1,
          pageSize: 1000,
        },
        resultField: 'data.data',
        labelField: 'trans',
        valueField: 'id',
        defaultValue: {
          id: ParentIdEnum.DEFAULT,
          parentId: -1,
          label: $t('sys.department.firstLevelDepartment'),
          value: ParentIdEnum.DEFAULT,
        },
      },
      defaultValue: ParentIdEnum.DEFAULT,
    },
    {
      fieldName: 'ancestors',
      label: $t('sys.department.ancestors'),
      component: 'Input',
      rules: z.string().max(200),
      componentProps: {
        disabled: true,
      },
      help: $t('sys.department.ancestorsHelpMessage'),
    },
    {
      fieldName: 'leader',
      label: $t('sys.department.leader'),
      component: 'Input',
      rules: z.string().max(20),
    },
    {
      fieldName: 'phone',
      label: $t('sys.department.phone'),
      component: 'Input',
      rules: z.string().max(18).optional(),
    },
    {
      fieldName: 'email',
      label: $t('sys.department.email'),
      component: 'Input',
      rules: z.string().max(70).optional(),
    },
    {
      fieldName: 'sort',
      label: $t('sys.department.sort'),
      component: 'InputNumber',
      rules: z.number().max(10_000),
      componentProps: {
        class: 'w-full',
      },
    },
    {
      fieldName: 'remark',
      label: $t('common.remark'),
      component: 'Input',
      rules: z.string().max(200).optional(),
    },
    {
      fieldName: 'status',
      label: $t('sys.department.status'),
      component: 'RadioButtonGroup',
      defaultValue: 1,
      componentProps: {
        options: [
          { label: $t('common.on'), value: 1 },
          { label: $t('common.off'), value: 0 },
        ],
      },
    },
  ],
};
