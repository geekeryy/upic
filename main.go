package main

import (
	"github.com/comeonjy/go-kit/pkg/xenv"
	"github.com/comeonjy/upic/cmd"
)

var EnvMap = map[string]string{
	xenv.AppName:     "upic",
	xenv.ApolloAppID: "upic",
	xenv.ApolloUrl:   "http://apollo.dev.jiangyang.me",
}

func main() {
	xenv.Init(EnvMap)
	cmd.Execute()
}
