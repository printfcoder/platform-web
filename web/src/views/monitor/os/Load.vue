<template>
    <el-container>
        <el-main style="padding-top: 0px; padding-left: 0px">
            <el-col :span="3">
                <el-card>
                    <el-form style="height: 150px">
                        <el-form-item label="Load1: ">
                            <span> {{ load1}}</span>
                        </el-form-item>
                        <el-form-item label="Load5: ">
                            <span> {{ load5}}</span>
                        </el-form-item>
                        <el-form-item label="Load15: ">
                            <span> {{ load15}}</span>
                        </el-form-item>
                    </el-form>
                </el-card>
            </el-col>
            <el-col :span="21">
                <el-card>
                    <div>
                        <ve-line
                                :height="'150px'"
                                :width="'100%'"
                                :extend="chartExtend"
                                :data="chartData"
                        ></ve-line>
                    </div>
                </el-card>
            </el-col>
        </el-main>
    </el-container>
</template>
<script lang="ts">
    import { Component, Prop, Watch } from 'vue-property-decorator';
    import MVue from '@/basic/MVue';

    import { LoadAvgStat } from '@/store/modules/os/types';

    @Component({
        components: {
        },
    })
    export default class Load extends MVue {
        private lastUpdateTime: Date = null;

        @Prop()
        private loadAvgStats?: [];

        private chartData = {
            columns: ['time', 'load1', 'load5', 'load15'],
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

        private load1 = '0';
        private load5 = '0';
        private load15 = '0';

        mounted() {
        }

        @Watch('loadAvgStats', { immediate: true, deep: true })
        asyncData(loadAvgStats: LoadAvgStat[]) {
            if (loadAvgStats != null) {
                this.chartData.rows = [];

                loadAvgStats.forEach((ld: LoadAvgStat) => {
                    let now = new Date();
                    let xAxisName = this.$xools.getTimeInterval(ld.time, now);

                    this.load1 = ld.load1.toFixed(4);
                    this.load5 = ld.load5.toFixed(4);
                    this.load15 = ld.load15.toFixed(4);

                    this.chartData.rows.push({
                        'time': xAxisName,
                        'load1': this.load1,
                        'load5': this.load5,
                        'load15': this.load15,
                    });
                });
            }
        }
    }
</script>
