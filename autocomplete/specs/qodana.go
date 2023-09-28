// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["qodana"] = model.Subcommand{
		Name:        []string{"qodana"},
		Description: `Run Qodana as fast as possible, with minimum effort required`,
		Options: []model.Option{{
			Name:         []string{"-h", "--help"},
			Description:  `Show help page for command`,
			IsPersistent: true,
		}, {
			Name:        []string{"-v", "--version"},
			Description: `Version for Qodana`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"init"},
			Description: `Configure project for Qodana`,
			Options: []model.Option{{
				Name:        []string{"-i", "--project-dir"},
				Description: `Root directory of the project to configure (default ".")`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFolders},
					Name:      "string",
				}},
			}},
		}, {
			Name:        []string{"scan"},
			Description: `Scan a project with Qodana`,
			Options: []model.Option{{
				Name:        []string{"-a", "--analysis-id"},
				Description: `Unique report identifier (GUID) to be used by Qodana Cloud`,
				Args: []model.Arg{{
					Name: "string",
				}},
			}, {
				Name:        []string{"-b", "--baseline"},
				Description: `Provide the path to an existing SARIF report to be used in the baseline state calculation`,
				Args: []model.Arg{{
					Name: "string",
				}},
			}, {
				Name:        []string{"--baseline-include-absent"},
				Description: `Include in the output report the results from the baseline run that are absent in the current run`,
			}, {
				Name:        []string{"--cache-dir"},
				Description: `Override cache directory (default <userCacheDir>/JetBrains/<linter>/cache)`,
				Args: []model.Arg{{
					Name: "string",
				}},
			}, {
				Name:        []string{"-c", "--changes"},
				Description: `Override the docker image to be used for the analysis`,
			}, {
				Name:        []string{"--clear-cache"},
				Description: `Clear the local Qodana cache before running the analysis`,
			}, {
				Name:        []string{"--disable-sanity"},
				Description: `Skip running the inspections configured by the sanity profile`,
			}, {
				Name:        []string{"-e", "--env"},
				Description: `Define additional environment variables for the Qodana container (you can use the flag multiple times). CLI is not reading full host environment variables and does not pass it to the Qodana container for security reasons`,
				Args: []model.Arg{{
					Name: "stringArray",
				}},
			}, {
				Name:        []string{"--fail-threshold"},
				Description: `Set the number of problems that will serve as a quality gate. If this number is reached, the inspection run is terminated with a non-zero exit code`,
				Args: []model.Arg{{
					Name: "string",
				}},
			}, {
				Name:        []string{"-l", "--linter"},
				Description: `Override linter to use`,
				Args: []model.Arg{{
					Name: "string",
				}},
			}, {
				Name:        []string{"--port"},
				Description: `Port to serve the report on (default 8080)`,
				Args: []model.Arg{{
					Name: "int",
				}},
			}, {
				Name:        []string{"--print-problems"},
				Description: `Print all found problems by Qodana in the CLI output`,
			}, {
				Name:        []string{"-n", "--profile-name"},
				Description: `Profile name defined in the project`,
				Args: []model.Arg{{
					Name: "string",
				}},
			}, {
				Name:        []string{"-p", "--profile-path"},
				Description: `Path to the profile file`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "string",
				}},
			}, {
				Name:        []string{"-i", "--project-dir"},
				Description: `Root directory of the inspected project (default ".")`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFolders},
					Name:      "string",
				}},
			}, {
				Name:        []string{"--property"},
				Description: `Set a JVM property to be used while running Qodana using the --property property.name=value1,value2,...,valueN notation`,
				Args: []model.Arg{{
					Name: "stringArray",
				}},
			}, {
				Name:        []string{"-o", "--results-dir"},
				Description: `Override directory to save Qodana inspection results to (default <userCacheDir>/JetBrains/<linter>/results)`,
				Args: []model.Arg{{
					Name: "string",
				}},
			}, {
				Name:        []string{"--run-promo"},
				Description: `Set to 'true' to have the application run the inspections configured by the promo profile; set to 'false' otherwise (default: 'true' only if Qodana is executed with the default profile)`,
				Args: []model.Arg{{
					Name:        "string",
					Suggestions: []model.Suggestion{{Name: []string{`true`}}, {Name: []string{`false`}}},
				}},
			}, {
				Name:        []string{"-s", "--save-report"},
				Description: `Generate HTML report (default true)`,
			}, {
				Name:        []string{"--script"},
				Description: `Override the run scenario (default "default")`,
				Args: []model.Arg{{
					Name: "string",
				}},
			}, {
				Name:        []string{"--send-report"},
				Description: `Send the inspection report to Qodana Cloud, requires the '--token' option to be specified`,
			}, {
				Name:        []string{"-w", "--show-report"},
				Description: `Serve HTML report on port`,
			}, {
				Name:        []string{"--skip-pull"},
				Description: `Skip pulling the latest Qodana container`,
			}, {
				Name:        []string{"-d", "--source-directory"},
				Description: `Directory inside the project-dir directory must be inspected. If not specified, the whole project is inspected`,
				Args: []model.Arg{{
					Name: "string",
				}},
			}, {
				Name:        []string{"--stub-profile"},
				Description: `Absolute path to the fallback profile file. This option is applied in case the profile was not specified using any available options`,
				Args: []model.Arg{{
					Name: "string",
				}},
			}, {
				Name:        []string{"-u", "--user"},
				Description: `User to run Qodana container as. Please specify user id – '$UID' or user id and group id $(id -u):$(id -g). Use 'root' to run as the root user (default: the current user)`,
				Args: []model.Arg{{
					Name: "string",
				}},
			}, {
				Name:        []string{"-v", "--volume"},
				Description: `Define additional volumes for the Qodana container (you can use the flag multiple times)`,
				Args: []model.Arg{{
					Name: "string",
				}},
			}},
		}, {
			Name:        []string{"show"},
			Description: `Show Qodana report`,
			Options: []model.Option{{
				Name:        []string{"-d", "--dir-only"},
				Description: `Open report directory only, don't serve it`,
			}, {
				Name:        []string{"-p", "--port"},
				Description: `Specify port to serve report at (default 8080)`,
				Args: []model.Arg{{
					Name: "int",
				}},
			}, {
				Name:        []string{"-i", "--project-dir"},
				Description: `Root directory of the inspected project (default ".")`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFolders},
					Name:      "string",
				}},
			}, {
				Name:        []string{"-r", "--report-dir"},
				Description: `Specify HTML report path (the one with index.html inside) (default <userCacheDir>/JetBrains/<linter>/results/report)`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFolders},
					Name:      "string",
				}},
			}},
		}, {
			Name:        []string{"view"},
			Description: `View SARIF files in CLI`,
			Options: []model.Option{{
				Name:        []string{"-f", "--sarif-file"},
				Description: `Path to the SARIF file (default "./qodana.sarif.json")`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "string",
				}},
			}},
		}},
	}
}