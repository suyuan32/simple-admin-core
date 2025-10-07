<script lang="ts" setup>
import type { AlertDialogCancelProps } from 'radix-vue';

import { AlertDialogCancel } from 'radix-vue';

const props = defineProps<AlertDialogCancelProps>();
</script>

<template>
  <AlertDialogCancel v-bind="props">
    <slot></slot>
  </AlertDialogCancel>
</template>
