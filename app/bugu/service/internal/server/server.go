package server

import (
	"time"

	"bugu/app/bugu/service/internal/conf"
	"github.com/go-kratos/kratos/contrib/registry/polaris/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	polarisConf "github.com/polarismesh/polaris-go/pkg/config"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewHTTPServer, NewRegistrar)

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := polarisConf.NewDefaultConfiguration(conf.Polaris.Addresses)

	r := polaris.NewRegistryWithConfig(
		c,
		polaris.WithNamespace(conf.Polaris.Namespace),
		polaris.WithServiceToken(conf.Polaris.Token),
		polaris.WithHealthy(true),
		polaris.WithTTL(30),
		polaris.WithTimeout(time.Second*15),
		polaris.WithHeartbeat(true),
	)

	return r
}
