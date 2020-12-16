package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

const CONFIG_NET = "tcp"
const DB_TYPE = "mysql"

type Cfg struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Hostname string `yaml:"hostname"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"db_name"`
	// Type     string `yaml:"type"`
}

type Interface interface {
	Connect()
	sqlx.QueryerContext
	sqlx.ExecerContext
}

type client struct {
	dbcfg *mysql.Config
	db    *sqlx.DB
}

func New() Interface {
	return &c
}

var c client

func init() {
	if cfg, err := readConfig(); err != nil {
		log.Fatal("Fatal error: Failed reading database config")
	} else {
		c.dbcfg = cfg
	}

	c.Connect()
}

func (c *client) Connect() {
	c.db = sqlx.MustOpen(DB_TYPE, c.dbcfg.FormatDSN())
	c.db.SetConnMaxLifetime(0)
	c.db.SetMaxIdleConns(3)
	c.db.SetMaxOpenConns(3)

	// make sure no connectivity issues
	if err := c.db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func (c *client) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return c.db.QueryContext(ctx, query, args)
}

func (c *client) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	return c.db.QueryxContext(ctx, query)
}

func (c *client) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	return c.db.QueryRowxContext(ctx, query)
}

func (c *client) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	// run within transaction
	return c.runTransaction(ctx, func(*sqlx.Tx) (sql.Result, error) {
		return c.db.ExecContext(ctx, query)
	})
}

// readConfig: read config into a db config and return
func readConfig() (*mysql.Config, error) {
	// read from config file
	viper.SetConfigName("datasources")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("Fatal error reading config file: %s \n", err)
	}

	//
	dbcfg := mysql.NewConfig()
	dbcfg.User = viper.GetString("cp.user")
	dbcfg.Passwd = viper.GetString("cp.password")
	dbcfg.DBName = viper.GetString("cp.db_name")
	dbcfg.Addr = viper.GetString("cp.hostname") + ":" + viper.GetString("cp.port")
	dbcfg.Net = CONFIG_NET
	dbcfg.ParseTime = true

	return dbcfg, nil
}

// runs the provided function in a transaction
func (q *client) runTransaction(ctx context.Context, logic func(tx *sqlx.Tx) (sql.Result, error)) (sql.Result, error) {
	// Start a transaction
	tx, err := c.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
	})
	if err != nil {
		log.Printf("ERROR: internal/db/RunTransaction: %s", err)
		return nil, err
	}

	// Run the logic
	var result sql.Result
	result, err = logic(tx)
	if err != nil {
		// Rollback
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Printf("ERROR: datamodels/db_svc.go/RunTransaction: %s [unable to rollback]", rollbackErr)
		}
		return nil, err
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		log.Printf("ERROR: datamodels/db_svc.go/RunTransaction: %s [Commit transaction]", err)
		return nil, err
	}

	return result, nil
}
