<script setup lang="ts">
import type { SelectLabelProps } from 'radix-vue';

import { cn } from '@vben-core/shared/utils';
import { SelectLabel } from 'radix-vue';

const props = defineProps<{ class?: any } & SelectLabelProps>();
</script>

<template>
  <SelectLabel :class="cn('px-2 py-1.5 text-sm font-semibold', props.class)">
    <slot></slot>
  </SelectLabel>
</template>
