package app

import "backend-assigment/common/postgresql"

// ConfigurationManager holds configuration settings for the application.
type ConfigurationManager struct {
	PostgreSqlConfig postgresql.Config
}

// NewConfigurationManager creates a new instance of ConfigurationManager.
func NewConfigurationManager() *ConfigurationManager {
	// Retrieve PostgreSQL configuration
	postgreSqlConfig := getPostgreSqlConfig()
	return &ConfigurationManager{PostgreSqlConfig: postgreSqlConfig}
}

// getPostgreSqlConfig retrieves PostgreSQL configuration.
func getPostgreSqlConfig() postgresql.Config {
	// Set PostgreSQL configuration values
	return postgresql.Config{
		Host:                  "localhost",
		Port:                  "5432",
		DbName:                "report",
		UserName:              "postgres",
		Password:              "postgres",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "30s",
	}
}
