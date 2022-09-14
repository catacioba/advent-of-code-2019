package ch23

import (
	"aoc/intcode"
	"aoc/util"
	"errors"
	"fmt"
	"time"
)

var errTimedOut = errors.New("channel read timed out")

func readWithTimeout[T any](ch chan T) (T, error) {
	select {
	case x := <-ch:
		{
			return x, nil
		}
	case <-time.After(time.Millisecond):
		return *new(T), errTimedOut
	}
}

func PartOne() {
	var programs [50]*intcode.Program

	for idx := 0; idx < 50; idx++ {
		programs[idx] = intcode.NewIntCodeProgramFromFile("ch23/input.txt")
		programs[idx].Input <- idx
		go programs[idx].Run()
	}

	for {
		for idx := 0; idx < 50; idx++ {
			p := programs[idx]

			dst, err := readWithTimeout(p.Output)
			if err != errTimedOut {
				x := <-p.Output
				y := <-p.Output

				fmt.Printf("Sending (%d, %d) to %d\n", x, y, dst)

				if dst == 255 {
					fmt.Printf("%d\n", y)
					return
				}
				programs[dst].Input <- x
				programs[dst].Input <- y
			}
			p.Input <- -1
		}
	}
}

func PartTwo() {
	var programs [50]*intcode.Program

	for idx := 0; idx < 50; idx++ {
		programs[idx] = intcode.NewIntCodeProgramFromFile("ch23/input.txt")
		programs[idx].Input <- idx
		go programs[idx].Run()
	}

	var nat util.Point
	previousNatY := -1

	idleComputers := make(map[int]int)
	for idx := 0; idx < 50; idx++ {
		idleComputers[idx] = 0
	}

	queues := make([]chan util.Point, 50)
	for idx := 0; idx < 50; idx++ {
		queues[idx] = make(chan util.Point, 100)
	}

	for {
		for idx := 0; idx < 50; idx++ {
			p := programs[idx]

			dst, receiveErr := readWithTimeout(p.Output)
			if receiveErr != errTimedOut {
				x := <-p.Output
				y := <-p.Output

				//fmt.Printf("Sending (%d, %d) to %d\n", x, y, dst)

				if dst == 255 {
					fmt.Println("Packet to NAT")
					nat = util.Point{X: x, Y: y}
				} else {
					queues[dst] <- util.Point{X: x, Y: y}
				}
			}
			in, sendErr := readWithTimeout(queues[idx])
			if sendErr != errTimedOut {
				p.Input <- in.X
				p.Input <- in.Y
			} else {
				p.Input <- -1
			}

			if sendErr == errTimedOut {
				idleComputers[idx] += 1
			} else {
				idleComputers[idx] = 0
			}
		}

		allIdle := true
		for _, v := range idleComputers {
			allIdle = allIdle && (v > 2)
		}

		if allIdle {
			fmt.Println("ALL IDLE")
			if previousNatY == nat.Y {
				fmt.Println(nat.Y)
				return
			}
			programs[0].Input <- nat.X
			programs[0].Input <- nat.Y
			previousNatY = nat.Y
		}
	}
}
