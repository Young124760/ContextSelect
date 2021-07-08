package readfile

import (
	"bufio"
	"fmt"
	"io"
	"my_code/selectContext/constant"
	"my_code/selectContext/writefile"
	"os"
	"strings"
)
//逐行读取文件
func Readthinf(filename string)  {
	target := constant.Target
	filepath := constant.Filepath
	mark := constant.Mark
	//打开文件
	file,err := os.Open(filename)
	if err != nil{
		fmt.Println("打开文件错误：",err)
	}
	//defer关闭文件句柄
	defer file.Close()
	reader := bufio.NewReader(file)
	//逐行读取
	for{
		str,err := reader.ReadString(mark)
		Judge(str,target,filepath)
		if err == io.EOF{
			break
		}
	}
}

//判断每行是否包含指定字段，如果包含则写入
func Judge(str string,target string,filepath string)  {
	num := strings.Count(str,target)
	if num == 0{
		writefile.Writethinf(str,filepath)
	}
}