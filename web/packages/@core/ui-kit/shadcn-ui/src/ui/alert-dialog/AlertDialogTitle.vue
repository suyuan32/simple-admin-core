<script lang="ts" setup>
import type { AlertDialogTitleProps } from 'radix-vue';

import { cn } from '@vben-core/shared/utils';
import { AlertDialogTitle, useForwardProps } from 'radix-vue';
import { computed } from 'vue';

const props = defineProps<{ class?: any } & AlertDialogTitleProps>();

const delegatedProps = computed(() => {
  const { class: _, ...delegated } = props;

  return delegated;
});

const forwardedProps = useForwardProps(delegatedProps);
</script>

<template>
  <AlertDialogTitle
    :class="
      cn('text-lg font-semibold leading-none tracking-tight', props.class)
    "
    v-bind="forwardedProps"
  >
    <slot></slot>
  </AlertDialogTitle>
</template>
