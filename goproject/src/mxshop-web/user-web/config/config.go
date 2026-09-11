package config
type UserSrvConfig struct{
	Host string `mapstructure:"host"`
	Prot int `mapstructuer:"prot"`
}
type ServerConfig struct{
	Name string `mapstructure:"name"`
	Prot int `mapstructuer:"prot"`
	UserSrvInfo UserSrvConfig `mapstructure:"user-srv"`

}
