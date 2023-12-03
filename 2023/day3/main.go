package main

func panicerr(err error, msg string) {
	if err != nil {
		panic(msg + ": " + err.Error())
	}
}

func main() {
	day_three_part_two()
}
