package main

import (
	"fmt"
	"math/rand"
	"time"
)

var ones = []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}

var tens = []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func randomInt() int {
	// Generate a random integer between 0 and 99999999999
	return r.Intn(999999999999)
}

func main() {
	var numChannels int
	fmt.Print("How many tasks do you want: ")
	_, err := fmt.Scanf("%d", &numChannels)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	chans := make([]chan string, numChannels)
	startTime := time.Now()
	fmt.Println("start : ", startTime)
	for i := 0; i < numChannels; i++ {
		chans[i] = make(chan string)

		go setNumber(chans[i])

	}
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	/*for i := 0; i < len(chans); i++ {
		num := <-chans[i]
		fmt.Println("result of ", i, " is :", num)

	}*/
	fmt.Println("End: ", endTime)
	fmt.Println("Elapsed Time: ", elapsedTime)
}
func numberToString(number int) string {
	if number < 20 {
		return ones[number]
	} else if number < 100 {
		return tens[(number-(number%10))/10] + " " + numberToString(number%10)
	} else if number < 1_000 {
		return numberToString(number/100) + " hundred " + numberToString((number % 100))
	} else if number < 1_000_000 {
		return numberToString(number/1_000) + " thousand " + numberToString(number%1_000)
	} else if number < 1_000_000_000 {
		return numberToString(number/1_000_000) + " million " + numberToString(number%1_000_000)
	} else if number < 1_000_000_000_000 {
		return numberToString(number/1_000_000_000) + " Billion " + numberToString(number%1_000_000_000)
	} else {
		return "The number is too big"
	}
}

func setNumber(strChan chan string) {
	var number = randomInt()

	if number == 0 {
		fmt.Println("Zero")
		return
	}
	defer close(strChan)
	strChan <- numberToString(number)

}
