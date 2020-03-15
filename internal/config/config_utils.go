package config

import (
	"github.com/subosito/gotenv"
	"path/filepath"
	"runtime"
)

const envFileName = ".env"

// RootDir returns root dir of project
func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "../..")
}

// Load .env file in root path of module
func LoadEnv() error {
	envFile := filepath.Join(RootDir(), envFileName)
	return gotenv.Load(envFile)

} // Load .env file in root path of module
func MustLoadEnv() {
	err := LoadEnv()
	if err != nil {
		panic(err)
	}
}
