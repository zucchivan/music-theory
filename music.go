// Package main implements a command-line utility for music
package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/go-music/music/key"
)

func init() {
	opt = &Options{}
	opt.Parse()
}

func main() {
	if len(opt.Key) > 0 {
		yaml, _ := yaml.Marshal(key.Of(opt.Key))
		fmt.Printf("%+v", string(yaml[:]))
	}

	os.Exit(0)

}

var (
	opt *Options
)

type Options struct {
	Key string
}

func (o *Options) Parse() {
	flag.StringVar(&o.Key, "k", "", "Key")
	flag.Parse()
}
