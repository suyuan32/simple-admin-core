import {
  type BaseDataResp,
  type BaseIDReq,
  type BaseIDsReq,
  type BaseListReq,
  type BaseResp,
} from '#/api/model/baseModel';
import { requestClient } from '#/api/request';

import { type ApiInfo, type ApiListResp } from './model/apiModel';

enum Api {
  CreateApi = '/sys-api/api/create',
  DeleteApi = '/sys-api/api/delete',
  GetApiById = '/sys-api/api',
  GetApiList = '/sys-api/api/list',
  UpdateApi = '/sys-api/api/update',
}

/**
 * @description: Get api list
 */

export const getApiList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<ApiListResp>>(Api.GetApiList, params);
};

/**
 *  @description: Create a new api
 */
export const createApi = (params: ApiInfo) => {
  return requestClient.post<BaseResp>(Api.CreateApi, params);
};

/**
 *  @description: Update the api
 */
export const updateApi = (params: ApiInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateApi, params);
};

/**
 *  @description: Delete apis
 */
export const deleteApi = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteApi, params);
};

/**
 *  @description: Get api By ID
 */
export const getApiById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<ApiInfo>>(Api.GetApiById, params);
};
