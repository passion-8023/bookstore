package main

import (
	"bookstore/internal/controller"
	_ "bookstore/internal/store"
	"bookstore/pkg/app"
	"bookstore/server"
	"bookstore/store/factory"
	"context"

	"github.com/gorilla/mux"
)

func main() {
	s, err := factory.New("mem") // 创建图书数据存储模块实例
	if err != nil {
		panic(err)
	}
	bookController := controller.NewBookController(s)
	router := &server.Router{
		Router:         mux.NewRouter(),
		BookController: bookController,
	}
	httpServer := server.NewHttpServer(router)

	newApp, err := app.NewApp(app.Context(context.Background()), app.Server(httpServer...), app.Version("v1.0.0"))
	if err != nil {
		panic(err)
	}

	if err = newApp.Run(); err != nil {
		panic(err)
	}
}
