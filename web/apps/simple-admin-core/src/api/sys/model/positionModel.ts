import { type BaseListResp } from '../../model/baseModel';

/**
 *  @description: Position info response
 */
export interface PositionInfo {
  id?: number;
  createdAt?: number;
  updatedAt?: number;
  status?: number;
  sort?: number;
  name?: string;
  code?: string;
  remark?: string;
}

/**
 *  @description: Position list response
 */

export type PositionListResp = BaseListResp<PositionInfo>;
