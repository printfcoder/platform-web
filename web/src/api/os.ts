import $axios from '@/utils/axios'

export function getIPGroup () {
    return $axios.get(`/api/v1/os/ip-group`)
}
