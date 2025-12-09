package logger

import (
	serverconfig "testops_copilot/internal/config"

	"github.com/NoOl01/golog/pkg/golog"
	"github.com/NoOl01/golog/pkg/golog/golog_config"
)

var Log golog.DefaultLogger

func InitLogger() {
	config := golog_config.Config{
		Format:  "${level} ${l} ${name} ${l} ${content} ${l} ${timestamp}",
		Literal: " | ",
		Debug:   serverconfig.Env.Debug,
	}

	Log = golog.Start(&config)
}
