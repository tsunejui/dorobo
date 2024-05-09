package update

import (
	"fmt"
	"os"

	"github.com/minio/selfupdate"
)

type Updater struct {
}

func New() *Updater {
	return &Updater{}
}

func (u *Updater) RUN(binary *os.File) error {
	if err := selfupdate.Apply(binary, selfupdate.Options{}); err != nil {
		return fmt.Errorf("failed to update self process, %v", err)
	}
	return nil
}
