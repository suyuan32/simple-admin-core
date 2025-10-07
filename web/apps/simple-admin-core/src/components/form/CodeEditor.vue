<script lang="ts" setup>
import { json } from '@codemirror/lang-json';
import { StreamLanguage } from '@codemirror/language';
import { yaml } from '@codemirror/legacy-modes/mode/yaml';
import { githubDark, githubLight } from '@uiw/codemirror-theme-github';
import { usePreferences } from '@vben/preferences';
import { useVModel } from '@vueuse/core';
import { onBeforeUpdate, ref } from 'vue';
import { Codemirror } from 'vue-codemirror';

const props = defineProps({
  value: { type: String, default: undefined },
  mode: {
    type: String,
    default: 'json',
  },
  autoFormat: { type: Boolean, default: true },
});

const emits = defineEmits(['update:value']);
const state = useVModel(props, 'value', emits, {
  defaultValue: props.value,
  passive: true,
});

const inputValue = ref<string>('');

onBeforeUpdate(() => {
  inputValue.value = props.value as string;
});

const preference = usePreferences();

const darkMode = preference.isDark.value;
const extensions: any = [];

if (darkMode) {
  extensions.push(githubDark);
} else {
  extensions.push(githubLight);
}

switch (props.mode) {
  case 'json': {
    extensions.push(json());
    break;
  }
  case 'yaml': {
    extensions.push(StreamLanguage.define(yaml));
    break;
  }
  default: {
    extensions.push(json(), StreamLanguage.define(yaml));
  }
}

function handleValueChange(v: string) {
  // if (props.mode == MODE.JSON) {
  //   v = v.replace(/(\r\n|\n|\r)/gm, '');
  // }
  emits('update:value', v);
}
</script>
<template>
  <Codemirror
    v-bind="$attrs"
    v-model="state"
    :autofocus="true"
    :extensions="extensions"
    :indent-with-tab="true"
    :style="{ minWidth: '500px' }"
    :tab-size="2"
    @change="handleValueChange"
  />
</template>
