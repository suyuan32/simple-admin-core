import { type BaseListResp } from '../../model/baseModel';

/**
 *  @description: MemberRank info response
 */
export interface MemberRankInfo {
  id?: number;
  createdAt?: number;
  updatedAt?: number;
  name?: string;
  description?: string;
  remark?: string;
}

/**
 *  @description: MemberRank list response
 */

export type MemberRankListResp = BaseListResp<MemberRankInfo>;
