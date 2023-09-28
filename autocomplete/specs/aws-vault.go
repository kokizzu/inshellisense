// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["aws-vault"] = model.Subcommand{
		Name:        []string{"aws-vault"},
		Description: `A vault for securely storing and accessing AWS credentials in development environments`,
		Options: []model.Option{{
			Name:        []string{"--help"},
			Description: `Show context-sensitive help (also try --help-long and --help-man)`,
		}, {
			Name:        []string{"--version"},
			Description: `Show application version`,
		}, {
			Name:        []string{"--debug"},
			Description: `Show debugging output`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"add"},
			Description: `Add credentials to the secure keystore`,
			Args: []model.Arg{{
				Name:        "profile",
				Description: `The profile you want to add`,
			}},
		}, {
			Name:        []string{"remove"},
			Description: `Remove credentials from the secure keystore`,
			Args: []model.Arg{{
				Name:      "profile",
				Generator: nil, // TODO: port over generator
			}},
			Options: []model.Option{{
				Name:        []string{"-f", "--force"},
				Description: `Force-remove the profile without a prompt`,
			}},
		}, {
			Name:        []string{"list"},
			Description: `List profiles, along with their credentials and sessions`,
			Options: []model.Option{{
				Name:        []string{"--profiles"},
				Description: `Show only the profile names`,
			}, {
				Name:        []string{"--sessions"},
				Description: `Show only the session names`,
			}, {
				Name:        []string{"--credentials"},
				Description: `Show only the profiles with stored credential`,
			}},
		}, {
			Name:        []string{"rotate"},
			Description: `Rotate credentials`,
			Args: []model.Arg{{
				Name:      "profile",
				Generator: nil, // TODO: port over generator
			}},
			Options: []model.Option{{
				Name:        []string{"-n", "--no-session"},
				Description: `Use master credentials, no session or role used`,
			}},
		}, {
			Name:        []string{"exec"},
			Description: `Execute a command with AWS credentials`,
			Args: []model.Arg{{
				Name:      "profile",
				Generator: nil, // TODO: port over generator
			}, {
				Name:      "command",
				IsCommand: true,
			}},
			Options: []model.Option{{
				Name:        []string{"-d", "--duration"},
				Description: `Duration of the temporary or assume-role session. Defaults to 1h`,
				Args: []model.Arg{{
					Name: "DURATION",
				}},
			}, {
				Name:        []string{"-n", "--no-session"},
				Description: `Use master credentials, no session or role used`,
			}, {
				Name:        []string{"--region"},
				Description: `The AWS region`,
				Args: []model.Arg{{
					Name: "REGION",
				}},
			}, {
				Name:        []string{"-t", "--mfa-token"},
				Description: `The MFA token to use`,
				Args: []model.Arg{{
					Name: "MFA-TOKEN",
				}},
			}},
		}, {
			Name:        []string{"export"},
			Description: `Export AWS credentials`,
			Args: []model.Arg{{
				Name:      "profile",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"clear"},
			Description: `Clear temporary credentials from the secure keystore`,
			Args: []model.Arg{{
				Name:      "profile",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"login"},
			Description: `Generate a login link for the AWS Console`,
			Args: []model.Arg{{
				Name:      "profile",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"help"},
			Description: `Show help about the command`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateHelp},
				Name:      "command",
			}},
		}},
	}
}