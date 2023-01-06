package main

import (
	"context"
	"log"

	"github.com/spf13/pflag"

	"github.com/denis96z/simple-version-tracker/worker/pkg/checker"
	"github.com/denis96z/simple-version-tracker/worker/pkg/config"
	"github.com/denis96z/simple-version-tracker/worker/pkg/logs"
	"github.com/denis96z/simple-version-tracker/worker/pkg/loop"
	"github.com/denis96z/simple-version-tracker/worker/pkg/storage"
)

func main() {
	confFileName := ""
	pflag.StringVarP(
		&confFileName,
		"config", "c", "",
		"config file path",
	)

	pflag.Parse()
	if confFileName == "" {
		log.Fatal(`no "config" param`)
	}

	conf, err := config.Load(confFileName)
	if err != nil {
		log.Fatal("failed to load config: ", err)
	}
	if err = logs.SetMinLevel(conf.Logger.MinLevel); err != nil {
		log.Fatal("failed to set min log level: ", err)
	}
	if err = logs.SetFileOutput(conf.Logger.FilePath); err != nil {
		log.Fatal("failed to set log file path: ", err)
	}

	logs.Debug(
		"[CONFIG]\n", conf.Dump(),
	)

	strg := storage.NewStorage(conf.Storage)
	if err = strg.Init(context.Background()); err != nil {
		log.Fatal("failed to init storage: ", err)
	}
	defer func() {
		if err = strg.Finit(context.Background()); err != nil {
			logs.Error("failed to finit storage: ", err)
		}
	}()

	lp := loop.NewLoop(
		conf.Loop, strg, checker.NewChecker(conf.Checker),
	)
	if err = lp.Run(context.Background()); err != nil {
		logs.Error("failed to run loop: ", err)
	}
}
