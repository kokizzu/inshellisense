// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["time"] = model.Subcommand{
		Name:        []string{"time"},
		Description: `Time how long a command takes!`,
		Args: []model.Arg{{
			IsCommand: true,
		}},
	}
}