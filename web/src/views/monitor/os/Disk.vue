<template>
    <el-container>
        <el-main style="padding-top: 0px; padding-left: 0px">
            <el-col :span="7">
                <el-card>
                    <el-form style="height: 186px">
                        <el-form-item label="Used: ">
                            <span> {{ used }} GB</span>
                        </el-form-item>
                        <el-form-item label="Free: ">
                            <span> {{ free }} GB</span>
                        </el-form-item>
                        <el-form-item label="Total: ">
                            <span> {{ total }} GB</span>
                        </el-form-item>
                    </el-form>
                </el-card>
            </el-col>
            <el-col :span="17">
                <el-card>
                    <div>
                        <div>
                            <div>
                                <ve-line
                                        :height="'186px'"
                                        :width="'100%'"
                                        :extend="chartExtend"
                                        :data="chartData"
                                        :settings="chartSettings"></ve-line>
                            </div>
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

    import { DiskUsage } from '@/store/modules/os/types';

    @Component({
        components: {},
    })
    export default class Disk extends MVue {
        private lastUpdateTime: Date = null;

        @Prop()
        private diskUsageStats: DiskUsage[];

        private byteToGB = 1024 * 1024 * 1024;

        private chartSettings = {
            stack: { 'disk': ['used', 'free'] },
            area: true,
        };

        private chartData = {
            columns: ['time', 'used', 'free'],
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

        private used = '0';
        private free = '0';
        private total = '0';

        @Watch('diskUsageStats', { immediate: true, deep: true })
        asyncData(diskUsages: DiskUsage[]) {
            if (diskUsages == null) {
                return;
            }

            this.chartData.rows = [];

            diskUsages.forEach((du: DiskUsage) => {
                let now = new Date();
                let xAxisName = this.$xools.getTimeInterval(du.time, now);

                this.used = (du.used / this.byteToGB).toFixed(1);
                this.free = (du.free / this.byteToGB).toFixed(1);
                this.total = (du.total / this.byteToGB).toFixed(1);

                this.chartData.rows.push({
                    'time': xAxisName,
                    'used': this.used,
                    'free': this.free,
                });
            });
        }
    }
</script>

<style scoped>
    .el-card .el-form {
        overflow: scroll;
    }
</style>
