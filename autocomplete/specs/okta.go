// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["okta"] = model.Subcommand{
		Name:        []string{"okta"},
		Description: `The Okta CLI is the easiest way to get started with Okta!`,
		Options: []model.Option{{
			Name:        []string{"--help", "-h"},
			Description: `Show help for Okta CLI`,
		}, {
			Name:        []string{"--version", "-V"},
			Description: `Print version information`,
		}, {
			Name:        []string{"--verbose"},
			Description: `Verbose logging`,
		}, {
			Name:        []string{"--batch"},
			Description: `Batch mode, will not prompt for user input`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"register"},
			Description: `Sign up for a new Okta account`,
			Options: []model.Option{{
				Name:        []string{"--company"},
				Description: `Company/organization used when registering a new Okta account`,
			}, {
				Name:        []string{"--email"},
				Description: `Email used when registering a new Okta account`,
			}, {
				Name:        []string{"--first-name"},
				Description: `First name used when registering a new Okta account`,
			}, {
				Name:        []string{"--last-name"},
				Description: `Last name used when registering a new Okta account`,
			}},
		}, {
			Name:        []string{"login"},
			Description: `Authorizes the Okta CLI tool`,
		}, {
			Name:        []string{"apps"},
			Description: `Manage Okta apps`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"config"},
				Description: `Show an Okta app's configuration`,
				Options: []model.Option{{
					Name:        []string{"--app"},
					Description: `The App ID`,
				}},
			}, {
				Name:        []string{"create"},
				Description: `Create a new Okta app`,
				Options: []model.Option{{
					Name:        []string{"--app-name"},
					Description: `Application name to be created, defaults to current directory name`,
				}, {
					Name:        []string{"--authorization-server-id"},
					Description: `Okta Authorization Server Id`,
				}, {
					Name:        []string{"--config-file"},
					Description: `Application config file`,
				}, {
					Name:        []string{"--redirect-uri"},
					Description: `OIDC Redirect URI`,
				}},
			}, {
				Name:        []string{"delete"},
				Description: `Deletes an Okta app`,
				Args: []model.Arg{{
					Name:        "appIds",
					Description: `List of application IDs to be deleted`,
				}},
				Options: []model.Option{{
					Name:        []string{"-f", "--force"},
					Description: `Deactivate and delete applications without confirmation`,
				}},
			}},
		}, {
			Name:        []string{"start"},
			Description: `Creates an Okta Sample Application`,
			Args: []model.Arg{{
				Name:        "name",
				Description: `The name of the sample app to create`,
				IsOptional:  true,
			}},
		}, {
			Name:        []string{"help"},
			Description: `Displays help information about the specified command`,
			Options: []model.Option{{
				Name:        []string{"register"},
				Description: `Sign up for a new Okta account`,
			}, {
				Name:        []string{"login"},
				Description: `Authorizes the Okta CLI tool`,
			}, {
				Name:        []string{"apps"},
				Description: `Manage Okta apps`,
			}, {
				Name:        []string{"start"},
				Description: `Creates an Okta Sample Application`,
			}, {
				Name:        []string{"help"},
				Description: `Displays help information about the specified command`,
			}, {
				Name:        []string{"generate-completion"},
				Description: `Generate bash/zsh completion script for Okta`,
			}},
		}, {
			Name:        []string{"generate-completion"},
			Description: `Generate bash/zsh completion script for Okta`,
		}},
	}
}