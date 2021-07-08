package viewpak

import (
	"fmt"
	"io/ioutil"
	"log"
	"my_code/selectContext/constant"
	"my_code/selectContext/writefile"
)

func ListFiles(dirname string) {
	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil{
		log.Fatal(err)
	}
	filepath := constant.FilenamePath
	for _, fi := range fileInfos {
		filename := dirname + "\\" + fi.Name()
		fmt.Printf("%s\n",filename)
		writefile.Writethinf(filename,filepath)
		if fi.IsDir() {
			//继续遍历fi这个目录
			ListFiles(filename)
		}
	}
}
