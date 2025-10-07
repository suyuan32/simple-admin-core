import {
  type BaseDataResp,
  type BaseIDReq,
  type BaseIDsReq,
  type BaseListReq,
  type BaseResp,
} from '#/api/model/baseModel';
import { requestClient } from '#/api/request';

import {
  type EmailProviderInfo,
  type EmailProviderListResp,
} from './model/emailProviderModel';

enum Api {
  CreateEmailProvider = '/sys-api/email_provider/create',
  DeleteEmailProvider = '/sys-api/email_provider/delete',
  GetEmailProviderById = '/sys-api/email_provider',
  GetEmailProviderList = '/sys-api/email_provider/list',
  UpdateEmailProvider = '/sys-api/email_provider/update',
}

/**
 * @description: Get email provider list
 */

export const getEmailProviderList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<EmailProviderListResp>>(
    Api.GetEmailProviderList,
    params,
  );
};

/**
 *  @description: Create a new email provider
 */
export const createEmailProvider = (params: EmailProviderInfo) => {
  return requestClient.post<BaseResp>(Api.CreateEmailProvider, params);
};

/**
 *  @description: Update the email provider
 */
export const updateEmailProvider = (params: EmailProviderInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateEmailProvider, params);
};

/**
 *  @description: Delete email providers
 */
export const deleteEmailProvider = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteEmailProvider, params);
};

/**
 *  @description: Get email provider By ID
 */
export const getEmailProviderById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<EmailProviderInfo>>(
    Api.GetEmailProviderById,
    params,
  );
};
