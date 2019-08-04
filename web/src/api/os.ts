import $axios from '@/utils/axios'

export function getIPGroup () {
    return $axios.get(`/api/v1/os/ip-group`)
}

export function getCPUTimes (ips: string[], startTime: Date, endTime: Date) {
    return $axios.get(`/api/v1/os/cpu/times?ips=` + ips.join() + '&startTime=' + startTime.toJSON() + '&endTime=' + endTime.toJSON())
}

export function getMemPercents (ips: string[], startTime: Date, endTime: Date) {
    return $axios.get(`/api/v1/os/mem/percents?ips=` + ips.join() + '&startTime=' + startTime.toJSON() + '&endTime=' + endTime.toJSON())
}

export function getLoadAvgStat (ips: string[], startTime: Date, endTime: Date) {
    return $axios.get(`/api/v1/os/load/avg-stat?ips=` + ips.join() + '&startTime=' + startTime.toJSON() + '&endTime=' + endTime.toJSON())
}

export function getDiskUsageStats (ips: string[], startTime: Date, endTime: Date) {
    return $axios.get(`/api/v1/os/disk/usage-stat?ips=` + ips.join() + '&startTime=' + startTime.toJSON() + '&endTime=' + endTime.toJSON())
}

export function getDiskIOStats (ips: string[], startTime: Date, endTime: Date) {
    return $axios.get(`/api/v1/os/disk/io-stat?ips=` + ips.join() + '&startTime=' + startTime.toJSON() + '&endTime=' + endTime.toJSON())
}
