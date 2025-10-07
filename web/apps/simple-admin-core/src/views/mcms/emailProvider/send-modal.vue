<script lang="ts" setup>
import { useVbenForm } from '#/adapter/form';
import { sendEmail } from '#/api/mcms/messageSender';
import { emailSenderFormSchemas } from '#/views/mcms/emailProvider/schemas';
import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';
import { message } from 'ant-design-vue';

defineOptions({
  name: 'EmailSenderFormModal',
});

const [Form, formApi] = useVbenForm({
  handleSubmit: onSubmit,
  schema: [...(emailSenderFormSchemas.schema as any)],
  showDefaultActions: false,
  layout: 'vertical',
});

async function onSubmit(values: any) {
  const result = await sendEmail(values);
  if (result.code === 0) {
    message.info($t('common.successful'));
  }
}

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
  onOpenChange(_isOpen: boolean) {},
  title: $t('mcms.email.sendEmail'),
});
</script>

<template>
  <Modal>
    <Form />
  </Modal>
</template>
