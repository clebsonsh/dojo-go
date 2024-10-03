package main

import "fmt"

// https://judge.beecrowd.com/pt/problems/view/1168
func main() {
	leds := map[string]int{
		"1": 2,
		"2": 5,
		"3": 5,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 3,
		"8": 7,
		"9": 6,
		"0": 6,
	}

	lines := 0
	fmt.Scanln(&lines)

	for i := 0; i < lines; i++ {
		var number string
		fmt.Scanln(&number)

		total := 0
		for _, n := range number {
			total += leds[string(n)]
		}

		fmt.Println(total, "leds")
	}
}
