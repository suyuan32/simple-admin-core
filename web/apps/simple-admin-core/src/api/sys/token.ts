import {
  type BaseDataResp,
  type BaseListReq,
  type BaseResp,
  type BaseUUIDReq,
  type BaseUUIDsReq,
} from '#/api/model/baseModel';
import { requestClient } from '#/api/request';

import { type TokenInfo, type TokenListResp } from './model/tokenModel';

enum Api {
  CreateToken = '/sys-api/token/create',
  DeleteToken = '/sys-api/token/delete',
  GetTokenById = '/sys-api/token',
  GetTokenList = '/sys-api/token/list',
  Logout = '/sys-api/token/logout',
  UpdateToken = '/sys-api/token/update',
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
