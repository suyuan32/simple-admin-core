import { type BaseListResp } from '../../model/baseModel';

/**
 *  @description: Member info response
 */
export interface MemberInfo {
  id?: string;
  createdAt?: number;
  updatedAt?: number;
  status?: number;
  username?: string;
  password?: string;
  nickname?: string;
  rankId?: number;
  mobile?: string;
  email?: string;
  avatar?: string;
  expiredAt?: number;
}

/**
 *  @description: Member list response
 */

export type MemberListResp = BaseListResp<MemberInfo>;
