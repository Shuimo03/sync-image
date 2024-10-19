package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync-image/config"
	"sync-image/docker"
	"sync-image/sync"
	"time"
)

var (
	configFile = flag.String("config", "config.yaml", "config file")
	authFile   = flag.String("auth", "", "auth file")
	logDir     = "logs" // 日志目录
)

func setupLogging() *os.File {
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Fatal("Failed to create log directory:", err)
	}

	currentDate := time.Now().Format("2006-01-02")
	logFilePath := filepath.Join(logDir, fmt.Sprintf("sync_image_%s.log", currentDate))

	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}

	log.SetOutput(logFile)
	log.Println("Logging started...")
	return logFile
}

func main() {

	flag.Parse()
	logFile := setupLogging()
	defer logFile.Close()

	cnf, err := config.LoadConfig(*configFile)
	if err != nil {
		log.Fatal("Load config error:", err)
	}

	authCfg, err := docker.LoadAuth(*authFile)
	if err != nil {
		log.Fatal("Load auth error:", err)
	}

	authStr, err := docker.EncodeAuthToBase64(authCfg)
	if err != nil {
		log.Fatal("Encode auth error:", err)
	}

	operator, err := docker.NewImageOperator()
	if err != nil {
		log.Fatal("Error initializing image operator:", err)
	}
	defer operator.Client.Close()

	imageSync := sync.ImageSync{
		Operator: operator,
		Config:   cnf,
	}

	if sync_image_err := imageSync.SyncImages(authStr); sync_image_err != nil {
		log.Fatal("Sync image error:", sync_image_err)
	}
}
