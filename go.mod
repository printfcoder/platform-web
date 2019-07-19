module github.com/micro-in-cn/platform-web

go 1.12

replace github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.0-20181115231424-8e868ca12c0f

replace github.com/golang/lint => github.com/golang/lint v0.0.0-20190227174305-8f45f776aaf1

require (
	contrib.go.opencensus.io/exporter/ocagent v0.5.0 // indirect
	github.com/Azure/azure-sdk-for-go v29.0.0+incompatible // indirect
	github.com/Azure/go-autorest v12.1.0+incompatible // indirect
	github.com/DataDog/dd-trace-go v1.14.0 // indirect
	github.com/DataDog/zstd v1.4.0 // indirect
	github.com/Jeffail/gabs v1.4.0 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/SAP/go-hdb v0.14.1 // indirect
	github.com/Shopify/sarama v1.22.1 // indirect
	github.com/StackExchange/wmi v0.0.0-20181212234831-e0a55b97c705 // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v0.0.0-20190522081930-582d16a078d0 // indirect
	github.com/aliyun/aliyun-oss-go-sdk v1.9.6 // indirect
	github.com/asaskevich/govalidator v0.0.0-20190424111038-f61b66f89f4a // indirect
	github.com/aws/aws-sdk-go v1.19.35 // indirect
	github.com/coredns/coredns v1.5.0 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20190321100706-95778dfbb74e // indirect
	github.com/denisenkom/go-mssqldb v0.0.0-20190515213511-eb9f6a1743f3 // indirect
	github.com/denverdino/aliyungo v0.0.0-20190410085603-611ead8a6fed // indirect
	github.com/digitalocean/godo v1.15.0 // indirect
	github.com/dnstap/golang-dnstap v0.0.0-20190521061535-1a0dab85b926 // indirect
	github.com/docker/spdystream v0.0.0-20181023171402-6480d4af844c // indirect
	github.com/elazarl/goproxy v0.0.0-20190421051319-9d40249d3c2f // indirect
	github.com/emicklei/go-restful v2.9.5+incompatible // indirect
	github.com/envoyproxy/go-control-plane v0.8.0 // indirect
	github.com/evanphx/json-patch v4.2.0+incompatible // indirect
	github.com/farsightsec/golang-framestream v0.0.0-20190425193708-fa4b164d59b8 // indirect
	github.com/fullsailor/pkcs7 v0.0.0-20190404230743-d7302db945fa // indirect
	github.com/gammazero/workerpool v0.0.0-20190521015540-3b91a70bc0a1 // indirect
	github.com/go-ldap/ldap v3.0.3+incompatible // indirect
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/go-openapi/jsonpointer v0.19.0 // indirect
	github.com/go-openapi/jsonreference v0.19.0 // indirect
	github.com/go-openapi/spec v0.19.0 // indirect
	github.com/go-openapi/swag v0.19.0 // indirect
	github.com/go-stomp/stomp v2.0.3+incompatible // indirect
	github.com/gocql/gocql v0.0.0-20190423091413-b99afaf3b163 // indirect
	github.com/gogo/googleapis v1.2.0 // indirect
	github.com/golang/protobuf v1.3.2
	github.com/google/gofuzz v1.0.0 // indirect
	github.com/google/uuid v1.1.1
	github.com/gophercloud/gophercloud v0.0.0-20190520235722-e87e5f90e7e6 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20190430165422-3e4dfb77656c // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.9.0 // indirect
	github.com/hashicorp/consul v1.5.1
	github.com/hashicorp/go-checkpoint v0.5.0 // indirect
	github.com/hashicorp/go-discover v0.0.0-20190522154730-8aba54d36e17 // indirect
	github.com/hashicorp/go-hclog v0.9.2 // indirect
	github.com/hashicorp/go-memdb v1.0.2 // indirect
	github.com/hashicorp/hil v0.0.0-20190212132231-97b3a9cdfa93 // indirect
	github.com/hashicorp/nomad/api v0.0.0-20190522160243-df84e07c1a46 // indirect
	github.com/hashicorp/raft v1.0.1 // indirect
	github.com/hashicorp/raft-boltdb v0.0.0-20171010151810-6e5ba93211ea // indirect
	github.com/hashicorp/vault v1.1.2 // indirect
	github.com/hashicorp/vault-plugin-auth-alicloud v0.5.1 // indirect
	github.com/hashicorp/vault-plugin-auth-azure v0.5.1 // indirect
	github.com/hashicorp/vault-plugin-auth-centrify v0.5.1 // indirect
	github.com/hashicorp/vault-plugin-auth-jwt v0.5.1 // indirect
	github.com/hashicorp/vault-plugin-auth-kubernetes v0.5.1 // indirect
	github.com/hashicorp/vault-plugin-secrets-ad v0.5.1 // indirect
	github.com/hashicorp/vault-plugin-secrets-alicloud v0.5.1 // indirect
	github.com/hashicorp/vault-plugin-secrets-azure v0.5.1 // indirect
	github.com/hashicorp/vault-plugin-secrets-gcp v0.5.2 // indirect
	github.com/hashicorp/vault-plugin-secrets-gcpkms v0.5.1 // indirect
	github.com/hashicorp/vault-plugin-secrets-kv v0.5.1 // indirect
	github.com/hashicorp/vault/api v1.0.2 // indirect
	github.com/hashicorp/vault/sdk v0.1.11 // indirect
	github.com/influxdata/influxdb v1.7.6 // indirect
	github.com/jarcoal/httpmock v1.0.4 // indirect
	github.com/joyent/triton-go v0.0.0-20190112182421-51ffac552869 // indirect
	github.com/keybase/go-crypto v0.0.0-20190416182011-b785b22cc757 // indirect
	github.com/klauspost/cpuid v1.2.1 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/lib/pq v1.1.1
	github.com/linode/linodego v0.8.1 // indirect
	github.com/lyft/protoc-gen-validate v0.0.14 // indirect
	github.com/mailru/easyjson v0.0.0-20190403194419-1ea4449da983 // indirect
	github.com/mholt/caddy v1.0.0 // indirect
	github.com/mholt/certmagic v0.5.1 // indirect
	github.com/micro/cli v0.2.0
	github.com/micro/go-config v1.1.0
	github.com/micro/go-log v0.1.0
	github.com/micro/go-micro v1.7.1-0.20190718142447-8f2585724c34
	github.com/micro/go-rcache v0.3.0 // indirect
	github.com/micro/util v0.2.0
	github.com/mitchellh/pointerstructure v0.0.0-20190430161007-f252a8fd71c8 // indirect
	github.com/munnerz/goautoneg v0.0.0-20190414153302-2ae31c8b6b30 // indirect
	github.com/opentracing/opentracing-go v1.1.0 // indirect
	github.com/openzipkin/zipkin-go-opentracing v0.3.5 // indirect
	github.com/rogpeppe/fastuuid v1.1.0 // indirect
	github.com/shirou/gopsutil v2.18.12+incompatible // indirect
	github.com/smartystreets/assertions v0.0.0-20190401211740-f487f9de1cd3 // indirect
	github.com/softlayer/softlayer-go v0.0.0-20190508182157-7c592eb2559c // indirect
	github.com/streadway/amqp v0.0.0-20190404075320-75d898a42a94 // indirect
	github.com/ugorji/go v1.1.4 // indirect
	github.com/vmware/govmomi v0.20.1 // indirect
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/zap v1.10.0
	golang.org/x/sys v0.0.0-20190710143415-6ec70d6a5542
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	k8s.io/api v0.0.0-20190515023547-db5a9d1c40eb // indirect
	k8s.io/client-go v11.0.0+incompatible // indirect
	k8s.io/gengo v0.0.0-20190327210449-e17681d19d3a // indirect
	k8s.io/kube-openapi v0.0.0-20190510232812-a01b7d5d6c22 // indirect
	k8s.io/utils v0.0.0-20190520173318-324c5df7d3f0 // indirect
	layeh.com/radius v0.0.0-20190322222518-890bc1058917 // indirect
	sigs.k8s.io/structured-merge-diff v0.0.0-20190521201008-1c46bef2e9c8 // indirect
)
