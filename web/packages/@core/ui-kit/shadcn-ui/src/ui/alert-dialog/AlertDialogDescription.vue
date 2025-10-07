<script lang="ts" setup>
import type { AlertDialogDescriptionProps } from 'radix-vue';

import { cn } from '@vben-core/shared/utils';
import { AlertDialogDescription, useForwardProps } from 'radix-vue';
import { computed } from 'vue';

const props = defineProps<{ class?: any } & AlertDialogDescriptionProps>();

const delegatedProps = computed(() => {
  const { class: _, ...delegated } = props;

  return delegated;
});

const forwardedProps = useForwardProps(delegatedProps);
</script>

<template>
  <AlertDialogDescription
    :class="cn('text-muted-foreground text-sm', props.class)"
    v-bind="forwardedProps"
  >
    <slot></slot>
  </AlertDialogDescription>
</template>
