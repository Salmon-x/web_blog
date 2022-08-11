package initialize

func init() {
	// 有序初始化
	go inItRedis()
	go inItDb()
}
