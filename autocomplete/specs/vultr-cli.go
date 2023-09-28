// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["vultr-cli"] = model.Subcommand{
		Name:        []string{"vultr-cli"},
		Description: `Official command line interface for the Vultr API`,
		Subcommands: []model.Subcommand{{
			Name:        []string{"account"},
			Description: `Retrieve information about your account`,
		}, {
			Name:        []string{"apps", "a"},
			Description: `Display all available applications`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"list", "l"},
				Description: `List applications`,
			}},
		}, {
			Name:        []string{"backups"},
			Description: `Display backups`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"get"},
				Description: `Get backup`,
				Args: []model.Arg{{
					Name:        "backupId",
					Description: `Backup ID`,
				}},
			}, {
				Name:        []string{"list", "l"},
				Description: `List backups`,
			}},
		}, {
			Name:        []string{"bare-metal", "bm"},
			Description: `Bare-metal is used to access bare metal server commands`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"app", "a"},
				Description: `App is used to access bare metal server application commands`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"change", "c"},
					Description: `Change a bare metal server's application`,
					Args: []model.Arg{{
						Name:        "bareMetalId",
						Description: `Bare Metal ID`,
					}, {
						Name:        "appID",
						Description: `Application ID`,
					}},
				}, {
					Name:        []string{"list"},
					Description: `Available apps a bare metal server can change to`,
					Args: []model.Arg{{
						Name:        "bareMetalId",
						Description: `Bare Metal ID`,
					}},
				}},
			}, {
				Name:        []string{"bandwidth", "b"},
				Description: `Get a bare metal server's bandwidth usage`,
				Args: []model.Arg{{
					Name:        "bareMetalId",
					Description: `Bare Metal ID`,
				}},
			}, {
				Name:        []string{"create", "c"},
				Description: `Create a bare metal server`,
				Options: []model.Option{{
					Name:        []string{"--os", "-o"},
					Description: `ID of the operating system that will be installed on the server`,
					Args: []model.Arg{{
						Name: "osId",
					}},
				}, {
					Name:        []string{"--persistent_pxe", "-x"},
					Description: `Enable persistent_pxe`,
					Args: []model.Arg{{
						Name:        "persistent",
						Suggestions: []model.Suggestion{{Name: []string{`true`}}, {Name: []string{`false`}}},
					}},
				}, {
					Name:        []string{"--plan", "-p"},
					Description: `ID of the plan that the server will subscribe to`,
					Args: []model.Arg{{
						Name: "planId",
					}},
				}, {
					Name:        []string{"--region", "-r"},
					Description: `ID of the region where the server will be created`,
					Args: []model.Arg{{
						Name: "regionId",
					}},
				}},
			}, {
				Name:        []string{"delete", "destroy"},
				Description: `Delete a bare metal server`,
				Args: []model.Arg{{
					Name:        "bareMetalId",
					Description: `Bare Metal ID`,
				}},
			}, {
				Name:        []string{"get"},
				Description: `Get a bare metal server by <bareMetalID>`,
				Args: []model.Arg{{
					Name:        "bareMetalId",
					Description: `Bare Metal ID`,
				}},
			}, {
				Name:        []string{"halt", "h"},
				Description: `Halt a bare metal server`,
				Args: []model.Arg{{
					Name:        "bareMetalId",
					Description: `Bare Metal ID`,
				}},
			}, {
				Name:        []string{"image", "i"},
				Description: `Image is used to access bare metal server image commands`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"change"},
					Description: `Change a bare metal server's image`,
					Args: []model.Arg{{
						Name:        "bareMetalId",
						Description: `Bare Metal ID`,
					}, {
						Name:        "imageID",
						Description: `Image ID`,
					}},
				}, {
					Name:        []string{"list"},
					Description: `Available images a bare metal server can change to`,
					Args: []model.Arg{{
						Name:        "bareMetalId",
						Description: `Bare Metal ID`,
					}},
				}},
			}, {
				Name:        []string{"ipv4"},
				Description: `List the IPv4 information of a bare metal server`,
			}, {
				Name:        []string{"ipv6"},
				Description: `List the IPv6 information of a bare metal server`,
			}, {
				Name:        []string{"list"},
				Description: `List all bare metal servers`,
			}, {
				Name:        []string{"os"},
				Description: `Os is used to access bare metal server operating system commands`,
			}, {
				Name:        []string{"reboot"},
				Description: `Reboot a bare metal server. This is a hard reboot, which means that the server is powered off, then back on`,
			}, {
				Name:        []string{"reinstall"},
				Description: `Reinstall the operating system on a bare metal server`,
			}, {
				Name:        []string{"start"},
				Description: `Start a bare metal server`,
			}, {
				Name:        []string{"user-data"},
				Description: `User-data is used to access bare metal server user-data commands`,
			}, {
				Name:        []string{"vnc"},
				Description: `Get a bare metal server's VNC url by <bareMetalID>`,
				Args: []model.Arg{{
					Name:        "bareMetalId",
					Description: `Bare Metal ID`,
				}},
			}},
		}, {
			Name:        []string{"block-storage"},
			Description: `Block storage commands`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"attach"},
				Description: `Attaches a block storage to an instance`,
			}, {
				Name:        []string{"create"},
				Description: `Create a new block storage`,
			}, {
				Name:        []string{"delete"},
				Description: `Delete a block storage`,
			}, {
				Name:        []string{"detach"},
				Description: `Detaches a block storage from an instance`,
			}, {
				Name:        []string{"get"},
				Description: `Retrieves a block storage`,
			}, {
				Name:        []string{"label"},
				Description: `Sets a label for a block storage`,
			}, {
				Name:        []string{"list"},
				Description: `Retrieves a list of active block storage`,
			}, {
				Name:        []string{"resize"},
				Description: `Resize a block storage`,
			}},
		}, {
			Name:        []string{"completion"},
			Description: `Generate the autocompletion script for the specified shell`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"bash"},
				Description: `Generate the autocompletion script for bash`,
			}, {
				Name:        []string{"fish"},
				Description: `Generate the autocompletion script for fish`,
			}, {
				Name:        []string{"powershell"},
				Description: `Generate the autocompletion script for powershell`,
			}, {
				Name:        []string{"zsh"},
				Description: `Generate the autocompletion script for zsh`,
			}},
		}, {
			Name:        []string{"dns"},
			Description: `Dns is used to access dns commands`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"domain"},
				Description: `Dns domain`,
			}, {
				Name:        []string{"record"},
				Description: `Dns record`,
			}},
		}, {
			Name:        []string{"firewall", "fw"},
			Description: `Firewall is used to access firewall commands`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"group"},
				Description: `Group is used to access firewall group commands`,
			}, {
				Name:        []string{"rule"},
				Description: `Rule is used to access firewall rule commands`,
			}},
		}, {
			Name:        []string{"help"},
			Description: `Help about any command`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"account"},
				Description: `Retrieve information about your account`,
			}, {
				Name:        []string{"apps"},
				Description: `Display all available applications`,
			}, {
				Name:        []string{"backups"},
				Description: `Display backups`,
			}, {
				Name:        []string{"bare-metal"},
				Description: `Bare-metal is used to access bare metal server commands`,
			}, {
				Name:        []string{"block-storage"},
				Description: `Block storage commands`,
			}, {
				Name:        []string{"completion"},
				Description: `Generate the autocompletion script for the specified shell`,
			}, {
				Name:        []string{"dns"},
				Description: `Dns is used to access dns commands`,
			}, {
				Name:        []string{"firewall"},
				Description: `Firewall is used to access firewall commands`,
			}, {
				Name:        []string{"help"},
				Description: `Help about any command`,
			}, {
				Name:        []string{"instance"},
				Description: `Commands to interact with instances on vultr`,
			}, {
				Name:        []string{"iso"},
				Description: `Iso is used to access iso commands`,
			}, {
				Name:        []string{"kubernetes"},
				Description: `Kubernetes is used to access kubernetes commands`,
			}, {
				Name:        []string{"load-balancer"},
				Description: `Load balancer commands`,
			}, {
				Name:        []string{"network"},
				Description: `Network interacts with network actions`,
			}, {
				Name:        []string{"object-storage"},
				Description: `Object storage commands`,
			}, {
				Name:        []string{"os"},
				Description: `Os is used to access os commands`,
			}, {
				Name:        []string{"plans"},
				Description: `Get information about Vultr plans`,
			}, {
				Name:        []string{"regions"},
				Description: `Get regions`,
			}, {
				Name:        []string{"reserved-ip"},
				Description: `Reserved-ip lets you interact with reserved-ip`,
			}, {
				Name:        []string{"script"},
				Description: `Startup script commands`,
			}, {
				Name:        []string{"snapshot"},
				Description: `Snapshot commands`,
			}, {
				Name:        []string{"ssh-key"},
				Description: `Ssh-key commands`,
			}, {
				Name:        []string{"user"},
				Description: `User commands`,
			}, {
				Name:        []string{"version"},
				Description: `Display current version of Vultr-cli`,
			}},
		}, {
			Name:        []string{"instance"},
			Description: `Commands to interact with instances on vultr`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"app"},
				Description: `Update application for an instance`,
			}, {
				Name:        []string{"backup"},
				Description: `List and create backup schedules for an instance`,
			}, {
				Name:        []string{"bandwidth"},
				Description: `Bandwidth for instance`,
			}, {
				Name:        []string{"create"},
				Description: `Create an instance`,
				Options: []model.Option{{
					Name:        []string{"--os", "-o"},
					Description: `ID of the operating system that will be installed on the server`,
					Args: []model.Arg{{
						Name: "osId",
					}},
				}, {
					Name:        []string{"--persistent_pxe", "-x"},
					Description: `Enable persistent_pxe`,
					Args: []model.Arg{{
						Name:        "persistent",
						Suggestions: []model.Suggestion{{Name: []string{`true`}}, {Name: []string{`false`}}},
					}},
				}, {
					Name:        []string{"--plan", "-p"},
					Description: `ID of the plan that the server will subscribe to`,
					Args: []model.Arg{{
						Name: "planId",
					}},
				}, {
					Name:        []string{"--region", "-r"},
					Description: `ID of the region where the server will be created`,
					Args: []model.Arg{{
						Name: "regionId",
					}},
				}},
			}, {
				Name:        []string{"delete", "destroy"},
				Description: `Delete/destroy an instance`,
				Args: []model.Arg{{
					Name:      "instanceID",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"get"},
				Description: `Get info about a specific instance`,
				Args: []model.Arg{{
					Name:      "instanceID",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"image"},
				Description: `Update image for an instance`,
			}, {
				Name:        []string{"ipv4"},
				Description: `List/create/delete ipv4 on instance`,
			}, {
				Name:        []string{"ipv6"},
				Description: `Commands for ipv6 on instance`,
			}, {
				Name:        []string{"iso"},
				Description: `Attach/detach ISOs to a given instance`,
			}, {
				Name:        []string{"label"},
				Description: `Label an instance`,
			}, {
				Name:        []string{"list"},
				Description: `List all available instances`,
			}, {
				Name:        []string{"os"},
				Description: `Update operating system for an instance`,
			}, {
				Name:        []string{"plan"},
				Description: `Update/list plans for an instance`,
			}, {
				Name:        []string{"reinstall"},
				Description: `Reinstall an instance`,
			}, {
				Name:        []string{"restart"},
				Description: `Restart an instance`,
			}, {
				Name:        []string{"restore"},
				Description: `Restore instance from backup/snapshot`,
			}, {
				Name:        []string{"reverse-dns"},
				Description: `Commands to handle reverse-dns on an instance`,
			}, {
				Name:        []string{"start"},
				Description: `Starts an instance`,
			}, {
				Name:        []string{"stop"},
				Description: `Stops an instance`,
			}, {
				Name:        []string{"tag"},
				Description: `Add/modify tag on instance`,
			}, {
				Name:        []string{"update-firewall-group"},
				Description: `Assign a firewall group to instance`,
			}, {
				Name:        []string{"user-data"},
				Description: `Commands to handle userdata on an instance`,
			}},
		}, {
			Name:        []string{"iso"},
			Description: `Iso is used to access iso commands`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"create"},
				Description: `Create iso from url`,
			}, {
				Name:        []string{"delete"},
				Description: `Delete a private iso`,
			}, {
				Name:        []string{"get"},
				Description: `Get private ISO <isoID>`,
				Args: []model.Arg{{
					Name: "isoID",
				}},
			}, {
				Name:        []string{"list"},
				Description: `List all private ISOs available`,
			}, {
				Name:        []string{"public"},
				Description: `List all public ISOs available`,
			}},
		}, {
			Name:        []string{"kubernetes", "k"},
			Description: `Kubernetes is used to access kubernetes commands`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"config"},
				Description: `Gets a kubernetes cluster's config`,
			}, {
				Name:        []string{"create"},
				Description: `Create kubernetes cluster`,
			}, {
				Name:        []string{"delete"},
				Description: `Delete a kubernetes cluster`,
			}, {
				Name:        []string{"delete-with-resources"},
				Description: `Delete a kubernetes cluster and related resources`,
			}, {
				Name:        []string{"get"},
				Description: `Retrieves a kubernetes cluster`,
			}, {
				Name:        []string{"list"},
				Description: `List kubernetes clusters`,
			}, {
				Name:        []string{"node-pool"},
				Description: `Node pools commands for a kubernetes cluster`,
			}, {
				Name:        []string{"update"},
				Description: `Updates a kubernetes cluster`,
			}, {
				Name:        []string{"versions"},
				Description: `Gets supported kubernetes versions`,
			}},
		}, {
			Name:        []string{"load-balancer", "lb"},
			Description: `Load balancer commands`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"create"},
				Description: `Create a load balancer`,
			}, {
				Name:        []string{"delete"},
				Description: `Deletes a load balancer`,
			}, {
				Name:        []string{"firewall-rule"},
				Description: `Get/list firewall rules for a load balancer`,
			}, {
				Name:        []string{"get"},
				Description: `Retrieves a load balancer`,
			}, {
				Name:        []string{"list"},
				Description: `Retrieves a list of active load balancers`,
			}, {
				Name:        []string{"rule"},
				Description: `Create/delete/list forwarding rules for a load balancer`,
			}, {
				Name:        []string{"update"},
				Description: `Updates a load balancer`,
			}},
		}, {
			Name:        []string{"network"},
			Description: `Network interacts with network actions`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"create"},
				Description: `Create a private network`,
			}, {
				Name:        []string{"delete"},
				Description: `Delete a private network`,
			}, {
				Name:        []string{"get"},
				Description: `Get a private network`,
			}, {
				Name:        []string{"list"},
				Description: `List all private networks`,
			}},
		}, {
			Name:        []string{"object-storage"},
			Description: `Object storage commands`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"create"},
				Description: `Create a new object storage subscription`,
			}, {
				Name:        []string{"delete"},
				Description: `Deletes an object storage subscription`,
			}, {
				Name:        []string{"get"},
				Description: `Retrieves a given object storage subscription`,
			}, {
				Name:        []string{"label"},
				Description: `Change the label for object storage subscription`,
			}, {
				Name:        []string{"list"},
				Description: `Retrieves a list of active object storage subscriptions`,
			}, {
				Name:        []string{"list-cluster"},
				Description: `Retrieves a list of object storage clusters`,
			}, {
				Name:        []string{"s3key-regenerate"},
				Description: `Regenerate the S3 API keys of an object storage subscription`,
			}},
		}, {
			Name:        []string{"os"},
			Description: `Os is used to access os commands`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"list"},
				Description: `List all available operating systems`,
			}},
		}, {
			Name:        []string{"plans", "p"},
			Description: `Get information about Vultr plans`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"list"},
				Description: `List plans`,
			}},
		}, {
			Name:        []string{"regions"},
			Description: `Get regions`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"availability"},
				Description: `List availability in region`,
			}, {
				Name:        []string{"list"},
				Description: `List regions`,
			}},
		}, {
			Name:        []string{"reserved-ip", "rip"},
			Description: `Reserved-ip lets you interact with reserved-ip`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"attach"},
				Description: `Attach a reservedIP to an instance`,
			}, {
				Name:        []string{"convert"},
				Description: `Convert IP address to reservedIP`,
			}, {
				Name:        []string{"create"},
				Description: `Create reservedIP`,
			}, {
				Name:        []string{"delete"},
				Description: `Delete a reserved ip`,
			}, {
				Name:        []string{"detach"},
				Description: `Detach a reservedIP to an instance`,
			}, {
				Name:        []string{"get"},
				Description: `Get a reserved IP`,
			}, {
				Name:        []string{"list"},
				Description: `List all reserved IPs`,
			}},
		}, {
			Name:        []string{"script", "ss"},
			Description: `Startup script commands`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"create"},
				Description: `Create a startup script`,
			}, {
				Name:        []string{"delete"},
				Description: `Delete a startup script`,
			}, {
				Name:        []string{"get"},
				Description: `Displays the contents of specified script`,
			}, {
				Name:        []string{"list"},
				Description: `List all startup scripts`,
			}, {
				Name:        []string{"update"},
				Description: `Update startup script`,
			}},
		}, {
			Name:        []string{"snapshot", "sn"},
			Description: `Snapshot commands`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"create"},
				Description: `Create a snapshot`,
			}, {
				Name:        []string{"create-url"},
				Description: `Create a snapshot from a URL`,
			}, {
				Name:        []string{"delete"},
				Description: `Delete a snapshot`,
			}, {
				Name:        []string{"get"},
				Description: `Get a snapshot`,
			}, {
				Name:        []string{"list"},
				Description: `List all snapshots`,
			}},
		}, {
			Name:        []string{"ssh-key", "ssh"},
			Description: `Ssh-key commands`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"create"},
				Description: `Create an SSH key`,
			}, {
				Name:        []string{"delete"},
				Description: `Delete an SSH key`,
			}, {
				Name:        []string{"get"},
				Description: `Get an SSH key`,
			}, {
				Name:        []string{"list"},
				Description: `List all SSH keys`,
			}, {
				Name:        []string{"update"},
				Description: `Update SSH key`,
			}},
		}, {
			Name:        []string{"user", "u"},
			Description: `User commands`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"create"},
				Description: `Create a user`,
			}, {
				Name:        []string{"delete"},
				Description: `Delete a user`,
			}, {
				Name:        []string{"get"},
				Description: `Get a user`,
			}, {
				Name:        []string{"list"},
				Description: `List all available users`,
			}, {
				Name:        []string{"update"},
				Description: `Update User`,
			}},
		}, {
			Name:        []string{"version"},
			Description: `Display current version of Vultr-cli`,
		}},
	}
}