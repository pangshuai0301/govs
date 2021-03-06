/*
 * Copyright 2016 Xiaomi Corporation. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 *
 * Authors:    Yu Bo <yubo@xiaomi.com>
 */
package main

import (
	"errors"
	"fmt"
	"os"
	"os/user"

	"github.com/dpvs/govs"
)

var (
	EACCES = errors.New("Permission denied (you must be root)")
	ECONN  = errors.New("cannot connection to dpvs server")
)

func main() {

	usr, err := user.Current()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if usr.Uid != "0" {
		fmt.Println(EACCES.Error())
		os.Exit(1)
	}

	handler()

	cmd := Cmd
	if cmd.Action == nil {
		return
	}
	if err := govs.Vs_dial(); err != nil {
		fmt.Println(ECONN.Error())
		return
	}
	defer govs.Vs_close()
	err = cmd.Action(&govs.CallOptions{Opt: govs.CmdOpt})
	if err != nil {
		fmt.Println(err)
		govs.Vs_close()
		if err.Error() != "worker or resource busy" {
			os.Exit(1)
		}
	}
}
