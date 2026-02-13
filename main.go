package main

import (
	hi "github.com/golangllm/hello/service-hi"
	say "github.com/golangllm/hello/service-say"
)

func main() {
	say.Say()
	hi.Hello()
}
