<script lang="ts" setup>
import type { SplitterGroupEmits, SplitterGroupProps } from 'radix-vue';
import type { HTMLAttributes } from 'vue';

import { cn } from '@vben-core/shared/utils';
import { SplitterGroup, useForwardPropsEmits } from 'radix-vue';
import { computed } from 'vue';

const props = defineProps<
  { class?: HTMLAttributes['class'] } & SplitterGroupProps
>();
const emits = defineEmits<SplitterGroupEmits>();

const delegatedProps = computed(() => {
  const { class: _, ...delegated } = props;
  return delegated;
});

const forwarded = useForwardPropsEmits(delegatedProps, emits);
</script>

<template>
  <SplitterGroup
    v-bind="forwarded"
    :class="
      cn(
        'flex h-full w-full data-[panel-group-direction=vertical]:flex-col',
        props.class,
      )
    "
  >
    <slot></slot>
  </SplitterGroup>
</template>
