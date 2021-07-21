package cmd

import (
	cfg "casorder/utils/config"

	"casorder/db"
	"casorder/utils/logging"
)

func Initialize() {

	// Initialize configuration
	cfg.Initialize("order_config")

	// Initialize log
	logging.Initialize()
	logger := logging.GetLogger()
	logger.Info("Starting Service")

	// Initialize database
	db.Initialize()
}
