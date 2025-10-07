import type { AxiosResponseHeaders, RequestClientOptions } from '@vben/request';

import { useAuthStore } from '#/store';
import { useAppConfig } from '@vben/hooks';
import { $t } from '@vben/locales';
import { preferences } from '@vben/preferences';
/**
 * 该文件可自行根据业务逻辑进行调整
 */
import {
  authenticateResponseInterceptor,
  errorMessageResponseInterceptor,
  RequestClient,
} from '@vben/request';
import { useAccessStore } from '@vben/stores';
import { cloneDeep } from '@vben/utils';

import { message } from 'ant-design-vue';
import JSONBigInt from 'json-bigint';

import { refreshTokenApi } from './core';

const { apiURL } = useAppConfig(import.meta.env, import.meta.env.PROD);

function createRequestClient(baseURL: string, options?: RequestClientOptions) {
  const client = new RequestClient({
    ...options,
    baseURL,
    transformResponse: (data: any, header: AxiosResponseHeaders) => {
      // storeAsString指示将BigInt存储为字符串，设为false则会存储为内置的BigInt类型
      return header.getContentType()?.toString().includes('application/json')
        ? cloneDeep(
            JSONBigInt({ storeAsString: true, strict: true }).parse(data),
          )
        : data;
    },
  });

  /**
   * 重新认证逻辑
   */
  async function doReAuthenticate() {
    console.warn('Access token or refresh token is invalid or expired. ');
    const accessStore = useAccessStore();
    const authStore = useAuthStore();
    if (authStore.isLoggingOut) {
      return; // 如果已经在登出中，跳过处理
    }
    accessStore.setAccessToken(null);
    if (
      preferences.app.loginExpiredMode === 'modal' &&
      accessStore.isAccessChecked
    ) {
      accessStore.setLoginExpired(true);
    } else {
      await authStore.logout();
    }
  }

  /**
   * 刷新token逻辑
   */
  async function doRefreshToken() {
    const accessStore = useAccessStore();
    const resp = await refreshTokenApi();
    const newToken = resp.data;
    accessStore.setAccessToken(newToken);
    return newToken;
  }

  function formatToken(token: null | string) {
    return token ? `Bearer ${token}` : null;
  }

  // 请求头处理
  client.addRequestInterceptor({
    fulfilled: async (config) => {
      const accessStore = useAccessStore();

      config.headers.Authorization = formatToken(accessStore.accessToken);
      config.headers['Accept-Language'] =
        preferences.app.locale === 'zh-CN' ? 'zh' : 'en';
      return config;
    },
  });

  // response数据解构
  client.addResponseInterceptor<any>({
    fulfilled: (response) => {
      const { data: responseData } = response;

      const { code, msg } = responseData;
      if (code !== undefined && code !== 0) {
        message.error(msg);
      }

      return responseData;

      throw Object.assign({}, response, { response });
    },
  });

  // token过期的处理
  client.addResponseInterceptor(
    authenticateResponseInterceptor({
      client,
      doReAuthenticate,
      doRefreshToken,
      enableRefreshToken: preferences.app.enableRefreshToken,
      formatToken,
    }),
  );

  // 通用的错误处理,如果没有进入上面的错误处理逻辑，就会进入这里
  client.addResponseInterceptor(
    errorMessageResponseInterceptor((msg: string, error) => {
      // 这里可以根据业务进行定制,你可以拿到 error 内的信息进行定制化处理，根据不同的 code 做不同的提示，而不是直接使用 message.error 提示 msg
      // 当前mock接口返回的错误字段是 error 或者 message
      const responseData = error?.response?.data ?? {};

      if (error.status !== 200) {
        let errMessage = '';

        switch (error.status) {
          case 400: {
            errMessage = $t(msg);
            break;
          }
          // 401: Not logged in
          // Jump to the login page if not logged in, and carry the path of the current page
          // Return to the current page after successful login. This step needs to be operated on the login page.
          case 401: {
            errMessage = $t('ui.fallback.http.unauthorized');
            break;
          }
          case 403: {
            errMessage = $t('sys.api.errMsg403');
            break;
          }
          // 404请求不存在
          case 404: {
            errMessage = $t('sys.api.errMsg404');
            break;
          }
          case 405: {
            errMessage = $t('sys.api.errMsg405');
            break;
          }
          case 408: {
            errMessage = $t('sys.api.errMsg408');
            break;
          }
          case 500: {
            errMessage = $t('sys.api.errMsg500');
            break;
          }
          case 501: {
            errMessage = $t('sys.api.errMsg501');
            break;
          }
          case 502: {
            errMessage = $t('sys.api.errMsg502');
            break;
          }
          case 503: {
            errMessage = $t('sys.api.errMsg503');
            break;
          }
          case 504: {
            errMessage = $t('sys.api.errMsg504');
            break;
          }
          case 505: {
            errMessage = $t('sys.api.errMsg505');
            break;
          }
          default:
        }

        message.error(errMessage);
        return;
      }

      // 如果没有错误信息，则会根据状态码进行提示
      if (responseData?.code !== 0) {
        message.error(responseData?.msg);
        return;
      }

      message.error(msg);
    }),
  );

  return client;
}

export const requestClient = createRequestClient(apiURL);

export const baseRequestClient = new RequestClient({ baseURL: apiURL });

export interface PageFetchParams {
  pageNo?: number;
  pageSize?: number;

  [key: string]: any;
}
