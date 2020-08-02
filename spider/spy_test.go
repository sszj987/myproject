package pa

import (
	"myproject/pa/spy"
	"testing"
)

func TestSpy(t *testing.T) {
	for i:=0;i<5;i++{
		url := spy.GetUrlPage(i)
		spy.GetToWork(url, `data-original="(?s:(.*?))"`)
	}
}