import { type BaseListResp } from '../../model/baseModel';

/**
 *  @description: Role info response
 */
export interface RoleInfo {
  id?: number;
  createdAt?: number;
  updatedAt?: number;
  status?: number;
  name?: string;
  code?: string;
  defaultRouter?: string;
  remark?: string;
  sort?: number;
}

/**
 *  @description: Role list response
 */

export type RoleListResp = BaseListResp<RoleInfo>;
