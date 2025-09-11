package setting

type Config struct {
	Mysql  MySQLSetting  `mapstructure:"mysql"`
	Server ServerSetting `mapstructure:"server"`
}

type MySQLSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Dbname   string `mapstructure:"dbname"`
}

type ServerSetting struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
