<template>
    <el-container>
        <el-main style="padding-top: 0px; padding-left: 0px">
            <el-card>
                <div>
                    <span style="float: right"> {{ lastUpdateTime && ($t('monitor.lastUpdated') + lastUpdateTime.toLocaleTimeString()) }}</span>
                    <div>
                        <v-chart
                                style="width: 100%"
                                :options="cpuLoadLinearOptions"
                                :autoresize="true"
                        />
                    </div>
                </div>
            </el-card>
        </el-main>
    </el-container>
</template>
<script lang="ts">
    import { Component } from 'vue-property-decorator';
    import MVue from '@/basic/MVue';

    // @ts-ignore
    import ECharts from 'vue-echarts';
    import 'echarts/lib/chart/line';
    import 'echarts/lib/component/polar';
    import 'echarts/theme/macarons';

    @Component({
        components: {
            'v-chart': ECharts,
        },
    })
    export default class Network extends MVue {
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
    }
</script>

<style scoped>

</style>
