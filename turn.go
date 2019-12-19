package turn

var x []string

func Push(str string)  {
	x = append(x ,str)
}

func Bob() string {

	var bus []string

	num := len(x)

	if num == 0{
		return ""
	}
	first := x[0]
	if num == 1 {
		x[0] = ""
		return first
	}

	for i:=1; i < num; i++ {
		bus = append(bus, x[i])
	}
	x = bus
	return first
}