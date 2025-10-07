import {
  type BaseDataResp,
  type BaseListReq,
  type BaseResp,
  type BaseUUIDReq,
  type BaseUUIDsReq,
} from '#/api/model/baseModel';
import { requestClient } from '#/api/request';

import { type TokenInfo, type TokenListResp } from '../sys/model/tokenModel';

enum Api {
  CreateToken = '/mms-api/token/create',
  DeleteToken = '/mms-api/token/delete',
  GetTokenById = '/mms-api/token',
  GetTokenList = '/mms-api/token/list',
  Logout = '/mms-api/token/logout',
  UpdateToken = '/mms-api/token/update',
}

/**
 * @description: Get token list
 */

export const getTokenList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<TokenListResp>>(
    Api.GetTokenList,
    params,
  );
};

/**
 *  @description: Create a new token
 */
export const createToken = (params: TokenInfo) => {
  return requestClient.post<BaseResp>(Api.CreateToken, params);
};

/**
 *  @description: Update the token
 */
export const updateToken = (params: TokenInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateToken, params);
};

/**
 *  @description: Delete tokens
 */
export const deleteToken = (params: BaseUUIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteToken, params);
};

/**
 *  @description: Get token By ID
 */
export const getTokenById = (params: BaseUUIDReq) => {
  return requestClient.post<BaseDataResp<TokenInfo>>(Api.GetTokenById, params);
};

/**
 *  @description: Force user log out
 */

export const logout = (id: string) => requestClient.post(Api.Logout, { id });
