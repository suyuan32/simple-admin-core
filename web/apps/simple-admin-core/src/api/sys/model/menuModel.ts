import type { RouteRecordStringComponent } from '@vben-core/typings';
import type { RouteMeta } from 'vue-router';

import type { BaseListResp } from '../../model/baseModel';

export interface RouteItem {
  id?: number;
  parentId?: number;
  path: string;
  component: any;
  meta: RouteMeta;
  name?: string;
  alias?: string | string[];
  redirect?: string;
  caseSensitive?: boolean;
  children?: RouteItem[];
  permission?: string;
}

/**
 *  author: ryan
 *  @description: Get menu by page
 */
export interface MenuPageResp {
  total: number;
  data: RouteItem[];
}

export interface MenuInfoPlain {
  id?: number;
  type?: number;
  trans?: string;
  parentId?: number;
  path?: string;
  name?: string;
  redirect?: string;
  component?: string;
  sort?: number;
  disabled?: boolean;
  createdAt?: number;
  updatedAt?: number;
  title?: string;
  icon?: string;
  hideMenu?: boolean;
  hideBreadcrumb?: boolean;
  ignoreKeepAlive?: boolean;
  hideTab?: boolean;
  frameSrc?: string;
  carryParam?: boolean;
  hideChildrenInMenu?: boolean;
  affix?: boolean;
  dynamicLevel?: number;
  realPath?: string;
  serviceName?: string;
  permission?: string;
}

/**
 *  author: ryan
 *  @description: menu list response model
 */
export type MenuPlainListResp = BaseListResp<MenuInfoPlain>;

/**
 * @description: Get menu return value
 */
export type RoleMenuResp = BaseListResp<RouteRecordStringComponent>;
