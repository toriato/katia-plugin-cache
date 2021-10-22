package main

import (
	"errors"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/toriato/katia"
)

const Context string = "cache"

var ErrUnsupportedSource = errors.New("지원하지 않는 캐시 종류입니다")

var Plugin = katia.Plugin{
	Name:        "katia-plugin-cache",
	Description: "메모리, Redis 등을 통한 키-값 캐시를 제공합니다",
	Author:      "Sangha Lee <totoriato@gmail.com>",
	Version:     [3]int{0, 1, 0},

	OnEnable: func(bot *katia.Bot, plugin *katia.Plugin) error {
		bot.Context.Set(Context, cache.New(time.Minute, time.Minute))
		return nil
	},
}
