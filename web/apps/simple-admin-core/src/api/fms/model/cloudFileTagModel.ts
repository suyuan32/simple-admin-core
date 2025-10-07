import { type BaseListResp } from '../../model/baseModel';

/**
 *  @description: CloudFileTag info response
 */
export interface CloudFileTagInfo {
  id?: number;
  createdAt?: number;
  updatedAt?: number;
  status?: number;
  name?: string;
  remark?: string;
}

/**
 *  @description: CloudFileTag list response
 */

export type CloudFileTagListResp = BaseListResp<CloudFileTagInfo>;
