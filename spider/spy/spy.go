package spy

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func GetUrlPage(page int) string {
	return "https://tieba.baidu.com/f?kw=dota2&ie=utf-8&cid=&tab=corearea&pn=" + strconv.Itoa(50*page)
}
func GetFileName(url string) string {
	stringSlice := strings.SplitAfter(url, "/")
	return stringSlice[len(stringSlice)-1]
}
func GetToWork(url string, compileString string) {
	var ch = make(chan int)

	html, err := Spy(url)
	if err != nil {
		fmt.Println(err)
	}

	imageList := GetImageList(html, compileString)
	for i, v := range imageList {
		go ImageHandler(i, ch, v)
	}

	for i:=1;i<=len(imageList);i++{
		fmt.Println(i,":", <-ch, " goroutine finish")
	}
	return
}


func Spy(url string) (html string, err error){
	resp, tempErr := http.Get(url)
	if tempErr != nil {
		err = tempErr
		return
	}
	defer resp.Body.Close()
	res := make([]byte, 1024)
	for {
		num, tempErr := resp.Body.Read(res)
		if tempErr != nil && tempErr != io.EOF {
			err = tempErr
			return
		}
		if num == 0 {
			break
		}
		html += string(res[:num])
	}
	return
}

func GetImageList(html ,compileString string) []string {
	reg := regexp.MustCompile(compileString)
	var result []string
	objSlice := reg.FindAllStringSubmatch(html, -1)
	for _, v := range objSlice {
		result = append(result, v[1])
	}

	fmt.Println("num:", len(result))
	return result
}

func ImageHandler(num int, ch chan int, url string) (err error){
	res, tempErr := Spy(url)
	if tempErr != nil {
		err = tempErr
		return
	}


	filePath := ".\\imageDownload\\" + GetFileName(url)
	file , tempErr := os.Create(filePath)
	defer file.Close()
	if tempErr != nil {
		err = tempErr
		return
	}

	file.Write([]byte(res))
	ch <- num+1
	return
}