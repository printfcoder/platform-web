import $axios from '@/utils/axios';

export function getServiceDetails() {
  return $axios.get(`/api/v1/b/service-details`);
}

export function postServiceRequest(req: object) {
  return $axios.post(`/api/v1/b/rpc`, req);
}
