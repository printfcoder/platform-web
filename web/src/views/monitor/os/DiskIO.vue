<template>
    <el-container>
        <el-main style="padding-top: 0px; padding-left: 0px">
            <el-col :span="5">
                <el-card>
                    <el-form style="height: 186px">
                        <el-form-item label="Reads In: ">
                            <span> {{ readsIn.toLocaleString()}}</span>
                        </el-form-item>
                        <el-form-item label="Writes Out: ">
                            <span> {{ writesOut.toLocaleString()}}</span>
                        </el-form-item>
                        <el-form-item label="Reads In/sec: ">
                            <span> {{ readsInSec}}</span>
                        </el-form-item>
                        <el-form-item label="Writes Out/sec: ">
                            <span> {{ writesOutSec}}</span>
                        </el-form-item>
                    </el-form>
                </el-card>
            </el-col>
            <el-col :span="14">
                <el-card>
                    <div>
                        <el-button class="modelBtn" @click="changeModel" type="text">{{currentModel}}</el-button>
                        <div>
                            <ve-line v-if="currentModel == 'IO' "
                                     :height="'186px'"
                                     :width="'100%'"
                                     :extend="chartExtend"
                                     :data="ioChartData"
                                     :settings="ioChartSettings"
                            ></ve-line>
                            <ve-line v-if="currentModel == 'Data' "
                                     :height="'186px'"
                                     :width="'100%'"
                                     :extend="chartExtend"
                                     :data="dataChartData"
                                     :settings="dataChartSettings"
                            ></ve-line>
                        </div>
                    </div>
                </el-card>
            </el-col>
            <el-col :span="5">
                <el-card>
                    <el-form style="height: 186px">
                        <el-form-item label="Data Read: ">
                            <span> {{ dataRead}} GB</span>
                        </el-form-item>
                        <el-form-item label="Data Written: ">
                            <span> {{ dataWritten }} GB</span>
                        </el-form-item>
                        <el-form-item label="Data Read/sec: ">
                            <span> {{ format(dataReadInSec) }}</span>
                        </el-form-item>
                        <el-form-item label="Data Written/sec: ">
                            <span> {{ format(dataWrittenSec) }}</span>
                        </el-form-item>
                    </el-form>
                </el-card>
            </el-col>
        </el-main>
    </el-container>
</template>
<script lang="ts">
    import { Component, Prop, Watch } from 'vue-property-decorator';
    import MVue from '@/basic/MVue';

    import { DiskIO } from '@/store/modules/os/types';

    @Component({
        components: {},
    })
    export default class Network extends MVue {
        private lastUpdateTime: Date = null;

        private currentModel = 'IO';
        private byteToGB = 1024 * 1024 * 1024;

        @Prop()
        private diskIOStats: DiskIO[];

        private dataChartData = {
            columns: ['time', 'dataRead', 'dataWritten'],
            rows: [],
        };

        private ioChartData = {
            columns: ['time', 'readsIn', 'writesOut'],
            rows: [],
        };

        private dataChartSettings = {
            stack: { 'diskIO': ['dataRead', 'dataWritten'] },
            area: true,
        };

        private ioChartSettings = {
            stack: { 'diskIO': ['readsIn', 'writesOut'] },
            area: true,
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

        private readsIn = 0;
        private writesOut = 0;
        private readsInSec = '';
        private writesOutSec = '';

        private dataRead = '';
        private dataWritten = '';
        private dataReadInSec = '';
        private dataWrittenSec = '';

        changeModel() {
            if (this.currentModel == 'IO') {
                this.currentModel = 'Data';
                return;
            } else {
                this.currentModel = 'IO';
                return;
            }
        }

        groupByTime(ios: DiskIO[]) {
            let result = [];
            ios.reduce(function(res: Map<Date, DiskIO>, ct: DiskIO) {
                if (!res.get(ct.time)) {
                    let dio = new DiskIO();
                    dio.time = ct.time;
                    dio.readCount = ct.readCount ? ct.readCount : 0;
                    dio.writeCount = ct.writeCount ? ct.writeCount : 0;
                    dio.readBytes = ct.readBytes ? ct.readBytes : 0;
                    dio.writeBytes = ct.writeBytes ? ct.writeBytes : 0;
                    res.set(ct.time, dio);
                    result.push(res.get(ct.time));
                } else {
                    res.get(ct.time).readCount += (ct.readCount ? ct.readCount : 0);
                    res.get(ct.time).writeCount += (ct.writeCount ? ct.writeCount : 0);
                    res.get(ct.time).readBytes += (ct.readBytes ? ct.readBytes : 0);
                    res.get(ct.time).writeBytes += (ct.writeBytes ? ct.writeBytes : 0);
                }

                return res;
            }, new Map<Date, DiskIO>());

            return result;
        }

        getDataInSec(ios: DiskIO[]) {
            let ret = [];

            if (ios.length > 2) {
                let last = ios[ios.length - 1];
                let lastButOne = ios[ios.length - 2];
                ret[0] = last.readCount - lastButOne.readCount;
                ret[1] = last.writeCount - lastButOne.writeCount;
                ret[2] = last.readBytes - lastButOne.readBytes;
                ret[3] = last.writeBytes - lastButOne.writeBytes;
            } else {
                ret[0] = 0;
                ret[1] = 0;
                ret[2] = 0;
                ret[3] = 0;
            }

            return ret;
        }

        format(input: number) {
            let output = '';
            if (input > this.byteToGB) {
                output = (input / this.byteToGB).toFixed(2) + ' GB';
            } else if (input > 1024 * 800) {
                output = (input / (1024 * 800)).toFixed(2) + ' MB';
            } else if (input > 1024) {
                output = (input / 1024).toFixed(2) + ' KB';
            } else {
                output = input + ' bytes';
            }

            return output;
        }

        @Watch('diskIOStats', { immediate: true, deep: true })
        asyncData(ios: DiskIO[]) {
            if (ios == null || ios.length == 0) {
                return;
            }

            let data = this.groupByTime(ios);
            let dataInSec = this.getDataInSec(data);

            this.ioChartData.rows = [];
            this.dataChartData.rows = [];

            this.readsInSec = dataInSec[0];
            this.writesOutSec = dataInSec[1];
            this.dataReadInSec = dataInSec[2];
            this.dataWrittenSec = dataInSec[3];

            let last = data[data.length - 1];
            this.readsIn = last.readCount;
            this.writesOut = last.writeCount;

            this.dataRead = (last.readBytes / this.byteToGB).toFixed(2);
            this.dataWritten = (last.writeBytes / this.byteToGB).toFixed(2);

            data.forEach((io: DiskIO, idx: number, me: DiskIO[]) => {
                let now = new Date();
                let xAxisName = this.$xools.getTimeInterval(io.time, now);

                let dataRead = 0;
                let dataWritten = 0;
                let readsIn = 0;
                let writesOut = 0;

                if (idx > 0) {
                    dataRead = me[idx].readBytes - me[idx - 1].readBytes;
                    dataWritten = me[idx].writeBytes - me[idx - 1].writeBytes;
                    readsIn = me[idx].readCount - me[idx - 1].readCount;
                    writesOut = me[idx].writeCount - me[idx - 1].writeCount;
                }

                this.dataChartData.rows.push({
                    'time': xAxisName,
                    'dataRead': dataRead,
                    'dataWritten': -dataWritten,
                });
                this.ioChartData.rows.push({
                    'time': xAxisName,
                    'readsIn': readsIn,
                    'writesOut': -writesOut,
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

    button.modelBtn {
        position: absolute;
        z-index: 1;
        padding-top: 2px;
    }

</style>
