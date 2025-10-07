import { type BaseListResp } from '../../model/baseModel';

/**
 *  @description: Dictionary info response
 */
export interface DictionaryInfo {
  id?: number;
  createdAt?: number;
  updatedAt?: number;
  title?: string;
  name?: string;
  status?: number;
  desc?: string;
}

/**
 *  @description: Dictionary list response
 */

export type DictionaryListResp = BaseListResp<DictionaryInfo>;
