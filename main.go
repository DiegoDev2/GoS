package main

import (
	"fmt"
	"gs/cli"
	"time"

	"github.com/briandowns/spinner"
)

func printBanner() {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Suffix = " Loading the GoS CLI..."
	s.Start()

	time.Sleep(2 * time.Second)
	s.Stop()

	printMascot()

	fmt.Printf("%s\n", "• Follow us on X: @diegodev5")
	fmt.Printf("%s\n", "• Don't forget to support the project")
}
func printMascot() {

	for i := 0; i < 3; i++ {
		fmt.Printf("\033[H\033[2J")
		if i%2 == 0 {
			fmt.Println("╭─────╮")
			fmt.Println("│ ◠ ◡ ◠ │  Let's create something unique!")
			fmt.Println("╰─────╯")
		} else {
			fmt.Println("╭─────╮")
			fmt.Println("│ > - <   What can we build today?")
			fmt.Println("╰─────╯")
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {

	printBanner()

	cli.Execute()
}
