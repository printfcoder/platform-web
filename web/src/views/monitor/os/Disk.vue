<template>
    <el-container>
        <el-main style="padding-top: 0px; padding-left: 0px">
            <el-card>
                <div>
                    <span style="float: right"> {{ lastUpdateTime && ($t('monitor.lastUpdated') + lastUpdateTime.toLocaleTimeString()) }}</span>
                    <div>
                        <v-chart
                                ref="diskUsageChart"
                                style="width: 100%; height: 240px"
                                :options="diskUsageLinearOptions"
                                :autoresize="true"
                        />
                    </div>
                </div>
            </el-card>
        </el-main>
    </el-container>
</template>
<script lang="ts">
    import { Component, Prop, Watch } from 'vue-property-decorator';
    import MVue from '@/basic/MVue';

    // @ts-ignore
    import ECharts from 'vue-echarts';
    import 'echarts/lib/chart/line';
    import 'echarts/lib/component/polar';
    import 'echarts/theme/macarons';
    import { DiskUsage } from '@/store/modules/os/types';

    @Component({
        components: {
            'v-chart': ECharts,
        },
    })
    export default class Disk extends MVue {
        private lastUpdateTime: Date = null;

        @Prop()
        private diskUsageStats: DiskUsage[];

        private byteToGB = 1024 * 1024 * 1024;

        private usedData = [];
        private freeData = [];
        private totalData = [];

        private diskUsageLinearOptions = {
            title: {},
            tooltip: {
                trigger: 'axis',
                axisPointer: {
                    type: 'cross',
                    label: {
                        backgroundColor: '#6a7985',
                    },
                },
                formatter: function(params) {
                    let res = '';
                    for (let i = 0, l = params.length; i < l; i++) {
                        res += '<div style="color:' + params[i].color + '">' + params[i].seriesName + ' : ' + params[i].value[1] + 'G : ' + params[i].value[2] + '%\</div>';
                    }
                    return res;
                },
            },
            color: ['#FF4041', '#00AFF5', '#3B3B3B'],
            legend: {
                data: ['used', 'free', 'total'],
                x: 0,
            },
            toolbox: {},
            grid: {
                left: '0%',
                right: '1%',
                bottom: '2%',
                containLabel: true,
            },
            xAxis: [
                {
                    type: 'category',
                    splitLine: {
                        show: false,
                    },
                    boundaryGap: false,
                },
            ],
            yAxis: [
                {
                    type: 'value',
                    splitLine: {
                        show: true,
                    },
                    axisLine: { show: false },
                    axisLabel: { show: false },
                },
            ],
            series: [
                {
                    name: 'used',
                    type: 'line',
                    areaStyle: {},
                    data: [],
                },
                {
                    name: 'free',
                    type: 'line',
                    areaStyle: {},
                    data: [],
                },
                {
                    name: 'total',
                    type: 'line',
                    areaStyle: {},
                    data: [],
                },
            ],
        };

        @Watch('diskUsageStats', { immediate: true, deep: true })
        asyncData(diskUsages: DiskUsage[]) {
            if (diskUsages == null) {
                return;
            }

            this.freeData = [];
            this.usedData = [];
            this.totalData = [];

            diskUsages.forEach((mp: DiskUsage) => {
                let now = new Date();
                let xAxisName = this.$xools.getTimeInterval(mp.time, now);

                this.freeData.push({
                    name: xAxisName,
                    value: [xAxisName + 's', (mp.free / this.byteToGB).toFixed(1), (100 - mp.usedPercent).toFixed(2)],
                });

                this.usedData.push({
                    name: xAxisName,
                    value: [xAxisName + 's', (mp.used / this.byteToGB).toFixed(1), mp.usedPercent.toFixed(2)],
                });

                this.totalData.push({
                    name: xAxisName,
                    value: [xAxisName + 's', (mp.total / this.byteToGB).toFixed(1), 100],
                });
            });

            let chart = this.$refs['diskUsageChart'];
            chart && chart.chart && chart.chart.setOption({
                series: [
                    {
                        data: this.usedData,
                    },
                    {
                        data: this.freeData,
                    },
                    {
                        data: this.totalData,
                    },
                ],
            });
        }
    }
</script>

<style scoped>

</style>
