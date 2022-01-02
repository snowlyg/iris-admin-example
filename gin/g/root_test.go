package g

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

var goPath = os.Getenv("GOPATH")
var projectPath = strings.ToLower(filepath.ToSlash(filepath.Join(goPath, "src/github.com/snowlyg/iris-admin-example/gin")))

func Test_InitRootDir(t *testing.T) {
	t.Run("测试项目根目录", func(t *testing.T) {
		InitRootDir()
		if strings.ToLower(filepath.ToSlash(Root)) != projectPath {
			t.Errorf("Root path want '%s' and get %s", projectPath, Root)
		}
	})
}
func Test_GetCurrentAbPath(t *testing.T) {
	t.Run("测试当前路径", func(t *testing.T) {
		path, err := getCurrentAbPath()
		if err != nil {
			t.Error(err)
			return
		}
		path = strings.ToLower(filepath.ToSlash(path))
		want := filepath.Join(projectPath, "g")
		if path != want {
			t.Errorf("path want '%s' and get %s", want, path)
		}
	})
}
func Test_GetCurrentAbPathByExecutable(t *testing.T) {
	t.Run("测试当前执行文件目录", func(t *testing.T) {
		path, err := getCurrentAbPathByExecutable()
		if err != nil {
			t.Error(err)
			return
		}
		if !strings.Contains(path, "Temp") && runtime.GOOS == "windows" {
			t.Errorf("path '%s' not contain Temp", path)
		}
		if !strings.Contains(path, "T") && runtime.GOOS == "darwin" {
			t.Errorf("path '%s' not contain Temp", path)
		}

	})
}
func Test_GetCurrentAbPathByCaller(t *testing.T) {
	t.Run("测试当前方法执行目录", func(t *testing.T) {
		path := getCurrentAbPathByCaller()
		path = strings.ToLower(filepath.ToSlash(path))
		want := filepath.Join(projectPath, "g")
		if strings.ToLower(filepath.ToSlash(path)) != want {
			t.Errorf(" path want '%s' and get %s", want, path)
		}
	})
}
