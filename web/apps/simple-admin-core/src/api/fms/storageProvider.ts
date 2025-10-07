import {
  type BaseDataResp,
  type BaseIDReq,
  type BaseIDsReq,
  type BaseListReq,
  type BaseResp,
} from '#/api/model/baseModel';
import { requestClient } from '#/api/request';

import {
  type StorageProviderInfo,
  type StorageProviderListResp,
} from './model/storageProviderModel';

enum Api {
  CreateStorageProvider = '/fms-api/storage_provider/create',
  DeleteStorageProvider = '/fms-api/storage_provider/delete',
  GetStorageProviderById = '/fms-api/storage_provider',
  GetStorageProviderList = '/fms-api/storage_provider/list',
  UpdateStorageProvider = '/fms-api/storage_provider/update',
}

/**
 * @description: Get storage provider list
 */

export const getStorageProviderList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<StorageProviderListResp>>(
    Api.GetStorageProviderList,
    params,
  );
};

/**
 *  @description: Create a new storage provider
 */
export const createStorageProvider = (params: StorageProviderInfo) => {
  return requestClient.post<BaseResp>(Api.CreateStorageProvider, params);
};

/**
 *  @description: Update the storage provider
 */
export const updateStorageProvider = (params: StorageProviderInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateStorageProvider, params);
};

/**
 *  @description: Delete storage providers
 */
export const deleteStorageProvider = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteStorageProvider, params);
};

/**
 *  @description: Get storage provider By ID
 */
export const getStorageProviderById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<StorageProviderInfo>>(
    Api.GetStorageProviderById,
    params,
  );
};
