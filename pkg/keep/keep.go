package keep

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cloudflare/tableflip"
)

type Keeper struct {
	closers []io.Closer
	upg     *tableflip.Upgrader
	isKeep  bool
}

func New(closers ...io.Closer) *Keeper {
	return &Keeper{
		closers: closers,
	}
}

func (k *Keeper) SetKeep(isKeep bool) *Keeper {
	k.isKeep = isKeep
	return k
}

// testing: kill -HUP $PID
func (k *Keeper) Run() error {
	fmt.Printf("current PID is: %d\n\n", os.Getpid())
	upg, err := tableflip.New(tableflip.Options{})
	if err != nil {
		return fmt.Errorf("failed to new tableflip: %v", err)
	}
	defer upg.Stop()
	k.upg = upg

	if k.isKeep {
		go k.keep()
	}
	if err := upg.Ready(); err != nil {
		return fmt.Errorf("failed to ready tableflip: %v", err)
	}
	<-upg.Exit()
	return nil
}

func (k *Keeper) Restart() error {
	if err := k.close(); err != nil {
		return fmt.Errorf("failed to close services: %v", err)
	}
	if err := k.upg.Upgrade(); err != nil {
		return fmt.Errorf("failed to upgrade process: %v", err)
	}
	return nil
}

func (k *Keeper) close() error {
	for _, closer := range k.closers {
		if err := closer.Close(); err != nil {
			log.Println(err)
		}
	}
	return nil
}

func (k *Keeper) keep() error {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP)
	for range sig {
		k.Restart()
	}
	return nil
}
