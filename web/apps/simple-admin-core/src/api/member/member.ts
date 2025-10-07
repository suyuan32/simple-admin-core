import {
  type BaseDataResp,
  type BaseListReq,
  type BaseResp,
  type BaseUUIDReq,
  type BaseUUIDsReq,
} from '#/api/model/baseModel';
import { requestClient } from '#/api/request';

import { type MemberInfo, type MemberListResp } from './model/memberModel';

enum Api {
  CreateMember = '/mms-api/member/create',
  DeleteMember = '/mms-api/member/delete',
  GetMemberById = '/mms-api/member',
  GetMemberList = '/mms-api/member/list',
  UpdateMember = '/mms-api/member/update',
}

/**
 * @description: Get member list
 */

export const getMemberList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<MemberListResp>>(
    Api.GetMemberList,
    params,
  );
};

/**
 *  @description: Create a new member
 */
export const createMember = (params: MemberInfo) => {
  return requestClient.post<BaseResp>(Api.CreateMember, params);
};

/**
 *  @description: Update the member
 */
export const updateMember = (params: MemberInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateMember, params);
};

/**
 *  @description: Delete members
 */
export const deleteMember = (params: BaseUUIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteMember, params);
};

/**
 *  @description: Get member By ID
 */
export const getMemberById = (params: BaseUUIDReq) => {
  return requestClient.post<BaseDataResp<MemberInfo>>(
    Api.GetMemberById,
    params,
  );
};
