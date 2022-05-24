package main

import (
	a "github.com/brxyxn/go_gpr_nclouds/backend/internal"
	u "github.com/brxyxn/go_gpr_nclouds/backend/utils"
)

func main() {
	a := a.App{}

	a.L = u.InitLogs("nclouds-api ")

	a.Setup()

	a.Run()
}
