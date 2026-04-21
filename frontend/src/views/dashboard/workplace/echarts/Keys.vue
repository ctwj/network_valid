<template>
  <div ref="chartRef" :style="{ height, width }"></div>
</template>
<script lang="ts">
import {defineComponent, onMounted, ref, Ref, watch} from 'vue';

import {useECharts} from '@/hooks/web/useECharts';
import echarts from "@/utils/lib/echarts";
import {basicProps} from './props.ts';

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
        legend: {
          data: ['创建', '激活']
        },
        grid: {left: '3%', right: '3%', top: '10%', bottom: "3%", containLabel: true},
        series: [
          {
            smooth: true,
            name: "创建",
            data: props.add,
            type: 'line',
            lineStyle: {
              width: 0
            },
            areaStyle: {
              opacity: 0.8,
              color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                {
                  offset: 0,
                  color: 'rgb(55, 162, 255)'
                },
                {
                  offset: 1,
                  color: 'rgb(116, 21, 219)'
                }
              ])
            },
          },
          {
            smooth: true,
            name: "激活",
            data: props.active,
            type: 'line',
            lineStyle: {
              width: 0
            },
            areaStyle: {
              opacity: 0.8,
              color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                {
                  offset: 0,
                  color: 'rgb(0, 221, 255)'
                },
                {
                  offset: 1,
                  color: 'rgb(77, 119, 255)'
                }
              ])
            },
          },
        ],
      });
    }
    onMounted(() => {
      reload()
    });
    watch(() => props.add, (n) => {
      console.log("改变", n)
      reload()
    })
    return {chartRef};
  },
});
</script>
