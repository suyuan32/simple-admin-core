<script lang="ts" setup>
import { useVModel } from '@vueuse/core';
import { DatePicker, FormItemRest, TimePicker } from 'ant-design-vue';
import dayjs from 'dayjs';
import { ref, watch } from 'vue';

const props = defineProps({
  timeMode: {
    type: String,
    default: 'datetime',
  },
  valueFormat: {
    type: String,
    default: 'unixmilli',
  },
  value: {
    type: Number,
    default: undefined,
  },
});

const emits = defineEmits(['change', 'update:value']);

const dateVal = ref<dayjs.Dayjs>();
const timeVal = ref<dayjs.Dayjs>();
const showTimePicker = ref<boolean>();

showTimePicker.value = props.timeMode === 'datetime';

const state = useVModel(props, 'value', emits, {
  defaultValue: props.value,
  passive: true,
});

watch(
  () => state.value,
  (v) => {
    if (v !== undefined) {
      if (props.valueFormat === 'unixmilli') {
        dateVal.value = dayjs(v);
        timeVal.value = dayjs(v);
      } else {
        dateVal.value = dayjs.unix(v);
        timeVal.value = dayjs.unix(v);
      }
    }
    emits('update:value', v);
    emits('change', v);
  },
);

function handleChange(v: any) {
  if (v === null) {
    state.value = undefined;
  } else {
    let dateTime = dayjs();
    if (dateVal.value !== undefined) {
      dateTime = dateVal.value?.clone();
    }
    if (props.timeMode === 'datetime') {
      if (timeVal.value !== undefined) {
        dateTime = dateTime
          .hour(timeVal.value.hour())
          .minute(timeVal.value.minute())
          .second(timeVal.value.second())
          .millisecond(0);
      }
    } else {
      dateTime = dateTime.hour(0).minute(0).second(0).millisecond(0);
    }

    state.value =
      props.valueFormat === 'unixmilli' ? dateTime.valueOf() : dateTime.unix();
  }
}
</script>
<template>
  <DatePicker
    v-model:value="dateVal"
    allow-clear
    v-bind="$attrs"
    @change="handleChange"
  />

  <FormItemRest>
    <TimePicker
      v-bind="$attrs"
      v-if="showTimePicker"
      v-model:value="timeVal"
      allow-clear
      class="ml-4"
      @change="handleChange"
    />
  </FormItemRest>
</template>
