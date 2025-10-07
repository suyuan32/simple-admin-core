<script setup lang="ts">
import type { DialogTitleProps } from 'radix-vue';

import { cn } from '@vben-core/shared/utils';
import { DialogTitle } from 'radix-vue';
import { computed } from 'vue';

const props = defineProps<{ class?: any } & DialogTitleProps>();

const delegatedProps = computed(() => {
  const { class: _, ...delegated } = props;

  return delegated;
});
</script>

<template>
  <DialogTitle
    :class="cn('text-foreground font-medium', props.class)"
    v-bind="delegatedProps"
  >
    <slot></slot>
  </DialogTitle>
</template>
