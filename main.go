package main

import (
	say "github.com/golangllm/hello/service-say"
	hi "github.com/golangllm/service-hi"
)

func main() {
	say.Say()
	hi.Hello()
}
