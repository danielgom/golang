package exportlib

import "runtime"

func GetGoVersion() string {
	return runtime.Version()
}
