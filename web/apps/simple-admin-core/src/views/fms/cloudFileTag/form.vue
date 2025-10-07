<script lang="ts" setup>
import type { CloudFileTagInfo } from '#/api/fms/model/cloudFileTagModel';

import { useVbenForm } from '#/adapter/form';
import { createCloudFileTag, updateCloudFileTag } from '#/api/fms/cloudFileTag';
import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';
import { message } from 'ant-design-vue';
import { ref } from 'vue';

import { dataFormSchemas } from './schemas';

defineOptions({
  name: 'CloudFileTagForm',
});

const record = ref();
const isUpdate = ref(false);
const gridApi = ref();

async function onSubmit(values: Record<string, any>) {
  const result = isUpdate.value
    ? await updateCloudFileTag(values as CloudFileTagInfo)
    : await createCloudFileTag(values as CloudFileTagInfo);
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
      title: isUpdate.value
        ? $t('fms.cloudFileTag.editCloudFileTag')
        : $t('fms.cloudFileTag.addCloudFileTag'),
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
