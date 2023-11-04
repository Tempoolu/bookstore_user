package main

import (
	"github.com/Tempoolu/bookstore_user/controllers"
	"github.com/Tempoolu/bookstore_user/discovery"
	"github.com/Tempoolu/bookstore_user/models"
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File
)

func main() {
	Cfg, _ = ini.Load("conf/user.ini")
	server, _ := Cfg.GetSection("server")
    if discovery.InitDiscovery(server.Key("etcd_ip").MustString("")) != nil {
        return
    }
    addr := server.Key("ip").MustString("") + ":" + server.Key("port").MustString("")
    discovery.Register("user", addr)
    discovery.Discovery("user")

	database, _ := Cfg.GetSection("database")
	models.InitDB(database.Key("user").MustString(""), database.Key("password").MustString(""), database.Key("address").MustString(""), database.Key("database").MustString(""))
	models.Migrate()

	e := gin.Default()
	controllers.InitRouter(e)
    // FIXME: have call now
    defer discovery.Unregister("user", addr)
	e.Run(":"+server.Key("port").MustString(""))
}
