package initialize

func Run() {
	cfg := LoadConfig()
	InitMysql(cfg)
	InitRoute(cfg)
}