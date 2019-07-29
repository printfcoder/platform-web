<template>
    <el-container>
        <el-main style="padding-top: 0px; padding-left: 0px;">
            <el-col :span="5">
                <el-card>
                    <el-form>
                        <el-form-item label="CPU: ">
                            <el-select v-model="cpu" placeholder="CPU">
                                <el-option
                                        v-for="item in cpuOptions"
                                        :key="item.name"
                                        :label="item.name"
                                        :value="item.name">
                                </el-option>
                            </el-select>
                        </el-form-item>
                        <el-form-item label="System: ">
                            <span> {{ systemData.length>0 ?systemData[systemData.length-1].value[1] : ''}}%</span>
                        </el-form-item>
                        <el-form-item label="User: ">
                            <span> {{ userData.length>0 ?userData[userData.length-1].value[1] : ''}}%</span>
                        </el-form-item>
                        <el-form-item label="Idle: ">
                            <span> {{ idleData.length>0 ?idleData[idleData.length-1].value[1] : ''}}%</span>
                        </el-form-item>
                    </el-form>
                </el-card>
            </el-col>
            <el-col :span="19">
                <el-card>
                    <div>
                        <span style="float: right"> {{ lastUpdateTime && ($t('monitor.lastUpdated') + lastUpdateTime.toLocaleTimeString()) }}</span>
                        <div>
                            <v-chart
                                    ref="cpuChart"
                                    style="width: 100%; height: 186px"
                                    :options="cpuLoadLinearOptions"
                                    :autoresize="true"
                            />
                        </div>
                    </div>
                </el-card>
            </el-col>
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
        private cpu = 'cpu-total';
        private cpuOptions = [];

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
                left: '1%',
                right: '1%',
                bottom: '2%',
                containLabel: true,
            },
            toolbox: {
                feature: {},
            },
            xAxis: {
                type: 'category',
                splitLine: {
                    show: false,
                },
                boundaryGap: false,
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
            let cpuTimesRet = [];
            cpuTimes.forEach(v => {
                if (v.cpu == this.cpu) {
                    cpuTimesRet.push(v);
                }
            });

            return cpuTimesRet;
        }

        collectCPU(cpuTimes: CPUTime[]) {
            cpuTimes.forEach(item => {
                if (this.cpuOptions.indexOf(item.cpu) == -1) {
                    this.cpuOptions.push({ name: item.cpu });
                }
            });
        }

        @Watch('cpuTimes', { immediate: true, deep: true })
        asyncData(cpuTimes: CPUTime[]) {
            if (cpuTimes != null && cpuTimes.length > 0) {
                this.collectCPU(cpuTimes);

                let cpuTimesShow = this.groupByTime(cpuTimes);
                this.systemData = [];
                this.userData = [];
                this.idleData = [];

                console.log(cpuTimesShow);

                cpuTimesShow.forEach((ct: CPUTime) => {
                    let total = ct.system + ct.user + ct.idle;
                    let now = new Date();
                    let xAxisName = this.$xools.getTimeInterval(ct.time, now);

                    this.systemData.push({
                        name: xAxisName,
                        value: [xAxisName + 's', ((ct.system / total) * 100).toFixed(2)],
                    });

                    this.userData.push({
                        name: xAxisName,
                        value: [xAxisName + 's', ((ct.user / total) * 100).toFixed(2)],
                    });

                    this.idleData.push({
                        name: xAxisName,
                        value: [xAxisName + 's', ((ct.idle / total) * 100).toFixed(2)],
                    });
                });

                let chart = this.$refs['cpuChart'];
                chart && chart.chart && chart.chart.setOption({
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

<style scoped>

</style>;
