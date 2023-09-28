// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["sw_vers"] = model.Subcommand{
		Name:        []string{"sw_vers"},
		Description: `Print macOS version information`,
		Options: []model.Option{{
			Name:        []string{"-productName"},
			Description: `Print product name`,
		}, {
			Name:        []string{"-productVersion"},
			Description: `Print product version`,
		}, {
			Name:        []string{"-buildVersion"},
			Description: `Print build version`,
		}},
	}
}