import { type BaseResp } from '#/api/model/baseModel';
import { requestClient } from '#/api/request';

enum Api {
  InitializeMMSDatabase = '/mms-api/init/database',
}

/**
 * @description: initialize the member management service database
 */

export const initializeMMSDatabase = () => {
  return requestClient.get<BaseResp>(Api.InitializeMMSDatabase);
};
