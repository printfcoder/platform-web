import $axios from '@/utils/axios';

export function getServices() {
    return $axios.get(`/v1/b/services`);
}

export function getService(name: string) {
    return $axios.get(`/v1/b/service/${name}`);
}

export function getWebServices() {
    return $axios.get(`/v1/b/web-services`);
}

export function getAPIGatewayServices() {
    return $axios.get(`/v1/b/api-gateway-services`);
}

export function getMicroServices() {
    return $axios.get(`/v1/b/micro-services`);
}
