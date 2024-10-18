package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/TiyaAnlite/FocotServicesCommon/envx"
	"github.com/duke-git/lancet/v2/netutil"
	"k8s.io/klog/v2"
	"net/http"
	"os"
	"time"
)

type config struct {
	Endpoint string `env:"ENDPOINT,required"`
}

var (
	cfg     = &config{}
	buffer  string
	scanner = bufio.NewScanner(os.Stdin)
	client  = netutil.NewHttpClientWithConfig(&netutil.HttpClientConfig{
		HandshakeTimeout: time.Second * 3,
		ResponseTimeout:  time.Second * 3,
		Compressed:       true,
	})
)

func init() {
	flag.Parse()
	envx.MustLoadEnv(cfg)
}

func main() {
	klog.Info("checking service...")
	resp, err := client.Get(cfg.Endpoint)
	if err != nil {
		klog.Fatalf("failed check service: %s", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		klog.Fatalf("failed check service with code: %d", resp.StatusCode)
	}
selectLoop:
	for {
		fmt.Print("Select function(h for help): ")
		inputScanner()
		switch buffer {
		case "h":
			fmt.Println("position: p\nbook: b\naction: a\nquit: q")
		case "p":
			position()
		case "b":
			book()
		case "a":
			action()
		case "q":
			break selectLoop
		}
	}
}
