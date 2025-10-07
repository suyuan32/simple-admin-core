import type { LayoutType } from '@vben/types';

import {
  defineOverridesPreferences,
  updatePreferences,
} from '@vben/preferences';

import { defineStore } from 'pinia';

import { getPublicSystemConfigurationList } from '#/api/sys/configuration';

interface DynamicConfig {
  systemName: string;
  systemLogo: string;
  showSettingButton: boolean;
  showNotice: boolean;
  layoutType: LayoutType;
  showBreadCrumb: boolean;
}

export const useDynamicConfigStore = defineStore('dynamic-config', {
  state: (): DynamicConfig => ({
    systemName: '',
    systemLogo: '',
    showSettingButton: true,
    showNotice: false,
    layoutType: 'sidebar-nav',
    showBreadCrumb: true,
  }),
  getters: {
    getSystemName(): string {
      return this.systemName;
    },
    getSystemLogo(): string {
      return this.systemLogo;
    },
  },
  actions: {
    async getDynamicConfigFromServer() {
      const config = await getPublicSystemConfigurationList();

      if (config.code === 0 && config.data.total !== 0) {
        for (const v of config.data.data) {
          if (v.key !== undefined) {
            switch (v.key) {
              case 'sys.ui.header.showNotice': {
                if (v.state === false) {
                  this.showNotice = false;
                  break;
                }
                this.showNotice = v.value !== undefined && v.value === 'true';
                break;
              }
              case 'sys.ui.layoutType': {
                if (v.state === false) {
                  this.layoutType = 'sidebar-nav';
                  break;
                }
                this.layoutType =
                  v.value !== undefined &&
                  (v.value === 'sidebar-nav' ||
                    v.value === 'sidebar-mixed-nav' ||
                    v.value === 'header-nav' ||
                    v.value === 'mixed-nav' ||
                    v.value === 'full-content')
                    ? (v.value as any)
                    : 'sidebar-nav';
                break;
              }
              case 'sys.ui.logo': {
                if (v.state === false) {
                  this.systemLogo = '';
                  break;
                }
                this.systemLogo = v.value === undefined ? '' : v.value;
                break;
              }
              case 'sys.ui.name': {
                if (v.state === false) {
                  this.systemName = '';
                  break;
                }
                this.systemName = v.value === undefined ? '' : v.value;
                break;
              }
              case 'sys.ui.showBreadCrumb': {
                if (v.state === false) {
                  this.showBreadCrumb = true;
                  break;
                }
                this.showBreadCrumb = !(
                  v.value !== undefined && v.value === 'false'
                );
                break;
              }
              case 'sys.ui.showSettingButton': {
                if (v.state === false) {
                  this.showSettingButton = true;
                  break;
                }
                this.showSettingButton =
                  v.value !== undefined && v.value === 'true';
                break;
              }
            }
          }
        }
        const overridesPreferences = defineOverridesPreferences({
          app: {
            enablePreferences: this.showSettingButton,
            name: this.systemName === '' ? 'Simple Admin' : this.systemName,
          },
          logo: {
            enable: true,
            source:
              this.systemLogo === ''
                ? 'https://simpleadmin-2024.oss-cn-shanghai.aliyuncs.com/logo.png'
                : this.systemLogo,
          },
          breadcrumb: {
            enable: this.showBreadCrumb,
          },
          tabbar: {
            keepAlive: true,
            persist: true,
          },
          widget: {
            notification: this.showNotice,
          },
        });
        updatePreferences(overridesPreferences);
      } else if (config.code === 0 && config.data.total === 0) {
        this.systemName = '';
        this.systemLogo = '';
        this.showSettingButton = true;
        this.showNotice = false;
        this.layoutType = 'sidebar-nav';
        this.showBreadCrumb = true;
      }
    },
  },
  persist: true,
});
