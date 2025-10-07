import { type BaseResp } from '#/api/model/baseModel';
import { requestClient } from '#/api/request';

enum Api {
  InitializeDatabase = '/fms-api/init/database',
}

/**
 * @description: initialize the file manager database
 */

export const initializeFileDatabase = () => {
  return requestClient.get<BaseResp>(Api.InitializeDatabase);
};
