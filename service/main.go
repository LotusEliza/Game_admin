package main

import (
	"flag"
	"fmt"
	"github.com/finalist736/gokit/config"
	"github.com/finalist736/gokit/database"
	"github.com/finalist736/gokit/logger"
	"github.com/finalist736/gokit/mainloop"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"math/rand"
	"os"
	"tcs"
	"tcs/machanics/daytime"
	"tcs/machanics/miner"
	"tcs/telegram"
	"time"
)

var configPath = flag.String("config", "config.ini", "config file path")

func main() {
	rand.Seed(time.Now().UnixNano())
	var err = config.Init(config.NewFileProvider(configPath))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "can't load config")
		os.Exit(1)
	}

	logger.ReloadLogs(config.DefaultConfig())

	logger.StdOut().Infof("Space Colony v%s", strategy.Version)

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

	err = goose.Up(database.GetDefaultSession().DB, "migrations")
	if err != nil {
		logger.StdErr().Errorf("migrations error: %s", err)
		os.Exit(1)
	}

	//items.Dump()
	telegram.RestartDutyRoutines()
	miner.Init()

	daytime.Init()

	err = telegram.Init(config.MustString("tgapikey"))
	if err != nil {
		logger.StdErr().Errorf("telegram bot init error: %s", err)
		os.Exit(1)
	}

	mainloop.Loop(stop, grace, config.DefaultConfig())

}

func stop() {
	telegram.Stop()
}

func grace() {
	telegram.Stop()
}
