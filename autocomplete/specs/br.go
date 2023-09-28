// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["br"] = model.Subcommand{
		Name:        []string{"broot"},
		Description: `Broot lets you explore file hierarchies with a tree-like view, manipulate files, launch actions, and define your own shortcuts. broot is best launched as br: this shell function gives you access to more commands, especially cd. The br shell function is interactively installed on first broot launch. Flags and options can be classically passed on launch but also written in the configuration file. Each flag has a counter-flag so that you can cancel at command line a flag which has been set in the configuration file`,
		Args: []model.Arg{{
			Name:        "ROOT",
			Description: `Sets the root directory`,
		}},
		Options: []model.Option{{
			Name:        []string{"--dates", "-d"},
			Description: `Show the last modified date of files and directories`,
			ExclusiveOn: []string{"--no-dates", "-D"},
		}, {
			Name:        []string{"--no-dates", "-D"},
			Description: `Don't show the last modified date`,
			ExclusiveOn: []string{"--dates", "-d"},
		}, {
			Name:        []string{"--only-folders", "-f"},
			Description: `Only show folders`,
			ExclusiveOn: []string{"--no-only-folders", "-F"},
		}, {
			Name:        []string{"--no-only-folders", "-F"},
			Description: `Show folders and files alike`,
			ExclusiveOn: []string{"--only-folders", "-f"},
		}, {
			Name:        []string{"--show-root-fs"},
			Description: `Show filesystem info on top`,
		}, {
			Name:        []string{"--show-git-info", "-g"},
			Description: `Show git statuses on files and stats of repository`,
			ExclusiveOn: []string{"--no-show-git-info", "-G"},
		}, {
			Name:        []string{"--no-show-git-info", "-G"},
			Description: `Don't show git statuses on files nor stats`,
			ExclusiveOn: []string{"--show-git-info", "-g"},
		}, {
			Name:        []string{"--git-status"},
			Description: `Only show files having an interesting git status, including hidden ones`,
		}, {
			Name:        []string{"--hidden", "-h"},
			Description: `Show hidden files`,
			ExclusiveOn: []string{"--no-hidden", "-H"},
		}, {
			Name:        []string{"--no-hidden", "-H"},
			Description: `Don't show hidden files`,
			ExclusiveOn: []string{"--hidden", "-h"},
		}, {
			Name:        []string{"--show-gitignored", "-i"},
			Description: `Show files which should be ignored according to git`,
			ExclusiveOn: []string{"--no-show-gitignored", "-I"},
		}, {
			Name:        []string{"--no-show-gitignored", "-I"},
			Description: `Don't show gitignored files`,
			ExclusiveOn: []string{"--show-gitignored", "-i"},
		}, {
			Name:        []string{"--permissions", "-p"},
			Description: `Show permissions with owner and group`,
			ExclusiveOn: []string{"--no-permissions", "-P"},
		}, {
			Name:        []string{"--no-permissions", "-P"},
			Description: `Don't show permissions`,
			ExclusiveOn: []string{"--permissions", "-p"},
		}, {
			Name:        []string{"--sizes", "-s"},
			Description: `Show the sizes of files and directories. When this mode is on, children aren't shown so that the biggest entries at the selected level can be sorted first`,
			ExclusiveOn: []string{"--no-sizes", "-S"},
		}, {
			Name:        []string{"--no-sizes", "-S"},
			Description: `Don't show sizes`,
			ExclusiveOn: []string{"--sizes", "-s"},
		}, {
			Name:        []string{"--sort-by-count"},
			Description: `Sort by count (only show one level of the tree)`,
			ExclusiveOn: []string{"--no-sort"},
		}, {
			Name:        []string{"--sort-by-date"},
			Description: `Sort by date (only show one level of the tree)`,
			ExclusiveOn: []string{"--no-sort"},
		}, {
			Name:        []string{"--sort-by-size"},
			Description: `Sort by size (only show one level of the tree)`,
			ExclusiveOn: []string{"--no-sort"},
		}, {
			Name:        []string{"--whale-spotting", "-w"},
			Description: `Sort by size, show ignored and hidden files`,
		}, {
			Name:        []string{"--no-sort"},
			Description: `Don't sort`,
			ExclusiveOn: []string{"--sort-by-count", "--sort-by-date", "--sort-by-size"},
		}, {
			Name:        []string{"--trim-root", "-t"},
			Description: `Trim the root: remove elements which would exceed the screen size. This removes the scrollbar`,
			ExclusiveOn: []string{"--no-trim-root", "-T"},
		}, {
			Name:        []string{"--no-trim-root", "-T"},
			Description: `Don't trim the root (still trim the deeper levels). A scrollbar may be used when there are too many elements to show on screen`,
			ExclusiveOn: []string{"--trim-root", "-t"},
		}, {
			Name:        []string{"--install"},
			Description: `Install or reinstall the br shell function`,
		}, {
			Name:        []string{"--get-root"},
			Description: `Ask for the current root of the remote broot`,
		}, {
			Name:        []string{"--color"},
			Description: `Controls styling of the output (default: auto). If set to auto, all styling is removed when output is piped`,
			Args: []model.Arg{{
				Suggestions: []model.Suggestion{{Name: []string{`yes`}}, {Name: []string{`no`}}, {Name: []string{`auto`}}},
				IsOptional:  true,
			}},
		}, {
			Name:        []string{"--help"},
			Description: `Prints a help page, with more or less the same content as this man page`,
		}, {
			Name:        []string{"--version", "-v"},
			Description: `Prints the version of broot`,
		}, {
			Name:        []string{"--outcmd"},
			Description: `Where to write a command if broot produces one`,
			Args: []model.Arg{{
				Name: "cmd-export-path",
			}},
		}, {
			Name:        []string{"--cmd", "-c"},
			Description: `Semicolon separated commands to execute on start of broot`,
			Args: []model.Arg{{
				Name: "commands",
			}},
		}, {
			Name:        []string{"--conf"},
			Description: `Semicolon separated paths to specific config files`,
			Args: []model.Arg{{
				Name: "conf",
			}},
		}, {
			Name:        []string{"--height"},
			Description: `Height to use if you don't want to fill the screen or for file export (by default the terminal height is used)`,
			Args: []model.Arg{{
				Name: "height",
			}},
		}, {
			Name:        []string{"--out", "-o"},
			Description: `Where to write the produced path, if any`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "file-export-path",
			}},
		}, {
			Name:        []string{"--set-install-state"},
			Description: `Set the installation state. This is mostly useful in installation scripts to override the normal installation process`,
			Args: []model.Arg{{
				Name:        "state",
				Suggestions: []model.Suggestion{{Name: []string{`undefined`}}, {Name: []string{`refused`}}, {Name: []string{`installed`}}},
			}},
		}, {
			Name:        []string{"--print-shell-function"},
			Description: `Print to stdout the br function for the given shell. This can be used in your own installation process overriden the standard one`,
			Args: []model.Arg{{
				Name:        "shell",
				Suggestions: []model.Suggestion{{Name: []string{`bash`}}, {Name: []string{`fish`}}, {Name: []string{`zsh`}}},
			}},
		}, {
			Name:        []string{"--listen"},
			Description: `Listen for commands`,
			Args: []model.Arg{{
				Name: "listen",
			}},
		}, {
			Name:        []string{"--send"},
			Description: `Send commands to a remote broot then quits`,
			Args: []model.Arg{{
				Name: "send",
			}},
		}},
	}
}