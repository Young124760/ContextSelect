package main

import (
	"bufio"
	"fmt"
	"io"
	"my_code/selectContext/constant"
	"my_code/selectContext/readfile"
	"my_code/selectContext/viewpak"
	"os"
	"time"
)

func main()  {
	//开始时间
	starttime := time.Now().UnixNano()
	dirname := constant.Dirname
	mark := constant.Mark
	viewpak.ListFiles(dirname)
	filenamepath := constant.FilenamePath
	file,err := os.Open(filenamepath)
	if err != nil{
		fmt.Println("打开文件错误：",err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for{
		str,err := reader.ReadString(mark)
		readfile.Readthinf(str)
		if err == io.EOF{
			break
		}
	}
	//计算所用时间
	endtime := time.Now().UnixNano()
	var totaltime float64
	totaltime = float64(endtime - starttime)/1000000000
	fmt.Printf("写入完毕,总共花时间为：%.2f 秒",totaltime)
}
