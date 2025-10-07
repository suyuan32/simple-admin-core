<script lang="ts">
import { oauthLoginCallback } from '#/api/sys/oauthProvider';
import { getPermCode } from '#/api/sys/user';
import { useAuthStore } from '#/store';
import { $t } from '@vben/locales';
import { preferences } from '@vben/preferences';
import { useAccessStore } from '@vben/stores';
import { message, notification } from 'ant-design-vue';
import { defineComponent, ref } from 'vue';
import { useRouter } from 'vue-router';

export default defineComponent({
  name: 'OauthCallbackPage',
  components: {},
  setup() {
    const router = useRouter();
    const query = ref<string>('');
    query.value += `?state=${router.currentRoute.value.query.state}`;
    query.value += `&code=${router.currentRoute.value.query.code}`;

    async function login(url: string) {
      try {
        const result = await oauthLoginCallback(url);
        const { token } = result;

        const accessStore = useAccessStore();
        const authStore = useAuthStore();
        // save token
        accessStore.setAccessToken(token);
        await authStore.fetchUserInfo();
        // 获取用户信息并存储到 accessStore 中
        const [fetchUserInfoResult, accessCodes] = await Promise.all([
          authStore.fetchUserInfo(),
          getPermCode(),
        ]);

        const userInfo = fetchUserInfoResult;

        accessStore.setAccessCodes(accessCodes.data);

        await router.push(
          userInfo?.homePath || preferences.app.defaultHomePath,
        );

        if (userInfo?.nickname) {
          notification.success({
            description: `${$t('authentication.loginSuccessDesc')}:${userInfo?.nickname}`,
            duration: 3,
            message: $t('authentication.loginSuccess'),
          });
        }
      } catch {
        message.error($t('sys.oauth.createAccount'), 5);
        router.replace('/auth/login');
      }
    }

    login(query.value);
    return {};
  },
});
</script>
<template>
  <div></div>
</template>
