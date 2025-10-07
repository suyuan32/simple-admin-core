<script lang="ts" setup>
import type { UserInfo } from '#/api/sys/model/userModel';

import { useVbenForm } from '#/adapter/form';
import { createUser, updateUser } from '#/api/sys/user';
import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';
import { message } from 'ant-design-vue';
import { ref } from 'vue';

import { dataFormSchemas } from './schema';

defineOptions({
  name: 'UserForm',
});

const record = ref();
const isUpdate = ref(false);
const gridApi = ref();

async function onSubmit(values: Record<string, any>) {
  if (
    isUpdate.value === false &&
    (values.password === undefined || values.password.length === 0)
  ) {
    message.error($t('sys.login.passwordPlaceholder'));
    return;
  }

  const result = isUpdate.value
    ? await updateUser(values as UserInfo)
    : await createUser(values as UserInfo);
  if (result.code === 0) {
    message.success(result.msg);
    gridApi.value.reload();
  }
}

const [Form, formApi] = useVbenForm({
  handleSubmit: onSubmit as any,
  schema: [...(dataFormSchemas.schema as any)],
  showDefaultActions: false,
  layout: 'vertical',
});

const [Modal, modalApi] = useVbenModal({
  fullscreenButton: false,
  onCancel() {
    modalApi.close();
  },
  onConfirm: async () => {
    const values = await formApi.submitForm();
    if (
      isUpdate.value === false &&
      (values.password === undefined || values.password.length === 0)
    ) {
      return;
    }
    modalApi.close();
  },
  onOpenChange(isOpen: boolean) {
    isUpdate.value = modalApi.getData()?.isUpdate;
    record.value = isOpen ? modalApi.getData()?.record || {} : {};
    gridApi.value = isOpen ? modalApi.getData()?.gridApi : null;
    if (isOpen) {
      formApi.setValues(record.value);
    }
    modalApi.setState({
      title: isUpdate.value ? $t('sys.user.editUser') : $t('sys.user.addUser'),
    });
  },
});

defineExpose(modalApi);
</script>
<template>
  <Modal>
    <Form />
  </Modal>
</template>
