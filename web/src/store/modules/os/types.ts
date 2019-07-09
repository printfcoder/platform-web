import { Error } from '@/store/basic/types'
import { State } from '@/store/state'

export class ipGroup {
    groupName: string;
    ips: string[];
}

export interface OSState extends State {
    ipGroups: ipGroup[];
    xError: Error
}
