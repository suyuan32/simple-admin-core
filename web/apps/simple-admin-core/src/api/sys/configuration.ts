import {
  type BaseDataResp,
  type BaseIDReq,
  type BaseIDsReq,
  type BaseResp,
} from '#/api/model/baseModel';
import { requestClient } from '#/api/request';

import {
  type ConfigurationInfo,
  type ConfigurationListReq,
  type ConfigurationListResp,
} from './model/configurationModel';

enum Api {
  CreateConfiguration = '/sys-api/configuration/create',
  DeleteConfiguration = '/sys-api/configuration/delete',
  GetConfigurationById = '/sys-api/configuration',
  GetConfigurationList = '/sys-api/configuration/list',
  GetPublicSystemConfigurationList = '/sys-api/configuration/system/list',
  UpdateConfiguration = '/sys-api/configuration/update',
}

/**
 * @description: Get configuration list
 */

export const getConfigurationList = (params: ConfigurationListReq) => {
  return requestClient.post<BaseDataResp<ConfigurationListResp>>(
    Api.GetConfigurationList,
    params,
  );
};

/**
 * @description: Get public system configuration list
 */

export const getPublicSystemConfigurationList = () => {
  return requestClient.get<BaseDataResp<ConfigurationListResp>>(
    Api.GetPublicSystemConfigurationList,
  );
};

/**
 *  @description: Create a new configuration
 */
export const createConfiguration = (params: ConfigurationInfo) => {
  return requestClient.post<BaseResp>(Api.CreateConfiguration, params);
};

/**
 *  @description: Update the configuration
 */
export const updateConfiguration = (params: ConfigurationInfo) => {
  return requestClient.post<BaseResp>(Api.UpdateConfiguration, params);
};

/**
 *  @description: Delete configurations
 */
export const deleteConfiguration = (params: BaseIDsReq) => {
  return requestClient.post<BaseResp>(Api.DeleteConfiguration, params);
};

/**
 *  @description: Get configuration By ID
 */
export const getConfigurationById = (params: BaseIDReq) => {
  return requestClient.post<BaseDataResp<ConfigurationInfo>>(
    Api.GetConfigurationById,
    params,
  );
};
