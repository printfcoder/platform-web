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
                <el-col :span="8">
                    <cpu :cpuTimes="cpuTimes"></cpu>
                </el-col>
                <el-col :span="8">
                    <memory></memory>
                </el-col>
                <el-col :span="8">
                    <network></network>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="8">
                    <disk></disk>
                </el-col>
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
    import Memory from './Memory.vue';
    import Network from './Network.vue';

    import { Component, Watch } from 'vue-property-decorator';
    import { State, Action } from 'vuex-class';


    import { Error } from '@/store/basic/types';
    import { CPUTime, IpGroup } from '@/store/modules/os/types';

    const namespace: string = 'monitorOS';

    @Component({
        components: {
            'cpu': CPU,
            'memory': Memory,
            'network': Network,
            'disk': Disk,
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

        @State(state => state.monitorOS.xError)
        xError?: string;

        @Action('getIPGroups', { namespace })
        getIPGroups: any;

        @Action('getCPUTimes', { namespace })
        getCPUTimes: any;

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

        private now = new Date(new Date().setSeconds(new Date().getSeconds() - 10));

        changeIP() {
            if (!this.serverIP) {
                return;
            }

            let go = () => {
                let startTime = new Date(this.now);
                this.now = new Date(this.now.setSeconds(this.now.getSeconds() + 1));
                this.getCPUTimes({
                    ips: [this.serverIP],
                    startTime: startTime,
                    endTime: this.now,
                });
            };

            clearInterval(this.currentInterval);

            go();

            this.currentInterval = setInterval(go, 2000);
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
