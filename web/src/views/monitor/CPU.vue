<template>
    <el-container>
        <el-main style="padding-top: 0px; padding-left: 0px">
            <el-card>
                <div>
                    <span style="float: right"> {{ lastUpdateTime && ($t('monitor.lastUpdated') + lastUpdateTime.toLocaleTimeString()) }}</span>
                    <div>
                        <v-chart
                                ref="cpuChart"
                                style="width: 100%"
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
                    show: false,
                },
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
                cpuTimes = this.groupByTime(cpuTimes);
                cpuTimes.forEach((ct: CPUTime) => {
                    this.systemData.push({ name: 'system', value: ct.system });
                    this.userData.push({ name: 'user', value: ct.user });
                    this.idleData.push({ name: 'idle', value: ct.idle });
                });
            }
        }
    }
</script>

<style scoped>

</style>
