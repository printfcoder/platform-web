import { Error, Service } from '@/store/basic/types'

export interface CallState {
    services: Service[];
    requestLoading: boolean;
    requestResult: Object;
    xError: Error;
}
