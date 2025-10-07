import { defineOverridesPreferences } from '@vben/preferences';

/**
 * @description 项目配置文件
 * 只需要覆盖项目中的一部分配置，不需要的配置不用覆盖，会自动使用默认配置
 * !!! 更改配置后请清空缓存，否则可能不生效
 */
export const overridesPreferences = defineOverridesPreferences({
  // overrides
  app: {
    name: import.meta.env.VITE_APP_TITLE,
    defaultAvatar:
      'https://simpleadmin-2024.oss-cn-shanghai.aliyuncs.com/logo.png',
    accessMode: 'backend',
    defaultHomePath: '/dashboard',
  },
  logo: {
    enable: true,
    source: 'https://simpleadmin-2024.oss-cn-shanghai.aliyuncs.com/logo.png',
  },
  tabbar: {
    keepAlive: true,
    persist: true,
  },
  widget: {
    notification: false,
  },
});
