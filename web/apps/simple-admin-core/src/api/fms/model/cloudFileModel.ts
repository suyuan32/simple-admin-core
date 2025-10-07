import type { BaseListResp } from '../../model/baseModel';

/**
 *  @description: CloudFile info response
 */
export interface CloudFileInfo {
  id?: string;
  createdAt?: number;
  updatedAt?: number;
  state?: boolean;
  name?: string;
  url?: string;
  size?: number;
  fileType?: number;
  userId?: string;
  providerId?: number;
  tagIds?: number[];
}

/**
 *  @description: CloudFile list response
 */

export type CloudFileListResp = BaseListResp<CloudFileInfo>;

/**
 *  @description: CloudFile deletion request
 */

export interface CloudFileDeleteReq {
  url: string;
}
