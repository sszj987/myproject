package obsesrver

type update interface {
	updateTemperature(int)
}

type server struct {
	temperature int
	observer []update
}

func(s *server) addObserver(observer update) {
	s.observer = append(s.observer, observer)
}

func(s *server) updateTemperature() {
	for i := range s.observer {
		s.observer[i].updateTemperature(s.temperature)
	}
}

type observer struct {
	temperature int
}

func(o *observer) updateTemperature(input int) {
	o.temperature = input
}
