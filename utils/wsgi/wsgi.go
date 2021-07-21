package wsgi

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"casorder/api"
	"casorder/db"
	"casorder/middlewares"
)

func Initialize() {

	flag.Parse()
	port := viper.GetString("server.port")
	ip := viper.GetString("server.ip")

	app := gin.Default() // create gin app

	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	gin.DisableConsoleColor()

	// Logging to a file.
	logPath := fmt.Sprintf("%v%v", viper.GetString("logging.logFolder"), "api.log")
	f, _ := os.Create(logPath)
	gin.DefaultWriter = io.MultiWriter(f)
	app.Use(gin.Recovery())

	app.Use(db.Inject(db.GetDB()))
	app.Use(middlewares.JWTMiddleware())
	api.ApplyRoutes(app) // apply api router
	var serverAddr = fmt.Sprintf("%v:%v", ip, port)
	app.Run(serverAddr) // listen to given port
}
