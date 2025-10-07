import { type BaseListResp } from '../../model/baseModel';

/**
 *  @description: Tag info response
 */
export interface TagInfo {
  id?: number;
  createdAt?: number;
  updatedAt?: number;
  status?: number;
  name?: string;
  remark?: string;
}

/**
 *  @description: Tag list response
 */

export type TagListResp = BaseListResp<TagInfo>;
