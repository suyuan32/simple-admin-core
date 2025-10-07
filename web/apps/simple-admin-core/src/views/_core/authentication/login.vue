<script lang="ts" setup>
import type { VbenFormSchema } from '#/adapter/form';
import type { BasicOption } from '@vben/types';

import { getCaptcha, getEmailCaptcha, getSmsCaptcha } from '#/api/sys/captcha';
import { oauthLogin } from '#/api/sys/oauthProvider';
import { useAuthStore } from '#/store';
import { AuthenticationLogin, z } from '@vben/common-ui';
import { $t } from '@vben/locales';
import { usePreferences } from '@vben/preferences';
import { Image } from 'ant-design-vue';
import { computed, h, ref } from 'vue';

defineOptions({ name: 'Login' });

const authStore = useAuthStore();

const { isDark } = usePreferences();

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

const imgPath = ref<string>('');
const captchaId = ref<string>('');
const msgType = ref<string>('');
const target = ref<string>('');

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
      formItemClass: 'col-span-2 items-baseline',
      dependencies: {
        triggerFields: ['selectLoginType'],
        trigger(values, _) {
          msgType.value = values.selectLoginType;
        },
      },
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
        if(values) {
          return values.selectLoginType === 'captcha';
        },
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
            values.selectLoginType === 'email'
              ? $t('sys.login.emailPlaceholder')
              : $t('sys.login.mobilePlaceholder');

          target.value = values.target;
        },
        if(values) {
          return values.selectLoginType !== 'captcha';
        },
        triggerFields: ['selectLoginType', 'target'],
      },
      formItemClass: 'col-span-2 items-baseline',
    },
    {
      component: 'VbenInputPassword',
      componentProps: {
        placeholder: $t('authentication.password'),
      },
      dependencies: {
        if(values) {
          return values.selectLoginType === 'captcha';
        },
        triggerFields: ['selectLoginType'],
      },
      fieldName: 'password',
      label: $t('authentication.password'),
      rules: z.string().min(6, { message: $t('authentication.passwordTip') }),
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

async function handleLogin(values: any) {
  switch (values.selectLoginType) {
    case 'captcha': {
      authStore
        .authLogin(
          {
            password: values.password,
            username: values.username,
            captcha: values.captcha,
            captchaId: captchaId.value,
          },
          'captcha',
        )
        .then(() => {})
        .catch(() => {
          getCaptchaData();
        });

      break;
    }
    case 'email': {
      authStore
        .authLogin(
          {
            captcha: values.captchaVerified,
            email: values.target,
          },
          'email',
        )
        .then(() => {})
        .catch(() => {
          getCaptchaData();
        });

      break;
    }
    case 'mobile': {
      authStore
        .authLogin(
          {
            captcha: values.captchaVerified,
            phoneNumber: values.target,
          },
          'mobile',
        )
        .then(() => {})
        .catch(() => {
          getCaptchaData();
        });

      break;
    }
    // No default
  }
}

const thirdPartyProviderList: any[] = [
  {
    icon: 'icon-park-outline:github',
    oauthProvider: 'github',
  },
  {
    icon: 'mingcute:google-fill',
    oauthProvider: 'google',
  },
];

async function handleOauthLogin(provider: string) {
  const result = await oauthLogin({
    state: `${new Date().getMilliseconds()}-${provider}`,
    provider,
  });
  if (result.code === 0) window.open(result.data.URL);
}
</script>

<template>
  <AuthenticationLogin
    :form-schema="formSchema"
    :loading="authStore.loginLoading"
    :show-code-login="false"
    :show-qrcode-login="false"
    :third-party-provider-list="thirdPartyProviderList"
    @oauth-login="handleOauthLogin"
    @submit="handleLogin"
  />
</template>
