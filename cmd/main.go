package main

func main() {
	//配置路由
	engin := InitApp()

	err := engin.Run(":8080")
	if err != nil {
		return
	}
}
