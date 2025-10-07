export type {
  AlertProps,
  BeforeCloseScope,
  IconType,
  PromptProps,
} from './alert';
export { useAlertContext } from './alert';
export { default as Alert } from './alert.vue';
export {
  clearAllAlerts,
  vbenAlert as alert,
  vbenConfirm as confirm,
  vbenPrompt as prompt,
} from './AlertBuilder';
