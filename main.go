// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
lightdev is a sample implementation of UPnP standard device, BinaryLight:1.

        NAME
        lightdev

        SYNOPSIS
        lightdev [OPTIONS]

        DESCRIPTION
        lightdev is a sample implmentation of UPnP Standardized DCP, BinaryLight:1

        OPTIONS
        -v : *level* Enable verbose output.

        RETURN VALUE
          Return EXIT_SUCCESS or EXIT_FAILURE
*/
package main

import (
	"os"
	"ubeybox/lightdev"
	//"github.com/labstack/echo"
)

func main() {

	//e := echo.New()

	dev, err := lightdev.NewLightDevice()
	if err != nil {
		os.Exit(1)
	}

	err = dev.Start()
	if err != nil {
		os.Exit(1)
	}
	defer dev.Stop()

//	e.Logger.Fatal(e.Start(":10001"))

}
