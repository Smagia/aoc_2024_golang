package day6

type Race struct {
	Time   int
	Record int
}

func NewRace(time int, record int) *Race {
	return &Race{
		Time:   time,
		Record: record,
	}
}

type Boat struct {
	Name       string
	TimesToWin []int
}

func NewBoat(name string) *Boat {
	return &Boat{
		Name:       name,
		TimesToWin: make([]int, 0),
	}
}

func (boat *Boat) RunBoat(race *Race) {
	for i := 0; i < race.Time; i++ {
		remainingTime := race.Time - i
		boatDistance := i * remainingTime
		if boatDistance > race.Record {
			boat.TimesToWin = append(boat.TimesToWin, i)
		}
	}
}
