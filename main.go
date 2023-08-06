package main

import (
	"github.com/hugplus/go-walker/cmd"
)

//go:generate swag init --parseDependency --parseDepth=6

// @title go-walker API
// @version 2.0.0
// @description 一个简单的脚手架
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}
