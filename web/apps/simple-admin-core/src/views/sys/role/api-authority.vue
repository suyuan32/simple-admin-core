<script lang="ts" setup>
import type { DataNode } from 'ant-design-vue/es/vc-tree/interface';

import type { BaseDataResp } from '#/api/model/baseModel';
import type { ApiInfo, ApiListResp } from '#/api/sys/model/apiModel';
import type { ApiAuthorityInfo } from '#/api/sys/model/authorityModel';

import { ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { message, Tree } from 'ant-design-vue';
import { clone, concat, isNumber, unique } from 'remeda';

import {
  createOrUpdateApiAuthority,
  getApiAuthority,
  getApiList,
} from '#/api/sys/authority';

defineOptions({
  name: 'ApiAuthorityModal',
});

const treeApiData = ref<DataNode[]>([]);

const checkedKeys = ref();
const roleId = ref<number>();
let tempApiList: BaseDataResp<ApiListResp> = {
  code: 0,
  msg: '',
  data: { total: 0, data: [] },
};

const [Modal, modalApi] = useVbenModal({
  fullscreenButton: false,
  onCancel() {
    modalApi.close();
  },
  onConfirm: async () => {
    const apiReqData: ApiAuthorityInfo[] = convertApiCheckedKeysToReq(
      checkedKeys.value,
      tempApiList.data.data,
    );
    const result = await createOrUpdateApiAuthority({
      roleId: Number(roleId.value),
      data: apiReqData,
    });
    if (result.code === 0) {
      message.success($t('common.successful'));
    }
    modalApi.close();
  },
  onOpenChange(isOpen: boolean) {
    roleId.value = isOpen ? modalApi.getData()?.roleId || {} : {};
    if (isOpen) {
      getApiData();
    }
  },
  title: $t('sys.authority.apiAuthority'),
});

async function getApiData() {
  try {
    treeApiData.value = [];
    const apiData = await getApiList({
      page: 1,
      pageSize: 10_000,
    });
    tempApiList = apiData;
    const dataConv = convertApiTreeData(apiData.data.data);
    for (const key in dataConv) {
      treeApiData.value.push(dataConv[key] as any);
    }
    const checkedData = await getApiAuthority({ id: roleId.value });
    checkedKeys.value =
      checkedData.data.data === null
        ? convertApiToCheckedKeys([], apiData.data.data)
        : convertApiToCheckedKeys(checkedData.data.data, apiData.data.data);
  } catch {
    // console.log(error);
  }
}

/**
 *  author: Ryan Su
 *  @description: this function is used to convert menu data into tree node data
 */

function convertApiTreeData(params: ApiInfo[]): DataNode[] {
  const finalData: DataNode[] = [];
  const apiData: DataNode[] = [];
  if (params.length === 0) {
    return apiData;
  }

  const apiMap = new Map<string, string>();
  const serviceMap = new Map<string, boolean>();
  for (const param of params) {
    apiMap.set(param.group, param.serviceName);
    serviceMap.set(param.serviceName, true);
  }

  for (const k of apiMap.keys()) {
    const apiTmp: DataNode = {
      title: k,
      key: k,
      children: [],
    };

    for (const param of params) {
      if (param.group === k) {
        apiTmp.children?.push({
          title: param.trans,
          key: param.id as number,
          disableCheckbox: param.isRequired,
        });
      }
    }

    apiData.push(apiTmp);
  }

  for (const k1 of serviceMap.keys()) {
    const svcTmp: DataNode = {
      title: k1,
      key: k1,
      children: [],
    };

    for (const apiDatum of apiData) {
      if (apiMap.get(apiDatum.title) === k1) {
        svcTmp.children?.push(clone(apiDatum));
      }
    }

    finalData.push(svcTmp);
  }

  return finalData;
}

/**
 *  author: Ryan Su
 *  @description: convert checked data into authorized data
 */
function convertApiCheckedKeysToReq(
  checked: number[],
  data: ApiInfo[],
): ApiAuthorityInfo[] {
  // delete string keys
  const pureDigit: number[] = [];
  for (const element of checked) {
    if (isNumber(element)) {
      pureDigit.push(element);
    }
  }

  // sort data
  data.sort((a, b) => {
    if (a.id !== undefined && b.id !== undefined) return a.id - b.id;
    return 1;
  });
  pureDigit.sort((a, b) => {
    return a - b;
  });
  // convert data
  const target: ApiAuthorityInfo[] = [];
  let j = 0;
  for (const datum of data) {
    if (datum.id === pureDigit[j]) {
      target.push({
        path: datum.path,
        method: datum.method,
      });
      j++;
    }
  }
  return target;
}

/**
 *  author: Ryan Su
 *  @description: this function is used to convert authorization api response into checked keys
 */

function convertApiToCheckedKeys(
  checked: ApiAuthorityInfo[],
  data: ApiInfo[],
): number[] {
  const dataMap = new Map();
  const authorityApis: number[] = [];
  const requiredAPIs: number[] = [];
  data.forEach((value, _key) => {
    if (value.isRequired) {
      requiredAPIs.push(value.id as number);
    }
  });

  for (const datum of data) {
    dataMap.set(datum.path + datum.method, datum.id);
  }
  for (const element of checked) {
    authorityApis.push(dataMap.get(element.path + element.method));
  }

  return unique(concat(authorityApis, requiredAPIs));
}

defineExpose(modalApi);
</script>
<template>
  <Modal>
    <Tree
      v-model:checked-keys="checkedKeys"
      :tree-data="treeApiData"
      checkable
      default-expand-all
    />
  </Modal>
</template>
