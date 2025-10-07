import type { VxeGridProps } from '#/adapter/vxe-table';
import type { VbenFormProps } from '@vben/common-ui';

import { updateCloudFile } from '#/api/fms/cloudFile';
import { getCloudFileTagList } from '#/api/fms/cloudFileTag';
import { getStorageProviderList } from '#/api/fms/storageProvider';
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
      title: $t('fms.cloudFile.name'),
      field: 'name',
    },
    {
      title: $t('component.cropper.preview'),
      field: 'url',
      cellRender: { name: 'CellImage' },
    },
    {
      title: $t('fms.cloudFile.size'),
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
      title: $t('fms.cloudFile.fileType'),
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
      title: $t('common.status'),
      field: 'state',
      slots: {
        default: (e) =>
          h(Switch, {
            checked: e.row.state,
            onClick: () => {
              const newStatus = !e.row.state;
              updateCloudFile({
                id: e.row.id,
                state: newStatus,
                providerId: e.row.providerId,
              }).then(() => {
                e.row.state = newStatus;
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
      label: $t('fms.cloudFile.name'),
      component: 'Input',
    },
    {
      fieldName: 'providerId',
      label: $t('fms.cloudFile.providerId'),
      component: 'Input',
    },
    {
      fieldName: 'tagIds',
      label: $t('fms.tag.tag'),
      component: 'ApiSelect',
      componentProps: {
        api: getCloudFileTagList,
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
    {
      fieldName: 'fileType',
      label: $t('fms.cloudFile.fileType'),
      defaultValue: 0,
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
      label: $t('fms.cloudFile.name'),
      component: 'Input',
    },
    {
      fieldName: 'url',
      label: $t('fms.cloudFile.url'),
      component: 'Input',
    },
    {
      fieldName: 'size',
      label: $t('fms.cloudFile.size'),
      component: 'InputNumber',
      componentProps: {
        class: 'w-full',
      },
    },
    {
      fieldName: 'fileType',
      label: $t('fms.cloudFile.fileType'),
      defaultValue: 0,
      component: 'Select',
      componentProps: {
        options: [
          { label: $t('common.all'), value: 0 },
          { label: $t('fms.file.other'), value: 1 },
          { label: $t('fms.file.image'), value: 2 },
          { label: $t('fms.file.video'), value: 3 },
          { label: $t('fms.file.audio'), value: 4 },
        ],
        class: 'w-full',
      },
    },
    {
      fieldName: 'userId',
      label: $t('fms.cloudFile.userId'),
      component: 'Input',
    },
    {
      fieldName: 'providerId',
      label: $t('fms.cloudFile.providerId'),
      component: 'ApiSelect',
      componentProps: {
        api: getStorageProviderList,
        params: {
          page: 1,
          pageSize: 1000,
        },
        resultField: 'data.data',
        labelField: 'name',
        valueField: 'id',
      },
    },
    {
      fieldName: 'tagIds',
      label: $t('fms.tag.tag'),
      component: 'ApiSelect',
      componentProps: {
        api: getCloudFileTagList,
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
    {
      fieldName: 'state',
      label: $t('fms.cloudFile.state'),
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
