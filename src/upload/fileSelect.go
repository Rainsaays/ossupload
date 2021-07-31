package upload

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

var filePathMap = make(map[string]string, 21)

func GetFiles(dirPath, dirPathFixed string) (a *map[string]string) {
	files, err := ioutil.ReadDir(dirPath)
	//判断是否为文件
	if err != nil {
		//fmt.Printf("fileone is %v,name is%v\r\n", dirPath, path.Base(dirPath))
		filePathMap[dirPath] = path.Base(dirPath)
		return &filePathMap
		os.Exit(0)
	}
	for _, file := range files {
		//判断是否目录，如果是目录在次运行本函数取文件
		if file.IsDir() {
			//fmt.Printf("dir is %v \r\n", path.Join(dirPath, file.Name()))
			GetFiles(path.Join(dirPath, file.Name()), dirPathFixed)
		} else {
			//取目录路径的简短目录上传
			shorthandDir, err := filepath.Rel(dirPathFixed, path.Join(dirPath, file.Name()))
			if err != nil {
				panic(err)
			}

			filePathMap[path.Join(dirPath, file.Name())] = shorthandDir
			//fmt.Printf("file is %v,name is%v\r\n", path.Join(dirPath, file.Name()), path.Base(path.Join(dirPath, file.Name())))
		}
	}
	return &filePathMap

}
