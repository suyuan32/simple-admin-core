import { type BaseListResp } from '../../model/baseModel';

/**
 *  @description: DictionaryDetail info response
 */
export interface DictionaryDetailInfo {
  id?: number;
  createdAt?: number;
  updatedAt?: number;
  status?: number;
  title?: string;
  key?: string;
  value?: string;
  dictionaryId?: number;
  sort?: number;
}

/**
 *  @description: DictionaryDetail list response
 */

export type DictionaryDetailListResp = BaseListResp<DictionaryDetailInfo>;

/**
 *  @description: Dictionary name request
 */
export interface DictionaryNameReq {
  name: string;
}
