// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["ansible-playbook"] = model.Subcommand{
		Name:        []string{"ansible-playbook"},
		Description: `Runs Ansible playbooks, executing the defined tasks on the targeted hosts`,
		Args: []model.Arg{{
			Name:        "playbook",
			Description: `Playbook(s)`,
			Generator:   nil, // TODO: port over generator
			IsVariadic:  true,
		}},
		Options: []model.Option{{
			Name:        []string{"--ask-vault-pass"},
			Description: `Ask for vault password`,
		}, {
			Name:        []string{"--flush-cache"},
			Description: `Clears the fact cache for every host in inventory`,
		}, {
			Name:        []string{"--force-handlers"},
			Description: `Run handlers even if a task fails`,
		}, {
			Name:        []string{"--list-hosts"},
			Description: `Outputs a list of matching hosts; does not execute`,
		}, {
			Name:        []string{"--list-tags"},
			Description: `List all available tags`,
		}, {
			Name:        []string{"--list-tasks"},
			Description: `List all tasks that would be executed`,
		}, {
			Name:        []string{"--skip-tags"},
			Description: `Only run plays and tasks whose tags do not match these values`,
			Args: []model.Arg{{
				Name: "tags",
			}},
		}, {
			Name:        []string{"--start-at-task"},
			Description: `Start the playbook at the task matching this name one-step-at-a-time`,
			Args: []model.Arg{{
				Name: "task name",
			}},
		}, {
			Name:        []string{"--step"},
			Description: `Execute one-step-at-a-time`,
		}, {
			Name:        []string{"--syntax-check"},
			Description: `Perform a syntax check on the playbook, but do not execute it`,
		}, {
			Name:        []string{"--vault-id"},
			Description: `Specify the vault identity to use`,
			Args: []model.Arg{{
				Name: "vault ID",
			}},
		}, {
			Name:        []string{"--vault-password-file"},
			Description: `Specify a vault password file`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "vault password file",
			}},
		}, {
			Name:        []string{"--version"},
			Description: `Show program's version number, config file location, configured module search path, module location and executable location`,
		}, {
			Name:        []string{"--check", "-C"},
			Description: `Don't make any changes; instead, try to predict some of the changes that may occur`,
		}, {
			Name:        []string{"--diff", "-D"},
			Description: `When changing (small) files and templates, show the differences in those files`,
		}, {
			Name:        []string{"--module-path", "-M"},
			Description: `Prepend colon-separated path(s) to module library`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFolders},
				Name:      "module path",
			}},
		}, {
			Name:        []string{"--extra-vars", "-e"},
			Description: `Set additional variables as key=value or YAML/JSON, if filename prepend with @`,
			Args: []model.Arg{{
				Name: "extra vars",
			}},
		}, {
			Name:        []string{"--forks", "-f"},
			Description: `Specify number of parallel processes to use`,
			Args: []model.Arg{{
				Name: "forks",
			}},
		}, {
			Name:        []string{"--help", "-h"},
			Description: `Show help for ansible`,
		}, {
			Name:        []string{"--inventory", "-i"},
			Description: `Specify inventory host path or comma separated host list`,
			Args: []model.Arg{{
				Name: "inventory",
			}},
		}, {
			Name:        []string{"--limit", "-l"},
			Description: `Limit selected hosts to an additional pattern`,
			Args: []model.Arg{{
				Name: "subset",
			}},
		}, {
			Name:        []string{"--tags", "-t"},
			Description: `Only run plays and tasks tagged with these values`,
			Args: []model.Arg{{
				Name: "tags",
			}},
		}, {
			Name:        []string{"--verbose", "-v"},
			Description: `Enable verbose mode`,
		}, {
			Name:        []string{"-vvv"},
			Description: `Enable very verbose mode`,
		}, {
			Name:        []string{"-vvvv"},
			Description: `Enable connection debug mode`,
		}, {
			Name:        []string{"--become-method"},
			Description: `Privilege escalation method to use`,
			Args: []model.Arg{{
				Name:        "become method",
				Suggestions: []model.Suggestion{{Name: []string{`sudo`}}, {Name: []string{`su`}}, {Name: []string{`pbrun`}}, {Name: []string{`pfexec`}}, {Name: []string{`doas`}}, {Name: []string{`dzdo`}}, {Name: []string{`ksu`}}, {Name: []string{`runas`}}, {Name: []string{`machinectl`}}},
			}},
		}, {
			Name:        []string{"--become-user"},
			Description: `Privilege escalation user to use`,
			Args: []model.Arg{{
				Name: "become user",
			}},
		}, {
			Name:        []string{"--ask-become-pass", "-K"},
			Description: `Prompt for privilege escalation password`,
		}, {
			Name:        []string{"--become", "-b"},
			Description: `Run operations with become`,
		}, {
			Name:        []string{"--private-key", "--key-file"},
			Description: `Use this fole to authenticate the connection`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "private key",
			}},
		}, {
			Name:        []string{"--scp-extra-args"},
			Description: `Extra arguments to pass to (only) scp`,
			Args: []model.Arg{{
				Name: "SCP extra args",
			}},
		}, {
			Name:        []string{"--sftp-extra-args"},
			Description: `Extra arguments to pass to (only) sftp`,
			Args: []model.Arg{{
				Name: "SFTP extra args",
			}},
		}, {
			Name:        []string{"-ssh-extra-args"},
			Description: `Extra arguments to pass to (only) ssh`,
			Args: []model.Arg{{
				Name: "SSH extra args",
			}},
		}, {
			Name:        []string{"--ssh-common-args"},
			Description: `Extra arguments to pass to sftp/scp/ssh`,
			Args: []model.Arg{{
				Name: "SSH common args",
			}},
		}, {
			Name:        []string{"--timeout", "-T"},
			Description: `Override the connection timeout in seconds`,
			Args: []model.Arg{{
				Name: "timeout",
			}},
		}, {
			Name:        []string{"--connection", "-c"},
			Description: `Connection type to use`,
			Args: []model.Arg{{
				Name: "connection type",
			}},
		}, {
			Name:        []string{"--ask-pass", "-k"},
			Description: `Ask for connection password`,
		}, {
			Name:        []string{"--user", "-u"},
			Description: `Connect as this user`,
			Args: []model.Arg{{
				Name: "user",
			}},
		}},
	}
}