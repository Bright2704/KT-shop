package main

import (
	"fmt"
	"os"

	"github.com/Bright2704/KT-shop-tutorial/config"
	"github.com/Bright2704/KT-shop-tutorial/pkg/databases"


	"github.com/Bright2704/KT-shop-tutorial/modules/servers"
)


func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	} 
}


func main() {
	cfg := config.LoadConfig(envPath())
	 //fmt.Println(cfg.App())

	db := databases.DbConnect(cfg.Db())
	defer db.Close()


	fmt.Println(db)
	
	servers.NewServer(cfg, db).Start()

}