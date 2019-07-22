<template>
    <el-container>
        <el-main style="padding-top: 0px; padding-left: 0px">
            <el-card>
                <div>
                    <span style="float: right"> {{ lastUpdateTime && ($t('monitor.lastUpdated') + lastUpdateTime.toLocaleTimeString()) }}</span>
                    <div>
                        <v-chart
                                ref="memChart"
                                style="width: 100%; height: 240px"
                                :options="memLinearOptions"
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

    import { MemPercent } from '@/store/modules/os/types';

    @Component({
        components: {
            'v-chart': ECharts,
        },
    })
    export default class Memory extends MVue {
        private lastUpdateTime: Date = null;

        @Prop()
        private memPercents?: MemPercent[];

        private byteToGB = 1024 * 1024 * 1024;

        private activeData = [];
        private compressedData = [];
        private inactiveData = [];
        private wiredData = [];
        private freeData = [];
        private totalData = [];

        private memLinearOptions = {
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
            legend: {
                data: ['active', 'inactive', 'compressed', 'wired', 'free'],
                x: 0,
            },
            toolbox: {},
            grid: {
                left: '1%',
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
                    name: 'active',
                    type: 'line',
                    stack: '总量',
                    areaStyle: {},
                    data: [],
                },
                {
                    name: 'inactive',
                    type: 'line',
                    stack: '总量',
                    areaStyle: {},
                    data: [],
                },
                {
                    name: 'compressed',
                    type: 'line',
                    stack: '总量',
                    areaStyle: {},
                    data: [],
                },
                {
                    name: 'wired',
                    type: 'line',
                    stack: '总量',
                    areaStyle: { normal: {} },
                    data: [],
                },
                {
                    name: 'free',
                    type: 'line',
                    stack: '总量',
                    areaStyle: { normal: {} },
                    data: [],
                },
            ],
        };

        @Watch('memPercents', { immediate: true, deep: true })
        asyncData(memPercents: MemPercent[]) {
            if (memPercents == null) {
                return;
            }

            memPercents.forEach((mp: MemPercent) => {
                if (this.activeData.length > 10) {
                    this.compressedData.shift();
                    this.inactiveData.shift();
                    this.wiredData.shift();
                    this.freeData.shift();
                    this.totalData.shift();
                }

                let now = new Date();
                let xAxisName = this.$xools.getTimeInterval(mp.time, now);

                this.activeData.push({
                    name: xAxisName,
                    value: [xAxisName + 's', (mp.activeBytes / this.byteToGB).toFixed(1), ((mp.activeBytes / mp.totalBytes) * 100).toFixed(2)],
                });

                this.compressedData.push({
                    name: xAxisName,
                    value: [xAxisName + 's', (mp.compressedBytes / this.byteToGB).toFixed(1), ((mp.compressedBytes / mp.totalBytes) * 100).toFixed(2)],
                });

                this.inactiveData.push({
                    name: xAxisName,
                    value: [xAxisName + 's', (mp.inactiveBytes / this.byteToGB).toFixed(1), ((mp.inactiveBytes / mp.totalBytes) * 100).toFixed(2)],
                });

                this.wiredData.push({
                    name: xAxisName,
                    value: [xAxisName + 's', (mp.wiredBytes / this.byteToGB).toFixed(1), ((mp.wiredBytes / mp.totalBytes) * 100).toFixed(2)],
                });

                this.freeData.push({
                    name: xAxisName,
                    value: [xAxisName + 's', (mp.freeBytes / this.byteToGB).toFixed(1), ((mp.freeBytes / mp.totalBytes) * 100).toFixed(2)],
                });
            });

            let chart = this.$refs['memChart'];
            chart && chart.chart && chart.chart.setOption({
                series: [
                    {
                        data: this.activeData,
                    },
                    {
                        data: this.compressedData,
                    },
                    {
                        data: this.inactiveData,
                    },
                    {
                        data: this.wiredData,
                    },
                    {
                        data: this.freeData,
                    },
                ],
            });
        }
    }
</script>

<style scoped>

</style>
