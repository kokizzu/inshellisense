// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["elm-json"] = model.Subcommand{
		Name:        []string{"elm-json"},
		Description: `Deal with your elm.json`,
		Options: []model.Option{{
			Name:        []string{"--help", "-h"},
			Description: `Prints help information`,
		}, {
			Name:        []string{"--offline"},
			Description: `Enable offline mode, which means no HTTP traffic will happen`,
		}, {
			Name:        []string{"--version", "-V"},
			Description: `Prints version information`,
		}, {
			Name:        []string{"--verbose", "-v"},
			Description: `Sets the level of verbosity`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"help"},
			Description: `Prints help information or the help of the given subcommand(s)`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateHelp},
				Name:      "subcommand",
			}},
		}, {
			Name:        []string{"install"},
			Description: `Install a package`,
			Args: []model.Arg{{
				Name:        "PACKAGE",
				Description: `Package to install, e.g. elm/core or elm/core@1.0.2 or elm/core@1`,
				Generator:   nil, // TODO: port over generator
			}, {
				Name:        "-- INPUT",
				Description: `The elm.json file to upgrade [default: elm.json]`,
				Generator:   nil, // TODO: port over generator
				IsOptional:  true,
			}},
			Options: []model.Option{{
				Name:        []string{"--help", "-h"},
				Description: `Prints help information`,
			}, {
				Name:        []string{"--test"},
				Description: `Install as a test-dependency`,
			}, {
				Name:        []string{"--version", "-V"},
				Description: `Prints version information`,
			}, {
				Name:        []string{"--yes"},
				Description: `Answer "yes" to all questions`,
			}},
		}, {
			Name:        []string{"new"},
			Description: `Create a new elm.json file`,
			Options: []model.Option{{
				Name:        []string{"--help", "-h"},
				Description: `Prints help information`,
			}, {
				Name:        []string{"--version", "-V"},
				Description: `Prints version information`,
			}},
		}, {
			Name:        []string{"tree"},
			Description: `List entire dependency graph as a tree`,
			Args: []model.Arg{{
				Name:        "PACKAGE",
				Description: `Limit output to show path to some (indirect) dependency`,
				Generator:   nil, // TODO: port over generator
			}, {
				Name:        "-- INPUT",
				Description: `The elm.json file to solve [default: elm.json]`,
				Generator:   nil, // TODO: port over generator
				IsOptional:  true,
			}},
			Options: []model.Option{{
				Name:        []string{"--help", "-h"},
				Description: `Prints help information`,
			}, {
				Name:        []string{"--test"},
				Description: `Promote test-dependencies to top-level dependencies`,
			}, {
				Name:        []string{"--version", "-V"},
				Description: `Prints version information`,
			}},
		}, {
			Name:        []string{"uninstall"},
			Description: `Uninstall a package`,
			Args: []model.Arg{{
				Name:        "PACKAGE",
				Description: `Package to uninstall, e.g. elm/html`,
				Generator:   nil, // TODO: port over generator
			}, {
				Name:        "-- INPUT",
				Description: `The elm.json file to upgrade [default: elm.json]`,
				Generator:   nil, // TODO: port over generator
				IsOptional:  true,
			}},
			Options: []model.Option{{
				Name:        []string{"--help", "-h"},
				Description: `Prints help information`,
			}, {
				Name:        []string{"--version", "-V"},
				Description: `Prints version information`,
			}, {
				Name:        []string{"--yes"},
				Description: `Answer "yes" to all questions`,
			}},
		}, {
			Name:        []string{"upgrade"},
			Description: `Bring your dependencies up to date`,
			Args: []model.Arg{{
				Name:        "INPUT",
				Description: `The elm.json file to upgrade [default: elm.json]`,
				Generator:   nil, // TODO: port over generator
				IsOptional:  true,
			}},
			Options: []model.Option{{
				Name:        []string{"--help", "-h"},
				Description: `Prints help information`,
			}, {
				Name:        []string{"--unsafe"},
				Description: `Allow major versions bumps`,
			}, {
				Name:        []string{"--version", "-V"},
				Description: `Prints version information`,
			}, {
				Name:        []string{"--yes"},
				Description: `Answer "yes" to all questions`,
			}},
		}},
	}
}