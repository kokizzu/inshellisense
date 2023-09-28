// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["git-cliff"] = model.Subcommand{
		Name:        []string{"git-cliff"},
		Description: `A highly customizable changelog generator ⛰️`,
		Args: []model.Arg{{
			Name:        "range",
			Description: `Sets the commit range to process`,
			Generator:   nil, // TODO: port over generator
		}},
		Options: []model.Option{{
			Name:        []string{"--verbose", "-v"},
			Description: `Increases the logging verbosity`,
		}, {
			Name:        []string{"--init", "-i"},
			Description: `Writes the default configuration file to cliff.toml`,
		}, {
			Name:        []string{"--latest", "-l"},
			Description: `Processes the commits starting from the latest tag`,
			ExclusiveOn: []string{"--current", "-u"},
		}, {
			Name:        []string{"--current"},
			Description: `Processes the commits that belong to the current tag`,
			ExclusiveOn: []string{"-l", "-u"},
		}, {
			Name:        []string{"--unreleased", "-u"},
			Description: `Processes the commits that do not belong to a tag`,
			ExclusiveOn: []string{"-l", "--current"},
		}, {
			Name:        []string{"--date-order"},
			Description: `Sorts the tags chronologically`,
		}, {
			Name:        []string{"--context"},
			Description: `Prints changelog context as JSON`,
		}, {
			Name:        []string{"--help", "-h"},
			Description: `Prints help information`,
		}, {
			Name:        []string{"--version", "-V"},
			Description: `Prints version information`,
		}, {
			Name:        []string{"--config", "-c"},
			Description: `Sets the configuration file`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "path",
			}},
		}, {
			Name:        []string{"--workdir", "-w"},
			Description: `Sets the working directory`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFolders},
				Name:      "path",
			}},
		}, {
			Name:        []string{"--repository", "-r"},
			Description: `Sets the git repository`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFolders},
				Name:      "path",
			}},
		}, {
			Name:        []string{"--include-path"},
			Description: `Sets the path to include related commits`,
			Args: []model.Arg{{
				Name:       "pattern",
				IsVariadic: true,
			}},
		}, {
			Name:        []string{"--exclude-path"},
			Description: `Sets the path to exclude related commits`,
			Args: []model.Arg{{
				Name:       "pattern",
				IsVariadic: true,
			}},
		}, {
			Name:        []string{"--with-commit"},
			Description: `Sets custom commit messages to include in the changelog`,
			Args: []model.Arg{{
				Name:       "msg",
				IsVariadic: true,
			}},
		}, {
			Name:        []string{"--prepend", "-p"},
			Description: `Prepends entries to the given changelog file`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "path",
			}},
		}, {
			Name:        []string{"--output", "-o"},
			Description: `Writes output to the given file`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "path",
			}},
		}, {
			Name:        []string{"--tag", "-t"},
			Description: `Sets the tag for the latest version`,
			Args: []model.Arg{{
				Name: "tag",
			}},
		}, {
			Name:        []string{"--body", "-b"},
			Description: `Sets the template for the changelog body`,
			Args: []model.Arg{{
				Name: "template",
			}},
		}, {
			Name:        []string{"--strip", "-s"},
			Description: `Strips the given parts from the changelog`,
			Args: []model.Arg{{
				Name: "part",
				Suggestions: []model.Suggestion{{
					Name: []string{`header`},
				}, {
					Name: []string{`footer`},
				}, {
					Name: []string{`all`},
				}},
			}},
		}, {
			Name:        []string{"--sort"},
			Description: `Sets sorting of the commits inside sections`,
			Args: []model.Arg{{
				Name: "sort",
				Suggestions: []model.Suggestion{{
					Name: []string{`oldest`},
				}, {
					Name: []string{`newest`},
				}},
			}},
		}},
	}
}