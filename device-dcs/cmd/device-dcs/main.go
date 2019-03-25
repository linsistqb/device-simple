// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2017-2018 Canonical Ltd
// Copyright (C) 2018 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

// This package provides a simple example of a device service.
package main

import (
	"github.com/edgexfoundry/device-dcs/driver"
	"github.com/edgexfoundry/device-dcs/pkg/startup"
)

const (
	version     string = "1.0.0"
	serviceName string = "device-dcs"
)

func main() {
	sd := driver.SimpleDriver{}
	startup.Bootstrap(serviceName, version, &sd)
}
