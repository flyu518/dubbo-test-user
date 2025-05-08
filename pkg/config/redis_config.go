package config

// RedisConfig Redis配置结构
type RedisConfig struct {
	Mode     string       `mapstructure:"mode" yaml:"mode" json:"mode"`
	Password string       `mapstructure:"password" yaml:"password" json:"password"`
	Single   RedisSingle  `mapstructure:"single" yaml:"single" json:"single"`
	Cluster  RedisCluster `mapstructure:"cluster" yaml:"cluster" json:"cluster"`
}

// RedisSingle 单节点配置
type RedisSingle struct {
	Addr string `mapstructure:"addr" yaml:"addr" json:"addr"`
	DB   int    `mapstructure:"db" yaml:"db" json:"db"`
}

// RedisCluster 集群配置
type RedisCluster struct {
	Addrs []string `mapstructure:"addrs" yaml:"addrs" json:"addrs"`
}
