package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"time"

	"github.com/hoangnhat/project/routers"

	"github.com/hoangnhat/project/dataservice"
	"github.com/hoangnhat/project/helpers"

	"github.com/gin-gonic/gin"
)

// *****************************************************************************
// Application Settings
// *****************************************************************************

// configuration contains the application settings
type configuration struct {
	Database dataservice.Info `json:"Database"`
	Session  helpers.Session  `json:"Session"`
}

// config the settings variable
var config = &configuration{}

func main() {
	//* write Log file
	dateTime := time.Now()
	logFile, err := os.Create("logs/log" + dateTime.Format("01-02-2006"))
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	log.SetOutput(gin.DefaultWriter)
	//* end write Log file

	//! connect DB
	helpers.Load("dataservice"+string(os.PathSeparator)+"configDB.json", config)
	// Configure the session cookie store
	helpers.Configure(config.Session)
	// Connect to database
	dataservice.InitDb(config.Database)
	//! end connect DB
	start := routers.SetRouter()
	if !start {
		panic("error. Cannot start server")
	}

}

// ParseJSON unmarshals bytes to structs
func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}
