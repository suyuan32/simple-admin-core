import {
  type BaseDataResp,
  type BaseIDReq,
  type BaseIDsReq,
  type BaseListReq,
  type BaseResp,
} from '#/api/model/baseModel';
import { requestClient } from '#/api/request';

import {
  type CloudFileTagInfo,
  type CloudFileTagListResp,
} from './model/cloudFileTagModel';

enum Api {
  CreateCloudFileTag = '/fms-api/cloud_file_tag/create',
  DeleteCloudFileTag = '/fms-api/cloud_file_tag/delete',
  GetCloudFileTagById = '/fms-api/cloud_file_tag',
  GetCloudFileTagList = '/fms-api/cloud_file_tag/list',
  UpdateCloudFileTag = '/fms-api/cloud_file_tag/update',
}

/**
 * @description: Get cloud file tag list
 */

export const getCloudFileTagList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<CloudFileTagListResp>>(
    Api.GetCloudFileTagList,
    params,
  );
};

/**
 *  @description: Create a new cloud file tag
 */
export const createCloudFileTag = (params: CloudFileTagInfo) => {
  return requestClient.post<BaseResp>(Api.CreateCloudFileTag, params);
};

/**
 *  @description: Update the cloud file tag
 */
export const updateCloudFileTag = (params: CloudFileTagInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateCloudFileTag, params);
};

/**
 *  @description: Delete cloud file tags
 */
export const deleteCloudFileTag = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteCloudFileTag, params);
};

/**
 *  @description: Get cloud file tag By ID
 */
export const getCloudFileTagById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<CloudFileTagInfo>>(
    Api.GetCloudFileTagById,
    params,
  );
};
