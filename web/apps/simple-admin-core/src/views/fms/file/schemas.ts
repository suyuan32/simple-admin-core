import type { VxeGridProps } from '#/adapter/vxe-table';

import { h } from 'vue';

import { type VbenFormProps } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { Switch } from 'ant-design-vue';

import { z } from '#/adapter/form';
import { setFileStatus } from '#/api/fms/file';
import { getTagList } from '#/api/fms/fileTag';

export const tableColumns: VxeGridProps = {
  columns: [
    {
      type: 'checkbox',
      width: 60,
    },
    {
      title: $t('fms.file.fileName'),
      field: 'name',
    },
    {
      title: $t('component.cropper.preview'),
      field: 'publicPath',
      cellRender: { name: 'CellImage' },
    },
    {
      title: $t('fms.file.fileType'),
      field: 'fileType',
      slots: {
        default: (e) => {
          switch (e.row.fileType) {
            case 2: {
              return $t('fms.file.image');
            }
            case 3: {
              return $t('fms.file.video');
            }
            case 4: {
              return $t('fms.file.audio');
            }
            default: {
              return $t('fms.file.other');
            }
          }
        },
      },
    },
    {
      title: $t('fms.file.fileSize'),
      field: 'size',
      slots: {
        default: (e) => {
          if (e.row.size > 1_073_741_824) {
            return `${(e.row.size / 1_073_741_824).toFixed(2)}GB`;
          } else if (e.row.size > 1_048_576) {
            return `${(e.row.size / 1_048_576).toFixed(2)}MB`;
          } else {
            return `${(e.row.size / 1024).toFixed(2)}KB`;
          }
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
              setFileStatus(e.row.id, newStatus).then(() => {
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
      fieldName: 'fileType',
      label: $t('fms.file.fileType'),
      rules: z.number().max(10).optional(),
      component: 'Select',
      componentProps: {
        options: [
          { label: $t('common.all'), value: 0 },
          { label: $t('fms.file.other'), value: 1 },
          { label: $t('fms.file.image'), value: 2 },
          { label: $t('fms.file.video'), value: 3 },
          { label: $t('fms.file.audio'), value: 4 },
        ],
      },
    },
    {
      fieldName: 'fileName',
      label: $t('fms.file.fileName'),
      component: 'Input',
      rules: z.string().max(50).optional(),
    },
    {
      fieldName: 'period',
      label: $t('common.createTime'),
      component: 'Input', // todo: range picker
    },
    {
      fieldName: 'fileTagIds',
      label: $t('fms.tag.tag'),
      component: 'ApiSelect',
      componentProps: {
        api: getTagList,
        params: {
          page: 1,
          pageSize: 1000,
          name: '',
        },
        resultField: 'data.data',
        labelField: 'name',
        valueField: 'id',
        multiple: true,
      },
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
      label: $t('fms.file.fileName'),
      component: 'Input',
    },
    {
      fieldName: 'tagIds',
      label: $t('fms.tag.tag'),
      component: 'ApiSelect',
      componentProps: {
        api: getTagList,
        params: {
          page: 1,
          pageSize: 1000,
          name: '',
        },
        resultField: 'data.data',
        labelField: 'name',
        valueField: 'id',
        multiple: true,
      },
    },
  ],
};
