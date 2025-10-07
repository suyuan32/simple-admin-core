import {
  type BaseDataResp,
  type BaseIDReq,
  type BaseIDsReq,
  type BaseListReq,
  type BaseResp,
} from '#/api/model/baseModel';
import { requestClient } from '#/api/request';

import {
  type MemberRankInfo,
  type MemberRankListResp,
} from '../member/model/memberRankModel';

enum Api {
  CreateMemberRank = '/mms-api/member_rank/create',
  DeleteMemberRank = '/mms-api/member_rank/delete',
  GetMemberRankById = '/mms-api/member_rank',
  GetMemberRankList = '/mms-api/member_rank/list',
  UpdateMemberRank = '/mms-api/member_rank/update',
}

/**
 * @description: Get member rank list
 */

export const getMemberRankList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<MemberRankListResp>>(
    Api.GetMemberRankList,
    params,
  );
};

/**
 *  @description: Create a new member rank
 */
export const createMemberRank = (params: MemberRankInfo) => {
  return requestClient.post<BaseResp>(Api.CreateMemberRank, params);
};

/**
 *  @description: Update the member rank
 */
export const updateMemberRank = (params: MemberRankInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateMemberRank, params);
};

/**
 *  @description: Delete member ranks
 */
export const deleteMemberRank = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteMemberRank, params);
};

/**
 *  @description: Get member rank By ID
 */
export const getMemberRankById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<MemberRankInfo>>(
    Api.GetMemberRankById,
    params,
  );
};
