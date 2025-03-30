package test

import (
	"ganxue-server/initialize"
	"os"
	"path/filepath"
	"testing"
)

func TestMain(m *testing.M) {
	// 获取项目根目录
	rootDir, err := filepath.Abs(filepath.Dir(filepath.Join("..", "config.yaml")))
	if err != nil {
		panic("无法获取项目根目录: " + err.Error())
	}

	// 设置工作目录为项目根目录
	err = os.Chdir(rootDir)
	if err != nil {
		panic("无法设置工作目录: " + err.Error())
	}

	initialize.InitAll()

	// 运行所有测试
	os.Exit(m.Run())
}
