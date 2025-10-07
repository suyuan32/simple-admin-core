import { type BaseListResp } from '../../model/baseModel';

/**
 *  @description: EmailLog info response
 */
export interface EmailLogInfo {
  id?: string;
  createdAt?: number;
  updatedAt?: number;
  target?: string;
  subject?: string;
  content?: string;
  sendStatus?: number;
  provider?: string;
}

/**
 *  @description: EmailLog list response
 */

export type EmailLogListResp = BaseListResp<EmailLogInfo>;
