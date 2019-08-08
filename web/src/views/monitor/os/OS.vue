<template>
    <el-container v-loading="pageLoading">
        <el-header>
            <el-card :height="60" :body-style="{ padding: '10px 10px 10px 20px'}">
                <el-row>
                    <el-col :span="4">
                        <el-select v-model="serverGroup" :placeholder="$t('base.group')" @change="changeGroup"
                                   clearable
                                   style="width:90%">
                            <el-option
                                    v-for="(item, idx) in ipGroups"
                                    :key="idx"
                                    :label="item.groupName"
                                    :value="item.groupName">
                            </el-option>
                        </el-select>
                    </el-col>
                    <el-col :span="3">
                        <el-select v-model="serverIP" :placeholder="$t('base.ip')" @change="changeIP"
                                   clearable>
                            <el-option
                                    v-for="(item, idx) in serverIPs"
                                    :key="idx"
                                    :label="item"
                                    :value="item">
                            </el-option>
                        </el-select>
                    </el-col>
                    <el-col :span="3" style="float: right;">
                        <el-button style="float: right;" :disabled="!(this.serverGroup && this.serverIP)"
                                   @click="changeIP">{{$t('base.refresh')}}
                        </el-button>
                    </el-col>
                </el-row>
            </el-card>
        </el-header>
        <el-main>
            <el-row>
                <el-col :span="24">
                    <load :loadAvgStats="loadAvgStats"></load>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="10">
                    <cpu :cpuTimes="cpuTimes"></cpu>
                </el-col>
                <el-col :span="14">
                    <memory :memPercents="memPercents"></memory>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="8">
                    <disk :diskUsageStats="diskUsageStats"></disk>
                </el-col>
                <el-col :span="16">
                    <disk-io :diskIOStats="diskIOStats"></disk-io>
                </el-col>
            </el-row>
            <el-row>
                <network></network>
            </el-row>
        </el-main>
    </el-container>
</template>

<style scoped>
    .el-header {
        padding: 0 20px 0 0;
        height: 70px !important;
    }

    .el-container > .el-main {
        padding: 0;
    }
</style>

<script lang="ts">
    import MVue from '@/basic/MVue';
    import CPU from './CPU.vue';
    import Disk from './Disk.vue';
    import DiskIO from './DiskIO.vue';
    import Load from './Load.vue';
    import Memory from './Memory.vue';
    import Network from './Network.vue';

    import { Component, Watch } from 'vue-property-decorator';
    import { State, Action } from 'vuex-class';


    import { Error } from '@/store/basic/types';
    import { CPUTime, DiskUsage, IpGroup, LoadAvgStat, MemPercent } from '@/store/modules/os/types';

    const namespace: string = 'monitorOS';

    @Component({
        components: {
            'cpu': CPU,
            'memory': Memory,
            'network': Network,
            'disk': Disk,
            'disk-io': DiskIO,
            'load': Load,
        },
    })
    export default class OS extends MVue {
        private serverGroup: string = '';
        private currentInterval: any;
        private serverIP: string = '';
        private serverIPs: string[] = [];
        private lastUpdateTime: Date = null;

        @State(state => state.monitorOS.loaded)
        loaded ?: boolean;

        @State(state => state.monitorOS.pageLoading)
        pageLoading?: boolean;

        @State(state => state.monitorOS.ipGroups)
        ipGroups?: IpGroup[];

        @State(state => state.monitorOS.cpuTimes)
        cpuTimes?: CPUTime[];

        @State(state => state.monitorOS.memPercents)
        memPercents?: MemPercent[];

        @State(state => state.monitorOS.loadAvgStats)
        loadAvgStats?: LoadAvgStat[];

        @State(state => state.monitorOS.diskUsageStats)
        diskUsageStats?: DiskUsage[];

        @State(state => state.monitorOS.diskIOStats)
        diskIOStats?: DiskIO[];

        @State(state => state.monitorOS.xError)
        xError?: string;

        @Action('getIPGroups', { namespace })
        getIPGroups: any;

        @Action('getCPUTimes', { namespace })
        getCPUTimes: any;

        @Action('getMemPercents', { namespace })
        getMemPercents: any;

        @Action('getLoadAvgStat', { namespace })
        getLoadAvgStat: any;

        @Action('getDiskUsageStats', { namespace })
        getDiskUsageStats: any;

        @Action('getDiskIOStats', { namespace })
        getDiskIOStats: any;

        created() {
            if (!this.loaded) {
                this.getIPGroups();
            }
        }

        mounted() {

        }

        changeGroup(name: string) {
            this.serverIPs = [];
            this.serverIP = '';

            if (name) {
                for (let i = 0; i < this.ipGroups.length; i++) {
                    if (this.ipGroups[i].groupName === name) {
                        this.serverIPs = this.ipGroups[i].ips;
                        break;
                    }
                }
            }
        }

        changeIP() {
            if (!this.serverIP) {
                clearInterval(this.currentInterval);
                return;
            }

            let go = () => {
                let startTime = new Date(new Date().setSeconds(new Date().getSeconds() - 15));
                let now = new Date();
                this.getCPUTimes({
                    ips: [this.serverIP],
                    startTime: startTime,
                    endTime: now,
                });
                this.getMemPercents({
                    ips: [this.serverIP],
                    startTime: startTime,
                    endTime: now,
                });
                this.getLoadAvgStat({
                    ips: [this.serverIP],
                    startTime: startTime,
                    endTime: now,
                });
                this.getDiskUsageStats({
                    ips: [this.serverIP],
                    startTime: startTime,
                    endTime: now,
                });
                this.getDiskIOStats({
                    ips: [this.serverIP],
                    startTime: startTime,
                    endTime: now,
                });
            };

            clearInterval(this.currentInterval);

            go();

            this.currentInterval = setInterval(go, 1000);
        }

        @Watch('xError')
        catchError(xError: Error) {
            if (xError) {
                clearInterval(this.currentInterval);
                // @ts-ignore
                this.$message.error('Oops, ' + xError.detail || xError);
            }
        }
    }
</script>
