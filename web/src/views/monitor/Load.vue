<template>
    <el-container>
        <el-main style="padding-top: 0px; padding-left: 0px">
            <el-card>
                <div>
                    <span style="float: right"> {{ lastUpdateTime && ($t('monitor.lastUpdated') + lastUpdateTime.toLocaleTimeString()) }}</span>
                    <div>
                        <v-chart
                                ref="loadChart"
                                style="width: 100%; height: 230px"
                                :options="loadLinearOptions"
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

    import { LoadAvgStat } from '@/store/modules/os/types';

    @Component({
        components: {
            'v-chart': ECharts,
        },
    })
    export default class Load extends MVue {
        private lastUpdateTime: Date = null;

        @Prop()
        private loadAvgStats?: [];

        private load1Data = [];
        private load5Data = [];
        private load15Data = [];

        private loadLinearOptions = {
            title: {},
            tooltip: {
                trigger: 'axis',
                formatter: function(params) {
                    let res = '';
                    for (let i = 0, l = params.length; i < l; i++) {
                        res += '<div style="color:' + params[i].color + '">' + params[i].seriesName + ' : ' + params[i].value[1] + '%\</div>';
                    }
                    return res;
                },
            },
            color: ['#FF4041', '#00AFF5', '#3B3B3B'],
            legend: {
                data: ['Load1', 'Load5', 'Load15'],
                x: 0,
            },
            grid: {
                left: '3%',
                right: '4%',
                bottom: '3%',
                containLabel: true,
            },
            toolbox: {
                feature: {},
            },
            xAxis: {
                type: 'time',
                splitLine: {
                    show: false,
                },
            },
            yAxis: {
                type: 'value',
                boundaryGap: [0, '100%'],
                splitLine: {
                    show: true,
                },
                axisLine: { show: false },
                axisLabel: { show: false },
            },
            series: [
                {
                    name: 'Load1',
                    type: 'line',
                    data: this.load1Data,
                },
                {
                    name: 'Load5',
                    type: 'line',
                    data: this.load5Data,
                },
                {
                    name: 'Load15',
                    type: 'line',
                    data: this.load15Data,
                },
            ],
        };

        mounted() {

        }

        @Watch('loadAvgStats', { immediate: true, deep: true })
        asyncData(loadAvgStats: LoadAvgStat[]) {
            if (loadAvgStats != null) {
                loadAvgStats.forEach((ld: LoadAvgStat) => {
                    let total = ld.load1 + ld.load5 + ld.load15;
                    if (this.load1Data.length > 20) {
                        this.load1Data.shift();
                        this.load5Data.shift();
                        this.load15Data.shift();
                    }

                    this.load1Data.push({
                        name: ld.time,
                        value: [ld.time, ld.load1.toFixed(4)],
                    });

                    this.load5Data.push({
                        name: ld.time,
                        value: [ld.time, ld.load5.toFixed(4)],
                    });

                    this.load15Data.push({
                        name: ld.time,
                        value: [ld.time, ld.load15.toFixed(4)],
                    });
                });
                let chart = this.$refs['loadChart'];
                chart && chart.chart && chart.chart.setOption({
                    series: [
                        {
                            data: this.load1Data,
                        },
                        {
                            data: this.load5Data,
                        },
                        {
                            data: this.load15Data,
                        },
                    ],
                });
            }
        }
    }
</script>

< style;
scoped >

</

style

>;
