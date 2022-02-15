package config

import (
	"context"
	"encoding/json"
	"log"

	"github.com/comeonjy/go-kit/pkg/xconfig"
	"github.com/comeonjy/go-kit/pkg/xconfig/apollo"
	"github.com/comeonjy/go-kit/pkg/xenv"
)

type Config struct {
	QiniuAK  string `json:"qiniu_ak"`
	QiniuSK  string `json:"qiniu_sk"`
	QiniuCDN string `json:"qiniu_cdn"`
}

func (c *Config) Validate() error {
	return nil
}

// Interface 对外暴露接口（用于功能扩展）
type Interface interface {
	Get() Config
	xconfig.ReloadConfigInterface
}

// 内部配置类载体
type config struct {
	xconfig.IConfig
}

// Get 获取配置
func (c *config) Get() Config {
	return c.LoadValue().(Config)
}

func (c *config) ReloadConfig() error {
	return c.ReLoad()
}

func storeHandler(data []byte) interface{} {
	conf := Config{}
	if err := json.Unmarshal(data, &conf); err != nil {
		log.Println(err)
	}
	if err := conf.Validate(); err != nil {
		log.Println(err)
	}
	return conf
}

func NewConfig(ctx context.Context) Interface {
	cfg := &config{
		xconfig.New(ctx,
			apollo.NewSource(xenv.GetEnv(xenv.ApolloUrl), xenv.GetEnv(xenv.ApolloAppID), xenv.GetApolloCluster("default"), xenv.GetApolloSecret(), xenv.GetApolloNamespace("grpc"), xenv.GetApolloNamespace("common")),
			storeHandler),
	}
	return cfg
}
