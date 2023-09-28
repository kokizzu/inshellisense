// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["iconv"] = model.Subcommand{
		Name:        []string{"iconv"},
		Description: `Character set conversion`,
		Args: []model.Arg{{
			Templates:  []model.Template{model.TemplateFilepaths},
			Name:       "inputfile",
			IsVariadic: true,
		}},
		Options: []model.Option{{
			Name:        []string{"--help"},
			Description: `Show help for iconv`,
		}, {
			Name:        []string{"--version"},
			Description: `Output version information and exit`,
		}, {
			Name:        []string{"-f", "--from-code"},
			Description: `Specifies the encoding of the input`,
			Args: []model.Arg{{
				Name:      "encoding",
				Generator: nil, // TODO: port over generator
			}},
			ExclusiveOn: []string{"-l", "--list"},
		}, {
			Name:        []string{"-t", "--to-code"},
			Description: `Specifies the encoding of the output`,
			Args: []model.Arg{{
				Name:      "encoding",
				Generator: nil, // TODO: port over generator
			}},
			ExclusiveOn: []string{"-l", "--list"},
		}, {
			Name:        []string{"-c"},
			Description: `Discard unconvertible characters`,
			ExclusiveOn: []string{"-l", "--list"},
		}, {
			Name:        []string{"-l", "--list"},
			Description: `List the supported encodings`,
			ExclusiveOn: []string{"-f", "--from-code", "-t", "--to-code", "--unicode-subst", "--byte-subst", "--widechar-subst"},
		}, {
			Name:        []string{"--unicode-subst"},
			Description: `Substitution for unconvertible Unicode characters`,
			Args: []model.Arg{{
				Name:        "FORMATSTRING",
				Description: `The formatstring must be a format string in the same format as for the printf command`,
			}},
			ExclusiveOn: []string{"-l", "--list"},
		}, {
			Name:        []string{"--byte-subst"},
			Description: `Substitution for unconvertible bytes`,
			Args: []model.Arg{{
				Name:        "FORMATSTRING",
				Description: `The formatstring must be a format string in the same format as for the printf command`,
			}},
			ExclusiveOn: []string{"-l", "--list"},
		}, {
			Name:        []string{"--widechar-subst"},
			Description: `Substitution for unconvertible wide characters`,
			Args: []model.Arg{{
				Name:        "FORMATSTRING",
				Description: `The formatstring must be a format string in the same format as for the printf command`,
			}},
			ExclusiveOn: []string{"-l", "--list"},
		}},
	}
}