package static

import (
	"encoding/json"
	"ganxue-server/global"
	"ganxue-server/model/answer_model"
	"ganxue-server/utils/bf"
	"ganxue-server/utils/db/mongodb"
	"ganxue-server/utils/log"
	"github.com/fsnotify/fsnotify"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"path/filepath"
	"strings"
)

func Init() {
	go func() {
		var err error
		global.WATCHER, err = fsnotify.NewWatcher()
		if err != nil {
			log.Error("创建文件监听失败", err)
			return
		}

		defer func(WATCHER *fsnotify.Watcher) {
			err := WATCHER.Close()
			if err != nil {
				log.Error("关闭文件监听失败", err)
			}
		}(global.WATCHER)

		err = global.WATCHER.Add("./static/golang")
		if err != nil {
			log.Error("添加文件监听失败", err)
			return
		}
		for {
			select {
			case event := <-global.WATCHER.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Debug("文件修改：", event.Name)
					fileChange(event.Name)
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					log.Debug("文件创建：", event.Name)
					fileCreate(event.Name)
				}
				if event.Op&fsnotify.Remove == fsnotify.Remove {
					log.Debug("文件删除：", event.Name)
					fileDelete(event.Name)
				}
			case err := <-global.WATCHER.Errors:
				log.Error("文件监听错误：", err)
			}
		}
	}()
}

func fileChange(path string) {
	// TODO: 文件内容被修改
	name, ext := fileType(path)
	val, err := os.ReadFile(path)
	if err != nil {
		log.Error("获取文件信息失败：", err)
		return
	}
	switch ext {
	case "md":
		content := bf.ToHTML(val)
		_err := mongodb.Update(global.MD, bson.M{"id": name}, bson.M{"$set": bson.M{"content": string(content)}})
		if _err != nil {
			log.Error("更新数据库失败：", _err)
			return
		}
	case "txt":
		_err := mongodb.Update(global.MD, bson.M{"id": name}, bson.M{"$set": bson.M{"code": string(val)}})
		if _err != nil {
			log.Error("更新数据库失败：", _err)
			return
		}
	case "json":
		var as answer_model.Answer
		if err := json.Unmarshal(val, &as); err != nil {
			log.Error("解析文件失败：", err)
			return
		}
		_err := mongodb.Update(global.ANSWER, bson.M{"id": name}, bson.M{"$set": bson.M{"input": as.Input, "output": as.Output}})
		if _err != nil {
			log.Error("更新数据库失败：", _err)
			return
		}
	}
}

func fileCreate(path string) {
	// TODO: 文件创建
}

func fileDelete(path string) {
	// TODO: 文件删除
}

// 文件分类
func fileType(path string) (string, string) {
	if path == "" {
		return "", ""
	}

	filename := filepath.Base(path)
	// 分割文件名和扩展名
	fileNameParts := strings.Split(filename, ".")
	if len(fileNameParts) > 1 {
		// 获取文件扩展名
		ext := fileNameParts[len(fileNameParts)-1]
		// 获取文件名
		fileName := strings.Join(fileNameParts[:len(fileNameParts)-1], ".")
		return fileName, ext
	} else {
		return filename, ""
	}
}
