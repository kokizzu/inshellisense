// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["encore"] = model.Subcommand{
		Name:        []string{"encore"},
		Description: `Encore is the fastest way of developing backend applications`,
		Options: []model.Option{{
			Name:        []string{"--verbose", "-v"},
			Description: `Verbose output`,
			Args: []model.Arg{{
				Name: "verbose",
			}},
			IsPersistent: true,
		}, {
			Name:         []string{"--help", "-h"},
			Description:  `Display help`,
			IsPersistent: true,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"app"},
			Description: `Commands to create and link Encore apps`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"clone"},
				Description: `Clone an Encore app to your computer`,
			}, {
				Name:        []string{"create"},
				Description: `Create a new Encore app`,
				Options: []model.Option{{
					Name:        []string{"--example"},
					Description: `URL to example code to use`,
					Args: []model.Arg{{
						Name: "example",
					}},
				}},
			}, {
				Name:        []string{"link"},
				Description: `Link an Encore app with the server`,
				Options: []model.Option{{
					Name:        []string{"--force", "-f"},
					Description: `Force link even if the app is already linked`,
				}},
			}},
		}, {
			Name:        []string{"auth"},
			Description: `Commands to authenticate with Encore`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"login"},
				Description: `Log in to Encore`,
				Options: []model.Option{{
					Name:        []string{"--auth-key", "-k"},
					Description: `Auth Key to use for login`,
					Args: []model.Arg{{
						Name: "auth-key",
					}},
				}},
			}, {
				Name:        []string{"logout"},
				Description: `Logs out the currently logged in user`,
			}, {
				Name:        []string{"signup"},
				Description: `Create a new Encore account`,
			}, {
				Name:        []string{"whoami"},
				Description: `Show the current logged in user`,
			}},
		}, {
			Name:        []string{"check"},
			Description: `Checks your application for compile-time errors using Encore's compiler`,
			Options: []model.Option{{
				Name:        []string{"--codegen-debug"},
				Description: `Dump generated code (for debugging Encore's code generation)`,
			}},
		}, {
			Name:        []string{"daemon"},
			Description: `Starts the encore daemon`,
			Options: []model.Option{{
				Name:        []string{"--foreground", "-f"},
				Description: `Start the daemon in the foreground`,
			}},
			Subcommands: []model.Subcommand{{
				Name:        []string{"env"},
				Description: `Prints Encore environment information`,
			}},
		}, {
			Name:        []string{"db"},
			Description: `Database management commands`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"conn-uri"},
				Description: `Outputs the database connection string`,
				Options: []model.Option{{
					Name:        []string{"--env", "-e"},
					Description: `Environment name to connect to (such as "prod")`,
					Args: []model.Arg{{
						Name: "env",
					}},
				}},
			}, {
				Name:        []string{"proxy"},
				Description: `Sets up a proxy tunnel to the database`,
				Options: []model.Option{{
					Name:        []string{"--env", "-e"},
					Description: `Environment name to connect to (such as "prod")`,
					Args: []model.Arg{{
						Name: "env",
					}},
				}, {
					Name:        []string{"--port", "-p"},
					Description: `Port to listen on (defaults to a random port)`,
					Args: []model.Arg{{
						Name: "port",
					}},
				}},
			}, {
				Name:        []string{"reset"},
				Description: `Resets the databases for the given services. Use --all to reset all databases`,
				Options: []model.Option{{
					Name:        []string{"--all"},
					Description: `Reset all services in the application`,
				}},
			}, {
				Name:        []string{"shell"},
				Description: `Connects to the database via psql shell`,
				Options: []model.Option{{
					Name:        []string{"--env", "-e"},
					Description: `Environment name to connect to (such as "prod")`,
					Args: []model.Arg{{
						Name: "env",
					}},
				}},
			}},
		}, {
			Name:        []string{"eject"},
			Description: `Eject provides ways to eject your application to migrate away from using Encore`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"docker"},
				Description: `Docker builds a portable docker image of your Encore application`,
				Options: []model.Option{{
					Name:        []string{"--base"},
					Description: `Base image to build from`,
					Args: []model.Arg{{
						Name: "base",
					}},
				}, {
					Name:        []string{"--push", "-p"},
					Description: `Push image to remote repository`,
				}},
			}},
		}, {
			Name:        []string{"gen"},
			Description: `Code generation commands`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"client"},
				Description: `Generates an API client for your app`,
				Options: []model.Option{{
					Name:        []string{"--env", "-e"},
					Description: `The environment to fetch the API for (defaults to the primary environment)`,
					Args: []model.Arg{{
						Name: "env",
					}},
				}, {
					Name:        []string{"--lang", "-l"},
					Description: `The language to generate code for ("typescript" and "go" are supported)`,
					Args: []model.Arg{{
						Name: "lang",
					}},
				}, {
					Name:        []string{"--output", "-o"},
					Description: `The filename to write the generated client code to`,
					Args: []model.Arg{{
						Templates: []model.Template{model.TemplateFilepaths},
						Name:      "output",
					}},
				}},
			}, {
				Name:        []string{"wrappers"},
				Description: `Generates user-facing wrapper code`,
			}},
		}, {
			Name:        []string{"logs"},
			Description: `Streams logs from your application`,
			Options: []model.Option{{
				Name:        []string{"--env", "-e"},
				Description: `Environment name to stream logs from (defaults to the production environment)`,
				Args: []model.Arg{{
					Name: "env",
				}},
			}, {
				Name:        []string{"--json"},
				Description: `Whether to print logs in raw JSON format`,
			}},
		}, {
			Name:        []string{"run"},
			Description: `Runs your application`,
			Options: []model.Option{{
				Name:        []string{"--debug"},
				Description: `Compile for debugging (disables some optimizations)`,
			}, {
				Name:        []string{"--listen"},
				Description: `Address to listen on (for example "0.0.0.0:4000")`,
				Args: []model.Arg{{
					Name: "listen",
				}},
			}, {
				Name:        []string{"--port", "-p"},
				Description: `Port to listen on`,
				Args: []model.Arg{{
					Name: "port",
				}},
			}, {
				Name:        []string{"--tunnel"},
				Description: `Create a tunnel to your machine for others to test against`,
			}, {
				Name:        []string{"--watch", "-w"},
				Description: `Watch for changes and live-reload`,
			}},
		}, {
			Name:        []string{"secret"},
			Description: `Secret management commands`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"set"},
				Description: `Sets a secret value`,
				Options: []model.Option{{
					Name:        []string{"--dev", "-d"},
					Description: `To set the secret for development use`,
				}, {
					Name:        []string{"--prod", "-p"},
					Description: `To set the secret for production use`,
				}},
			}},
		}, {
			Name:        []string{"test"},
			Description: `Tests your application`,
		}, {
			Name:        []string{"version"},
			Description: `Reports the current version of the encore application`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"update"},
				Description: `Checks for an update of encore and, if one is available, runs the appropriate command to update it`,
			}},
		}, {
			Name:        []string{"vpn"},
			Description: `VPN management commands`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"start"},
				Description: `Sets up a secure connection to private environments`,
			}, {
				Name:        []string{"status"},
				Description: `Determines the status of the VPN connection`,
			}, {
				Name:        []string{"stop"},
				Description: `Stops the VPN connection`,
			}},
		}, {
			Name:        []string{"help"},
			Description: `Help about any command`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"app"},
				Description: `Commands to create and link Encore apps`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"clone"},
					Description: `Clone an Encore app to your computer`,
				}, {
					Name:        []string{"create"},
					Description: `Create a new Encore app`,
				}, {
					Name:        []string{"link"},
					Description: `Link an Encore app with the server`,
				}},
			}, {
				Name:        []string{"auth"},
				Description: `Commands to authenticate with Encore`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"login"},
					Description: `Log in to Encore`,
				}, {
					Name:        []string{"logout"},
					Description: `Logs out the currently logged in user`,
				}, {
					Name:        []string{"signup"},
					Description: `Create a new Encore account`,
				}, {
					Name:        []string{"whoami"},
					Description: `Show the current logged in user`,
				}},
			}, {
				Name:        []string{"check"},
				Description: `Checks your application for compile-time errors using Encore's compiler`,
			}, {
				Name:        []string{"daemon"},
				Description: `Starts the encore daemon`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"env"},
					Description: `Prints Encore environment information`,
				}},
			}, {
				Name:        []string{"db"},
				Description: `Database management commands`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"conn-uri"},
					Description: `Outputs the database connection string`,
				}, {
					Name:        []string{"proxy"},
					Description: `Sets up a proxy tunnel to the database`,
				}, {
					Name:        []string{"reset"},
					Description: `Resets the databases for the given services. Use --all to reset all databases`,
				}, {
					Name:        []string{"shell"},
					Description: `Connects to the database via psql shell`,
				}},
			}, {
				Name:        []string{"eject"},
				Description: `Eject provides ways to eject your application to migrate away from using Encore`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"docker"},
					Description: `Docker builds a portable docker image of your Encore application`,
				}},
			}, {
				Name:        []string{"gen"},
				Description: `Code generation commands`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"client"},
					Description: `Generates an API client for your app`,
				}, {
					Name:        []string{"wrappers"},
					Description: `Generates user-facing wrapper code`,
				}},
			}, {
				Name:        []string{"logs"},
				Description: `Streams logs from your application`,
			}, {
				Name:        []string{"run"},
				Description: `Runs your application`,
			}, {
				Name:        []string{"secret"},
				Description: `Secret management commands`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"set"},
					Description: `Sets a secret value`,
				}},
			}, {
				Name:        []string{"test"},
				Description: `Tests your application`,
			}, {
				Name:        []string{"version"},
				Description: `Reports the current version of the encore application`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"update"},
					Description: `Checks for an update of encore and, if one is available, runs the appropriate command to update it`,
				}},
			}, {
				Name:        []string{"vpn"},
				Description: `VPN management commands`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"start"},
					Description: `Sets up a secure connection to private environments`,
				}, {
					Name:        []string{"status"},
					Description: `Determines the status of the VPN connection`,
				}, {
					Name:        []string{"stop"},
					Description: `Stops the VPN connection`,
				}},
			}},
		}},
	}
}