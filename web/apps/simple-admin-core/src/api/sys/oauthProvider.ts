import type { LoginResp } from '#/api/sys/model/userModel';

import {
  type BaseDataResp,
  type BaseIDReq,
  type BaseIDsReq,
  type BaseListReq,
  type BaseResp,
} from '#/api/model/baseModel';
import { requestClient } from '#/api/request';

import {
  type OauthLoginReq,
  type OauthProviderInfo,
  type OauthProviderListResp,
  type RedirectResp,
} from './model/oauthProviderModel';

enum Api {
  CreateOauthProvider = '/sys-api/oauth_provider/create',
  DeleteOauthProvider = '/sys-api/oauth_provider/delete',
  GetOauthProviderById = '/sys-api/oauth_provider',
  GetOauthProviderList = '/sys-api/oauth_provider/list',
  OauthLogin = '/sys-api/oauth/login',
  OauthLoginCallback = '/sys-api/oauth/login/callback',
  UpdateOauthProvider = '/sys-api/oauth_provider/update',
}

/**
 * @description: Get oauth provider list
 */

export const getOauthProviderList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<OauthProviderListResp>>(
    Api.GetOauthProviderList,
    params,
  );
};

/**
 *  @description: Create a new oauth provider
 */
export const createOauthProvider = (params: OauthProviderInfo) => {
  return requestClient.post<BaseResp>(Api.CreateOauthProvider, params);
};

/**
 *  @description: Update the oauth provider
 */
export const updateOauthProvider = (params: OauthProviderInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateOauthProvider, params);
};

/**
 *  @description: Delete oauth providers
 */
export const deleteOauthProvider = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteOauthProvider, params);
};

/**
 *  @description: Get oauth provider By ID
 */
export const getOauthProviderById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<OauthProviderInfo>>(
    Api.GetOauthProviderById,
    params,
  );
};

/**
 *  author: Ryan Su
 *  @description: oauth log in
 */
export const oauthLogin = (params: OauthLoginReq) => {
  return requestClient.post<RedirectResp>(Api.OauthLogin, params);
};

/**
 *  author: Ryan Su
 *  @description: oauth log in callback
 */
export const oauthLoginCallback = (URL: string) => {
  return requestClient.get<LoginResp>(Api.OauthLoginCallback + URL);
};
