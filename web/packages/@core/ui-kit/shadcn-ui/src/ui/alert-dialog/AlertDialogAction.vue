<script lang="ts" setup>
import type { AlertDialogActionProps } from 'radix-vue';

import { AlertDialogAction } from 'radix-vue';

const props = defineProps<AlertDialogActionProps>();
</script>

<template>
  <AlertDialogAction v-bind="props">
    <slot></slot>
  </AlertDialogAction>
</template>
