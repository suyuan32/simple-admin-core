import { type BaseListResp } from '#/api/model/baseModel';

/**
 *  @description: User info response
 */
export interface UserInfo {
  id?: string;
  createdAt?: number;
  updatedAt?: number;
  status?: number;
  username?: string;
  nickname?: string;
  password?: string;
  description?: string;
  homePath?: string;
  roleIds?: number[];
  mobile?: string;
  email?: string;
  avatar?: string;
  departmentId?: number;
  positionIds?: number[];
}

/**
 *  @description: User list response
 */

export type UserListResp = BaseListResp<UserInfo>;

export interface UserProfile {
  avatar: string;
  nickname: string;
  email: string;
  mobile: string;
}

export interface ChangePasswordReq {
  oldPassword: string;
  newPassword: string;
}

/**
 * @description: Login interface parameters
 */
export interface LoginReq {
  username: string;
  password: string;
  captcha: string;
  captchaId: string;
}

/**
 * @description: Login by email interface parameters
 */
export interface LoginByEmailReq {
  email: string;
  captcha: string;
}

/**
 * @description: Login by sms interface parameters
 */
export interface LoginBySmsReq {
  phoneNumber: string;
  captcha: string;
}

/**
 * @description: Register interface parameters
 */
export interface RegisterReq {
  username: string;
  password: string;
  email: string;
  captcha: string;
  captchaId: string;
}

/**
 * @description: Register by email interface parameters
 */
export interface RegisterByEmailReq {
  username: string;
  password: string;
  email: string;
  captcha: string;
}

/**
 * @description: Register by sms interface parameters
 */
export interface RegisterBySmsReq {
  username: string;
  password: string;
  phoneNumber: string;
  captcha: string;
}

/**
 * @description: Login interface return value
 */
export interface LoginResp {
  userId: number | string;
  token: string;
  expire?: number;
}

/**
 * @description: Get user information return value
 */
export interface GetUserInfoModel {
  // 用户id
  userId: number | string;
  // 用户名
  username: string;
  // 昵称
  nickname: string;
  // 头像
  avatar: string;
  // 介绍
  desc?: string;
  // 用户首页
  homePath: string;
  // 用户角色
  roles?: string[];
  // 用户角色名称
  roleName: string[];
  // 部门名称
  departmentName?: string;
  // 真实名称
  realName?: string;
}

/**
 * @description: Reset password by email request
 */

export interface ResetByEmailInfo {
  email: string;
  captcha: string;
  password: string;
}

/**
 * @description: Reset password by sms request
 */

export interface ResetBySmsInfo {
  phoneNumber: string;
  captcha: string;
  password: string;
}
