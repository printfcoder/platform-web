<template>
    <el-container>
        <el-main style="padding-top: 0px; padding-left: 0px;">
            <el-col :span="7">
                <el-card>
                    <el-form style="height: 186px">
                        <el-form-item label="CPU: ">
                            <el-select :size="'small'" v-model="cpu" placeholder="CPU" style="width: 70%"
                                       @change="changeCPU">
                                <el-option
                                        v-for="item in cpuOptions"
                                        :key="item"
                                        :label="item"
                                        :value="item">
                                </el-option>
                            </el-select>
                        </el-form-item>
                        <el-form-item label="System: ">
                            <span> {{ system }}%</span>
                        </el-form-item>
                        <el-form-item label="User: ">
                            <span> {{ user }}%</span>
                        </el-form-item>
                        <el-form-item label="Idle: ">
                            <span> {{ idle }}%</span>
                        </el-form-item>
                    </el-form>
                </el-card>
            </el-col>
            <el-col :span="17">
                <el-card>
                    <div>
                        <div>
                            <ve-line
                                    :height="'186px'"
                                    :width="'100%'"
                                    :extend="chartExtend"
                                    :data="chartData"
                            ></ve-line>
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

    import { CPUTime } from '@/store/modules/os/types';

    @Component({
        components: {},
    })
    export default class CPU extends MVue {
        private lastUpdateTime: Date = null;

        @Prop()
        private cpuTimes?: [];

        private chartData = {
            columns: ['time', 'system', 'user', 'idle'],
            rows: [],
        };

        private chartExtend = {
            series: {
                showSymbol: false,
            },
            grid: {
                left: '0%',
                right: '0%',
                top: '20%',
                bottom: '2%',
                containLabel: true,
            },
        };

        private system = '';
        private user = '';
        private idle = '';
        private cpu = 'cpu-total';
        private cpuOptions = [this.cpu];

        mounted() {

        }

        changeCPU(v) {
            this.cpu = v;
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
                    this.cpuOptions.push(item.cpu);
                }
            });
        }

        @Watch('cpuTimes', { immediate: true, deep: true })
        asyncData(cpuTimes: CPUTime[]) {
            if (cpuTimes != null && cpuTimes.length > 0) {
                this.collectCPU(cpuTimes);

                let cpuTimesShow = this.groupByTime(cpuTimes);
                this.chartData.rows = [];

                cpuTimesShow.forEach((ct: CPUTime) => {
                    let total = ct.system + ct.user + ct.idle;
                    let now = new Date();
                    let xAxisName = this.$xools.getTimeInterval(ct.time, now);

                    this.system = ((ct.system / total) * 100).toFixed(2);
                    this.user = ((ct.user / total) * 100).toFixed(2);
                    this.idle = ((ct.idle / total) * 100).toFixed(2);

                    this.chartData.rows.push({
                        'time': xAxisName,
                        'system': this.system,
                        'user': this.user,
                        'idle': this.idle,
                    });
                });
            }
        }
    }
</script>

<style scoped>
    .el-form-item {
        margin-bottom: 10px;
    }

    .el-card .el-form {
        overflow: scroll;
    }
</style>
