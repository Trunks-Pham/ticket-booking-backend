package settings

type Config struct {
	Server         ServerSetting         `mapstructure:"server"`
	PostgreSql     PostgreSqlSetting     `mapstructure:"postgreSql"`
	Authentication AuthenticationSetting `mapstructure:"authentication"`
}

type ServerSetting struct {
	Port int `mapstructure:"port"`
}

type PostgreSqlSetting struct {
	Host     string `mapstructure:"host"`
	Dbname   string `mapstructure:"dbname"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	Sslmode  string `mapstructure:"sslmode"`
}

type AuthenticationSetting struct {
	JwtScretKey string `mapstructure:"jwtScretKey"`
}
