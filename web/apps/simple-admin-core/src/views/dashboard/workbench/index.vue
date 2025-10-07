<script lang="ts" setup>
import type {
  WorkbenchProjectItem,
  WorkbenchQuickNavItem,
} from '@vben/common-ui';

import { useRouter } from 'vue-router';

import { WorkbenchQuickNav } from '@vben/common-ui';
import { $t } from '@vben/locales';
// import { useUserStore } from '@vben/stores';
import { openWindow } from '@vben/utils';

import { Card, TypographyTitle } from 'ant-design-vue';

// const userStore = useUserStore();

// 同样，这里的 url 也可以使用以 http 开头的外部链接
const quickNavItems: WorkbenchQuickNavItem[] = [
  {
    title: $t('sys.route.userManagementTitle'),
    icon: 'ant-design:user-outlined',
    color: '#bf0c2c',
    url: '/user',
  },
  {
    title: $t('sys.route.roleManagementTitle'),
    icon: 'eos-icons:role-binding-outlined',
    color: '#e18525',
    url: '/role',
  },
  {
    title: $t('sys.route.menuManagementTitle'),
    icon: 'ep:menu',
    color: '#3fb27f',
    url: '/menu',
  },
  {
    title: $t('sys.route.apiManagementTitle'),
    icon: 'ant-design:api-outlined',
    color: '#4daf1bc9',
    url: '/api',
  },
  {
    title: $t('sys.route.dictionaryManagementTitle'),
    icon: 'ant-design:book-outlined',
    color: '#cc00ff',
    url: '/dictionary',
  },
  {
    title: $t('sys.route.oauthManagement'),
    icon: 'ant-design:unlock-filled',
    color: '#0099ff',
    url: '/oauth',
  },
];

const router = useRouter();

// 这是一个示例方法，实际项目中需要根据实际情况进行调整
// This is a sample method, adjust according to the actual project requirements
function navTo(nav: WorkbenchProjectItem | WorkbenchQuickNavItem) {
  if (nav.url?.startsWith('http')) {
    openWindow(nav.url);
    return;
  }
  if (nav.url?.startsWith('/')) {
    router.push(nav.url).catch((error) => {
      console.error('Navigation failed:', error);
    });
  } else {
    console.warn(`Unknown URL for navigation item: ${nav.title} -> ${nav.url}`);
  }
}
</script>

<template>
  <div class="p-5">
    <!--    <WorkbenchHeader-->
    <!--      :avatar="userStore.userInfo?.avatar || preferences.app.defaultAvatar"-->
    <!--    >-->
    <!--      <template #title>-->
    <!--        早安, {{ userStore.userInfo?.realName }}, 开始您一天的工作吧！-->
    <!--      </template>-->
    <!--      &lt;!&ndash;      <template #description> 今日晴，20℃ - 32℃！ </template>&ndash;&gt;-->
    <!--    </WorkbenchHeader>-->

    <div class="mt-5 flex flex-col lg:flex-row">
      <div class="mr-4 w-full lg:w-3/5">
        <Card :title="$t('sys.sys.version')">
          <TypographyTitle :level="5"> Simple Admin v1.7.2 </TypographyTitle>
        </Card>
      </div>
      <div class="w-full lg:w-2/5">
        <WorkbenchQuickNav
          :items="quickNavItems"
          class="mt-5 lg:mt-0"
          title="快捷导航"
          @click="navTo"
        />
      </div>
    </div>
  </div>
</template>
