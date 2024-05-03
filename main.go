package main

import "github.com/wanminny/admin/cmd"

//go:generate go build -ldflags "-X github.com/wanminny/admin/cmd.VERSION=v0.1"

func main() {
	cmd.Start()
}
