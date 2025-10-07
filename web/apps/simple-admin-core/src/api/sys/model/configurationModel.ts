import { type BaseListResp } from '../../model/baseModel';

/**
 *  @description: Configuration info response
 */
export interface ConfigurationInfo {
  id?: number;
  createdAt?: number;
  updatedAt?: number;
  sort?: number;
  state?: boolean;
  name?: string;
  key?: string;
  value?: string;
  category?: string;
  remark?: string;
}

/**
 *  @description: Configuration list response
 */

export type ConfigurationListResp = BaseListResp<ConfigurationInfo>;

export interface ConfigurationListReq {
  page: number;
  pageSize: number;
  category?: string;
}
