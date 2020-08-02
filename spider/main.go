package pa

import "myproject/pa/spy"

func main() {
	for i:=0;i<5;i++{
		url := spy.GetUrlPage(i)
		spy.GetToWork(url, `data-original="(?s:(.*?))"`)
	}
}

