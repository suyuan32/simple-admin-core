<script lang="ts" setup>
import type { BasicOption, Recordable } from '@vben/types';

import { ref } from 'vue';
import { useRouter } from 'vue-router';

import {
  AuthenticationForgetPassword,
  type VbenFormSchema,
  z,
} from '@vben/common-ui';
import { $t } from '@vben/locales';

import { getEmailCaptcha, getSmsCaptcha } from '#/api/sys/captcha';
import { resetPasswordByEmail, resetPasswordBySms } from '#/api/sys/user';

defineOptions({ name: 'ForgetPassword' });

const route = useRouter();

const loginType: BasicOption[] = [
  {
    label: $t('sys.login.mobile'),
    value: 'mobile',
  },
  {
    label: $t('sys.login.email'),
    value: 'email',
  },
];

const emailOrPhonePlaceholder = ref('');
const target = ref<string>('');
const msgType = ref<string>('');
const loading = ref(false);

async function handleSendCaptcha(): Promise<boolean> {
  if (msgType.value === 'email') {
    const result = await getEmailCaptcha({ email: target.value });
    return result.code === 0;
  } else {
    const result = await getSmsCaptcha({ phoneNumber: target.value });
    return result.code === 0;
  }
}

const formSchema: VbenFormSchema[] = [
  {
    component: 'VbenSelect',
    componentProps: {
      options: loginType,
    },
    label: 'Login Type',
    fieldName: 'selectLoginType',
    rules: z.string().optional().default('email'),
    dependencies: {
      triggerFields: ['selectLoginType'],
      trigger(values, _) {
        msgType.value = values.selectLoginType;
      },
    },
    formItemClass: 'col-span-2 items-baseline',
  },
  {
    component: 'VbenInput',
    componentProps: {
      placeholder: emailOrPhonePlaceholder,
    },
    label: 'Target',
    fieldName: 'target',
    dependencies: {
      trigger(values, _) {
        emailOrPhonePlaceholder.value =
          values.selectLoginType === 'email'
            ? $t('sys.login.emailPlaceholder')
            : $t('sys.login.mobilePlaceholder');

        target.value = values.target;
      },
      // 只有指定的字段改变时，才会触发
      triggerFields: ['selectLoginType', 'target'],
    },
    rules: z.string().min(5),
    formItemClass: 'col-span-2 items-baseline',
  },
  {
    component: 'VbenPinInput',
    componentProps: {
      createText: (countdown: number) => {
        return countdown > 0
          ? $t('authentication.sendText', [countdown])
          : $t('authentication.sendCode');
      },
      placeholder: $t('authentication.code'),
      codeLength: 5,
      handleSendCode: handleSendCaptcha,
    },
    fieldName: 'captchaVerified',
    label: $t('authentication.code'),
    rules: z
      .string()
      .length(5, { message: $t('sys.login.captchaRequired') })
      .max(5),
    formItemClass: 'col-span-2 items-baseline',
  },
  {
    component: 'VbenInputPassword',
    componentProps: {
      passwordStrength: true,
      placeholder: $t('authentication.password'),
    },
    renderComponentContent() {
      return {
        strengthText: () => $t('authentication.passwordStrength'),
      };
    },
    fieldName: 'password',
    label: $t('authentication.password'),
    rules: z.string().min(6, { message: $t('authentication.passwordTip') }),
    formItemClass: 'col-span-2 items-baseline',
  },
  {
    component: 'VbenInputPassword',
    componentProps: {
      placeholder: $t('authentication.confirmPassword'),
    },
    dependencies: {
      rules(values) {
        const { password } = values;
        return z
          .string({ required_error: $t('authentication.passwordTip') })
          .min(1, { message: $t('authentication.passwordTip') })
          .refine((value) => value === password, {
            message: $t('authentication.confirmPasswordTip'),
          });
      },
      triggerFields: ['password'],
    },
    fieldName: 'confirmPassword',
    label: $t('authentication.confirmPassword'),
    formItemClass: 'col-span-2 items-baseline',
  },
];

async function handleSubmit(value: Recordable<any>) {
  if (value.selectLoginType === 'email') {
    const result = await resetPasswordByEmail({
      email: value.target,
      captcha: value.captchaVerified,
      password: value.password,
    });

    if (result.code === 0) {
      route.replace('/auth/login');
    }
  } else {
    const result = await resetPasswordBySms({
      phoneNumber: value.target,
      captcha: value.captchaVerified,
      password: value.password,
    });

    if (result.code === 0) {
      route.replace('/auth/login');
    }
  }
}
</script>

<template>
  <AuthenticationForgetPassword
    :form-schema="formSchema"
    :loading="loading"
    :sub-title="$t('sys.login.resetPassword')"
    :submit-button-text="$t('common.resetText')"
    @submit="handleSubmit"
  />
</template>
