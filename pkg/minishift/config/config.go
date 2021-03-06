/*
Copyright (C) 2017 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var Config *InstanceConfig

type InstanceConfig struct {
	FilePath string `json:"-"`
	OcPath   string
}

// Create new object with data if file exists or
// Create json file and return object if doesn't exists
func NewInstanceConfig(path string) (*InstanceConfig, error) {
	cfg := new(InstanceConfig)
	cfg.FilePath = path

	// Check json file existence
	_, err := os.Stat(cfg.FilePath)
	if os.IsNotExist(err) {
		if errWrite := cfg.Write(); errWrite != nil {
			return nil, errWrite
		}
	} else {
		if errRead := cfg.read(); errRead != nil {
			return nil, errRead
		}
	}

	return cfg, nil
}

func (cfg *InstanceConfig) Write() error {
	jsonData, err := json.MarshalIndent(cfg, "", "\t")
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(cfg.FilePath, jsonData, 0644); err != nil {
		return err
	}

	return nil
}

func (cfg *InstanceConfig) Delete() error {
	if err := os.Remove(cfg.FilePath); err != nil {
		return err
	}

	return nil
}

func (cfg *InstanceConfig) read() error {
	raw, err := ioutil.ReadFile(cfg.FilePath)
	if err != nil {
		return err
	}

	json.Unmarshal(raw, &cfg)
	return nil
}
