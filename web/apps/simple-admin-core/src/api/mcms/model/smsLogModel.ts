import { type BaseListResp } from '../../model/baseModel';

/**
 *  @description: SmsLog info response
 */
export interface SmsLogInfo {
  id?: string;
  createdAt?: number;
  updatedAt?: number;
  phoneNumber?: string;
  content?: string;
  sendStatus?: number;
  provider?: string;
}

/**
 *  @description: SmsLog list response
 */

export type SmsLogListResp = BaseListResp<SmsLogInfo>;
