import type { DataNode } from 'ant-design-vue/es/vc-tree/interface';

import type { Recordable } from '@vben/types';

import { arrayToTree } from 'performant-array-to-tree';
import { forEachObj, map, pick } from 'remeda';

import { ParentIdEnum } from '#/enums/common';

export interface buildNodeOption {
  labelField: string;
  idKeyField: string;
  valueField: string;
  parentKeyField: string;
  defaultValue?: object | string;
  childrenKeyField: string;
}

export function buildDataNode(data: any, options: buildNodeOption): DataNode[] {
  const treeNodeData = map(data, (obj) => {
    const tmpData = pick(obj as any, [
      options.labelField,
      options.idKeyField,
      options.valueField,
      options.parentKeyField,
    ]);

    forEachObj(tmpData, (value, key) => {
      if (key === options.labelField) {
        tmpData.title = value;
        // eslint-disable-next-line @typescript-eslint/no-dynamic-delete
        delete tmpData[key];
      } else if (key === options.valueField) {
        tmpData.key = tmpData[key];
        if (key !== options.idKeyField && key !== options.parentKeyField) {
          // eslint-disable-next-line @typescript-eslint/no-dynamic-delete
          delete tmpData[key];
        }
      }
    });

    if (
      tmpData[options.parentKeyField] === 0 ||
      tmpData[options.parentKeyField] === ParentIdEnum.DEFAULT
    ) {
      tmpData[options.parentKeyField] = null;
    }

    return tmpData;
  });

  const treeConv = arrayToTree(treeNodeData, {
    id: options.idKeyField,
    parentId: options.parentKeyField,
    childrenField: options.childrenKeyField,
    dataField: null,
  });

  // add default label
  if (options.defaultValue) {
    treeConv.push(options.defaultValue as DataNode);
  }
  return treeConv as DataNode[];
}

// buildTreeNode returns treeData for tree select from data
export function buildTreeNode(
  data: any,
  options: buildNodeOption,
): Recordable<any>[] {
  const treeNodeData = map(data, (obj) => {
    const tmpData = pick(obj as any, [
      options.labelField,
      options.idKeyField,
      options.valueField,
      options.parentKeyField,
    ]);

    forEachObj(tmpData, (value, key) => {
      if (key === options.labelField) {
        tmpData.label = value;
        // eslint-disable-next-line @typescript-eslint/no-dynamic-delete
        delete tmpData[key];
      } else if (key === options.valueField) {
        tmpData.value = tmpData[key];
        if (key !== options.idKeyField && key !== options.parentKeyField) {
          // eslint-disable-next-line @typescript-eslint/no-dynamic-delete
          delete tmpData[key];
        }
      }
    });

    if (
      tmpData[options.parentKeyField] === 0 ||
      tmpData[options.parentKeyField] === ParentIdEnum.DEFAULT
    ) {
      tmpData[options.parentKeyField] = null;
    }

    return tmpData;
  });

  const treeConv = arrayToTree(treeNodeData, {
    id: options.idKeyField,
    parentId: options.parentKeyField,
    childrenField: options.childrenKeyField,
    dataField: null,
  });

  // add default label
  if (options.defaultValue) {
    treeConv.push(options.defaultValue as DataNode);
  }
  return treeConv as Recordable<any>[];
}
