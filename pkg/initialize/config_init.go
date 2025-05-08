// config/manager.go
package initialize

import (
	"fmt"
	"sync"
	"user/pkg/config"
	"user/pkg/util"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"gopkg.in/yaml.v3"
)

var once sync.Once
var client config_client.IConfigClient
var configCenterConfig *config.ConfigCenterConfig

func InitConfig(centerConfig *config.ConfigCenterConfig, config *config.Config) error {
	var innerErr error
	once.Do(func() {
		configCenterConfig = centerConfig
		address, port := util.SplitHostPort(configCenterConfig.Address)

		serverConfigs := []constant.ServerConfig{
			*constant.NewServerConfig(address, port),
		}
		clientConfig := *constant.NewClientConfig(
			constant.WithTimeoutMs(5000),
			constant.WithNotLoadCacheAtStart(true),
			constant.WithLogLevel("warn"),
			constant.WithUsername(configCenterConfig.Username),
			constant.WithPassword(configCenterConfig.Password),
			constant.WithNamespaceId(configCenterConfig.Namespace),
		)
		var err error
		client, err = clients.NewConfigClient(vo.NacosClientParam{
			ServerConfigs: serverConfigs,
			ClientConfig:  &clientConfig,
		})
		if err != nil {
			innerErr = fmt.Errorf("初始化 Nacos 客户端失败: %v", err)
			return
		}

		err = loadConfig(config)
		if err != nil {
			innerErr = fmt.Errorf("加载配置失败: %v", err)
			return
		}

		err = listenConfig(config)
		if err != nil {
			innerErr = fmt.Errorf("监听配置失败: %v", err)
			return
		}
	})

	return innerErr
}

// loadConfig 加载配置
func loadConfig(config *config.Config) error {
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: configCenterConfig.ServiceDataID,
		Group:  configCenterConfig.Group,
	})
	if err != nil {
		return fmt.Errorf("获取配置失败: %v", err)
	}
	err = updateConfig(content, config)
	if err != nil {
		return fmt.Errorf("更新配置失败: %v", err)
	}
	return nil
}

// listenConfig 监听配置
func listenConfig(config *config.Config) error {
	err := client.ListenConfig(vo.ConfigParam{
		DataId: configCenterConfig.ServiceDataID,
		Group:  configCenterConfig.Group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("配置文件已修改，重新加载...")
			err := updateConfig(data, config)
			if err != nil {
				fmt.Printf("更新配置失败: %v \n", err)
			}
		},
	})
	if err != nil {
		return fmt.Errorf("监听配置失败: %v", err)
	}
	return nil
}

// updateConfig 更新配置
func updateConfig(content string, config *config.Config) error {
	err := yaml.Unmarshal([]byte(content), config)
	if err != nil {
		return fmt.Errorf("解析配置失败: %v", err)
	}
	fmt.Printf("更新配置: %v \n", content)
	return nil
}
