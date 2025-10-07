import {
  type BaseDataResp,
  type BaseIDReq,
  type BaseIDsReq,
  type BaseListReq,
  type BaseResp,
} from '#/api/model/baseModel';
import { requestClient } from '#/api/request';

import {
  type SmsProviderInfo,
  type SmsProviderListResp,
} from './model/smsProviderModel';

enum Api {
  CreateSmsProvider = '/sys-api/sms_provider/create',
  DeleteSmsProvider = '/sys-api/sms_provider/delete',
  GetSmsProviderById = '/sys-api/sms_provider',
  GetSmsProviderList = '/sys-api/sms_provider/list',
  UpdateSmsProvider = '/sys-api/sms_provider/update',
}

/**
 * @description: Get sms provider list
 */

export const getSmsProviderList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<SmsProviderListResp>>(
    Api.GetSmsProviderList,
    params,
  );
};

/**
 *  @description: Create a new sms provider
 */
export const createSmsProvider = (params: SmsProviderInfo) => {
  return requestClient.post<BaseResp>(Api.CreateSmsProvider, params);
};

/**
 *  @description: Update the sms provider
 */
export const updateSmsProvider = (params: SmsProviderInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateSmsProvider, params);
};

/**
 *  @description: Delete sms providers
 */
export const deleteSmsProvider = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteSmsProvider, params);
};

/**
 *  @description: Get sms provider By ID
 */
export const getSmsProviderById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<SmsProviderInfo>>(
    Api.GetSmsProviderById,
    params,
  );
};
