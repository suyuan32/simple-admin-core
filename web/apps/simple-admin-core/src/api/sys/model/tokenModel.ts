import { type BaseListResp } from '../../model/baseModel';

/**
 *  @description: Token info response
 */
export interface TokenInfo {
  id?: string;
  createdAt?: number;
  updatedAt?: number;
  status?: number;
  uuid?: string;
  token?: string;
  source?: string;
  expiredAt?: number;
  username?: string;
}

/**
 *  @description: Token list response
 */

export type TokenListResp = BaseListResp<TokenInfo>;
