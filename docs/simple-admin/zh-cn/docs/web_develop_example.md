## Example web 端

### 首先添加 API

> 添加 model 到 simple-admin-backend-ui/src/api/sys/model 中
src/api/sys/model/exampleModel.ts
```typescript
/**
 *  author: Ryan Su
 *  @description: example requst
 */
export interface HelloReq {
  name: string;
}

```
> 添加 example.ts 到 simple-admin-backend-ui/src/api/sys
```typescript
import { defHttp } from '/@/utils/http/axios';
import { ErrorMessageMode } from '/#/axios';
import { BaseResp } from '/@/api/model/baseModel';
import { HelloReq } from './model/exampleModel';

enum Api {
    Hello = '/sys-api/example/hello',
}

/**
 * @description: Get hello msg
 */

export const Hello = (params: HelloReq, mode: ErrorMessageMode = 'modal') => {
    return defHttp.post<BaseResp>(
        { url: Api.Hello, params: params },
        {
            errorMessageMode: mode,
        },
    );
};

```

添加 view :  src/view/example/index.vue

```vue
<template>
  <PageWrapper>
    <p>{{ resp }}</p>
    <AForm
      :model="name"
      name="basic"
      :label-col="{ span: 8 }"
      :wrapper-col="{ span: 16 }"
      autocomplete="off"
    >
      <AFormItem
        label="Name"
        name="name"
        :rules="[{ required: true, message: 'Please input your username!' }]"
      >
        <a-input v-model:value="name" />
      </AFormItem>

      <AFormItem :wrapper-col="{ offset: 8, span: 16 }">
        <a-button type="primary" @click="SayHello">Submit</a-button>
      </AFormItem>
    </AForm>
  </PageWrapper>
</template>
<script lang="ts" setup>
  import { PageWrapper } from '/@/components/Page';
  import { ref } from 'vue';
  import { Hello } from '/@/api/sys/example';

  const name = ref<string>('');
  const resp = ref<string>('');

  async function SayHello() {
    const result = await Hello({ name: name.value }, 'message');
    resp.value = 'Hello ' + result.msg;
    console.log(result);
  }
</script>
```

注意 **await Hello({ name: name.value }, 'message')** \
># message 模式显示效果如下
![example](../../assets/example_validator_message_mode.png)
># modal 模式显示效果如下
![example](../../assets/example_validator_modal_mode.png)



> 由于默认需要支持两种语言，所以要分别设置 src/locals/zh-CN/routes/system.ts  和  src/locals/en/routes/system.ts 

![example](../../assets/example_zh_title.png)
![example](../../assets/example_en_title.png)

> 推荐使用 i18n 插件，可以直接复制路径 

![I18n](../../assets/i18n_ext.png)
![I18n](../../assets/copy_translation_path.png)

> 新增菜单

![Menu](../../assets/add_example_menu.png)

> 添加菜单权限

![Menu](../../assets/add_example_authority.png)

> API 的介绍同样最好设置中英文

![Example](../../assets/example_api_desc_title_en.png)
![Example](../../assets/example_api_desc_title_zh.png)

> 新增API

![Example](../../assets/add_example_api_zh.png)

> 添加API权限

![Example](../../assets/add_example_authority_zh.png)

> 测试页面

![Example](../../assets/example_page.png)

