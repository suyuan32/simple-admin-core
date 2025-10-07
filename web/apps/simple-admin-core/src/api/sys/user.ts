import { requestClient } from '#/api/request';

import {
  type BaseDataResp,
  type BaseListReq,
  type BaseResp,
  type BaseUUIDReq,
  type BaseUUIDsReq,
} from '../model/baseModel';
import {
  type ChangePasswordReq,
  type GetUserInfoModel,
  type LoginByEmailReq,
  type LoginBySmsReq,
  type LoginReq,
  type LoginResp,
  type RegisterByEmailReq,
  type RegisterBySmsReq,
  type RegisterReq,
  type ResetByEmailInfo,
  type ResetBySmsInfo,
  type UserInfo,
  type UserListResp,
  type UserProfile,
} from './model/userModel';

enum Api {
  ChangePassword = '/sys-api/user/change_password',
  CreateUser = '/sys-api/user/create',
  DeleteUser = '/sys-api/user/delete',
  GetPermCode = '/sys-api/user/perm',
  GetUserById = '/sys-api/user',
  GetUserInfo = '/sys-api/user/info',
  GetUserList = '/sys-api/user/list',
  Login = '/sys-api/user/login',
  LoginByEmail = '/sys-api/user/login_by_email',
  LoginBySms = '/sys-api/user/login_by_sms',
  Logout = '/sys-api/user/logout',
  Profile = '/sys-api/user/profile',
  Register = '/sys-api/user/register',
  RegisterByEmail = '/sys-api/user/register_by_email',
  RegisterBySms = '/sys-api/user/register_by_sms',
  ResetPasswordByEmail = '/sys-api/user/reset_password_by_email',
  ResetPasswordBySms = '/sys-api/user/reset_password_by_sms',
  UpdateUser = '/sys-api/user/update',
}

/**
 * @description: Get user list
 */

export const getUserList = (params: BaseListReq) => {
  return requestClient.post<BaseDataResp<UserListResp>>(
    Api.GetUserList,
    params,
  );
};

/**
 *  @description: Create a new user
 */
export const createUser = (params: UserInfo) => {
  return requestClient.post<BaseResp>(Api.CreateUser, params);
};

/**
 *  @description: Update the user
 */
export const updateUser = (params: UserInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateUser, params);
};

/**
 *  @description: Delete users
 */
export const deleteUser = (params: BaseUUIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteUser, params);
};

/**
 *  @description: Get user By ID
 */
export const getUserById = (params: BaseUUIDReq) => {
  return requestClient.post<BaseDataResp<UserInfo>>(Api.GetUserById, params);
};

/**
 * @description: User login api
 */
export function login(params: LoginReq) {
  return requestClient.post<BaseDataResp<LoginResp>>(Api.Login, params);
}

/**
 * @description: User login by email api
 */
export function loginByEmail(params: LoginByEmailReq) {
  return requestClient.post<BaseDataResp<LoginResp>>(Api.LoginByEmail, params);
}

/**
 * @description: User login by sms api
 */
export function loginBySms(params: LoginBySmsReq) {
  return requestClient.post<BaseDataResp<LoginResp>>(Api.LoginBySms, params);
}

/**
 * @description: User register api
 */
export function register(params: RegisterReq) {
  return requestClient.post<BaseResp>(Api.Register, params);
}

/**
 * @description: User register by email api
 */
export function registerByEmail(params: RegisterByEmailReq) {
  return requestClient.post<BaseResp>(Api.RegisterByEmail, params);
}

/**
 * @description: User register by Sms api
 */
export function registerBySms(params: RegisterBySmsReq) {
  return requestClient.post<BaseResp>(Api.RegisterBySms, params);
}

/**
 * @description: Get user's basic info
 */

export function getUserInfo() {
  return requestClient.get<BaseDataResp<GetUserInfoModel>>(Api.GetUserInfo);
}

export function getPermCode() {
  return requestClient.get<BaseDataResp<string[]>>(Api.GetPermCode);
}

export function doLogout() {
  return requestClient.get(Api.Logout);
}

/**
 *  author: Ryan Su
 *  @description: Get user profile
 */
export function getUserProfile() {
  return requestClient.get<BaseDataResp<UserProfile>>(Api.Profile);
}

/**
 *  author: Ryan Su
 *  @description: update user profile
 */
export function updateProfile(params: UserProfile) {
  return requestClient.post<BaseResp>(Api.Profile, params);
}

/**
 *  author: Ryan Su
 *  @description: change user password
 */

export function changePassword(params: ChangePasswordReq) {
  return requestClient.post<BaseResp>(Api.ChangePassword, params);
}

/**
 *  author: Ryan Su
 *  @description: reset user password by email
 */

export function resetPasswordByEmail(params: ResetByEmailInfo) {
  return requestClient.post<BaseResp>(Api.ResetPasswordByEmail, params);
}

/**
 *  author: Ryan Su
 *  @description: reset user password by email
 */

export function resetPasswordBySms(params: ResetBySmsInfo) {
  return requestClient.post<BaseResp>(Api.ResetPasswordBySms, params);
}
