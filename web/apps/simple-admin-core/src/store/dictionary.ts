import type { DefaultOptionType } from 'ant-design-vue/lib/select';

import { ref } from 'vue';

import { defineStore } from 'pinia';

import { GetDictionaryDetailByDictionaryName } from '#/api/sys/dictionaryDetail';

interface DictionaryDataDefaultOptionType extends DefaultOptionType {
  status?: null | number | string;
}

interface DictionaryData {
  data: DictionaryDataDefaultOptionType[];
}

const requestCache = new Map<
  string,
  { promise: Promise<DictionaryData | null>; timestamp: number }
>();

const CACHE_TIME = 500; // 500ms

export const useDictionaryStore = defineStore('dictionary', {
  state: () => {
    return {
      data: JSON.stringify([...new Map<string, DictionaryData>()]),
    };
  },
  actions: {
    // Get dictionary info
    async getDictionary(name: string, isCache = true) {
      const mapData: Map<string, DictionaryData> = new Map(
        JSON.parse(this.data),
      );

      if (isCache && mapData.has(name)) {
        return mapData.get(name);
      } else {
        const cacheEntry = requestCache.get(name);

        // Check if we should use the cache and if the request is recent (within 500ms)
        if (cacheEntry && Date.now() - cacheEntry.timestamp < CACHE_TIME) {
          return cacheEntry.promise;
        }

        // If the request is not cached or isCache is false, perform the request
        const request = this.fetchDictionaryData(name);

        // Cache the new request promise with a timestamp
        requestCache.set(name, { promise: request, timestamp: Date.now() });

        // After 500ms, clear the request cache entry
        setTimeout(() => {
          requestCache.delete(name);
        }, 500);

        return request;
      }
    },
    async fetchDictionaryData(name: string) {
      const mapData: Map<string, DictionaryData> = new Map(
        JSON.parse(this.data),
      );
      const result = await GetDictionaryDetailByDictionaryName({ name });
      if (result.code === 0) {
        const dataConv = ref<DefaultOptionType[]>([]);

        for (let i = 0; i < result.data.data.length; i++) {
          const resultData: any = result.data.data[i];
          if (resultData !== undefined) {
            dataConv.value.push({
              label:
                resultData.trans === '' ? resultData.title : resultData.trans,
              value: resultData.value,
              status: resultData.status,
            });
          }
        }

        const dictData: DictionaryData = { data: dataConv.value };
        mapData.set(name, dictData);
        this.data = JSON.stringify([...mapData.entries()]);
        return dictData;
      } else {
        return null;
      }
    },

    // remove the dictionary in storage
    removeDictionary(name: string) {
      const mapData = new Map(JSON.parse(this.data));
      if (mapData.has(name)) {
        mapData.delete(name);
      }
      this.data = JSON.stringify([...mapData.entries()]);
    },

    // remove all the dictionary in storage
    clear() {
      const mapData = new Map();
      this.data = JSON.stringify([...mapData.entries()]);
    },
  },
  persist: {
    storage: sessionStorage,
  },
});
