// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["mvn"] = model.Subcommand{
		Name:        []string{"mvn"},
		Description: `Maven - a Java based project management and comprehension tool`,
		Args: []model.Arg{{
			Name:        "goal/phase",
			Description: `Goal or phase to execute`,
			Suggestions: []model.Suggestion{{Name: []string{`pre-clean`}}, {Name: []string{`clean`}}, {Name: []string{`post-clean`}}, {Name: []string{`validate`}}, {Name: []string{`initialize`}}, {Name: []string{`generate-sources`}}, {Name: []string{`process-sources`}}, {Name: []string{`generate-resources`}}, {Name: []string{`process-resources`}}, {Name: []string{`compile`}}, {Name: []string{`process-classes`}}, {Name: []string{`generate-test-sources`}}, {Name: []string{`process-test-sources`}}, {Name: []string{`generate-test-resources`}}, {Name: []string{`process-test-resources`}}, {Name: []string{`test-compile`}}, {Name: []string{`process-test-classes`}}, {Name: []string{`test`}}, {Name: []string{`prepare-package`}}, {Name: []string{`package`}}, {Name: []string{`pre-integration-test`}}, {Name: []string{`integration-test`}}, {Name: []string{`post-integration-test`}}, {Name: []string{`verify`}}, {Name: []string{`install`}}, {Name: []string{`deploy`}}, {Name: []string{`pre-site`}}, {Name: []string{`site`}}, {Name: []string{`post-site`}}, {Name: []string{`site-deploy`}}},
			IsVariadic:  true,
		}},
		Options: []model.Option{{
			Name:        []string{"--also-make", "-am"},
			Description: `Also build projects required by project list`,
		}, {
			Name:        []string{"--also-make-dependents", "-amd"},
			Description: `Also build projects that depend on projects in the project list`,
		}, {
			Name:        []string{"--batch-mode", "-B"},
			Description: `Run in non-interactive (batch)`,
		}, {
			Name:        []string{"--builder", "-b"},
			Description: `Specify the build strategy to use`,
			Args: []model.Arg{{
				Name: "id of build strategy",
			}},
		}, {
			Name:        []string{"--strict-checksums", "-C"},
			Description: `Fail if checksums do not match`,
		}, {
			Name:        []string{"--lax-checksums", "-c"},
			Description: `Warn if checksums do not match`,
		}, {
			Name:        []string{"--color"},
			Description: `Specify the color mode of the output`,
			Args: []model.Arg{{
				Name:        "color mode",
				Suggestions: []model.Suggestion{{Name: []string{`always`}}, {Name: []string{`never`}}, {Name: []string{`auto`}}},
			}},
		}, {
			Name:        []string{"--check-plugin-updates", "-cpu"},
			Description: `Ineffective. Only kept for backward compatibility`,
		}, {
			Name:        []string{"--define", "-D"},
			Description: `Define a system property`,
			Args: []model.Arg{{
				Name: "system property",
			}},
		}, {
			Name:        []string{"--errors", "-e"},
			Description: `Produce execution error messages`,
		}, {
			Name:        []string{"--encrypt-master-password", "-emp"},
			Description: `Encrypt the master security password`,
			Args: []model.Arg{{
				Name: "master password",
			}},
		}, {
			Name:        []string{"--encrypt-password", "-ep"},
			Description: `Encrypt the server password`,
			Args: []model.Arg{{
				Name: "server password",
			}},
		}, {
			Name:        []string{"--file", "-f"},
			Description: `Force the use of an alternate POM file (or directory with pom.xml)`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths, model.TemplateFolders},
				Name:      "path",
			}},
		}, {
			Name:        []string{"--fail-at-end", "-fae"},
			Description: `Only fail the build afterwards; allow all non-impacted builds to continue`,
		}, {
			Name:        []string{"--fail-fast", "-ff"},
			Description: `Stop at first failure in reactorized builds`,
		}, {
			Name:        []string{"--fail-never", "-fn"},
			Description: `Never fail the build, regardless of project result`,
		}, {
			Name:        []string{"--global-settings", "-gs"},
			Description: `Specify the global settings file to use`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "path",
			}},
		}, {
			Name:        []string{"--global-toolchains", "-gt"},
			Description: `Specify the global toolchains file to use`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "path",
			}},
		}, {
			Name:        []string{"--help", "-h"},
			Description: `Display help information`,
		}, {
			Name:        []string{"--log-file", "-l"},
			Description: `Specify the file to log to`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "path",
			}},
		}, {
			Name:        []string{"--legacy-local-repository", "-llr"},
			Description: `Use the Maven2 legacy local repository behaviour`,
		}, {
			Name:        []string{"--non-recursive", "-N"},
			Description: `Do not recurse into sub-projects`,
		}, {
			Name:        []string{"--no-plugin-registry", "-npr"},
			Description: `Ineffective. Only kept for backward compatibility`,
		}, {
			Name:        []string{"--no-plugin-updates", "-npu"},
			Description: `Ineffective. Only kept for backward compatibility`,
		}, {
			Name:        []string{"--no-snapshot-updates", "-nsu"},
			Description: `Suppress SNAPSHOT updates`,
		}, {
			Name:        []string{"--no-transfer-progress", "-ntp"},
			Description: `Do not display transfer progress when downloading or uploading`,
		}, {
			Name:        []string{"--offline", "-o"},
			Description: `Work offline`,
		}, {
			Name:        []string{"--activate-profiles", "-P"},
			Description: `Activate the specified profiles (comma delimited)`,
			Args: []model.Arg{{
				Name: "profiles",
			}},
		}, {
			Name:        []string{"--projects", "-pl"},
			Description: `Specify the projects to build`,
			Args: []model.Arg{{
				Name: "project list",
			}},
		}, {
			Name:        []string{"--quiet", "-q"},
			Description: `Quiet output - only shows errors`,
		}, {
			Name:        []string{"--resume-from", "-rf"},
			Description: `Resume from the specified project`,
			Args: []model.Arg{{
				Name: "project",
			}},
		}, {
			Name:        []string{"--settings", "-s"},
			Description: `Specify the user settings file to use`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "path",
			}},
		}, {
			Name:        []string{"--toolchains", "-t"},
			Description: `Specify the toolchains file to use`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "path",
			}},
		}, {
			Name:        []string{"--threads", "-T"},
			Description: `Specify the number of threads to use`,
			Args: []model.Arg{{
				Name: "threads",
			}},
		}, {
			Name:        []string{"--update-snapshots", "-U"},
			Description: `Forces a check for missing releases and updated snapshots on remote repositories`,
		}, {
			Name:        []string{"--update-plugins", "-up"},
			Description: `Ineffective. Only kept for backward compatibility`,
		}, {
			Name:        []string{"--version", "-v"},
			Description: `Display version information`,
		}, {
			Name:        []string{"--show-version", "-V"},
			Description: `Display version information`,
		}, {
			Name:        []string{"--debug", "-X"},
			Description: `Produce execution debug output`,
		}},
	}
}