package main

import (
	"cowait/core/cloud"
)

func main() {
	cloud := cloud.App()
	cloud.Run()
}
