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
	"os/user"

	"github.com/dpvs/govs"
	"github.com/yubo/gotool/flags"
)

var (
	EACCES = errors.New("Permission denied (you must be root)")
	ECONN  = errors.New("cannot connection to dpvs server")
)

func main() {
	//	flags.Parse()
	handler()

	usr, err := user.Current()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if usr.Uid != "0" {
		fmt.Println(EACCES.Error())
		return
	}

	cmd := flags.Cmd
	if cmd.Action != nil {
		err := govs.Vs_dial()
		if err != nil {
			fmt.Println(ECONN.Error())
			return
		}

		defer govs.Vs_close()
		cmd.Action(&govs.CallOptions{Opt: govs.CmdOpt,
			Args: flags.OthersCmd.Args()})
	} else {
		//flags.Usage()
		//	fmt.Println("error")
	}
	/*
		switch {
		case govs.FirstCmd.ADD:
			fmt.Println("add")
		default:
			fmt.Println("error!!!")
	*/

}
