import type {
  BaseDataResp,
  BaseListReq,
  BaseResp,
  BaseUUIDReq,
  BaseUUIDsReq,
} from '#/api/model/baseModel';

import type {
  CloudFileDeleteReq,
  CloudFileInfo,
  CloudFileListResp,
} from './model/cloudFileModel';

import { requestClient } from '#/api/request';

enum Api {
  CreateCloudFile = '/fms-api/cloud_file/create',
  DeleteCloudFile = '/fms-api/cloud_file/delete',
  DeleteCloudFileByUrl = '/fms-api/cloud_file/delete_by_url',
  GetCloudFileById = '/fms-api/cloud_file',
  GetCloudFileList = '/fms-api/cloud_file/list',
  UpdateCloudFile = '/fms-api/cloud_file/update',
  uploadFile = '/fms-api/cloud_file/upload',
}

/**
 * @description: Get cloud file list
 */

export const getCloudFileList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<CloudFileListResp>>(
    Api.GetCloudFileList,
    params,
  );
};

/**
 *  @description: Create a new cloud file
 */
export const createCloudFile = (params: CloudFileInfo) => {
  return requestClient.post<BaseResp>(Api.CreateCloudFile, params);
};

/**
 *  @description: Update the cloud file
 */
export const updateCloudFile = (params: CloudFileInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateCloudFile, params);
};

/**
 *  @description: Delete cloud files
 */
export const deleteCloudFile = (params: BaseUUIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteCloudFile, params);
};

/**
 *  @description: Get cloud file By ID
 */
export const getCloudFileById = (params: BaseUUIDReq) => {
  return requestClient.post<BaseDataResp<CloudFileInfo>>(
    Api.GetCloudFileById,
    params,
  );
};

/**
 * @description: Upload interface
 */
export function uploadCloudFile(file: File, provider: string = '') {
  return requestClient.upload(Api.uploadFile, { file, provider });
}

/**
 *  @description: Delete cloud file by url
 */
export const deleteCloudFileByUrl = (params: CloudFileDeleteReq) => {
  return requestClient.post<BaseResp>(Api.DeleteCloudFileByUrl, params);
};
