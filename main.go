package main

import (
	"flag"
	"fmt"
	"github.com/Opelord/jenkins_test/task4"
)

func main() {
	fileName := flag.String("f", "example_data", "name of the file containing")
	flag.Parse()

	fmt.Println(task4.SolvePart1(*fileName))
}
