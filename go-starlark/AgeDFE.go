package main

import (
	"fmt"
	"github.com/starlight-go/starlight"
	"go.starlark.net/resolve"
	"go.starlark.net/starlark"
	"log"
)


func failOnError(err error, message string) {
	if err != nil {
		log.Panicf("%s - %s", err, message)
	}
}

func plainStarlark() {
	fakeDB := make(map[string]float64)
	fakeDB["ioutcome"] = 8
	fakeDB["age"] = 25
	fakeDB["totac1"] = 15
	fakeDB["acthr"] = 5
	fakeDB["everot"] = 6

	// Map of all global variables can be created automatically from the given database

	globals := starlark.StringDict{
		"ioutcome": starlark.Float(fakeDB["ioutcome"]),
		"age": starlark.Float(fakeDB["age"]),
		"totac1": starlark.Float(fakeDB["totac1"]),
		"acthr": starlark.Float(fakeDB["acthr"]),
		"everot":  starlark.Float(fakeDB["everot"]),
		"notinthere": starlark.Float(12.45),
	}

	// By default starlark doesn't allow floating point values
	resolve.AllowFloat = true
	name := "go-starlark/AgeDFE"
	slThread := &starlark.Thread{Name:"Thread1"}
	globals, err := starlark.ExecFile(slThread, name + ".star",nil,globals)
	failOnError(err,"Compilation error")
	funcName := globals[name]
	v ,err1 := starlark.Call(slThread, funcName, nil, nil)
	valDict := v.(*starlark.Dict)
	failOnError(err1,"Run time error")

	for _, k := range valDict.Keys() {
		key := k.(starlark.String).GoString()
		value, _, err  := valDict.Get(starlark.String(key))
		failOnError(err, "Not in dictionary")
		valInt, ok := value.(starlark.Int).Int64()
		if ok {
			fakeDB[key] = float64(valInt)
		}
	}
	fmt.Println(fakeDB["bacthr"])
}

func main() {
	fakeDB := make(map[string]float64)
	fakeDB["ioutcome"] = 8
	fakeDB["age"] = 25
	fakeDB["totac1"] = 15
	fakeDB["acthr"] = 5
	fakeDB["everot"] = 6

	// Map of all global variables can be created automatically from the given database

	globals := map[string]interface{}{
		"ioutcome": fakeDB["ioutcome"],
		"age": fakeDB["age"],
		"totac1": fakeDB["totac1"],
		"acthr": fakeDB["acthr"],
		"everot":  fakeDB["everot"],
		"notinthere": 12.45,
	}

	name := "AgeDFE"
	sl := starlight.New("go-starlark")
	globalsRet, err := sl.Run(name + ".star", globals)
	failOnError(err,"Compilation error")
	funcName := globalsRet[name].(*starlark.Function)
	slThread := &starlark.Thread{Name:"Thread1"}
	v ,err1 := starlark.Call(slThread, funcName, nil, nil)
	valDict := v.(*starlark.Dict)
	failOnError(err1,"Run time error")

	for _, k := range valDict.Keys() {
		key := k.(starlark.String).GoString()
		value, _, err  := valDict.Get(starlark.String(key))
		failOnError(err, "Not in dictionary")
		valInt, ok := value.(starlark.Int).Int64()
		if ok {
			fakeDB[key] = float64(valInt)
		}
	}
	fmt.Println(fakeDB["bacthr"])
}