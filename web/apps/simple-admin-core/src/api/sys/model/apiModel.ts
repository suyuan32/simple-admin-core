import { type BaseListResp } from '../../model/baseModel';

/**
 *  @description: Api info response
 */
export interface ApiInfo {
  id?: number;
  createdAt?: number;
  updatedAt?: number;
  trans?: string;
  path: string;
  description?: string;
  group: string;
  method: string;
  isRequired: boolean;
  serviceName: string;
}

/**
 *  @description: Api list response
 */

export type ApiListResp = BaseListResp<ApiInfo>;
