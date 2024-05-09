package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/cloudflare/tableflip"
)

func main() {
	fmt.Println(os.Getpid())
	upg, _ := tableflip.New(tableflip.Options{})
	defer upg.Stop()

	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGHUP)
		for range sig {
			upg.Upgrade()
		}
	}()

	// Listen must be called before Ready
	ln, err := upg.Listen("tcp", "localhost:8080")
	fmt.Println(err)
	defer ln.Close()

	go http.Serve(ln, nil)

	if err := upg.Ready(); err != nil {
		panic(err)
	}

	<-upg.Exit()
}
