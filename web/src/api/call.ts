import $axios from '@/utils/axios';

let $qs = require('qs');

export function getServiceDetails() {
    return $axios.get(`/api/v1/b/service-details`);
}

export function postServiceRequest(req: object) {
   // let postData = $qs.stringify(req);
    return $axios.post(`/api/v1/b/rpc`, req);
}
