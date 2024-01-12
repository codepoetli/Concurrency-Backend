package files

import (
	"errors"
	"os"
)

// PathExists 判断是否存在路径
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if errors.Is(os.ErrNotExist, err) {
		return false, nil
	}
	return false, err
}
