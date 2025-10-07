<script lang="ts" setup>
import type { SelectValue } from 'ant-design-vue/es/select';
import type {
  DefaultOptionType,
  FilterFunc,
} from 'ant-design-vue/lib/vc-select/Select';
import type { PropType } from 'vue';

import { get } from '#/utils/object';
import { $t } from '@vben/locales';
import { useVModel } from '@vueuse/core';
import { Select } from 'ant-design-vue';
import { isFunction, omit } from 'remeda';
import { computed, ref, unref, watch } from 'vue';

type OptionsItem = {
  [name: string]: any;
  disabled?: boolean;
  label?: string;
  value?: string;
};

const props = defineProps({
  value: {
    type: [Array, Object, String, Number] as PropType<SelectValue>,
    default: undefined,
  },
  numberToString: {
    type: Boolean,
  },
  api: {
    type: Function as PropType<(arg?: any) => Promise<any>>,
    default: null,
  },
  // api params
  params: {
    type: Object,
    default: () => {},
  },
  // support xxx.xxx.xx
  resultField: {
    type: String,
    default: 'data.data',
  },
  labelField: {
    type: String,
    default: 'label',
  },
  valueField: {
    type: String,
    default: 'value',
  },
  immediate: {
    type: Boolean,
    default: true,
  },
  alwaysLoad: {
    type: Boolean,
    default: false,
  },
  appendOptions: {
    type: Array<OptionsItem>,
    default: [],
  },
  // search
  showSearch: {
    type: Boolean,
    default: false,
  },
  searchField: {
    type: String,
    default: '',
  },
  optionFilterProp: {
    type: String,
    default: 'label',
  },
  multiple: {
    type: Boolean,
    default: false,
  },
});

const emits = defineEmits(['optionsChange', 'update:value']);
const state = useVModel(props, 'value', emits, {
  defaultValue: props.value,
  passive: true,
});

const optionsData = ref<OptionsItem[]>([]);
const loading = ref(false);
// 首次是否加载过了
const isFirstLoaded = ref(false);
const useSearch = props.showSearch;
const searchFun = ref<any>();
const filterOption = ref<boolean | FilterFunc<DefaultOptionType> | undefined>();
const optionFilterProps = ref<string>();
const mode = props.multiple ? 'multiple' : undefined;
const selectPlaceholder = ref<string>($t('common.chooseText'));

if (useSearch) {
  selectPlaceholder.value = $t('common.inputText');
}

if (useSearch) {
  searchFun.value = searchFetch;
  filterOption.value = false;
} else {
  filterOption.value = true;
  optionFilterProps.value = props.optionFilterProp;
}

const getOptions = computed(() => {
  const { labelField, valueField, numberToString, appendOptions } = props;
  const res: OptionsItem[] = [];

  if (appendOptions.length > 0) {
    appendOptions.forEach((item: any) => {
      res.push(item);
    });
  }

  optionsData.value.forEach((item: any) => {
    const value = item[valueField];
    res.push({
      ...omit(item, [labelField, valueField]),
      label: item[labelField],
      value: numberToString ? `${value}` : value,
      disabled: item.disabled || false,
    });
  });
  return res;
});

watch(
  () => props.params,
  () => {
    if (!useSearch) {
      !unref(isFirstLoaded) && fetch();
    }
  },
  { deep: true, immediate: props.immediate },
);

async function fetch() {
  const api = props.api;
  if (!api || !isFunction(api) || loading.value) return;
  optionsData.value = [];
  try {
    loading.value = true;

    const res = await api(props.params);
    isFirstLoaded.value = true;
    if (Array.isArray(res)) {
      optionsData.value = res;
      emitChange();
      return;
    }
    if (props.resultField) {
      optionsData.value = get(res, props.resultField) || [];
    }
    emitChange();
  } catch (error) {
    console.warn(error);
    // reset status
    isFirstLoaded.value = false;
  } finally {
    loading.value = false;
  }
}

async function searchFetch(value: string) {
  const api = props.api;
  if (!api || !isFunction(api) || loading.value) return;
  optionsData.value = [];
  try {
    loading.value = true;

    const searchParam: any = {};

    if (props.searchField !== undefined) {
      searchParam[props.searchField] = value;
    }

    searchParam.page = 1;
    searchParam.pageSize = 10;

    const res = await api(searchParam);
    if (Array.isArray(res)) {
      optionsData.value = res;
      emitChange();
      return;
    }
    if (props.resultField) {
      optionsData.value = get(res, props.resultField) || [];
    }

    emitChange();
  } catch (error) {
    console.warn(error);
  } finally {
    loading.value = false;
  }
}

async function handleFetch(visible: boolean) {
  if (visible && !useSearch) {
    if (props.alwaysLoad) {
      await fetch();
    } else if (!props.immediate && !unref(isFirstLoaded)) {
      await fetch();
    }
  }
}

function emitChange() {
  emits('optionsChange', unref(getOptions));
}
</script>
<template>
  <Select
    @dropdown-visible-change="handleFetch"
    v-bind="$attrs"
    v-model:value="state"
    :filter-option="filterOption"
    :mode="mode"
    :option-filter-prop="optionFilterProps"
    :options="getOptions"
    :placeholder="selectPlaceholder"
    :show-arrow="false"
    :show-search="true"
    class="w-full"
    @search="searchFun"
  />
</template>
