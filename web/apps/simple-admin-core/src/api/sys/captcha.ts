import { requestClient } from '#/api/request';

import { type BaseDataResp } from '../model/baseModel';
import {
  type CaptchaResp,
  type GetEmailCaptchaReq,
  type GetSmsCaptchaReq,
} from './model/captcha';

enum Api {
  GetCaptcha = '/sys-api/captcha',
  SendEmailCaptcha = '/sys-api/captcha/email',
  SendSmsCaptcha = '/sys-api/captcha/sms',
}

/**
 * @description: Get captcha api
 */
export function getCaptcha() {
  return requestClient.get<BaseDataResp<CaptchaResp>>(Api.GetCaptcha);
}

/**
 * @description: Send email captcha
 */
export function getEmailCaptcha(params: GetEmailCaptchaReq) {
  return requestClient.post<BaseDataResp<string>>(Api.SendEmailCaptcha, params);
}

/**
 * @description: Send sms captcha
 */

export function getSmsCaptcha(params: GetSmsCaptchaReq) {
  return requestClient.post<BaseDataResp<string>>(Api.SendSmsCaptcha, params);
}
