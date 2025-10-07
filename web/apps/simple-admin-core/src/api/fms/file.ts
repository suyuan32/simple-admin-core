import type {
  BaseDataResp,
  BaseListReq,
  BaseResp,
  BaseUUIDsReq,
} from '../model/baseModel';
import type {
  FileDeleteReq,
  FileListResp,
  updateFileInfoReq,
} from './model/fileModel';

import { requestClient } from '#/api/request';
// import { type UploadApiResp } from '@/api/sys/model/uploadModel';
// import { type AxiosProgressEvent } from 'axios';

enum Api {
  DeleteFile = '/fms-api/file/delete',
  DeleteFileByUrl = '/fms-api/file/delete_by_url',
  DownloadFile = '/fms-api/file/download',
  GetFileList = '/fms-api/file/list',
  SetFileStatus = '/fms-api/file/status',
  UpdateFileInfo = '/fms-api/file/update',
  uploadFile = '/fms-api/upload',
}

/**
 * @description: Upload interface
 */
export function uploadFile(file: File) {
  return requestClient.upload(Api.uploadFile, { file });
}

/**
 * @description: Get file list
 */

export const getFileList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<FileListResp>>(
    Api.GetFileList,
    params,
  );
};

/**
 *  author: ryan
 *  @description: update file info
 */
export const updateFileInfo = (params: updateFileInfoReq) => {
  return requestClient.post<BaseResp>(Api.UpdateFileInfo, params);
};

/**
 *  author: Ryan Su
 *  @description: delete files
 */
export const deleteFile = (params: BaseUUIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteFile, params);
};

/**
 *  author: Ryan Su
 *  @description: set file's status
 */
export const setFileStatus = (id: string, status: number) =>
  requestClient.post(Api.SetFileStatus, { id, status });

/**
 *  author: Ryan Su
 *  @description: download file
 */

export const downloadFile = (id: number) =>
  requestClient.download(`${Api.DownloadFile}/${id}`);

/**
 *  author: Ryan Su
 *  @description: delete files
 */
export const deleteFileByUrl = (params: FileDeleteReq) => {
  return requestClient.post<BaseResp>(Api.DeleteFileByUrl, params);
};
