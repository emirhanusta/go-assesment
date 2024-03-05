package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

// GetConnectionPool creates a connection pool to PostgreSQL database using the provided configuration.
func GetConnectionPool(context context.Context, config Config) *pgxpool.Pool {
	// Construct connection string
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable statement_cache_mode=describe pool_max_conns=%s pool_max_conn_idle_time=%s",
		config.Host,
		config.Port,
		config.UserName,
		config.Password,
		config.DbName,
		config.MaxConnections,
		config.MaxConnectionIdleTime)

	// Parse configuration string
	connConfig, parseConfigErr := pgxpool.ParseConfig(connString)
	if parseConfigErr != nil {
		panic(parseConfigErr)
	}

	// Connect to PostgreSQL database
	conn, err := pgxpool.ConnectConfig(context, connConfig)
	if err != nil {
		log.Error("Unable to connect to database: %v\n", err)
		panic(err)
	}

	return conn
}
