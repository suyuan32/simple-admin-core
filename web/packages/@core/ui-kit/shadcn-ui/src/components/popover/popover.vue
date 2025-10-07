<script setup lang="ts">
import type { ClassType } from '@vben-core/typings';
import type {
  PopoverContentProps,
  PopoverRootEmits,
  PopoverRootProps,
} from 'radix-vue';

import { useForwardPropsEmits } from 'radix-vue';
import { computed } from 'vue';

import {
  PopoverContent,
  Popover as PopoverRoot,
  PopoverTrigger,
} from '../../ui';

interface Props extends PopoverRootProps {
  class?: ClassType;
  contentClass?: ClassType;
  contentProps?: PopoverContentProps;
  triggerClass?: ClassType;
}

const props = withDefaults(defineProps<Props>(), {});

const emits = defineEmits<PopoverRootEmits>();

const delegatedProps = computed(() => {
  const {
    class: _cls,
    contentClass: _,
    contentProps: _cProps,
    triggerClass: _tClass,
    ...delegated
  } = props;

  return delegated;
});

const forwarded = useForwardPropsEmits(delegatedProps, emits);
</script>

<template>
  <PopoverRoot v-bind="forwarded">
    <PopoverTrigger :class="triggerClass">
      <slot name="trigger"></slot>

      <PopoverContent
        :class="contentClass"
        class="side-content z-popup"
        v-bind="contentProps"
      >
        <slot></slot>
      </PopoverContent>
    </PopoverTrigger>
  </PopoverRoot>
</template>
