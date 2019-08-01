<template>
    <el-container>
        <el-main style="padding-top: 0px; padding-left: 0px">
            <el-col :span="5">
                <el-card>
                    <el-form style="height: 186px">
                        <el-form-item label="Act: ">
                            <span> {{ formatDBAndPercent(active) }}</span>
                        </el-form-item>
                        <el-form-item label="Ina: ">
                            <span> {{ formatDBAndPercent(inactive) }}</span>
                        </el-form-item>
                        <el-form-item label="Cpsed: ">
                            <span> {{ formatDBAndPercent(compressed) }}</span>
                        </el-form-item>
                        <el-form-item label="Wired: ">
                            <span> {{ formatDBAndPercent(wired) }}</span>
                        </el-form-item>
                        <el-form-item label="Free: ">
                            <span> {{ formatDBAndPercent(free) }}</span>
                        </el-form-item>
                    </el-form>
                </el-card>
            </el-col>
            <el-col :span="19">
                <el-card>
                    <div>
                        <ve-line
                                :height="'186px'"
                                :width="'100%'"
                                :extend="chartExtend"
                                :data="chartData"
                                :settings="chartSettings"></ve-line>
                    </div>
                </el-card>
            </el-col>
        </el-main>
    </el-container>
</template>
<script lang="ts">
    import { Component, Prop, Watch } from 'vue-property-decorator';
    import MVue from '@/basic/MVue';

    import { MemPercent } from '@/store/modules/os/types';

    @Component({
        components: {},
    })
    export default class Memory extends MVue {
        private lastUpdateTime: Date = null;

        @Prop()
        private memPercents?: MemPercent[];

        private byteToGB = 1024 * 1024 * 1024;

        private chartSettings = {
            stack: { 'mem': ['active', 'inactive', 'compressed', 'wired', 'free'] },
            area: true,
        };

        private chartData = {
            columns: ['time', 'active', 'inactive', 'compressed', 'wired', 'free'],
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

        private active = '0';
        private inactive = '0';
        private compressed = '0';
        private wired = '0';
        private free = '0';
        private total = '0';

        formatDBAndPercent(val) {
            if (this.total == '0') {
                return '';
            }

            return val + 'GB, ' + (val / Number.parseInt(this.total) * 100).toFixed(2) + '%';
        }

        @Watch('memPercents', { immediate: true, deep: true })
        asyncData(memPercents: MemPercent[]) {
            if (memPercents == null) {
                return;
            }

            this.chartData.rows = [];

            memPercents.forEach((mp: MemPercent) => {
                let now = new Date();
                let xAxisName = this.$xools.getTimeInterval(mp.time, now) + 's';

                this.active = (mp.activeBytes / this.byteToGB).toFixed(1);
                this.inactive = (mp.inactiveBytes / this.byteToGB).toFixed(1);
                this.compressed = (mp.compressedBytes / this.byteToGB).toFixed(1);
                this.wired = (mp.wiredBytes / this.byteToGB).toFixed(1);
                this.free = (mp.freeBytes / this.byteToGB).toFixed(1);
                this.total = (mp.totalBytes / this.byteToGB).toFixed(1);
                this.chartData.rows.push({
                    'time': xAxisName,
                    'active': this.active,
                    'inactive': this.inactive,
                    'compressed': this.compressed,
                    'wired': this.wired,
                    'free': this.free,
                });
            });
        }
    }
</script>

<style scoped>
    .el-form-item {
        margin-bottom: 0px;
    }

    .el-card .el-form {
        overflow: scroll;
    }
</style>
