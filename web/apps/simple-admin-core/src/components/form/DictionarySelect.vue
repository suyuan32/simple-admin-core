<script lang="ts" setup>
import type { DefaultOptionType } from 'ant-design-vue/lib/select';

import { useDictionaryStore } from '#/store/dictionary';
import { useVModel } from '@vueuse/core';
import { Select } from 'ant-design-vue';
import { onMounted, ref, watch } from 'vue';

const props = defineProps({
  dictionaryName: {
    type: String,
    default: '',
  },
  value: {
    type: [String, Number],
    default: undefined,
  },
  cache: {
    type: Boolean,
    default: true,
  },
});

const emits = defineEmits(['update:value']);

const loading = ref(false);
const options = ref<DefaultOptionType[]>();

const state = useVModel(props, 'value', emits, {
  defaultValue: props.value,
  passive: true,
});

onMounted(() => {
  handleFetch();
});

watch(
  () => state.value,
  (v) => {
    emits('update:value', v);
  },
);

async function handleFetch() {
  loading.value = true;
  const dictStore = useDictionaryStore();
  const dictData = await dictStore.getDictionary(
    props.dictionaryName,
    props.cache,
  );
  if (dictData !== null && dictData !== undefined) {
    options.value = dictData.data.filter((el) => {
      return el.status === 1;
    });
  }
  loading.value = false;
}
</script>
<template>
  <Select
    v-model:value="state"
    :options="options"
    class="w-full"
    v-bind="$attrs"
  />
</template>
