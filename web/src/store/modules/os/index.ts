import { MutationTree, ActionTree } from 'vuex'
import * as TYPES from '../../mutation-types'
import { Error } from '@/store/basic/types'
import { OSState, ipGroup } from './types'
import { getIPGroup } from '@/api/os'

const namespaced: boolean = true

const state: OSState = {
    ipGroups: [],
    requestLoading: false,
    xError: null
}

const mutations: MutationTree<any> = {
    [TYPES.SET_API_OS_IP_GROUPS] (state: OSState, { ipGroupsKV }): void {
        let ipGroups = []
        for (let k in ipGroupsKV) {
            let ipg: ipGroup = {
                groupName: k,
                ips: ipGroupsKV[k]
            }
            ipGroups.push(ipg)
        }
        state.ipGroups = ipGroups
        state.requestLoading = false
    },
    [TYPES.SET_FRAME_DATA_ERROR] (state: OSState, error: Error): void {
        state.xError = error
    }
}

const actions: ActionTree<any, any> = {
    async getIPGroups ({ commit }) {
        commit(TYPES.SET_FRAME_DATA_LOADING, true)

        const res: Ajax.AjaxResponse = await getIPGroup()
        commit(TYPES.SET_API_OS_IP_GROUPS, {
            ipGroupsKV: res.data
        })
    }
}

export default {
    namespaced,
    state,
    mutations,
    actions
}
