package main

import (
	"net/http"

	"github.com/zonesan/clog"
)

func main() {
	clog.Info("starting service amount agent, VERSION:", Version)

	router := createRouter()

	//clog.SetLogLevel(clog.LOG_LEVEL_DEBUG)
	clog.Info("listening on port 8080...")
	clog.Fatal(http.ListenAndServe(":8080", router))
}

func init() {
	clog.SetLogLevel(clog.LOG_LEVEL_INFO)
}
