import type { BaseListResp } from '../../model/baseModel';

/**
 *  author: Ryan Su
 *  @description: file info response
 */
export interface fileInfo {
  id: string;
  createdAt?: number;
  name: string;
  fileType: string;
  size: number;
  path: string;
  publicPath: string;
  tagIds: number[];
}

/**
 *  author: Ryan Su
 *  @description: file list response
 */

export type FileListResp = BaseListResp<fileInfo>;

/**
 *  author: Ryan Su
 *  @description: change status request
 */
export interface changeStatusReq {
  id: string;
  status: boolean;
}

/**
 *  author: Ryan Su
 *  @description: update file info request
 */
export interface updateFileInfoReq {
  id: string;
  name: string;
  tagIds: number[];
}

/**
 *  @description: file deletion request
 */

export interface FileDeleteReq {
  url: string;
}
