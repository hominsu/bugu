module github.com/hominsu/bugu/app/bugu/service

go 1.18

require (
	entgo.io/ent v0.10.1
	github.com/go-kratos/kratos/contrib/registry/consul/v2 v2.0.0-20220604031450-874d4c3edcf5
	github.com/go-kratos/kratos/v2 v2.3.1
	github.com/go-redis/redis/v8 v8.11.5
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang-jwt/jwt/v4 v4.4.1
	github.com/google/uuid v1.3.0
	github.com/google/wire v0.5.0
	github.com/gorilla/handlers v1.5.1
	github.com/hashicorp/consul/api v1.13.0
	github.com/hominsu/bugu v1.0.0
	github.com/hominsu/bugu/pkg v0.0.0-20220608071005-03611bdc29ab
	google.golang.org/grpc v1.47.0
	google.golang.org/protobuf v1.28.0
)

require (
	ariga.io/atlas v0.3.7-0.20220303204946-787354f533c3 // indirect
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/agext/levenshtein v1.2.1 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/armon/go-metrics v0.3.10 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/envoyproxy/protoc-gen-validate v0.6.7 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/felixge/httpsnoop v1.0.1 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/go-kratos/aegis v0.1.2 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.5 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/go-playground/form/v4 v4.2.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/btree v1.0.0 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.3 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.0.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/hcl/v2 v2.10.0 // indirect
	github.com/hashicorp/serf v0.9.6 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-wordwrap v0.0.0-20150314170334-ad45545899c7 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/shirou/gopsutil/v3 v3.21.8 // indirect
	github.com/tklauser/go-sysconf v0.3.9 // indirect
	github.com/tklauser/numcpus v0.3.0 // indirect
	github.com/zclconf/go-cty v1.8.0 // indirect
	go.opentelemetry.io/otel v1.7.0 // indirect
	go.opentelemetry.io/otel/trace v1.7.0 // indirect
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e // indirect
	golang.org/x/mod v0.5.1 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sync v0.0.0-20220513210516-0976fa681c29 // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220519153652-3a47de7e79bd // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/hominsu/bugu => ../../../
	github.com/hominsu/bugu/pkg => ../../../pkg
)
