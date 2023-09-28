// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["create-completion-spec"] = model.Subcommand{
		Name:        []string{"create-completion-spec"},
		Description: `Setup fig folder and create spec with the given name`,
		Args: []model.Arg{{
			Name: "name",
		}},
		Options: []model.Option{{
			Name:        []string{"--here"},
			Description: `Set if the spec should be created in the current folder`,
		}, {
			Name:        []string{"-h", "--help"},
			Description: `Display help for command`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"help"},
			Description: `Display help for command`,
			Args: []model.Arg{{
				Name:       "command",
				IsOptional: true,
			}},
		}},
	}
}