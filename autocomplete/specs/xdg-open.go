// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["xdg-open"] = model.Subcommand{
		Name:        []string{"xdg-open"},
		Description: `Opens a file or URL in the user's preferred application`,
		Args: []model.Arg{{
			Templates: []model.Template{model.TemplateFilepaths},
			Name:      "FILE or URL",
		}},
		Options: []model.Option{{
			Name:        []string{"--help"},
			Description: `Show command synopsis`,
		}, {
			Name:        []string{"--manual"},
			Description: `Show manual page`,
		}, {
			Name:        []string{"--version"},
			Description: `Show the xdg-utils version information`,
		}},
	}
}