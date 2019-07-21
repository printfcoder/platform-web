<template>
    <el-container>
        <el-main style="padding-top: 0px; padding-left: 0px;">
            <el-card>
                <div>
                    <span style="float: right"> {{ lastUpdateTime && ($t('monitor.lastUpdated') + lastUpdateTime.toLocaleTimeString()) }}</span>
                    <div>
                        <v-chart
                                ref="cpuChart"
                                style="width: 100%;"
                                :options="cpuLoadLinearOptions"
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

    import { CPUTime } from '@/store/modules/os/types';

    @Component({
        components: {
            'v-chart': ECharts,
        },
    })
    export default class CPU extends MVue {
        private lastUpdateTime: Date = null;

        @Prop()
        private cpuTimes?: [];

        private systemData = [];
        private userData = [];
        private idleData = [];

        private cpuLoadLinearOptions = {
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
                data: ['System', 'User', 'Idle'],
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
                    name: 'System',
                    type: 'line',
                    data: this.systemData,
                },
                {
                    name: 'User',
                    type: 'line',
                    data: this.userData,
                },
                {
                    name: 'Idle',
                    type: 'line',
                    data: this.idleData,
                },
            ],
        };

        mounted() {

        }

        groupByTime(cpuTimes: CPUTime[]) {
            let result = [];
            cpuTimes.reduce(function(res: Map<Date, CPUTime>, ct: CPUTime) {
                if (!res.get(ct.time)) {
                    res.set(ct.time, new CPUTime(ct.user, ct.system, ct.idle, ct.time));
                    result.push(res.get(ct.time));
                }
                res.get(ct.time).user += ct.user;
                res.get(ct.time).system += ct.system;
                res.get(ct.time).idle += ct.idle;
                return res;
            }, new Map<Date, CPUTime>());

            return result;
        }


        @Watch('cpuTimes', { immediate: true, deep: true })
        asyncData(cpuTimes: CPUTime[]) {
            if (cpuTimes != null) {
                let cpuTimesShow = this.groupByTime(cpuTimes);
                cpuTimesShow.forEach((ct: CPUTime) => {
                    let total = ct.system + ct.user + ct.idle;
                    if (this.systemData.length > 10) {
                        this.systemData.shift();
                        this.userData.shift();
                        this.idleData.shift();
                    }

                    this.systemData.push({
                        name: ct.time,
                        value: [ct.time, ((ct.system / total) * 100).toFixed(2)],
                    });

                    this.userData.push({
                        name: ct.time,
                        value: [ct.time, ((ct.user / total) * 100).toFixed(2)],
                    });

                    this.idleData.push({
                        name: ct.time,
                        value: [ct.time, ((ct.idle / total) * 100).toFixed(2)],
                    });
                });
                let chart = this.$refs['cpuChart'];
               /* let systemLast = this.systemData[this.systemData.length - 1].value[1] + '%';
                let userLast = this.userData[this.userData.length - 1].value[1] + '%';
                let idleLast = this.idleData[this.idleData.length - 1].value[1] + '%'; */

                chart && chart.chart && chart.chart.setOption({
                    /* legend: {
                         data: [systemLast, userLast, idleLast],
                         x: 0,
                     },*/
                    series: [
                        {
                            //  name: systemLast,
                            data: this.systemData,
                        },
                        {
                            // name: userLast,
                            data: this.userData,
                        },
                        {
                            //   name: idleLast,
                            data: this.idleData,
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
