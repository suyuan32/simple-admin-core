<script lang="ts" setup>
import type { VbenFormSchema } from '@vben/common-ui';
import type { BasicOption } from '@vben/types';

import { getCaptcha, getEmailCaptcha, getSmsCaptcha } from '#/api/sys/captcha';
import { register, registerByEmail, registerBySms } from '#/api/sys/user';
import { AuthenticationRegister, z } from '@vben/common-ui';
import { $t } from '@vben/locales';
import { usePreferences } from '@vben/preferences';
import { Image, message } from 'ant-design-vue';
import { computed, h, ref } from 'vue';
import { useRouter } from 'vue-router';

defineOptions({ name: 'Register' });

const loading = ref(false);
const router = useRouter();
const { isDark } = usePreferences();

const imgPath = ref<string>('');
const captchaId = ref<string>('');
const msgType = ref<string>('');
const target = ref<string>('');

const loginType: BasicOption[] = [
  {
    label: $t('sys.login.captcha'),
    value: 'captcha',
  },
  {
    label: $t('sys.login.mobile'),
    value: 'mobile',
  },
  {
    label: $t('sys.login.email'),
    value: 'email',
  },
];

// get captcha
async function getCaptchaData() {
  const captcha = await getCaptcha();
  if (captcha.code === 0) {
    captchaId.value = captcha.data.captchaId;
    imgPath.value = captcha.data.imgPath;
  }
}

getCaptchaData();

async function handleSendCaptcha(): Promise<boolean> {
  if (msgType.value === 'email') {
    const result = await getEmailCaptcha({ email: target.value });
    return result.code === 0;
  } else {
    const result = await getSmsCaptcha({ phoneNumber: target.value });
    return result.code === 0;
  }
}

const emailOrPhonePlaceholder = ref('');

const formSchema = computed((): VbenFormSchema[] => {
  return [
    {
      component: 'VbenSelect',
      componentProps: {
        options: loginType,
      },
      label: 'Login Type',
      fieldName: 'selectLoginType',
      rules: z.string().optional().default('captcha'),
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
        placeholder: $t('authentication.usernameTip'),
      },
      fieldName: 'username',
      label: $t('authentication.username'),
      rules: z.string().min(1, { message: $t('authentication.usernameTip') }),
      dependencies: {
        triggerFields: ['selectLoginType'],
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
            values.selectLoginType === 'email' ||
            values.selectLoginType === 'captcha'
              ? $t('sys.login.emailPlaceholder')
              : $t('sys.login.mobilePlaceholder');

          target.value = values.target;
        },
        triggerFields: ['selectLoginType', 'target'],
      },
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
    {
      fieldName: 'captcha',
      label: $t('authentication.password'),
      dependencies: {
        if(values) {
          return values.selectLoginType === 'captcha';
        },
        triggerFields: ['selectLoginType'],
      },
      component: 'VbenInput',
      componentProps: {
        maxlength: 5,
        placeholder: $t('sys.login.smsPlaceholder'),
      },
      formItemClass: 'col-span-1 items-baseline',
      rules: z
        .string()
        .length(5, { message: $t('sys.login.captchaRequired') })
        .max(5),
    },
    {
      fieldName: 'captchaImg',
      component: h(Image),
      componentProps: {
        src: imgPath.value,
        width: 120,
        height: 40,
        preview: false,
        onClick: getCaptchaData,
        style: {
          backgroundColor: isDark.value ? '#eee' : 'transparent',
        },
      },
      formItemClass: 'col-span-1 items-baseline',
      dependencies: {
        if(values) {
          return values.selectLoginType === 'captcha';
        },
        // 只有指定的字段改变时，才会触发
        triggerFields: ['selectLoginType'],
      },
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
      dependencies: {
        if(values) {
          return values.selectLoginType !== 'captcha';
        },
        // 只有指定的字段改变时，才会触发
        triggerFields: ['selectLoginType'],
      },
      fieldName: 'captchaVerified',
      label: $t('authentication.code'),
      rules: z
        .string()
        .length(5, { message: $t('sys.login.captchaRequired') })
        .max(5),
      formItemClass: 'col-span-2 items-baseline',
    },
  ];
});

async function handleRegister(values: any) {
  loading.value = true;
  switch (values.selectLoginType) {
    case 'captcha': {
      register({
        password: values.password,
        username: values.username,
        captcha: values.captcha,
        captchaId: captchaId.value,
        email: target.value,
      })
        .then((data) => {
          if (data.code === 0) {
            message.info($t('sys.login.signupSuccessTitle'));
            router.push('/auth/login');
          } else if (data.code !== 0) {
            getCaptchaData();
          }
        })
        .catch(() => {
          getCaptchaData();
        })
        .finally(() => {
          loading.value = false;
        });

      break;
    }
    case 'email': {
      registerByEmail({
        captcha: values.captchaVerified,
        username: values.username,
        password: values.password,
        email: values.target,
      })
        .then((data) => {
          if (data.code === 0) {
            message.info($t('sys.login.signupSuccessTitle'));
            router.push('/auth/login');
          } else if (data.code !== 0) {
            getCaptchaData();
          }
        })
        .catch(() => {
          getCaptchaData();
        })
        .finally(() => {
          loading.value = false;
        });

      break;
    }
    case 'mobile': {
      registerBySms({
        captcha: values.captchaVerified,
        phoneNumber: values.target,
        username: values.username,
        password: values.password,
      })
        .then((data) => {
          if (data.code === 0) {
            message.info($t('sys.login.signupSuccessTitle'));
            router.push('/auth/login');
          } else if (data.code !== 0) {
            getCaptchaData();
          }
        })
        .catch(() => {
          getCaptchaData();
        })
        .finally(() => {
          loading.value = false;
        });

      break;
    }
    // No default
  }
}
</script>

<template>
  <AuthenticationRegister
    :form-schema="formSchema"
    :loading="loading"
    @submit="handleRegister"
  />
</template>
