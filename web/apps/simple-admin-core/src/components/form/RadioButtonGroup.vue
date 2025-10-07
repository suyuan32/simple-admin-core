<script lang="ts" setup>
import type { PropType } from 'vue';

import { useVModel } from '@vueuse/core';
import { Radio } from 'ant-design-vue';
import { isString } from 'remeda';
import { computed } from 'vue';

type OptionsItem = {
  disabled?: boolean;
  label: string;
  value: boolean | number | string;
};
type RadioItem = OptionsItem | string;

const props = defineProps({
  value: {
    type: [String, Number, Boolean] as PropType<boolean | number | string>,
    default: undefined,
  },
  options: {
    type: Array as PropType<RadioItem[]>,
    default: () => [],
  },
});

const emits = defineEmits(['update:value', 'optionsChange']);
const state = useVModel(props, 'value', emits, {
  defaultValue: props.value,
  passive: true,
});
const RadioGroup = Radio.Group;
const RadioButton = Radio.Button;

// Processing options value
const getOptions = computed((): OptionsItem[] => {
  const { options } = props;
  if (!options || options?.length === 0) return [];

  const isStringArr = options.some((item) => isString(item));
  if (!isStringArr) return options as OptionsItem[];

  return options.map((item) => ({ label: item, value: item })) as OptionsItem[];
});
</script>
<template>
  <RadioGroup v-model:value="state" button-style="solid">
    <template v-for="item in getOptions" :key="`${item.value}`">
      <RadioButton :disabled="item.disabled" :value="item.value">
        {{ item.label }}
      </RadioButton>
    </template>
  </RadioGroup>
</template>
