package static

import (
	"encoding/json"
	"ganxue-server/global"
	"ganxue-server/model/answer_model"
	"ganxue-server/model/md_model"
	"ganxue-server/utils/db/mongodb"
	"ganxue-server/utils/log"
	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/yaml.v3"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type Static map[string]struct {
	M md_model.Markdown
	A answer_model.Answer
}

func Init() {
	data := make(Static)
	golang, err := getAllFilesRecursive("./static/golang/")
	if err != nil {
		log.Error("获取文件失败", err)
	}
	for _, file := range golang {
		if err := getFileContent(&data, file); err != nil {
			log.Error("获取文件内容失败", err)
		}
	}
	c, err := getAllFilesRecursive("./static/C/")
	if err != nil {
		log.Error("获取文件失败", err)
	}
	for _, file := range c {
		if err := getFileContent(&data, file); err != nil {
			log.Error("获取文件内容失败", err)
		}
	}
	cpp, err := getAllFilesRecursive("./static/Cpp/")
	if err != nil {
		log.Error("获取文件失败", err)
	}
	for _, file := range cpp {
		if err := getFileContent(&data, file); err != nil {
			log.Error("获取文件内容失败", err)
		}
	}

	for _, v := range data {
		if v.M.ID != "" {
			if err := mongodb.Update(global.MD, bson.M{"id": v.M.ID}, bson.M{"$set": v.M}, true); err != nil {
				log.Error("插入数据库失败", err)
			}
		}
		if v.A.ID != "" {
			if err := mongodb.Update(global.ANSWER, bson.M{"id": v.A.ID}, bson.M{"$set": v.A}, true); err != nil {
				log.Error("插入数据库失败", err)
			}
		}
	}

	catalogueFile, err := os.ReadFile("./static/catalogue.yaml")
	if err != nil {
		log.Panic("read catalogue.yaml fail", err)
	}
	if err := yaml.Unmarshal(catalogueFile, &global.CATALOGUE); err != nil {
		log.Panic("unmarshal catalogue.yaml fail", err)
	}

}

// getAllFilesRecursive 获取所有文件
func getAllFilesRecursive(root string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			files = append(files, path) // 包含完整路径
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

// getFileContent 获取文件内容
func getFileContent(data *Static, path string) error {
	name, ext := fileType(path)
	val, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	t := (*data)[name]
	switch ext {
	case "md":
		//content := bf.ToHTML(val)
		t.M.Content = string(val)
		t.M.ID = name
	case "txt":
		t.M.Code = string(val)
		t.M.ID = name
	case "json":
		if err := json.Unmarshal(val, &t.A); err != nil {
			return err
		}
		t.A.ID = name
	}
	(*data)[name] = t
	return nil
}

// 文件分类 返回 上级目录名 文件名 文件扩展名
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
