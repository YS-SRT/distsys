package config

type SystemConf struct {
	Addr       string `mapstructure:"addr" json:"addr" toml:"addr"`
	Mode       string `mapstructure:"mode" json:"mode" toml:"mode"`
	StaticPath string `mapstructure:"static-path" json:"static-path" toml:"static-path"`
}

type PostgresConf struct {
	DSN          string `mapstructure:"dsn" json:"dsn" toml:"dsn"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" toml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" toml:"max-open-conns"`
	LogMode      string `mapstructure:"log-mode" json:"logMode" toml:"log-mode"`
	LogZap       bool   `mapstructure:"log-zap" json:"logZap" toml:"log-zap"`
	EncryCode    string `mapstructure:"encry-code" json:"encry-code" toml:"encry-code"`
}

type ZapConf struct {
	Level         string `mapstructure:"level" json:"level" toml:"level"`
	Format        string `mapstructure:"format" json:"format" toml:"format"`
	Prefix        string `mapstructure:"prefix" json:"prefix" toml:"prefix"`
	Director      string `mapstructure:"director" json:"director"  toml:"director"`
	ShowLine      bool   `mapstructure:"show-line" json:"showLine" toml:"showLine"`
	EncodeLevel   string `mapstructure:"encode-level" json:"encodeLevel" toml:"encode-level"`
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktraceKey" toml:"stacktrace-key"`
	LogInConsole  bool   `mapstructure:"log-in-console" json:"logInConsole" toml:"log-in-console"`
}

type MinioConf struct {
	AccessKey    string `mapstructure:"access-key" json:"access-key" toml:"access-key"`
	AccessSecret string `mapstructure:"access-secret" json:"access-secret" toml:"access-secret"`
	EndPoint     string `mapstructure:"endpoint" json:"endpoint" toml:"endpoint"`
	UseSSL       bool   `mapstructure:"use-ssl" json:"use-ssl" toml:"use-ssl"`
	Bucket       string `mapstructure:"bucket" json:"bucket" toml:"bucket"`
	PublicPrefix string `mapstructure:"public-prefix" json:"public-prefix" toml:"public-prefix"`
}

type RedisConf struct {
	DSN        string `mapstructure:"dsn" json:"dsn" toml:"dsn"`
	ClusterDNS string `mapstructure:"cluster-dsn" json:"cluster-dsn" toml:"cluster-dsn"`
}

type JWTConf struct {
	AccessKey             string   `mapstructure:"access-key" json:"access-key" toml:"access-key"`
	RefreshKey            string   `mapstructure:"refresh-key" json:"refresh-key" toml:"refresh-key"`
	AccessExpiredMinutes  int      `mapstructure:"access-expired-minutes" json:"access-expired-minutes" toml:"access-expired-minutes"`
	RefreshExpiredMinutes int      `mapstructure:"refresh-expired-minutes" json:"refresh-expired-minutes" toml:"refresh-expired-minutes"`
	ExcludePaths          []string `mapstructure:"exclude-paths" json:"exclude-paths" toml:"rexclude-paths"`
	ECDSAPrivateKey       string   `mapstructure:"private-key" json:"private-key" toml:"private-key"`
	ECDSAPublicKey        string   `mapstructure:"public-key" json:"public-key" toml:"public-key"`
}

type CasbinConf struct {
	UserGroupPType      string `mapstructure:"user-group-ptype" json:"user-group-ptype" toml:"user-group-ptype"`
	ResourceGroupPType  string `mapstructure:"resource-group-ptype" json:"resource-group-ptype" toml:"resource-group-ptype"`
	UserPrefix          string `mapstructure:"user-prefix" json:"user-prefix" toml:"user-prefix"`
	UserGroupPrefix     string `mapstructure:"user-group-prefix" json:"user-group-prefix" toml:"user-group-prefix"`
	ResourceGroupPrefix string `mapstructure:"resource-group-prefix" json:"resource-group-prefix" toml:"resource-group-prefix"`
	GroupDefaultParam   string `mapstructure:"group-default-param" json:"group-default-param" toml:"group-default-param"`
}

type ConfigDoc struct {
	System    SystemConf   `mapstructure:"system" json:"system" toml:"system"`
	Postgres  PostgresConf `mapstructure:"postgres" json:"postgres" toml:"postgres"`
	Zap       ZapConf      `mapstructure:"zap" json:"zap" toml:"zap"`
	MinioAuth MinioConf    `mapstructure:"minio" json:"minio" toml:"minio"`
	Redis     RedisConf    `mapstructure:"redis" json:"redis" toml:"redis"`
	JWT       JWTConf      `mapstructure:"jwt" json:"jwt" toml:"jwt"`
	Casbin    CasbinConf   `mapstructure:"casbin" json:"casbin" toml:"casbin"`
}
