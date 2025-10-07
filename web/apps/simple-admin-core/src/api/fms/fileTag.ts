import {
  type BaseDataResp,
  type BaseIDReq,
  type BaseIDsReq,
  type BaseListReq,
  type BaseResp,
} from '#/api/model/baseModel';
import { requestClient } from '#/api/request';

import { type TagInfo, type TagListResp } from './model/fileTagModel';

enum Api {
  CreateTag = '/fms-api/file_tag/create',
  DeleteTag = '/fms-api/file_tag/delete',
  GetTagById = '/fms-api/file_tag',
  GetTagList = '/fms-api/file_tag/list',
  UpdateTag = '/fms-api/file_tag/update',
}

/**
 * @description: Get tag list
 */

export const getTagList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<TagListResp>>(Api.GetTagList, params);
};

/**
 *  @description: Create a new tag
 */
export const createTag = (params: TagInfo) => {
  return requestClient.post<BaseResp>(Api.CreateTag, params);
};

/**
 *  @description: Update the tag
 */
export const updateTag = (params: TagInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateTag, params);
};

/**
 *  @description: Delete tags
 */
export const deleteTag = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteTag, params);
};

/**
 *  @description: Get tag By ID
 */
export const getTagById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<TagInfo>>(Api.GetTagById, params);
};
