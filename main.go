package main

func main() {
	InitRouter()
	r := GetRouter()
	err := r.Run()
	if err != nil {
		return
	} // 监听并在 0.0.0.0:8080 上启动服务
}
