import { Error } from '@/store/basic/types'
import { State } from '@/store/state'

export class IpGroup {
    groupName: string;
    ips: string[];
}

export class CPUTime {
    cpu: string;
    user: number;
    system: number;
    idle: number;
    time: Date;
    ip: string;
    nodeName: string;

    constructor (user: number, system: number, idle: number, time: Date) {
        this.user = user
        this.system = system
        this.idle = idle
        this.time = time
    }
}

export interface OSState extends State {
    ipGroups: IpGroup[];
    cpuTimes: CPUTime[];
    xError: Error
}
