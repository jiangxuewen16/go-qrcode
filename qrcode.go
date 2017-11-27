package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/skip2/go-qrcode"
)

var inFile *string = flag.String("i", "", "二维码列表文件,例如：D:/test/qrcode.text")
var outFileDir *string = flag.String("o", "", "生成二维码输出目录,例如：D:/test/")
var size *int = flag.Int("s", 256, "二维码像素(px)")


func QR(url string) error {
	name := GetRandomString(32) + ".png"
	fileName := *outFileDir + "/" + name
	err := qrcode.WriteFile(url, qrcode.Medium, *size, fileName)
	if err != nil {
		fmt.Println("======转换二维码失败：")
		return err
	}
	return nil
}

//生成随机字符串
func GetRandomString(lens int) string{
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lens; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//读取文件的数据
func readValues(inFile string) (values []string, err error) {
	fmt.Println("======输入文件夹：",inFile)
	file, err := os.Open(inFile)
	if err != nil {
		fmt.Println("======打开文件失败：", inFile)
		return
	}
	defer file.Close()

	br := bufio.NewReader(file)

	values = make([]string, 0)

	for {
		line, isPrefix, err1 := br.ReadLine()

		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			return
		}

		if isPrefix {
			fmt.Println("======行数据太长，无法解析")
		}

		fmt.Println("======数据：",string(line))
		values = append(values, string(line))
	}
	return
}

func pathExists(path string) (bool, error) {
	fmt.Println("======输出文件夹:", path)
	_,err := os.Stat(path)
	if err != nil {
		return true,nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {

	flag.Parse()

	if inFile == nil || *inFile == "" {
		fmt.Println("请输入-i的参数,例如：-i [输入文件]")
		return
	}

	if outFileDir == nil || *outFileDir == "" {
		fmt.Println("请输入-o的参数,例如：-o [二维码的输出文件夹]")
		return
	}

	if boo,_ := pathExists(*outFileDir); boo {
		fmt.Println("文件夹不存在：",*outFileDir)
		return
	}

	values, err := readValues(*inFile)
	if err == nil {
		for _,value := range values {
			err := QR(value)
			if err != nil {
				fmt.Println("======转换二维码失败：",err)
			}
		}
	}

}
