package main

import (
	"codly/datastore"
	"codly/router"
)

func main() {
	db, err := datastore.InitDB()
	if err != nil {
		return
	}
	var r = router.InitRouter(db)
	err2 := r.Run()
	if err2 != nil {
		return
	} // 监听并在 0.0.0.0:8080 上启动服务
}
