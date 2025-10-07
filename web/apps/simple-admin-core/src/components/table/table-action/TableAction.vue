<script setup lang="ts">
import type { ButtonType } from 'ant-design-vue/es/button';

import type { ActionItem, PopConfirm } from './types';

import { computed, type PropType, toRaw } from 'vue';

import { useAccess } from '@vben/access';
import { isBoolean, isFunction } from '@vben/utils';

import { Icon } from '@iconify/vue';
import {
  Button,
  Dropdown,
  Menu,
  Popconfirm,
  Space,
  Tooltip,
} from 'ant-design-vue';

const props = defineProps({
  actions: {
    type: Array as PropType<ActionItem[]>,
    default() {
      return [];
    },
  },
  dropDownActions: {
    type: Array as PropType<ActionItem[]>,
    default() {
      return [];
    },
  },
  divider: {
    type: Boolean,
    default: true,
  },
});

const MenuItem = Menu.Item;

const { hasAccessByCodes } = useAccess();
function isIfShow(action: ActionItem): boolean {
  const ifShow = action.ifShow;

  let isIfShow = true;

  if (isBoolean(ifShow)) {
    isIfShow = ifShow;
  }
  if (isFunction(ifShow)) {
    isIfShow = ifShow(action);
  }
  return isIfShow;
}

const getActions = computed(() => {
  return (toRaw(props.actions) || [])
    .filter((action) => {
      return (
        (hasAccessByCodes(action.auth || []) ||
          (action.auth || []).length === 0) &&
        isIfShow(action)
      );
    })
    .map((action) => {
      const { popConfirm } = action;
      return {
        // getPopupContainer: document.body,
        type: 'link' as ButtonType,
        ...action,
        ...popConfirm,
        onConfirm: popConfirm?.confirm,
        onCancel: popConfirm?.cancel,
        enable: !!popConfirm,
      };
    });
});
const getDropdownList = computed((): any[] => {
  return (toRaw(props.dropDownActions) || [])
    .filter((action) => {
      return (
        (hasAccessByCodes(action.auth || []) ||
          (action.auth || []).length === 0) &&
        isIfShow(action)
      );
    })
    .map((action, index) => {
      const { label, popConfirm } = action;
      return {
        ...action,
        ...popConfirm,
        onConfirm: popConfirm?.confirm,
        onCancel: popConfirm?.cancel,
        text: label,
        divider:
          index < props.dropDownActions.length - 1 ? props.divider : false,
      };
    });
});
const getPopConfirmProps = (attrs: PopConfirm) => {
  const originAttrs: any = attrs;
  delete originAttrs.icon;
  if (attrs.confirm && isFunction(attrs.confirm)) {
    originAttrs.onConfirm = attrs.confirm;
    delete originAttrs.confirm;
  }
  if (attrs.cancel && isFunction(attrs.cancel)) {
    originAttrs.onCancel = attrs.cancel;
    delete originAttrs.cancel;
  }

  delete originAttrs.tooltip;
  return originAttrs;
};
const getButtonProps = (action: ActionItem) => {
  let colorClass = '';
  switch (action.color) {
    case 'error': {
      colorClass = 'text-red-500 hover:!text-red-400';
      break;
    }
    case 'success': {
      colorClass = 'text-green-500 hover:!text-green-400';
      break;
    }
    case 'warning': {
      colorClass = 'text-yellow-500 hover:!text-yellow-400';
      break;
    }
    default: {
      colorClass = 'text-blue-500 hover:!text-blue-400';
      break;
    }
  }

  const btnProps = {
    type: action.type || 'primary',
    ...action,
    class: colorClass,
  };
  delete btnProps.icon;
  delete btnProps.tooltip;
  return btnProps;
};
const handleMenuClick = (e: any) => {
  const action = getDropdownList.value[e.key];
  if (action.onClick && isFunction(action.onClick)) {
    action.onClick();
  }
};
</script>

<template>
  <div class="m-table-action">
    <Space :size="2">
      <template v-for="(action, index) in getActions" :key="index">
        <Popconfirm
          v-if="action.popConfirm"
          v-bind="getPopConfirmProps(action.popConfirm)"
        >
          <template v-if="action.popConfirm.icon" #icon>
            <Tooltip :title="action.tooltip">
              <Icon :icon="action.popConfirm.icon" />
            </Tooltip>
          </template>
          <Tooltip :title="action.tooltip">
            <Button :style="{}" title="" v-bind="getButtonProps(action)">
              <template v-if="action.icon" #icon>
                <Icon :icon="action.icon" />
              </template>
              {{ action.label }}
            </Button>
          </Tooltip>
        </Popconfirm>
        <Tooltip v-else :title="action.tooltip">
          <Button v-bind="getButtonProps(action)" @click="action.onClick">
            <template v-if="action.icon" #icon>
              <Icon :icon="action.icon" />
            </template>
            {{ action.label }}
          </Button>
        </Tooltip>
      </template>
    </Space>

    <Dropdown v-if="getDropdownList.length > 0" :trigger="['hover']">
      <slot name="more">
        <Button size="small" type="link">
          <template #icon>
            <Icon class="icon-more" icon="ant-design:more-outlined" />
          </template>
        </Button>
      </slot>
      <template #overlay>
        <Menu @click="handleMenuClick">
          <MenuItem v-for="(action, index) in getDropdownList" :key="index">
            <template v-if="action.popConfirm">
              <Popconfirm v-bind="getPopConfirmProps(action.popConfirm)">
                <template v-if="action.popConfirm.icon" #icon>
                  <Icon :icon="action.popConfirm.icon" />
                </template>
                <div>
                  <Icon v-if="action.icon" :icon="action.icon" />
                  <span class="ml-1">{{ action.text }}</span>
                </div>
              </Popconfirm>
            </template>
            <template v-else>
              <Icon v-if="action.icon" :icon="action.icon" />
              {{ action.label }}
            </template>
          </MenuItem>
        </Menu>
      </template>
    </Dropdown>
  </div>
</template>
<style lang="less">
.m-table-action {
  .ant-btn > .iconify + span,
  .ant-btn > span + .iconify {
    margin-inline-start: 8px;
  }

  .ant-btn > .iconify {
    display: inline-flex;
    align-items: center;
    width: 1em;
    height: 1em;
    font-style: normal;
    line-height: 0;
    color: inherit;
    text-align: center;
    text-transform: none;
    vertical-align: -0.125em;
    text-rendering: optimizelegibility;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
  }
}
</style>
