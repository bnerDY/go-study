package main

import "flag"
import "fmt"
import "os"

type Cmd Struct {
	helpFlag bool
	versionFlag bool
	cpOption string
	class string
	args[] []string
}