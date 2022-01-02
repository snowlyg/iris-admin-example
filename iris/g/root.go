package g

import (
	"os"
	"path"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strings"

	"github.com/snowlyg/helper/str"
)

// InitRootDir 初始化程序执行路径
var (
	Root string
)

type Empty struct{}

// InitRootDir 初始化程序执行路径
func InitRootDir() error {
	var err error
	if Root, err = getCurrentAbPath(); err != nil {
		return err
	}
	// 获取当前包名，排除当前包名路径
	packageName := filepath.Base(reflect.TypeOf(Empty{}).PkgPath())
	Root = filepath.ToSlash(Root)
	if b, _ := regexp.MatchString(str.Join("/", packageName, "$"), Root); b {
		Root = filepath.Dir(Root)
	}
	if b, _ := regexp.MatchString("/tests$", Root); b {
		Root = filepath.Dir(Root)
	}
	return nil
}

// getCurrentAbPath 当前路径
func getCurrentAbPath() (string, error) {
	dir, err := getCurrentAbPathByExecutable()
	if err != nil {
		return "", err
	}
	tmpDir, err := filepath.EvalSymlinks(os.TempDir())
	if err != nil {
		return "", err
	}
	if strings.Contains(dir, tmpDir) {
		return getCurrentAbPathByCaller(), nil
	}
	return dir, nil
}

//getCurrentAbPathByExecutable 当前执行文件目录
func getCurrentAbPathByExecutable() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	res, err := filepath.EvalSymlinks(filepath.Dir(exePath))
	if err != nil {
		return "", err
	}
	return res, nil
}

//getCurrentAbPathByCaller 当前方法执行目录
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
