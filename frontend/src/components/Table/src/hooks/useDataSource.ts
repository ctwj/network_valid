import {ref, ComputedRef, unref, computed, onMounted, watchEffect, watch} from 'vue';
import type {BasicTableProps} from '../types/table';
import type {PaginationProps} from '../types/pagination';
import {isBoolean, isFunction, isArray} from '@/utils/is';
import {APISETTING} from '../const';

export function useDataSource(
  propsRef: ComputedRef<BasicTableProps>,
  {getPaginationInfo, setPagination, setLoading, tableData},
  emit
) {
  const dataSourceRef = ref<Recordable[]>([]);

  watchEffect(() => {
    tableData.value = unref(dataSourceRef);
  });

  watch(
    () => unref(propsRef).dataSource,
    () => {
      const {dataSource}: any = unref(propsRef);
      dataSource && (dataSourceRef.value = dataSource);
    },
    {
      immediate: true,
    }
  );

  const getRowKey = computed(() => {
    const {rowKey}: any = unref(propsRef);
    return rowKey
      ? rowKey
      : () => {
        return 'key';
      };
  });

  const getDataSourceRef = computed(() => {
    const dataSource = unref(dataSourceRef);
    if (!dataSource || dataSource.length === 0) {
      return unref(dataSourceRef);
    }
    return unref(dataSourceRef);
  });

  async function fetch(opt?, nowPage = 0) {
    try {
      setLoading(true);
      let {request, pagination, beforeRequest, afterRequest}: any = unref(propsRef);
      if (!request) return;
      //组装分页信息
      const pageField = APISETTING.pageField;
      const sizeField = APISETTING.sizeField;
      const totalField = APISETTING.totalField;
      const listField = APISETTING.listField;

      let pageParams = {};
      const {page = 1, pageSize = 10} = unref(getPaginationInfo) as PaginationProps;

      if ((isBoolean(pagination) && !pagination) || isBoolean(getPaginationInfo)) {
        pageParams = {};
      } else {
        pageParams[pageField] = (opt && opt[pageField]) || page;
        pageParams[sizeField] = pageSize;
      }

      let params = {
        ...pageParams,
      };
      if (beforeRequest && isFunction(beforeRequest)) {
        // The params parameter can be modified by outsiders
        params = (await beforeRequest(params)) || params;
      }
      const res = await request(params);

      const resultTotal = res[totalField] || 0;
      let currentPage = res[pageField];
      if (nowPage !== 0) {
        currentPage = nowPage
      }
      // 如果数据异常，需获取正确的页码再次执行
      if (resultTotal) {
        if (page > resultTotal) {
          pagination = {
            [pageField]: resultTotal,
          }
          if (nowPage !== 0) {
            pagination['page'] = nowPage
          }
          setPagination(pagination);
          fetch(opt);
        }
      }
      let resultInfo = res[listField] ? res[listField] : [];
      if (afterRequest && isFunction(afterRequest)) {
        // can modify the data returned by the interface for processing
        resultInfo = (await afterRequest(resultInfo)) || resultInfo;
      }
      dataSourceRef.value = resultInfo;
      setPagination({
        [pageField]: currentPage,
        [totalField]: resultTotal,
      });
      if (opt && opt[pageField]) {
        setPagination({
          [pageField]: opt[pageField] || 1,
        });
      }
      emit('fetch-success', {
        items: unref(resultInfo),
        resultTotal,
      });
    } catch (error) {
      console.error(error);
      emit('fetch-error', error);
      dataSourceRef.value = [];
      // setPagination({
      //   pageCount: 0,
      // });
    } finally {
      setLoading(false);
    }
  }

  onMounted(() => {
    setTimeout(() => {
      fetch();
    }, 16);
  });

  function setTableData(values) {
    dataSourceRef.value = values;
  }

  function getDataSource(): any[] {
    return getDataSourceRef.value;
  }

  async function reload(opt?) {
    await fetch(opt);
  }
  async function fetchData(opt?) {
    await fetch(opt, 1);
  }

  return {
    fetch,
    getRowKey,
    getDataSourceRef,
    getDataSource,
    setTableData,
    reload,
    fetchData
  };
}
