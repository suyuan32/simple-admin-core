import { type BaseListResp } from '../../model/baseModel';

/**
 *  @description: EmailProvider info response
 */
export interface EmailProviderInfo {
  id?: number;
  createdAt?: number;
  updatedAt?: number;
  name?: string;
  authType?: number;
  emailAddr?: string;
  password?: string;
  hostName?: string;
  identify?: string;
  secret?: string;
  port?: number;
  tls?: boolean;
  isDefault?: boolean;
}

/**
 *  @description: EmailProvider list response
 */

export type EmailProviderListResp = BaseListResp<EmailProviderInfo>;
