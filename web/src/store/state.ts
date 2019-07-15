import { Error } from '@/store/basic/types'

export interface State {
    requestLoading: boolean;
    xError: Error;
}

const state: State = {
    requestLoading: false,
    xError: null
}

export default state
