<script lang="ts" setup>
import { useVModel } from '@vueuse/core';
import { RangePicker } from 'ant-design-vue';
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
    type: Array<number>,
    default: undefined,
  },
});

const emits = defineEmits(['change', 'update:value']);

const dateVal = ref<[dayjs.Dayjs, dayjs.Dayjs]>([dayjs(), dayjs()]);
const showTimePicker = ref<boolean>();

showTimePicker.value = props.timeMode === 'datetime';

const state = useVModel(props, 'value', emits, {
  defaultValue: props.value,
  passive: true,
});

watch(
  () => state.value,
  (v) => {
    if (v !== undefined && dateVal.value.length >= 2) {
      dateVal.value =
        props.valueFormat === 'unixmilli'
          ? [dayjs(v[0]), dayjs(v[1])]
          : [dayjs.unix(v[0] as number), dayjs.unix(v[1] as number)];
    }
    emits('update:value', v);
    emits('change', v);
  },
);

function handleChange(v: any) {
  if (v === null) {
    state.value = [0, 0];
  } else {
    const dateTime: any = [dayjs(), dayjs()];
    if (dateVal.value !== undefined) {
      dateTime[0] = dateVal.value[0]?.clone();
      dateTime[1] = dateVal.value[1]?.clone();
    }

    if (props.timeMode !== 'datetime') {
      dateTime[0] = dateTime[0].hour(0).minute(0).second(0).millisecond(0);
      dateTime[1] = dateTime[1].hour(0).minute(0).second(0).millisecond(0);
    }

    state.value =
      props.valueFormat === 'unixmilli'
        ? [dateTime[0].valueOf(), dateTime[1].valueOf()]
        : [dateTime[0].unix(), dateTime[1].unix()];
  }
}
</script>
<template>
  <RangePicker
    v-bind="$attrs"
    v-model:value="dateVal"
    :show-time="showTimePicker"
    allow-clear
    @change="handleChange"
  />
</template>
