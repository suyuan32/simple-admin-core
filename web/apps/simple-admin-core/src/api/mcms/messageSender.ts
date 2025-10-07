import { type BaseResp } from '#/api/model/baseModel';
import { requestClient } from '#/api/request';

import { type SendEmailReq, type SendSmsReq } from './model/messageModel';

enum Api {
  SendEmail = '/sys-api/email/send',
  SendSms = '/sys-api/sms/send',
}

/**
 * @description: Send Email
 */

export const sendEmail = (params: SendEmailReq) => {
  return requestClient.post<BaseResp>(Api.SendEmail, params);
};

/**
 * @description: Send Sms
 */

export const sendSms = (params: SendSmsReq) => {
  return requestClient.post<BaseResp>(Api.SendSms, params);
};
