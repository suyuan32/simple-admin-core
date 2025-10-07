<script setup lang="ts">
import type { ThirdPartyLoginIcon } from './types';

import { Icon } from '@iconify/vue';
import { $t } from '@vben/locales';
import { VbenIconButton } from '@vben-core/shadcn-ui';

defineOptions({
  name: 'ThirdPartyLogin',
});

const props = defineProps({
  iconList: {
    type: Array<ThirdPartyLoginIcon>,
    default: [
      {
        icon: 'icon-park-outline:github',
        oauthProvider: 'github',
      },
    ],
  },
});

const emits = defineEmits(['oauthLogin']);

function handleOauthLogin(provider: string) {
  emits('oauthLogin', provider);
}
</script>

<template>
  <div class="w-full sm:mx-auto md:max-w-md">
    <div class="mt-4 flex items-center justify-between">
      <span class="border-input w-[35%] border-b dark:border-gray-600"></span>
      <span class="text-muted-foreground text-center text-xs uppercase">
        {{ $t('authentication.thirdPartyLogin') }}
      </span>
      <span class="border-input w-[35%] border-b dark:border-gray-600"></span>
    </div>

    <div class="mt-4 flex flex-wrap justify-center">
      <VbenIconButton
        v-for="item in props.iconList"
        :key="item.icon"
        class="mb-3"
        @click="handleOauthLogin(item.oauthProvider)"
      >
        <Icon :icon="item.icon" />
      </VbenIconButton>
    </div>
  </div>
</template>
