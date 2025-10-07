import type { BaseDataResp } from '#/api/model/baseModel';
import type {
  GetUserInfoModel,
  LoginByEmailReq,
  LoginBySmsReq,
  LoginReq,
  LoginResp,
} from '#/api/sys/model/userModel';

import {
  doLogout,
  getPermCode,
  getUserInfo,
  login,
  loginByEmail,
  loginBySms,
} from '#/api/sys/user';
import { $t } from '#/locales';
import { LOGIN_PATH } from '@vben/constants';
import { preferences } from '@vben/preferences';
import { resetAllStores, useAccessStore, useUserStore } from '@vben/stores';
import { notification } from 'ant-design-vue';
import { defineStore } from 'pinia';
import { isArray } from 'remeda';
import { ref } from 'vue';
import { useRouter } from 'vue-router';

export const useAuthStore = defineStore('auth', () => {
  const accessStore = useAccessStore();
  const userStore = useUserStore();
  const router = useRouter();
  const isLoggingOut = ref(false);
  const loginLoading = ref(false);

  /**
   * 异步处理登录操作
   * Asynchronously handle the login process
   * @param params 登录表单数据
   * @param loginType
   * @param onSuccess
   */
  async function authLogin(
    params: LoginByEmailReq | LoginBySmsReq | LoginReq,
    loginType: 'captcha' | 'email' | 'mobile',
    onSuccess?: () => Promise<void> | void,
  ) {
    // 异步处理用户登录操作并获取 accessToken
    let userInfo: GetUserInfoModel = {
      avatar: '',
      homePath: '',
      nickname: '',
      roleName: [],
      userId: '',
      username: '',
      departmentName: '',
      realName: '',
    };

    try {
      loginLoading.value = true;

      let resp: BaseDataResp<LoginResp>;

      switch (loginType) {
        case 'captcha': {
          resp = await login(params as LoginReq);
          break;
        }
        case 'email': {
          resp = await loginByEmail(params as LoginByEmailReq);
          break;
        }
        case 'mobile': {
          resp = await loginBySms(params as LoginBySmsReq);
          break;
        }
        // No default
      }

      const { data } = resp;

      // 如果成功获取到 accessToken
      if (data.token) {
        accessStore.setAccessToken(data.token);

        // 获取用户信息并存储到 accessStore 中
        const [fetchUserInfoResult, accessCodes] = await Promise.all([
          fetchUserInfo(),
          getPermCode(),
        ]);

        userInfo = fetchUserInfoResult;

        accessStore.setAccessCodes(accessCodes.data);

        if (accessStore.loginExpired) {
          accessStore.setLoginExpired(false);
        } else {
          onSuccess
            ? await onSuccess?.()
            : await router.push(
                userInfo?.homePath || preferences.app.defaultHomePath,
              );
        }

        if (userInfo?.nickname) {
          notification.success({
            description: `${$t('authentication.loginSuccessDesc')}:${userInfo?.nickname}`,
            duration: 3,
            message: $t('authentication.loginSuccess'),
          });
        }
      }
    } finally {
      loginLoading.value = false;
    }

    return {
      userInfo,
    };
  }

  async function logout(redirect: boolean = true) {
    if (isLoggingOut.value) {
      return; // 如果正在登出，直接返回，避免重复登出
    }
    isLoggingOut.value = true;
    try {
      await doLogout();
    } catch {
      // 不做任何处理
    }
    resetAllStores();
    accessStore.setLoginExpired(false);

    // 回登录页带上当前路由地址
    await router.replace({
      path: LOGIN_PATH,
      query: redirect
        ? {
            redirect: encodeURIComponent(router.currentRoute.value.fullPath),
          }
        : {},
    });
    isLoggingOut.value = false;
  }

  async function fetchUserInfo() {
    let userInfo: GetUserInfoModel;
    const result = await getUserInfo();
    // eslint-disable-next-line prefer-const
    userInfo = result.data;
    if (
      userInfo.avatar === undefined ||
      userInfo.avatar === null ||
      userInfo.avatar === ''
    ) {
      userInfo.avatar = preferences.app.defaultAvatar;
    }
    userInfo.realName = userInfo.nickname;
    userStore.setUserInfo(userInfo as any);
    return userInfo;
  }

  function $reset() {
    loginLoading.value = false;
  }

  const elementPermissionList = ref<string[]>([]);

  function hasElementPermission(
    value?: string | string[],
    condition: 'AND' | 'OR' = 'OR',
  ): boolean {
    if (!value) {
      return false;
    }

    if (condition === 'OR') {
      if (isArray(value)) {
        for (const e of value) {
          if (elementPermissionList.value.includes(e)) {
            return true;
          }
        }
      } else {
        if (elementPermissionList.value.includes(value)) {
          return true;
        }
      }
    } else {
      if (isArray(value)) {
        for (const e of value) {
          if (!elementPermissionList.value.includes(e)) {
            return false;
          }
        }
      } else {
        if (elementPermissionList.value.includes(value)) {
          return true;
        }
      }
    }

    return false;
  }

  return {
    $reset,
    authLogin,
    fetchUserInfo,
    loginLoading,
    logout,
    isLoggingOut,
    elementPermissionList,
    hasElementPermission,
  };
});
