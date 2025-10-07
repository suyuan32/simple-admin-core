<script setup lang="ts">
import { useScrollLock } from '@vben-core/composables';
import { inject } from 'vue';

useScrollLock();
const id = inject('DISMISSABLE_DRAWER_ID');
</script>
<template>
  <div :data-dismissable-drawer="id" class="bg-overlay z-popup inset-0"></div>
</template>
