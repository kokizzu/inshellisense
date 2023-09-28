// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["hyper"] = model.Subcommand{
		Name:        []string{"hyper"},
		Description: `Hyper is an Electron-based terminal`,
		Args: []model.Arg{{
			Templates: []model.Template{model.TemplateFolders},
		}},
		Options: []model.Option{{
			Name:        []string{"-h", "--help"},
			Description: `Output usage information`,
		}, {
			Name:        []string{"-v", "--verbose"},
			Description: `Verbose mode (disabled by default)`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"install", "i"},
			Description: `Install a plugin`,
			Args: []model.Arg{{
				Name: "plugin",
			}},
		}, {
			Name:        []string{"docs", "d", "h", "home"},
			Description: `Open the npm page of a plugin`,
			Args: []model.Arg{{
				Name: "plugin",
			}},
		}, {
			Name:        []string{"help"},
			Description: `Display help`,
		}, {
			Name:        []string{"list", "ls"},
			Description: `List installed plugins`,
		}, {
			Name:        []string{"list-remote", "lsr", " ls-remote"},
			Description: `List plugins available on npm`,
		}, {
			Name:        []string{"search", "s"},
			Description: `Search for plugins on npm`,
			Args: []model.Arg{{
				Name:        "query",
				Description: `Your search query`,
				IsOptional:  true,
			}},
		}, {
			Name:        []string{"uninstall", "u", "rm", "remove"},
			Description: `Uninstall plugin`,
			Args: []model.Arg{{
				Name:        "plugin",
				Description: `Plugin to uninstall`,
				Generator:   nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"version"},
			Description: `Show version`,
		}},
	}
}