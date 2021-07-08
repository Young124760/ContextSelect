package writefile

import (
	"bufio"
	"fmt"
	"os"
)

func Writethinf(str string,filepath string){
	//选择操作类型，append插入
	file,err := os.OpenFile(filepath,os.O_WRONLY | os.O_APPEND,0666)
	if err != nil{
		fmt.Println("打开文件错误：",err)
		return
	}
	//关闭file句柄
	defer file.Close()
	//写入str
	writer := bufio.NewWriter(file)
	writer.WriteString(str)
	writer.Flush()
}