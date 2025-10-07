<script lang="ts" setup>
import type { MenuInfoPlain } from '#/api/sys/model/menuModel';

import { useVbenForm } from '#/adapter/form';
import { createMenu, updateMenu } from '#/api/sys/menu';
import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';
import { message } from 'ant-design-vue';
import { ref } from 'vue';

import { dataFormSchemas } from './schemas';

defineOptions({
  name: 'MenuForm',
});

const record = ref();
const isUpdate = ref(false);
const gridApi = ref();

async function onSubmit(values: Record<string, any>) {
  values.id = isUpdate.value ? values.id : 0;
  if (values.menuType === 2) {
    values.hideMenu = true;
  }

  const result = isUpdate.value
    ? await updateMenu(values as MenuInfoPlain)
    : await createMenu(values as MenuInfoPlain);
  if (result.code === 0) {
    message.success(result.msg);
    gridApi.value.reload();
  }
}

const [Form, formApi] = useVbenForm({
  handleSubmit: onSubmit,
  schema: [...(dataFormSchemas.schema as any)],
  showDefaultActions: false,
  layout: 'vertical',
  commonConfig: {
    // 所有表单项
    componentProps: {
      class: 'w-full',
    },
  },
  wrapperClass: 'grid-cols-2',
});

const [Modal, modalApi] = useVbenModal({
  fullscreenButton: false,
  onCancel() {
    modalApi.close();
  },
  onConfirm: async () => {
    const validationResult = await formApi.validate();
    if (validationResult.valid) {
      await formApi.submitForm();
      modalApi.close();
    }
  },
  onOpenChange(isOpen: boolean) {
    isUpdate.value = modalApi.getData()?.isUpdate;
    record.value = isOpen ? modalApi.getData()?.record || {} : {};
    gridApi.value = isOpen ? modalApi.getData()?.gridApi : null;
    if (isOpen) {
      formApi.setValues(record.value);
    }
    modalApi.setState({
      title: isUpdate.value ? $t('sys.menu.editMenu') : $t('sys.menu.addMenu'),
    });
  },
});

defineExpose(modalApi);
</script>
<template>
  <Modal class="w-1/2">
    <Form />
  </Modal>
</template>
