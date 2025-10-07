<script lang="ts" setup>
import type { TransferItem } from 'ant-design-vue/lib/transfer';

import { get } from '#/utils/object';
import { Transfer } from 'ant-design-vue';
import { isFunction, omit } from 'remeda';
import { computed, onMounted, ref, unref, watch } from 'vue';

const props = defineProps({
  value: { type: Array<string>, default: undefined },
  api: {
    type: Function,
    default: null,
  },
  params: { type: Object, default: undefined },
  dataSource: { type: Array, default: undefined },
  immediate: {
    type: Boolean,
    default: false,
  },
  afterFetch: { type: Function, default: undefined },
  resultField: {
    type: String,
    default: 'data.data',
  },
  labelField: {
    type: String,
    default: 'title',
  },
  valueField: {
    type: String,
    default: 'key',
  },
  showSearch: { type: Boolean, default: false },
  disabled: { type: Boolean, default: false },
  filterOption: {
    type: Function,
    default: undefined,
  },
  selectedKeys: { type: Array<string>, default: [] },
  showSelectAll: { type: Boolean, default: false },
  targetKeys: { type: Array<string>, default: [] },
});

const emit = defineEmits(['optionsChange', 'change', 'update:value']);

const _dataSource = ref<TransferItem[]>([]);

const getdataSource = computed(() => {
  const { labelField, valueField } = props;

  // eslint-disable-next-line unicorn/no-array-reduce
  return unref(_dataSource).reduce((prev, next) => {
    if (next) {
      prev.push({
        ...omit(next, [labelField, valueField]),
        title: next[labelField],
        key: next[valueField],
      });
    }
    return prev;
  }, [] as TransferItem[]);
});

const getTargetKeys = computed<string[]>(() => {
  if (Array.isArray(props.value)) {
    return props.value;
  }
  if (Array.isArray(props.targetKeys)) {
    return props.targetKeys;
  }
  return [];
});

onMounted(() => {
  props.immediate && fetch();
});

watch(
  () => props.params,
  () => {
    fetch();
  },
  { deep: true },
);

async function fetch() {
  const api = props.api;
  if (!api || !isFunction(api)) {
    if (Array.isArray(props.dataSource)) {
      _dataSource.value = props.dataSource as any;
    }
    return;
  }
  _dataSource.value = [];
  try {
    const res = await api(props.params);
    if (Array.isArray(res)) {
      _dataSource.value = res;
      emitChange();
      return;
    }
    if (props.resultField) {
      _dataSource.value = get(res, props.resultField) || [];
    }
    emitChange();
  } catch (error) {
    console.warn(error);
  }
}

function emitChange() {
  emit('optionsChange', unref(getdataSource));
}
</script>

<template>
  <Transfer
    v-bind="$attrs"
    :data-source="getdataSource"
    :disabled="disabled"
    :filter-option="filterOption as any"
    :render="(item: any) => item.title"
    :selected-keys="selectedKeys"
    :show-search="showSearch"
    :show-select-all="showSelectAll"
    :target-keys="getTargetKeys"
  />
</template>
