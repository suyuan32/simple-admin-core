import type { ApiListResp } from '#/api/sys/model/apiModel';

import {
  type BaseDataResp,
  type BaseIDReq,
  type BaseResp,
} from '#/api/model/baseModel';
import { requestClient } from '#/api/request';

import {
  type ApiAuthorityReq,
  type ApiAuthorityResp,
  type ApiListReq,
  type MenuAuthorityInfo,
} from './model/authorityModel';

enum Api {
  CreateOrUpdateApiAuthority = '/sys-api/authority/api/create_or_update',
  CreateOrUpdateMenuAuthority = '/sys-api/authority/menu/create_or_update',
  GetApiList = '/sys-api/api/list',
  GetRoleApiList = '/sys-api/authority/api/role',
  GetRoleMenuList = '/sys-api/authority/menu/role',
}

/**
 *  author: Ryan Su
 *  @description: this function is used to get api list for authorization
 */

export const getApiList = (params: ApiListReq) => {
  return requestClient.post<BaseDataResp<ApiListResp>>(Api.GetApiList, params);
};

/**
 * @description: Get api authorization list
 */

export const getApiAuthority = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<ApiAuthorityResp>>(
    Api.GetRoleApiList,
    params,
  );
};

/**
 *  author: ryan
 *  @description: create or update api authorization
 */
export const createOrUpdateApiAuthority = (params: ApiAuthorityReq) => {
  return requestClient.post<BaseResp>(Api.CreateOrUpdateApiAuthority, params);
};

/**
 *  author: Ryan Su
 *  @description:
 */

export const createOrUpdateMenuAuthority = (params: MenuAuthorityInfo) => {
  return requestClient.post<BaseResp>(Api.CreateOrUpdateMenuAuthority, params);
};

/**
 *  author: Ryan Su
 *  @description: get role's menu authorization ids
 */

export const getMenuAuthority = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<MenuAuthorityInfo>>(
    Api.GetRoleMenuList,
    params,
  );
};
