<script setup lang="ts">
import type { DialogTitleProps } from 'radix-vue';

import { cn } from '@vben-core/shared/utils';
import { DialogTitle, useForwardProps } from 'radix-vue';
import { computed } from 'vue';

const props = defineProps<{ class?: any } & DialogTitleProps>();

const delegatedProps = computed(() => {
  const { class: _, ...delegated } = props;

  return delegated;
});

const forwardedProps = useForwardProps(delegatedProps);
</script>

<template>
  <DialogTitle
    v-bind="forwardedProps"
    :class="
      cn('text-lg font-semibold leading-none tracking-tight', props.class)
    "
  >
    <slot></slot>
  </DialogTitle>
</template>
