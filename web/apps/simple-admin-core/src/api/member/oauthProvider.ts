import {
  type BaseDataResp,
  type BaseIDReq,
  type BaseIDsReq,
  type BaseListReq,
  type BaseResp,
} from '#/api/model/baseModel';
import { requestClient } from '#/api/request';

import {
  type OauthProviderInfo,
  type OauthProviderListResp,
} from '../sys/model/oauthProviderModel';

enum Api {
  CreateOauthProvider = '/mms-api/oauth_provider/create',
  DeleteOauthProvider = '/mms-api/oauth_provider/delete',
  GetOauthProviderById = '/mms-api/oauth_provider',
  GetOauthProviderList = '/mms-api/oauth_provider/list',
  UpdateOauthProvider = '/mms-api/oauth_provider/update',
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
