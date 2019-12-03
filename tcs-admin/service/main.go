package main

import (
	"flag"
	"fmt"
	"github.com/finalist736/gokit/config"
	"github.com/finalist736/gokit/database"
	"github.com/finalist736/gokit/logger"
	"github.com/finalist736/gokit/mainloop"
	"github.com/finalist736/gokit/webserver"
	"github.com/gocraft/web"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"math/rand"
	"os"
	"tcs/machanics/daytime"
	"tcs/tcs-admin/handlers"
	"tcs/tcs-admin/tcsctx"
	"time"
)

var configPath = flag.String("config", "config.ini", "config file path")

func main() {
	rand.Seed(time.Now().UnixNano())
	var err = config.Init(config.NewFileProvider(configPath))
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't load config")
		os.Exit(1)
	}

	logger.ReloadLogs(config.DefaultConfig())

	logger.StdOut().Infof("Tagoreans Crash Site Admin panel")

	dbconf := database.DBConfig{}
	dbconf.Dsn = config.MustString("dsn")
	dbconf.Driver = "postgres"
	dbconf.LifeTime = time.Minute * 5
	dbconf.Stream = logger.DatabaseStream()
	dbconf.MaxIdleConns = 1
	dbconf.MaxOpenConns = 10
	err = database.Add(&dbconf)
	if err != nil {
		logger.StdErr().Errorf("db open error: %s", err)
		os.Exit(1)
	}

	err = goose.Up(database.GetDefaultSession().DB, "../migrations")
	if err != nil {
		logger.StdErr().Errorf("migrations error: %s", err)
		os.Exit(1)
	}

	daytime.Init()

	router := web.New(tcsctx.Ctx{})
	router.Middleware((*tcsctx.Ctx).Cors)

	router.Get("/time", handlers.DayTimeGet)

	router.Get("/players", handlers.PlayersGet)
	router.Get("/player", handlers.PlayerGet)
	router.Get("/player/inventory", handlers.PlayersInventory)
	router.Get("/player/wallet", handlers.PlayersWallet)
	router.Post("/player/wallet", handlers.PlayersWalletSet)
	router.Post("/player/name", handlers.PlayerNameSet)
	router.Post("/player/inventory/add", handlers.PlayerInventorySet)
	router.Post("/player/inventory/remove", handlers.PlayerInventoryRemove)
	router.Post("/player/inventory/update", handlers.PlayerInventoryUpdate)

	router.Post("/player/location/update", handlers.PlayerLocationSet)
	router.Post("/player/air/update", handlers.PlayerAirSet)
	router.Post("/player/balloon/update", handlers.PlayerBalloonSet)
	router.Post("/player/console/update", handlers.PlayerConsoleSet)
	router.Get("/equipments", handlers.EquipmentsGet)
	router.Post("/equipments/remove", handlers.EquipmentsRemove)
	router.Post("/equipments/add", handlers.EquipmentsSet)
	router.Post("/equipment/update", handlers.EquipmentUpdate)
	router.Get("/player/norms", handlers.PlayerNormsGet)
	router.Get("/player/norm", handlers.PlayerNormGet)
	router.Post("/player/norms/remove", handlers.PlayerNormRemove)
	router.Post("/player/norms/add", handlers.PlayerNormSet)
	router.Post("/player/norms/update", handlers.PlayerNormUpdate)

	webserver.Start(router, config.MustString("port"))

	mainloop.Loop(stop, grace, config.DefaultConfig())

}

func stop() {
	webserver.Stop()
}

func grace() {
	webserver.Grace(configPath)
}
