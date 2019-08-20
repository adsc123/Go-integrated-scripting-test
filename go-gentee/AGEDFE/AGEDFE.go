package main

import (
	"fmt"
	"github.com/gentee/gentee"
	"github.com/gentee/gentee/core"
	"log"
	"strconv"
	"strings"
)

func failOnError(err error, message string) {
	if err != nil {
		log.Panicf("%s - %s", err, message)
	}
}

func main() {
	// Fake data structure to showcase getting values from a database
	fakeDB := make(map[string]float64)
	fakeDB["ioutcome"] = 8
	fakeDB["age"] = 25
	fakeDB["totac1"] = 15
	fakeDB["acthr"] = 5
	fakeDB["everot"] = 6

	// Compiling an example function
	g := gentee.New()
	unitID, err := g.CompileFile("go-gentee/AGEDFE/AGEDFE.g")
	failOnError(err, "Compilation failure")

	// Getting all constant strings from source code
	names := g.Unit(unitID).Lexeme[0].Strings
	parameters := make([]string, 0)

	for _, name :=  range(names) {
		// For every expected parameter
		if strings.Contains(name, "--"){
			// Add parameter to slice
			parameter := strings.Trim(name, "--")
			parameters = append(parameters, name)
			// Get parameter value from 'database' and add to slice
			parValue := strconv.FormatFloat(fakeDB[parameter], 'f', -1, 64)
			parameters = append(parameters, parValue)
		}
	}

	// Pass data through to Gentee
	g.CmdLine(parameters...)
	// Run gentee code with the given data
	result, runErr := g.Run(unitID)
	failOnError(runErr, "Runtime error")

	// Get keys and values from return map
	returnMap := result.(*core.Map)
	returnValues := returnMap.Data
	// keys are the names of variables to save to
	keys := returnMap.Keys
	// for each key
	for _, key := range(keys){
		fmt.Println(key)
		// Convert value (may need to be done more cleverly than this but for the moment constant types will do)
		returnValue := float64(returnValues[key].(int64))
		// Set correct element in database to return value
		fakeDB[key] = returnValue
		fmt.Println(fakeDB[key])
	}
}