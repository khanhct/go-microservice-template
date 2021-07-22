package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"github.com/spf13/viper"

	"casorder/cmd"
	"casorder/db"
	"casorder/utils/mgrpc"
)

type Command struct {
}

func (c *Command) MigrateDB() {
	fmt.Println("Migrating database")

	_db := db.GetDB()
	driver, _ := mysql.WithInstance(_db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"mysql",
		driver,
	)
	m.Steps(2)
}

func (c *Command) DowngrateDB() {
	fmt.Println("Dowgrating database")

	_db := db.GetDB()
	driver, _ := mysql.WithInstance(_db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"mysql",
		driver,
	)
	m.Steps(-2)
}

func (c *Command) TestGrpc() {
	fmt.Println("Test Grpc Server")
	grpcAddr := fmt.Sprintf("%v:%v", viper.GetString("grpc.host"), viper.GetString("grpc.port"))
	mgrpc.TestClient(grpcAddr)
}

func main() {
	cmd.Initialize()

	app := kingpin.New("modular", "My modular application.")
	app.Version("1.0.0")

	dbMigrate := app.Command("db-migrate", "Migrate database")
	dbDowngrate := app.Command("db-downgrate", "Downgrate database")
	testGrpc := app.Command("test-grpc", "Test grpc server")

	checkConfigCmd := app.Command("config", "Check if the config files are valid or not.")
	configFiles := checkConfigCmd.Arg(
		"config-files",
		"The config files to check.",
	).Required().ExistingFiles()

	command := Command{}

	parsedCmd := kingpin.MustParse(app.Parse(os.Args[1:]))
	switch parsedCmd {
	case checkConfigCmd.FullCommand():
		if *configFiles == nil {
			os.Exit(0)
		}
	case dbMigrate.FullCommand():
		command.MigrateDB()

	case dbDowngrate.FullCommand():
		command.DowngrateDB()
	case testGrpc.FullCommand():
		command.TestGrpc()
	}
}
