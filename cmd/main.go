package main

import (
	"github.com/andrelopes-code/simple-picpay-golang/cfg"
	"github.com/andrelopes-code/simple-picpay-golang/router"
)

// @title Simple PicPay API
func main() {
	cfg.Init()
	router.Init()
}
