package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// Make a channel as mutex
var c = make(chan int, 1)
var global_num int = 0

func calculate(n int, start int, wg *sync.WaitGroup) {
	var sum int = 0
	var x, y float64
	for i := start; i < start+n; i++ {
		//random source create
		source := rand.NewSource(time.Now().UnixNano())
		generator := rand.New(source)
		y = generator.Float64()
		x = generator.Float64()
		if (x*x + y*y) <= 1 {
			sum++
		}
	}
	<-c //lock
	global_num += sum
	c <- 1 //unlock
	wg.Done()
}
func main() {
	// get the step number
	arg_num := len(os.Args)
	if arg_num < 2 {
		fmt.Println("please enter a number of run step, ex, ./pi 100000")
		return
	}

	// wait all routine finish
	var wg sync.WaitGroup

	cpu_num := runtime.NumCPU()

	wg.Add(cpu_num)
	step, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	run_step := step / cpu_num

	c <- 1
	for i := 0; i < cpu_num; i++ {
		go calculate(run_step, i*run_step, &wg)
	}
	wg.Wait()
	//time.Sleep(time.Second * 5)
	fmt.Printf("pi = %.7f\n", 4.0*float64(global_num)/float64(step))
}
