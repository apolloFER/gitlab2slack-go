package main

import (
	"net/http"
	"runtime"
	"log"
	"fmt"
	flags "github.com/jessevdk/go-flags"
)

var options struct {
	Listen string `short:"l" long:"listen" default:"0.0.0.0:5000" description:"IP:port to listen on"`
	Domain string `short:"d" long:"domain" description:"Slack domain" required:"true"`
	Token string `short:"t" long:"token" description:"Slack incoming webhook token" required:"true"`
	Channels []string `short:"c" long:"channels" description:"A slice of channels" required:"true"`
}

var slack Slack

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	_, err := flags.Parse(&options)

	if err != nil {
		log.Fatalln(err)
	}

	slack.HookUrl = fmt.Sprintf("https://%s.slack.com/services/hooks/incoming-webhook?token=%s", options.Domain, options.Token)
	slack.Channels = options.Channels

	http.HandleFunc("/gitlab", PostOnly(gitlabHandler))

	if err := http.ListenAndServe(options.Listen, nil); err != nil {
		log.Fatalln(err)
	}
}
