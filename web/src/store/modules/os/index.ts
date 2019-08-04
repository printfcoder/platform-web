import { MutationTree, ActionTree } from 'vuex'
import * as TYPES from '../../mutation-types'
import { Error } from '@/store/basic/types'
import { OSState, IpGroup } from './types'
import { getIPGroup, getCPUTimes, getMemPercents, getLoadAvgStat, getDiskUsageStats, getDiskIOStats } from '@/api/os'

const namespaced: boolean = true

const state: OSState = {
    cpuTimes: [],
    diskUsageStats: [],
    diskIOStats: [],
    memPercents: [],
    loadAvgStats: [],
    ipGroups: [],
    requestLoading: false,
    xError: null
}

const mutations: MutationTree<any> = {
    [TYPES.SET_API_OS_IP_GROUPS] (state: OSState, { ipGroupsKV }): void {
        let ipGroups = []
        for (let k in ipGroupsKV) {
            let ipg: IpGroup = {
                groupName: k,
                ips: ipGroupsKV[k]
            }
            ipGroups.push(ipg)
        }
        state.ipGroups = ipGroups
        state.requestLoading = false
    },
    [TYPES.SET_API_OS_CPU_TIMES] (state: OSState, { cpuTimes }): void {
        state.cpuTimes = cpuTimes
        state.requestLoading = false
    },
    [TYPES.SET_API_OS_MEM_PERCENTS] (state: OSState, { memPercents }): void {
        state.memPercents = memPercents
        state.requestLoading = false
    },
    [TYPES.SET_API_OS_LOAD_AVG_STATS] (state: OSState, { loadAvgStats }): void {
        state.loadAvgStats = loadAvgStats
        state.requestLoading = false
    },
    [TYPES.SET_API_OS_DISK_USAGE_STATS] (state: OSState, { diskUsageStats }): void {
        state.diskUsageStats = diskUsageStats
        state.requestLoading = false
    },
    [TYPES.SET_API_OS_DISK_IO_STATS] (state: OSState, { diskIOStats }): void {
        state.diskIOStats = diskIOStats
        state.requestLoading = false
    },
    [TYPES.SET_FRAME_DATA_ERROR] (state: OSState, error: Error): void {
        state.xError = error
    }
}

const actions: ActionTree<any, any> = {
    async getIPGroups ({ commit }) {
        const res: Ajax.AjaxResponse = await getIPGroup()
        commit(TYPES.SET_API_OS_IP_GROUPS, {
            ipGroupsKV: res.data
        })
    },

    async getCPUTimes ({ commit }, { ips, startTime, endTime }) {
        const res: Ajax.AjaxResponse = await getCPUTimes(ips, startTime, endTime)
        commit(TYPES.SET_API_OS_CPU_TIMES, {
            cpuTimes: res.data
        })
    },

    async getMemPercents ({ commit }, { ips, startTime, endTime }) {
        const res: Ajax.AjaxResponse = await getMemPercents(ips, startTime, endTime)
        commit(TYPES.SET_API_OS_MEM_PERCENTS, {
            memPercents: res.data
        })
    },

    async getLoadAvgStat ({ commit }, { ips, startTime, endTime }) {
        const res: Ajax.AjaxResponse = await getLoadAvgStat(ips, startTime, endTime)
        commit(TYPES.SET_API_OS_LOAD_AVG_STATS, {
            loadAvgStats: res.data
        })
    },

    async getDiskUsageStats ({ commit }, { ips, startTime, endTime }) {
        const res: Ajax.AjaxResponse = await getDiskUsageStats(ips, startTime, endTime)
        commit(TYPES.SET_API_OS_DISK_USAGE_STATS, {
            diskUsageStats: res.data
        })
    },

    async getDiskIOStats ({ commit }, { ips, startTime, endTime }) {
        const res: Ajax.AjaxResponse = await getDiskIOStats(ips, startTime, endTime)
        commit(TYPES.SET_API_OS_DISK_IO_STATS, {
            diskIOStats: res.data
        })
    }
}

export default {
    namespaced,
    state,
    mutations,
    actions
}
