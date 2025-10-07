<script lang="ts" setup>
import { useVbenForm } from '#/adapter/form';
import { sendSms } from '#/api/mcms/messageSender';
import { smsSenderFormSchemas } from '#/views/mcms/smsProvider/schemas';
import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';
import { message } from 'ant-design-vue';

defineOptions({
  name: 'SmsSenderFormModal',
});

const [Form, formApi] = useVbenForm({
  handleSubmit: onSubmit,
  schema: [...(smsSenderFormSchemas.schema as any)],
  showDefaultActions: false,
  layout: 'vertical',
});

async function onSubmit(values: any) {
  const result = await sendSms(values);
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
  title: $t('mcms.sms.sendSms'),
});
</script>

<template>
  <Modal>
    <Form />
  </Modal>
</template>
