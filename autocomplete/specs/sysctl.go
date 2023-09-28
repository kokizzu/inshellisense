// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["sysctl"] = model.Subcommand{
		Name:        []string{"sysctl"},
		Description: `Get or set kernel state`,
		Args: []model.Arg{{
			Name:       "Variable names (and values if available)",
			Generator:  nil, // TODO: port over generator
			IsOptional: true,
			IsVariadic: true,
		}},
		Options: []model.Option{{
			Name:        []string{"-A", "-a"},
			Description: `List all the currently available non-opaque values`,
		}, {
			Name:        []string{"-b"},
			Description: `Force the value of the variable(s) to be output in raw, binary format`,
		}, {
			Name:        []string{"-d"},
			Description: `Print the description of the variable instead of its value`,
		}, {
			Name:        []string{"-e"},
			Description: `Separate the name and the value of the variable(s) with '='`,
			Args: []model.Arg{{
				Name:       "variable",
				IsVariadic: true,
			}},
		}, {
			Name:        []string{"-h"},
			Description: `Format output for human, rather than machine, readability`,
		}, {
			Name:        []string{"-i"},
			Description: `Ignore unknown OIDs`,
		}, {
			Name:        []string{"-N"},
			Description: `Show only variable names, not their values`,
		}, {
			Name:        []string{"-n"},
			Description: `Show only variable values, not their names`,
		}, {
			Name:        []string{"-o"},
			Description: `Show opaque variables (which are normally suppressed)`,
		}, {
			Name:        []string{"-q"},
			Description: `Suppress some warnings generated by sysctl to standard error`,
		}, {
			Name:        []string{"-X"},
			Description: `Equivalent to -x -a (for compatibility)`,
		}, {
			Name:        []string{"-x"},
			Description: `As -o, but prints a hex dump of the entire value instead of just the first few bytes`,
		}},
	}
}