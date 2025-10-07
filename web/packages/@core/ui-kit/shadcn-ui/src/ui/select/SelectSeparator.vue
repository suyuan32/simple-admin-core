<script setup lang="ts">
import type { SelectSeparatorProps } from 'radix-vue';

import { cn } from '@vben-core/shared/utils';
import { SelectSeparator } from 'radix-vue';
import { computed } from 'vue';

const props = defineProps<{ class?: any } & SelectSeparatorProps>();

const delegatedProps = computed(() => {
  const { class: _, ...delegated } = props;

  return delegated;
});
</script>

<template>
  <SelectSeparator
    v-bind="delegatedProps"
    :class="cn('bg-muted -mx-1 my-1 h-px', props.class)"
  />
</template>
