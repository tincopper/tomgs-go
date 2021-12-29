module tomgs-go

require (
	dubbo.apache.org/dubbo-go/v3 v3.0.0-rc2.0.20211009125508-c22d9c3e529f
	github.com/Joker/jade v1.0.0 // indirect
	github.com/Shopify/goreferrer v0.0.0-20181106222321-ec9c9a553398 // indirect
	github.com/ajg/form v1.5.1 // indirect
	github.com/apache/dubbo-go-hessian2 v1.9.3
	github.com/argoproj/argo v2.5.2+incompatible
	github.com/argoproj/pkg v0.9.1
	github.com/aymerick/raymond v2.0.2+incompatible // indirect
	github.com/common-nighthawk/go-figure v0.0.0-20210622060536-734e95fb86be
	github.com/docker/spdystream v0.0.0-20181023171402-6480d4af844c // indirect
	github.com/donnie4w/go-logger v0.0.0-20170827050443-4740c51383f4
	github.com/dubbogo/gost v1.11.17
	github.com/dubbogo/triple v1.0.6-0.20210909153707-3620c8d2d97c
	github.com/eapache/queue v1.1.0
	github.com/eknkc/amber v0.0.0-20171010120322-cdade1c07385 // indirect
	github.com/elazarl/go-bindata-assetfs v1.0.1
	github.com/elazarl/goproxy v0.0.0-20200315184450-1f3cb6622dad // indirect
	github.com/flosch/pongo2 v0.0.0-20190707114632-bbf5a6c351f4 // indirect
	github.com/gavv/monotime v0.0.0-20190418164738-30dba4353424 // indirect
	github.com/go-pg/pg/v10 v10.10.6
	github.com/golang/glog v1.0.0
	github.com/golang/protobuf v1.5.2
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/googleapis/gnostic v0.4.1 // indirect
	github.com/gorilla/schema v1.1.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.6.0
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/imkira/go-interpol v1.1.0 // indirect
	github.com/iris-contrib/blackfriday v2.0.0+incompatible // indirect
	github.com/iris-contrib/formBinder v5.0.0+incompatible // indirect
	github.com/iris-contrib/go.uuid v2.0.0+incompatible // indirect
	github.com/iris-contrib/httpexpect v0.0.0-20180314041918-ebe99fcebbce // indirect
	github.com/jhump/protoreflect v1.10.1
	github.com/jteeuwen/go-bindata v3.0.7+incompatible // indirect
	github.com/juju/errors v0.0.0-20190930114154-d42613fe1ab9 // indirect
	github.com/kataras/golog v0.0.10 // indirect
	github.com/kataras/iris v11.1.1+incompatible
	github.com/klauspost/compress v1.9.7 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.5 // indirect
	github.com/microcosm-cc/bluemonday v1.0.2 // indirect
	github.com/moul/http2curl v1.0.0 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/shima-park/agollo v1.1.7
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/viper v1.7.1
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/yalp/jsonpath v0.0.0-20180802001716-5cc68e5049a0 // indirect
	github.com/yudai/gojsondiff v1.0.0 // indirect
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	github.com/yudai/pp v2.0.1+incompatible // indirect
	go.uber.org/zap v1.19.1
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5
	google.golang.org/genproto v0.0.0-20210903162649-d08c68adba83
	google.golang.org/grpc v1.41.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.18.1
	k8s.io/apimachinery v0.18.1
	k8s.io/client-go v0.18.1
	k8s.io/kubectl v0.18.1
	k8s.io/utils v0.0.0-20200327001022-6496210b90e8 // indirect
)

//replace dubbo.apache.org/dubbo-go/v3 => ../github.com/apache/dubbo-go

go 1.13
