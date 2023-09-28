// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["arch"] = model.Subcommand{
		Name:        []string{"arch"},
		Description: `Print architecture type or run select architecture`,
		Args: []model.Arg{{
			Templates:  []model.Template{model.TemplateFilepaths},
			Name:       "program",
			IsCommand:  true,
			IsVariadic: true,
		}},
		Options: []model.Option{{
			Name:        []string{"-32"},
			Description: `Add the native 32-bit architecture to the list of architectures`,
		}, {
			Name:        []string{"-64"},
			Description: `Add the native 64-bit architecture to the list of architectures`,
		}, {
			Name: []string{"-arch"},
			Args: []model.Arg{{
				Name: "arch_name",
				Suggestions: []model.Suggestion{{
					Name:        []string{`i386`},
					Description: `32-bit intel`,
				}, {
					Name:        []string{`x86_64`},
					Description: `64-bit intel`,
				}, {
					Name:        []string{`x86_64h`},
					Description: `64-bit intel (haswell)`,
				}, {
					Name:        []string{`arm64`},
					Description: `64-bit arm`,
				}, {
					Name:        []string{`arm64e`},
					Description: `64-bit arm (Apple Silicon)`,
				}},
			}},
			ExclusiveOn: []string{"-i386", "-x86_64", "-x86_64h", "-arm64", "-arm64e"},
		}, {
			Name:        []string{"-i386"},
			Description: `32-bit intel`,
			ExclusiveOn: []string{"-arch"},
		}, {
			Name:        []string{"-x86_64"},
			Description: `64-bit intel`,
			ExclusiveOn: []string{"-arch"},
		}, {
			Name:        []string{"-x86_64h"},
			Description: `64-bit intel (haswell)`,
			ExclusiveOn: []string{"-arch"},
		}, {
			Name:        []string{"-arm64"},
			Description: `64-bit arm`,
			ExclusiveOn: []string{"-arch"},
		}, {
			Name:        []string{"-arm64e"},
			Description: `64-bit arm (Apple Silicon)`,
			ExclusiveOn: []string{"-arch"},
		}, {
			Name:        []string{"-c"},
			Description: `Clear the environment that will be passed to the command`,
		}, {
			Name:        []string{"-d"},
			Description: `Delete the named environment variable from the command's environment`,
			Args: []model.Arg{{
				Name: "envname",
			}},
		}, {
			Name:        []string{"-e"},
			Description: `Assign the given value to the variable in the command's environment`,
			Args: []model.Arg{{
				Name: "envname=value",
			}},
		}, {
			Name:        []string{"-h"},
			Description: `Print a help message and exit`,
		}},
	}
}