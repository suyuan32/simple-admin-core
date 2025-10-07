import {
  type BaseDataResp,
  type BaseIDReq,
  type BaseIDsReq,
  type BaseListReq,
  type BaseResp,
} from '#/api/model/baseModel';
import { requestClient } from '#/api/request';

import {
  type DictionaryDetailInfo,
  type DictionaryDetailListResp,
  type DictionaryNameReq,
} from './model/dictionaryDetailModel';

enum Api {
  CreateDictionaryDetail = '/sys-api/dictionary_detail/create',
  DeleteDictionaryDetail = '/sys-api/dictionary_detail/delete',
  GetDictionaryDetailByDictionaryName = '/sys-api/dict',
  GetDictionaryDetailById = '/sys-api/dictionary_detail',
  GetDictionaryDetailList = '/sys-api/dictionary_detail/list',
  UpdateDictionaryDetail = '/sys-api/dictionary_detail/update',
}

/**
 * @description: Get dictionary detail list
 */

export const getDictionaryDetailList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<DictionaryDetailListResp>>(
    Api.GetDictionaryDetailList,
    params,
  );
};

/**
 *  @description: Create a new dictionary detail
 */
export const createDictionaryDetail = (params: DictionaryDetailInfo) => {
  return requestClient.post<BaseResp>(Api.CreateDictionaryDetail, params);
};

/**
 *  @description: Update the dictionary detail
 */
export const updateDictionaryDetail = (params: DictionaryDetailInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateDictionaryDetail, params);
};

/**
 *  @description: Delete dictionary details
 */
export const deleteDictionaryDetail = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteDictionaryDetail, params);
};

/**
 *  @description: Get dictionary detail By ID
 */
export const getDictionaryDetailById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<DictionaryDetailInfo>>(
    Api.GetDictionaryDetailById,
    params,
  );
};

/**
 *  @description: Get dictionary detail By Dictionary name
 */
export const GetDictionaryDetailByDictionaryName = (
  params: DictionaryNameReq,
) => {
  return requestClient.get<BaseDataResp<DictionaryDetailListResp>>(
    `${Api.GetDictionaryDetailByDictionaryName}/${params.name}`,
  );
};
