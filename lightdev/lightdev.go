// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lightdev

import (
	"github.com/cybergarage/go-net-upnp/net/upnp"
	"fmt"
	"io/ioutil"
)

const (
	DefaultTarget = "Living"
	DefaultStatus = true
)

type LightDevice struct {
	*upnp.Device
	Target string
	Status bool
}

const descxml = "./description.xml"

func ReadFile(filename string) (string) {
	bytes, err := ioutil.ReadFile(filename)
	fmt.Printf("read byte :%s\n",bytes)
	if err != nil {
		return ""
	}

	return string(bytes)
}

func NewLightDevice() (*LightDevice, error) {

	desc := ReadFile(descxml)

	fmt.Printf("desc file : %s\n",desc)
	dev, err := upnp.NewDeviceFromDescription(desc)
	if err != nil {
		return nil, err
	}

	lightDev := &LightDevice{
		Device: dev,
		Target: DefaultTarget,
		Status: DefaultStatus,
	}

	pHttpHandler := NewHttpHandler()
	dev.HTTPListener = *pHttpHandler

	return lightDev, nil
}

func (self *LightDevice) ActionRequestReceived(action *upnp.Action) upnp.Error {

	fmt.Printf("ActionRequestReceived : %+v\n",*action)
	return upnp.NewErrorFromCode(upnp.ErrorOptionalActionNotImplemented)
}
