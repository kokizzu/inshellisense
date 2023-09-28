// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["defaults"] = model.Subcommand{
		Name:        []string{"defaults"},
		Description: `Command line interface to a user's defaults`,
		Subcommands: []model.Subcommand{{
			Name:        []string{"read"},
			Description: `Shows defaults`,
			Args: []model.Arg{{
				Name: "domain",
				Suggestions: []model.Suggestion{{
					Name:        []string{`-globalDomain`},
					Description: `Global domain`,
				}, {
					Name:        []string{`-app`},
					Description: `Application name`,
				}},
				Generator: nil, // TODO: port over generator
			}, {
				Name: "key",
			}},
		}, {
			Name:        []string{"write"},
			Description: `Writes key for domain`,
			Args: []model.Arg{{
				Name: "domain",
				Suggestions: []model.Suggestion{{
					Name:        []string{`-globalDomain`},
					Description: `Global domain`,
				}, {
					Name:        []string{`-app`},
					Description: `Application name`,
				}},
				Generator: nil, // TODO: port over generator
			}, {
				Name: "key",
			}, {
				Name: "value",
			}},
		}, {
			Name:        []string{"delete"},
			Description: `Deletes domain or key in domain`,
			Args: []model.Arg{{
				Name: "domain",
				Suggestions: []model.Suggestion{{
					Name:        []string{`-globalDomain`},
					Description: `Global domain`,
				}, {
					Name:        []string{`-app`},
					Description: `Application name`,
				}},
				Generator: nil, // TODO: port over generator
			}, {
				Name: "key",
			}},
		}, {
			Name:        []string{"rename"},
			Description: `Renames old_key to new_key`,
			Args: []model.Arg{{
				Name: "domain",
				Suggestions: []model.Suggestion{{
					Name:        []string{`-globalDomain`},
					Description: `Global domain`,
				}, {
					Name:        []string{`-app`},
					Description: `Application name`,
				}},
				Generator: nil, // TODO: port over generator
			}, {
				Name: "old_key",
			}, {
				Name: "new_key",
			}},
		}, {
			Name:        []string{"domains"},
			Description: `Lists all domains`,
		}, {
			Name:        []string{"find"},
			Description: `Lists all entries containing word`,
			Args: []model.Arg{{
				Name:        "word",
				Description: `The word to search for`,
			}},
		}, {
			Name:        []string{"help"},
			Description: `Show help text`,
		}, {
			Name:        []string{"read-type"},
			Description: `Shows the type for the given domain, key`,
			Args: []model.Arg{{
				Name: "domain",
				Suggestions: []model.Suggestion{{
					Name:        []string{`-globalDomain`},
					Description: `Global domain`,
				}, {
					Name:        []string{`-app`},
					Description: `Application name`,
				}},
				Generator: nil, // TODO: port over generator
			}, {
				Name: "key",
			}},
		}},
	}
}