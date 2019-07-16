import { MutationTree, ActionTree } from 'vuex'
import * as TYPES from '../../mutation-types'
import { Error } from '@/store/basic/types'
import { OSState, IpGroup } from './types'
import { getIPGroup, getCPUTimes } from '@/api/os'

const namespaced: boolean = true

const state: OSState = {
    cpuTimes: [],
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
    }
}

export default {
    namespaced,
    state,
    mutations,
    actions
}
