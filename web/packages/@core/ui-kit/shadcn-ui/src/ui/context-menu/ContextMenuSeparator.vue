<script setup lang="ts">
import type { ContextMenuSeparatorProps } from 'radix-vue';

import { cn } from '@vben-core/shared/utils';
import { ContextMenuSeparator } from 'radix-vue';
import { computed } from 'vue';

const props = defineProps<{ class?: any } & ContextMenuSeparatorProps>();

const delegatedProps = computed(() => {
  const { class: _, ...delegated } = props;

  return delegated;
});
</script>

<template>
  <ContextMenuSeparator
    v-bind="delegatedProps"
    :class="cn('bg-border -mx-1 my-1 h-px', props.class)"
  />
</template>
