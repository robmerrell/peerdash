package main

import (
	"fmt"
	"github.com/robmerrell/comandante"
	"github.com/robmerrell/peerdash/cmds"
	"github.com/robmerrell/peerdash/config"
	"github.com/robmerrell/peerdash/models"
	"github.com/robmerrell/peerdash/updaters"
	"os"
)

func main() {
	// get the environment for the config
	appEnv := ""
	env := os.Getenv("PEERDASH_ENV")
	switch env {
	case "dev", "test", "prod":
		appEnv = env
	default:
		appEnv = "dev"
	}

	config.LoadConfig(appEnv)

	// make the package level database connection
	if err := models.ConnectToDB(config.String("database.host"), config.String("database.db")); err != nil {
		fmt.Fprintln(os.Stderr, "Could not connect to the database")
		os.Exit(1)
	}
	defer models.CloseDB()

	bin := comandante.New("peerdash", "Peercoin dashboard")
	bin.IncludeHelp()

	// setup (database, indexes)

	// run web service

	// update peercoin prices
	updateCoinPrices := comandante.NewCommand("update_coin_prices", "Get updated peercoin prices", cmds.UpdateAction(&updaters.CoinPrice{}))
	updateCoinPrices.Documentation = cmds.UpdateCoinPricesDoc
	bin.RegisterCommand(updateCoinPrices)

	// update forum posts

	// update reddit stories

	// update network info

	if err := bin.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
