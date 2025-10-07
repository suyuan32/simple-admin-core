import { type VbenFormProps } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { z } from '#/adapter/form';

export const dataFormSchemas: VbenFormProps = {
  schema: [
    {
      fieldName: 'avatar',
      label: $t('sys.user.avatar'),
      component: 'ImageUpload',
      componentProps: {
        accept: ['png', 'jpeg', 'jpg'],
        maxSize: 2,
        maxNumber: 1,
        multiple: false,
        provider: 'cloud-default',
      },
    },
    {
      fieldName: 'nickname',
      label: $t('sys.user.nickname'),
      component: 'Input',
      rules: z.string().max(40),
    },
    {
      fieldName: 'mobile',
      label: $t('sys.login.mobile'),
      component: 'Input',
      rules: z.string().max(20).optional(),
    },
    {
      fieldName: 'email',
      label: $t('sys.login.email'),
      component: 'Input',
      rules: z.string().email(),
    },
  ],
};

export const changePasswordFormSchemas: VbenFormProps = {
  schema: [
    {
      fieldName: 'oldPassword',
      label: $t('sys.user.oldPassword'),
      component: 'Input',
      rules: z.string().max(40).min(6),
    },
    {
      fieldName: 'newPassword',
      label: $t('sys.user.newPassword'),
      component: 'Input',
      rules: z.string().max(40).min(6),
    },
  ],
};
