<script lang="ts" setup>
import { ref } from 'vue';

import { Page } from '@vben/common-ui';
import { $t } from '@vben/locales';

import {
  Button,
  Card,
  Col,
  Input,
  message,
  Row,
  TypographyTitle,
} from 'ant-design-vue';

import { initializeFileDatabase } from '#/api/fms/initialize';
import { initializeMMSDatabase } from '#/api/member/initialize';
// api
import {
  initializeJobDatabase,
  initializeMcmsDatabase,
  initialzeCoreDatabase,
} from '#/api/sys/initialize';

defineOptions({
  name: 'InitializationPage',
});
const coreInitButtonLoading = ref<boolean>(false);
const fileInitButtonLoading = ref<boolean>(false);
const mmsInitButtonLoading = ref<boolean>(false);
const jobInitButtonLoading = ref<boolean>(false);
const mcmsInitButtonLoading = ref<boolean>(false);
const customInitButtonLoading = ref<boolean>(false);
const customInitUrl = ref<string>('http://localhost');
const customInitPort = ref<string>('9100');
const customInitService = ref<string>('');

async function initCustomDatabase() {
  const serviceName: string =
    customInitService.value === '' ? '' : `/${customInitService.value}`;
  customInitButtonLoading.value = true;
  window.open(
    `${customInitUrl.value}:${customInitPort.value}${
      serviceName
    }/init/database`,
    '_blank',
  );
  customInitButtonLoading.value = false;
}

async function initCoreDatabase() {
  coreInitButtonLoading.value = true;
  const result = await initialzeCoreDatabase().finally(() => {
    coreInitButtonLoading.value = false;
  });
  if (result.code === 0) {
    message.success(result.msg, 3);
  }
}

async function initFileDatabase() {
  fileInitButtonLoading.value = true;
  const result = await initializeFileDatabase().finally(() => {
    fileInitButtonLoading.value = false;
  });
  if (result.code === 0) {
    message.success(result.msg, 3);
  }
}

async function initMMSDatabase() {
  mmsInitButtonLoading.value = true;
  const result = await initializeMMSDatabase().finally(() => {
    mmsInitButtonLoading.value = false;
  });
  if (result.code === 0) {
    message.success(result.msg, 3);
  }
}

async function initJobDatabase() {
  jobInitButtonLoading.value = true;
  const result = await initializeJobDatabase().finally(() => {
    jobInitButtonLoading.value = false;
  });
  if (result.code === 0) {
    message.success(result.msg, 3);
  }
}

async function initMcmsDatabase() {
  mcmsInitButtonLoading.value = true;
  const result = await initializeMcmsDatabase().finally(() => {
    mcmsInitButtonLoading.value = false;
  });
  if (result.code === 0) {
    message.success(result.msg, 3);
  }
}
</script>
<template>
  <Page>
    <Row :gutter="[16, 16]">
      <Col :span="6">
        <Card :hoverable="true" :title="$t('sys.init.initCoreDatabase')">
          <Button
            class="w-full"
            href="https://github.com/suyuan32/simple-admin-core"
            type="link"
          >
            Core Github
          </Button>

          <Button
            :loading="coreInitButtonLoading"
            class="w-full"
            type="primary"
            @click="initCoreDatabase"
          >
            {{ $t('common.start') }}
          </Button>
        </Card>
      </Col>
      <Col :span="6">
        <Card :hoverable="true" :title="$t('sys.init.initFileDatabase')">
          <Button
            class="w-full"
            href="https://github.com/suyuan32/simple-admin-file"
            type="link"
          >
            File Manager Github
          </Button>
          <Button
            :loading="fileInitButtonLoading"
            class="w-full"
            type="primary"
            @click="initFileDatabase"
          >
            {{ $t('common.start') }}
          </Button>
        </Card>
      </Col>
      <Col :span="6">
        <Card :hoverable="true" :title="$t('sys.init.initMMSDatabase')">
          <Button
            class="w-full"
            href="https://github.com/suyuan32/simple-admin-member-api"
            type="link"
          >
            Member Github
          </Button>
          <Button
            :loading="mmsInitButtonLoading"
            class="w-full"
            type="primary"
            @click="initMMSDatabase"
          >
            {{ $t('common.start') }}
          </Button>
        </Card>
      </Col>
      <Col :span="6">
        <Card :hoverable="true" :title="$t('sys.init.initJobDatabase')">
          <Button
            class="w-full"
            href="https://github.com/suyuan32/simple-admin-job"
            type="link"
          >
            Job Github
          </Button>
          <Button
            :loading="jobInitButtonLoading"
            class="w-full"
            type="primary"
            @click="initJobDatabase"
          >
            {{ $t('common.start') }}
          </Button>
        </Card>
      </Col>
      <Col :span="6">
        <Card :hoverable="true" :title="$t('sys.init.initMcmsDatabase')">
          <Button
            class="w-full"
            href="https://github.com/suyuan32/simple-admin-message-center"
            type="link"
          >
            Mcms Github
          </Button>
          <Button
            :loading="mcmsInitButtonLoading"
            class="w-full"
            type="primary"
            @click="initMcmsDatabase"
          >
            {{ $t('common.start') }}
          </Button>
        </Card>
      </Col>
    </Row>
    <Row class="pt-2">
      <Col :span="12">
        <Card :hoverable="true" :title="$t('sys.init.initCustom')">
          <TypographyTitle :level="5">
            {{ $t('sys.init.initUrl') }}
          </TypographyTitle>
          <p>
            <Input v-model:value="customInitUrl" />
          </p>
          <TypographyTitle :level="5">
            {{ $t('sys.init.initPort') }}
          </TypographyTitle>
          <p>
            <Input v-model:value="customInitPort" />
          </p>
          <TypographyTitle :level="5">
            {{ $t('sys.init.initService') }}
          </TypographyTitle>
          <p>
            <Input
              v-model:value="customInitService"
              :placeholder="$t('sys.init.initOptional')"
            />
          </p>
          <TypographyTitle :level="5">
            {{ $t('sys.init.initRedirect') }}
          </TypographyTitle>
          <p>
            <Button
              :loading="customInitButtonLoading"
              class="w-full"
              type="primary"
              @click="initCustomDatabase"
            >
              {{ $t('common.start') }}
            </Button>
          </p>
        </Card>
      </Col>
    </Row>
  </Page>
</template>
