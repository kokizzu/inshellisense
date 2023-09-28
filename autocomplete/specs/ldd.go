// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["ldd"] = model.Subcommand{
		Name:        []string{"ldd"},
		Description: `Print shared library dependencies`,
		Args: []model.Arg{{
			Templates:  []model.Template{model.TemplateFilepaths},
			Name:       "exe",
			IsVariadic: true,
		}},
		Options: []model.Option{{
			Name:        []string{"--help"},
			Description: `Show help for ldd`,
		}, {
			Name:        []string{"--version"},
			Description: `Print version information and exit`,
		}, {
			Name:        []string{"-d", "--data-relocs"},
			Description: `Process data relocations`,
		}, {
			Name:        []string{"-r", "--function-relocs"},
			Description: `Process data and function relocations`,
		}, {
			Name:        []string{"-u", "--unused"},
			Description: `Print unused direct dependencies`,
		}, {
			Name:        []string{"-v", "--verbose"},
			Description: `Print all information`,
		}},
	}
}