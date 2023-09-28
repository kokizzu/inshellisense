// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["wd"] = model.Subcommand{
		Name:        []string{"wd"},
		Description: `Warp to directories without using cd`,
		Args: []model.Arg{{
			Name:        "point",
			Description: `Warp point to the specified directory`,
			Generator:   nil, // TODO: port over generator
		}, {
			Name:        "path",
			Description: `Appended path`,
			IsOptional:  true,
		}},
		Options: []model.Option{{
			Name:        []string{"-v", "--version"},
			Description: `Print version`,
		}, {
			Name:        []string{"-d", "--debug"},
			Description: `Exit after execution with exit codes (for testing)`,
		}, {
			Name:        []string{"-c", "--config"},
			Description: `Specify config file (default ~/.warprc)`,
			Args: []model.Arg{{
				Name: "file",
			}},
		}, {
			Name:        []string{"-q", "--quiet"},
			Description: `Suppress all output`,
		}, {
			Name:        []string{"-f", "--force"},
			Description: `Allows overwriting without warning (for add & clean)`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"add"},
			Description: `Adds the current working directory to your warp points`,
			Args: []model.Arg{{
				Name:        "point",
				Description: `Name of the warp point to be created`,
				IsOptional:  true,
			}},
		}, {
			Name:        []string{"rm"},
			Description: `Removes the given warp point`,
			Args: []model.Arg{{
				Name:        "point",
				Description: `Name of the warp point to be removed`,
				Generator:   nil, // TODO: port over generator
				IsOptional:  true,
			}},
		}, {
			Name:        []string{"show"},
			Description: `Print path to given warp point`,
			Args: []model.Arg{{
				Name:        "point",
				Description: `Name of the warp point`,
				Generator:   nil, // TODO: port over generator
				IsOptional:  true,
			}},
		}, {
			Name:        []string{"list"},
			Description: `Print all stored warp points`,
		}, {
			Name:        []string{"ls"},
			Description: `Show files from given warp point (ls)`,
			Args: []model.Arg{{
				Name:        "point",
				Description: `Name of the warp point`,
				Generator:   nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"path"},
			Description: `Show the path to given warp point (pwd)`,
			Args: []model.Arg{{
				Name:        "point",
				Description: `Name of the warp point`,
				Generator:   nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"clean"},
			Description: `Remove points warping to nonexistent directories (will prompt unless --force is used)`,
		}, {
			Name:        []string{"help"},
			Description: `Shows help for wd`,
		}},
	}
}