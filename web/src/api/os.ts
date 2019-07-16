import $axios from '@/utils/axios'

export function getIPGroup () {
    return $axios.get(`/api/v1/os/ip-group`)
}

export function getCPUTimes (ips: string[], startTime: Date, endTime: Date) {
    return $axios.get(`/api/v1/os/cpu/times?ips=` + ips.join() + '&startTime=' + startTime.toJSON() + '&endTime=' + endTime.toJSON())
}
