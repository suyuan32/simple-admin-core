import { type BaseResp } from '#/api/model/baseModel';
import { requestClient } from '#/api/request';

enum Api {
  InitializeDatabase = '/sys-api/core/init/database',
  InitializeJobDatabase = '/sys-api/core/init/job_database',
  InitializeMcmsDatabase = '/sys-api/core/init/mcms_database',
}

/**
 * @description: initialize the core database
 */

export const initialzeCoreDatabase = () => {
  return requestClient.get<BaseResp>(Api.InitializeDatabase);
};

/**
 * @description: initialize the job management service database
 */

export const initializeJobDatabase = () => {
  return requestClient.get<BaseResp>(Api.InitializeJobDatabase);
};

/**
 * @description: initialize the message center management service database
 */

export const initializeMcmsDatabase = () => {
  return requestClient.get<BaseResp>(Api.InitializeMcmsDatabase);
};
