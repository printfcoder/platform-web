export function mergeAddressAndPort (address: string, port: number) {
    if (port == null) {
        return address
    }

    return address + ':' + port
}
