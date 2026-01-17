package filemgr

import (
	"fmt"
	"os"
	"path/filepath"
)

func CurrentDayPath() (string, error) {
	dataDir, err := DataDir()
	if err != nil {
		return "", fmt.Errorf("failed to get current day path: %w", err)
	}
	return filepath.Join(dataDir, "currentday.json"), nil
}

func DataDir() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("failed to get executable path: %w", err)
	}

	exeDir := filepath.Dir(exePath)
	dataDir := filepath.Join(exeDir, "data")

	err = os.MkdirAll(dataDir, 0755)
	if err != nil {
		return "", fmt.Errorf("failed to create data directory: %w", err)
	}

	return dataDir, nil
}
