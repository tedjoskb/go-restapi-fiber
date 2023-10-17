package config

import (
	"path/filepath"
	"runtime"
)

var (

	// Get current file full path from runtime
	_, b, _, _ = runtime.Caller(0)

	//root folder of this project
	ProjectRootPath = filepath.Join(filepath.Dir(b), "../")
)
