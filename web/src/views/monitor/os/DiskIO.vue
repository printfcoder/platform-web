<template>
    <el-container>
        <el-main style="padding-top: 0px; padding-left: 0px">
            <el-col :span="5">
                <el-card>
                    <el-form style="height: 186px">
                        <el-form-item label="Reads In: ">
                            <span> {{ readsIn}}</span>
                        </el-form-item>
                        <el-form-item label="Writes Out: ">
                            <span> {{ writesOut}}</span>
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
            <el-col :span="5">
                <el-card>
                    <el-form style="height: 186px">
                        <el-form-item label="Data Read: ">
                            <span> {{ dataRead}}</span>
                        </el-form-item>
                        <el-form-item label="Data Written: ">
                            <span> {{ dataWritten }}</span>
                        </el-form-item>
                        <el-form-item label="Data Read/sec: ">
                            <span> {{ dataReadInSec }}</span>
                        </el-form-item>
                        <el-form-item label="Data Written/sec: ">
                            <span> {{ dataWrittenSec }}</span>
                        </el-form-item>
                    </el-form>
                </el-card>
            </el-col>
        </el-main>
    </el-container>
</template>
<script lang="ts">
    import { Component, Prop } from 'vue-property-decorator';
    import MVue from '@/basic/MVue';

    import { DiskIO } from '@/store/modules/os/types';

    @Component({
        components: {},
    })
    export default class Network extends MVue {
        private lastUpdateTime: Date = null;

        private currentModel = 'IO';

        @Prop()
        private diskIOStats: DiskIO[];

        private chartData = {
            columns: ['time', 'readsIn', 'writesOut'],
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

        private writesOut = '';
        private readsIn = '';
        private writesOutSec = '';
        private readsInSec = '';

        private dataRead = '';
        private dataWritten = '';
        private dataReadInSec = '';
        private dataWrittenSec = '';

        changeModel() {
            if (this.currentModel == 'IO') {
                this.currentModel = 'Data';
                this.chartData.columns = ['time', 'dataRead', 'dataWritten'];
                return;
            } else {
                this.currentModel = 'IO';
                this.chartData.columns = ['time', 'readsIn', 'writesOut'];
                return;
            }
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
