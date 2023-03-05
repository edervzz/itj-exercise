package main

import "itj-code-exercise/core"

func main() {

	addresses := []string{}
	names := []string{}

	addresses, names = core.Input()

	core.Process(addresses, names)

}
