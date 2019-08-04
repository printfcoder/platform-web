import { Error } from '@/store/basic/types';
import { State } from '@/store/state';

export class IpGroup {
    groupName: string;
    ips: string[];
}

export class OSBase {
    public ip: string;
    public nodeName: string;
    public time: Date;
}

export class CPUTime extends OSBase {
    cpu: string;
    user: number;
    system: number;
    idle: number;

    constructor(user: number, system: number, idle: number, time: Date) {
        super();
        this.user = user;
        this.system = system;
        this.idle = idle;
        this.time = time;
    }
}

export class MemPercent extends OSBase {
    activeBytes: number;
    compressedBytes: number;
    inactiveBytes: number;
    wiredBytes: number;
    freeBytes: number;
    totalBytes: number;
}

export class DiskUsage extends OSBase {
    total: number;
    free: number;
    used: number;
    usedPercent: number;
}

export class DiskIO extends OSBase {
    readCount: number;
    writeCount: number;
    readBytes: number;
    writeBytes: number;
    readTime: number;
    writeTime: number;
    ioTime: number;
    name: string;
}

export class LoadAvgStat extends OSBase {
    load1: number;
    load5: number;
    load15: number;

    constructor(load1: number, load5: number, load15: number, time: Date) {
        super();
        this.load1 = load1;
        this.load5 = load5;
        this.load15 = load15;
        this.time = time;
    }
}

export interface OSState extends State {
    ipGroups: IpGroup[];
    cpuTimes: CPUTime[];
    diskUsageStats: DiskUsage[],
    diskIOStats: DiskIO[],
    memPercents: MemPercent[];
    loadAvgStats: LoadAvgStat[],
    xError: Error
}
