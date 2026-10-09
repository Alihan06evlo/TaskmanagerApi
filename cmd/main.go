package main

import (
	"flag"

	"taskmanagerApi/internal/apiserver"
	"taskmanagerApi/internal/core/logger"
	taskHandler "taskmanagerApi/internal/handler"
	taskRepo "taskmanagerApi/internal/repository/task"
	taskService "taskmanagerApi/internal/service/task"
	"taskmanagerApi/internal/store"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {

	logger.Init()
	flag.Parse()
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		logrus.WithError(err).Fatal("[Server] decode config error")
	}
	db := store.New()
	err = db.Open(config.BindAddr)
	repo := taskRepo.NewRepository(db)
	srv := taskService.NewTaskService(repo)
	h := taskHandler.NewHandler(srv)
		
	s := apiserver.New(config)
	s.ConfigureRouter(h)
	if err := s.Run(); err != nil {
		logrus.WithError(err).Fatal("[Server] application stopped ") 
	}


}
