package initiative

func Run() {
	InitLoadConfig()
	db := InitMySQL()
	InitRouter(db)
}
