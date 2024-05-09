package process

import (
	"dorobo/pkg/keep"
	"dorobo/pkg/update"
	"fmt"
	"os"
)

type Manager struct {
	keeper  *keep.Keeper
	updater *update.Updater
}

func New(keeper *keep.Keeper, updater *update.Updater) *Manager {
	return &Manager{
		keeper:  keeper,
		updater: updater,
	}
}

func (m *Manager) Upgrade(file *os.File) error {
	if err := m.updater.RUN(file); err != nil {
		return fmt.Errorf("failed to update binary file: %v", err)
	}

	if err := m.keeper.Restart(); err != nil {
		return fmt.Errorf("failed to restart process: %v", err)
	}
	return nil
}

func (m *Manager) RunKeeper() error {
	if err := m.keeper.SetKeep(true).Run(); err != nil {
		return fmt.Errorf("failed to run keeper: %v", err)
	}
	return nil
}
