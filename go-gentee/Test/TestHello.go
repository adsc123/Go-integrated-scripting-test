package main

import (
	"fmt"
	"github.com/gentee/gentee/core"
	"log"
	"strings"
	"sync"

	"github.com/gentee/gentee"
)

func failOnError(err error, message string) {
	if err != nil {
		log.Panicf("%s - %s", err, message)
	}
}

func main() {
	g := gentee.New()
	// dat, err := ioutil.ReadFile("TestHello.g")
	// failOnError(err, "Couldn't read file")
	// print(string(dat))

	// unitID, err := g.Compile(string(dat), "C:\\Users\\Cooka4\\Documents\\go-gentee\\a\\t")
	unitID, err := g.CompileFile("go-gentee/Test/TestHello.g")
	failOnError(err, "Failed to compile file")
	// print(unitID)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		names := g.Unit(unitID).Lexeme[0].Strings
		fmt.Println(names)
		g.CmdLine(``)
		result1, err1 := g.Run(unitID)
		failOnError(err1, "SomethingToo")
		resArr, ok := result1.(*core.Array)
		if ok {
			resArrString := resArr.String()
			resString := resArrString[1:len(resArrString) - 1]
			result := strings.Split(resString, " ")
			fmt.Println(result)
		}
		wg.Done()
	}()
	args := []string{`--name`,`Adam`,`-t`,`7`}
	g.CmdLine(args...)
	result, err := g.Run(unitID)
	wg.Wait()
	failOnError(err, "Something")

	fmt.Println(result)

}
