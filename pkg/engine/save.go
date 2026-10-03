package engine

import (
	"os"

	"github.com/adrg/xdg"
)

func SaveGame(fileDir string, data []byte) error {
	saveDir, err := xdg.DataFile(fileDir)
	if err != nil {
		return err
	}

	return os.WriteFile(saveDir, data, 0644)
}

func LoadGame(fileDir string) ([]byte, error) {
	saveDir, err := xdg.DataFile(fileDir)
	if err != nil {
		return nil, err
	}

	return os.ReadFile(saveDir)
}
