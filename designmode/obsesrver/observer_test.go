package obsesrver

import (
	"fmt"
	"testing"
)

func TestObserver (t *testing.T) {
	server := new(server)

	observer1 := new(observer)
	server.addObserver(observer1)

	server.temperature = 100
	server.updateTemperature()
	fmt.Println(server.temperature == observer1.temperature)

	server.temperature = 200
	server.updateTemperature()
	fmt.Println(server.temperature == observer1.temperature)
}