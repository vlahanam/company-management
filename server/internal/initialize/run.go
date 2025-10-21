package initialize

func Run() {
	cfg := LoadConfig()
	db := InitMysql(cfg)
	InitRoute(cfg, db)
}