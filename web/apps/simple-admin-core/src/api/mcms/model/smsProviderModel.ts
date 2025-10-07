import { type BaseListResp } from '../../model/baseModel';

/**
 *  @description: SmsProvider info response
 */
export interface SmsProviderInfo {
  id?: number;
  createdAt?: number;
  updatedAt?: number;
  name?: string;
  secretId?: string;
  secretKey?: string;
  region?: string;
  isDefault?: boolean;
}

/**
 *  @description: SmsProvider list response
 */

export type SmsProviderListResp = BaseListResp<SmsProviderInfo>;
