<script setup lang="ts">
import type { DropdownMenuLabelProps } from 'radix-vue';

import { cn } from '@vben-core/shared/utils';
import { DropdownMenuLabel, useForwardProps } from 'radix-vue';
import { computed } from 'vue';

const props = defineProps<
  { class?: any; inset?: boolean } & DropdownMenuLabelProps
>();

const delegatedProps = computed(() => {
  const { class: _, ...delegated } = props;

  return delegated;
});

const forwardedProps = useForwardProps(delegatedProps);
</script>

<template>
  <DropdownMenuLabel
    v-bind="forwardedProps"
    :class="
      cn('px-2 py-1.5 text-sm font-semibold', inset && 'pl-8', props.class)
    "
  >
    <slot></slot>
  </DropdownMenuLabel>
</template>
