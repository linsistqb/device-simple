// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018 Canonical Ltd
// Copyright (C) 2018 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

// This package provides a simple example implementation of
// a ProtocolDriver interface.
//
package driver

import (
	"fmt"
	dsModels "github.com/edgexfoundry/device-dcs/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/logging"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
	"time"
)

type SimpleDriver struct {
	lc           logger.LoggingClient
	asyncCh      chan<- *dsModels.AsyncValues
	randomDevices map[string]*randomDevice
}

// DisconnectDevice handles protocol-specific cleanup when a device
// is removed.
func (s *SimpleDriver) DisconnectDevice(deviceName string, protocols map[string]models.ProtocolProperties) error {
	return nil
}

// Initialize performs protocol-specific initialization for the device
// service.
func (s *SimpleDriver) Initialize(lc logger.LoggingClient, asyncCh chan<- *dsModels.AsyncValues) error {
	s.lc = lc
	s.asyncCh = asyncCh
	s.randomDevices = make(map[string]*randomDevice)
	return nil
}

// HandleReadCommands triggers a protocol Read operation for the specified device.
func (s *SimpleDriver) HandleReadCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []dsModels.CommandRequest) (res []*dsModels.CommandValue, err error) {

	rd, ok := s.randomDevices[deviceName]
	if !ok {
		rd = newRandomDevice()
		s.randomDevices[deviceName] = rd
	}

	res = make([]*dsModels.CommandValue, len(reqs))
	now := time.Now().UnixNano() / int64(time.Millisecond)

	for i, req := range reqs {
		t := req.DeviceResource.Properties.Value.Type
		v, err := rd.value(t)
		if err != nil {
			return nil, err
		}
		var cv *dsModels.CommandValue
		switch t {
		case "Int8":
			cv, _ = dsModels.NewInt8Value(&reqs[i].RO, now, int8(v))
		case "Int16":
			cv, _ = dsModels.NewInt16Value(&reqs[i].RO, now, int16(v))
		case "Int32":
			cv, _ = dsModels.NewInt32Value(&reqs[i].RO, now, int32(v))
		}
		res[i] = cv
	}

	return res, nil
}

// HandleWriteCommands passes a slice of CommandRequest struct each representing
// a ResourceOperation for a specific device resource.
// Since the commands are actuation commands, params provide parameters for the individual
// command.
func (s *SimpleDriver) HandleWriteCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []dsModels.CommandRequest,
	params []*dsModels.CommandValue) error {

	rd, ok := s.randomDevices[deviceName]
	if !ok {
		rd = newRandomDevice()
		s.randomDevices[deviceName] = rd
	}

	for _, param := range params {
		switch param.RO.Object {
		case "Min_Int8":
			v, err := param.Int8Value()
			if err != nil {
				return fmt.Errorf("RandomDriver.HandleWriteCommands: %v", err)
			}
			if v < defMinInt8 {
				return fmt.Errorf("RandomDriver.HandleWriteCommands: minimum value %d of %T must be int between %d ~ %d", v, v, defMinInt8, defMaxInt8)
			}

			rd.minInt8 = int64(v)
		case "Max_Int8":
			v, err := param.Int8Value()
			if err != nil {
				return fmt.Errorf("RandomDriver.HandleWriteCommands: %v", err)
			}
			if v > defMaxInt8 {
				return fmt.Errorf("RandomDriver.HandleWriteCommands: maximum value %d of %T must be int between %d ~ %d", v, v, defMinInt8, defMaxInt8)
			}

			rd.maxInt8 = int64(v)
		case "Min_Int16":
			v, err := param.Int16Value()
			if err != nil {
				return fmt.Errorf("RandomDriver.HandleWriteCommands: %v", err)
			}
			if v < defMinInt16 {
				return fmt.Errorf("RandomDriver.HandleWriteCommands: minimum value %d of %T must be int between %d ~ %d", v, v, defMinInt16, defMaxInt16)
			}

			rd.minInt16 = int64(v)
		case "Max_Int16":
			v, err := param.Int16Value()
			if err != nil {
				return fmt.Errorf("RandomDriver.HandleWriteCommands: %v", err)
			}
			if v > defMaxInt16 {
				return fmt.Errorf("RandomDriver.HandleWriteCommands: maximum value %d of %T must be int between %d ~ %d", v, v, defMinInt16, defMaxInt16)
			}

			rd.maxInt16 = int64(v)
		case "Min_Int32":
			v, err := param.Int32Value()
			if err != nil {
				return fmt.Errorf("RandomDriver.HandleWriteCommands: %v", err)
			}
			if v < defMinInt32 {
				return fmt.Errorf("RandomDriver.HandleWriteCommands: minimum value %d of %T must be int between %d ~ %d", v, v, defMinInt32, defMaxInt32)
			}

			rd.minInt32 = int64(v)
		case "Max_Int32":
			v, err := param.Int32Value()
			if err != nil {
				return fmt.Errorf("RandomDriver.HandleWriteCommands: %v", err)
			}
			if v > defMaxInt32 {
				return fmt.Errorf("RandomDriver.HandleWriteCommands: maximum value %d of %T must be int between %d ~ %d", v, v, defMinInt32, defMaxInt32)
			}

			rd.maxInt32 = int64(v)
		default:
			return fmt.Errorf("RandomDriver.HandleWriteCommands: there is no matched device resource for %s", param.String())
		}
	}

	return nil
}

// Stop the protocol-specific DS code to shutdown gracefully, or
// if the force parameter is 'true', immediately. The driver is responsible
// for closing any in-use channels, including the channel used to send async
// readings (if supported).
func (s *SimpleDriver) Stop(force bool) error {
	s.lc.Debug(fmt.Sprintf("SimpleDriver.Stop called: force=%v", force))
	return nil
}
