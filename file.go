package gokit

import "os"

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}

		return false
	}

	return !info.IsDir()
}
