<template>
  <div ref="chartRef" :style="{ height, width }"></div>
</template>
<script lang="ts">
import {defineComponent, onMounted, ref, Ref, watch} from 'vue';

import {useECharts} from '@/hooks/web/useECharts';

import {basicProps} from './props';

export default defineComponent({
  props: basicProps,
  setup(props) {
    const chartRef = ref<HTMLDivElement | null>(null);
    const {setOptions} = useECharts(chartRef as Ref<HTMLDivElement>);
    let reload = () => {
      setOptions({
        tooltip: {
          trigger: 'axis',
          axisPointer: {
            lineStyle: {
              width: 1,
              color: '#019680',
            },
          },
        },
        legend: {
          data: ['注册', '登录']
        },
        xAxis: {
          type: 'category',
          boundaryGap: true,
          data: props.range,
          splitLine: {
            show: true,
            lineStyle: {
              width: 1,
              type: 'solid',
              color: 'rgba(226,226,226,0.5)',
            },
          },
          axisTick: {
            show: false,
          },
        },
        yAxis: [
          {
            type: 'value',
            splitNumber: 4,
            axisTick: {
              show: false,
            },
            splitArea: {
              show: true,
              areaStyle: {
                color: ['rgba(255,255,255,0.2)', 'rgba(226,226,226,0.2)'],
              },
            },
          },
        ],
        grid: {left: '3%', right: '3%', top: '10%', bottom: "3%", containLabel: true},
        series: [
          {
            name: "登录",
            data: props.login,
            type: 'bar',
            itemStyle: {
              color: '#5ab1ef',
            },
            barGap: "10%"
          },
          {
            name: "注册",
            data: props.register,
            type: 'bar',
            itemStyle: {
              color: 'rgb(77, 119, 255)',
            },
            barGap: "10%"
          },
        ],
      });
    }
    onMounted(() => {
      reload()
    });
    watch(() => props.login, (n) => {
      console.log("改变", n)
      reload()
    })
    return {chartRef};
  },
});
</script>
