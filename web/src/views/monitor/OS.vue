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

        <cpu></cpu>
    </el-container>
</template>

<style scoped>
    .echarts {
        width: 100%;
        height: 100%;
    }

    .rowName {
        font-weight: 400;
        color: #1f2f3d;
    }

    .el-header {
        padding: 0 20px 0 0;
        height: 70px !important;
    }

    .el-card__body {
        padding: 10px 10px 10px 20px !important;
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



    import { Service, Node, Error } from '@/store/basic/types';
    import { Stats } from '@/store/modules/stats/types';
    import { ipGroup } from '@/store/modules/os/types';

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

        private cpuItems = [
            {
                name: 'system',
                key: 'system',
                formatter: (date: number) => {
                    return new Date(date * 1000).toLocaleString();
                },
                value: '',
            },
            {
                name: 'user',
                key: 'user',
                value: '',
                formatter: (date: number) => {
                    // @ts-ignore
                    return this.$xools.secondsToHHMMSS(((new Date() - date * 1000) / 1000).toFixed(0));
                },
            },
            {
                name: 'idle',
                key: 'idle',
                value: '',
                formatter: (memory: number) => {
                    return memory;
                },
            },
            {
                name: 'threads',
                key: 'threads',
                value: '',
                formatter: (threads: number) => {
                    return threads;
                },
            },
            {
                name: 'processes',
                key: 'processes',
                value: '',
                formatter: (gc: number) => {
                    return gc;
                },
            },
        ];

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
                type: 'category',
                boundaryGap: false,
                data: [],
            },
            yAxis: {
                type: 'value',
            },
            series: [
                {
                    name: 'System',
                    type: 'line',
                    data: [],
                },
                {
                    name: 'User',
                    type: 'line',
                    data: [],
                },
                {
                    name: 'Idle',
                    type: 'line',
                    data: [],
                },
            ],
        };

        private cpuData = {
            'system': 0,
            'user': 0,
            'idle': 0,
            'threads': 0,
            'processes': 0,
        };

        @State(state => state.monitorOS.loaded)
        loaded ?: boolean;

        @State(state => state.monitorOS.pageLoading)
        pageLoading?: boolean;

        @State(state => state.monitorOS.ipGroups)
        ipGroups?: ipGroup[];

        @State(state => state.monitorOS.xError)
        xError?: string;

        @Action('getIPGroups', { namespace })
        getIPGroups: any;

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
                return;
            }

            let go = () => {
                // this.getStats({name: this.serviceName, address: this.serviceNode})
            };

            go();

            this.currentInterval = setInterval(go, 5000);
        }

        clean() {

        }

        @Watch('xError')
        catchError(xError: Error) {
            if (xError) {
                clearInterval(this.currentInterval);
                // @ts-ignore
                this.$message.error('Oops, ' + xError.detail || xError);
            }
        }

        @Watch('currentNodeStats', { immediate: true, deep: true })
        asyncData(data: Stats) {

        }
    }
</script>
