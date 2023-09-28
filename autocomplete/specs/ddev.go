// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["ddev"] = model.Subcommand{
		Name:        []string{"ddev"},
		Description: `DDEV-Local local development environment`,
		Options: []model.Option{{
			Name:         []string{"--json-output", "-j"},
			Description:  `If true, user-oriented output will be in JSON format`,
			IsPersistent: true,
		}, {
			Name:         []string{"--help", "-h"},
			Description:  `Display help`,
			IsPersistent: true,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"auth"},
			Description: `A collection of authentication commands`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"ssh"},
				Description: `Add ssh key authentication to the ddev-ssh-auth container`,
				Options: []model.Option{{
					Name:        []string{"--ssh-key-path", "-d"},
					Description: `Full path to ssh key directory`,
					Args: []model.Arg{{
						Name: "ssh-key-path",
					}},
				}},
			}},
		}, {
			Name:        []string{"clean"},
			Description: `Removes items ddev has created`,
			Options: []model.Option{{
				Name:        []string{"--all", "-a"},
				Description: `Clean all ddev projects`,
			}, {
				Name:        []string{"--dry-run"},
				Description: `Run the clean command without deleting`,
			}},
		}, {
			Name:        []string{"composer"},
			Description: `Executes a composer command within the web container`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"create"},
				Description: `Executes 'composer create-project' within the web container with the arguments and flags provided`,
				Options: []model.Option{{
					Name:        []string{"--yes", "-y"},
					Description: `Yes - skip confirmation prompt`,
				}},
			}, {
				Name:        []string{"create-project"},
				Description: ``,
			}},
		}, {
			Name:        []string{"config"},
			Description: `Create or modify a ddev project configuration in the current directory`,
			Options: []model.Option{{
				Name:        []string{"--additional-fqdns"},
				Description: `A comma-delimited list of FQDNs for the project`,
				Args: []model.Arg{{
					Name: "additional-fqdns",
				}},
			}, {
				Name:        []string{"--additional-hostnames"},
				Description: `A comma-delimited list of hostnames for the project`,
				Args: []model.Arg{{
					Name: "additional-hostnames",
				}},
			}, {
				Name:        []string{"--apptype"},
				Description: `Provide the project type (one of backdrop, drupal10, drupal6, drupal7, drupal8, drupal9, laravel, magento, magento2, php, shopware6, typo3, wordpress). This is autodetected and this flag is necessary only to override the detection. This is the same as --project-type and is included only for backwards compatibility`,
				Args: []model.Arg{{
					Name: "apptype",
				}},
			}, {
				Name:        []string{"--auto"},
				Description: `Automatically run config without prompting`,
			}, {
				Name:        []string{"--bind-all-interfaces"},
				Description: `Bind host ports on all interfaces, not just on localhost network interface`,
			}, {
				Name:        []string{"--composer-root"},
				Description: `Overrides the default composer root directory for the web service`,
				Args: []model.Arg{{
					Name: "composer-root",
				}},
			}, {
				Name:        []string{"--composer-root-default"},
				Description: `Unsets a web service composer root directory override`,
			}, {
				Name:        []string{"--composer-version"},
				Description: `Specify override for composer version in web container. This may be "", "1", "2", "2.2", "stable", "preview", "snapshot" or a specific version`,
				Args: []model.Arg{{
					Name: "composer-version",
				}},
			}, {
				Name:        []string{"--create-docroot"},
				Description: `Prompts ddev to create the docroot if it doesn't exist`,
			}, {
				Name:        []string{"--database"},
				Description: `Specify the database type:version to use. Defaults to mariadb:10.4`,
				Args: []model.Arg{{
					Name: "database",
				}},
			}, {
				Name:        []string{"--db-image"},
				Description: `Sets the db container image`,
				Args: []model.Arg{{
					Name: "db-image",
				}},
			}, {
				Name:        []string{"--db-image-default"},
				Description: `Sets the default db container image for this ddev version`,
			}, {
				Name:        []string{"--db-working-dir"},
				Description: `Overrides the default working directory for the db service`,
				Args: []model.Arg{{
					Name: "db-working-dir",
				}},
			}, {
				Name:        []string{"--db-working-dir-default"},
				Description: `Unsets a db service working directory override`,
			}, {
				Name:        []string{"--dba-image"},
				Description: `Sets the dba container image`,
				Args: []model.Arg{{
					Name: "dba-image",
				}},
			}, {
				Name:        []string{"--dba-image-default"},
				Description: `Sets the default dba container image for this ddev version`,
			}, {
				Name:        []string{"--dba-working-dir"},
				Description: `Overrides the default working directory for the dba service`,
				Args: []model.Arg{{
					Name: "dba-working-dir",
				}},
			}, {
				Name:        []string{"--dba-working-dir-default"},
				Description: `Unsets a dba service working directory override`,
			}, {
				Name:        []string{"--dbimage-extra-packages"},
				Description: `A comma-delimited list of Debian packages that should be added to db container when the project is started`,
				Args: []model.Arg{{
					Name: "dbimage-extra-packages",
				}},
			}, {
				Name:        []string{"--default-container-timeout"},
				Description: `Default time in seconds that ddev waits for all containers to become ready on start`,
				Args: []model.Arg{{
					Name: "default-container-timeout",
				}},
			}, {
				Name:        []string{"--disable-settings-management"},
				Description: `Prevent ddev from creating or updating CMS settings files`,
			}, {
				Name:        []string{"--docroot"},
				Description: `Provide the relative docroot of the project, like 'docroot' or 'htdocs' or 'web', defaults to empty, the current directory`,
				Args: []model.Arg{{
					Name: "docroot",
				}},
			}, {
				Name:        []string{"--fail-on-hook-fail"},
				Description: `Decide whether 'ddev start' should be interrupted by a failing hook`,
			}, {
				Name:        []string{"--host-db-port"},
				Description: `The db container's localhost-bound port`,
				Args: []model.Arg{{
					Name: "host-db-port",
				}},
			}, {
				Name:        []string{"--host-dba-port"},
				Description: `The dba (PHPMyAdmin) container's localhost-bound port, if exposed via bind-all-interfaces`,
				Args: []model.Arg{{
					Name: "host-dba-port",
				}},
			}, {
				Name:        []string{"--host-https-port"},
				Description: `The web container's localhost-bound https port`,
				Args: []model.Arg{{
					Name: "host-https-port",
				}},
			}, {
				Name:        []string{"--host-webserver-port"},
				Description: `The web container's localhost-bound port`,
				Args: []model.Arg{{
					Name: "host-webserver-port",
				}},
			}, {
				Name:        []string{"--http-port"},
				Description: `The router HTTP port for this project`,
				Args: []model.Arg{{
					Name: "http-port",
				}},
			}, {
				Name:        []string{"--https-port"},
				Description: `The router HTTPS port for this project`,
				Args: []model.Arg{{
					Name: "https-port",
				}},
			}, {
				Name:        []string{"--image-defaults"},
				Description: `Sets the default web, db, and dba container images`,
			}, {
				Name:        []string{"--mailhog-https-port"},
				Description: `Router port to be used for mailhog access (https)`,
				Args: []model.Arg{{
					Name: "mailhog-https-port",
				}},
			}, {
				Name:        []string{"--mailhog-port"},
				Description: `Router port to be used for mailhog access`,
				Args: []model.Arg{{
					Name: "mailhog-port",
				}},
			}, {
				Name:        []string{"--mutagen-enabled"},
				Description: `Enable mutagen asynchronous update of project in web container`,
			}, {
				Name:        []string{"--nfs-mount-enabled"},
				Description: `Enable NFS mounting of project in container`,
			}, {
				Name:        []string{"--ngrok-args"},
				Description: `Provide extra args to ngrok in ddev share`,
				Args: []model.Arg{{
					Name: "ngrok-args",
				}},
			}, {
				Name:        []string{"--no-project-mount"},
				Description: `Whether or not to skip mounting project code into the web container`,
			}, {
				Name:        []string{"--nodejs-version"},
				Description: `Specify the nodejs version to use if you don't want the default NodeJS 16`,
				Args: []model.Arg{{
					Name: "nodejs-version",
				}},
			}, {
				Name:        []string{"--omit-containers"},
				Description: `A comma-delimited list of container types that should not be started when the project is started`,
				Args: []model.Arg{{
					Name: "omit-containers",
				}},
			}, {
				Name:        []string{"--php-version"},
				Description: `The version of PHP that will be enabled in the web container`,
				Args: []model.Arg{{
					Name: "php-version",
				}},
			}, {
				Name:        []string{"--phpmyadmin-https-port"},
				Description: `Router port to be used for PHPMyAdmin (dba) container access (https)`,
				Args: []model.Arg{{
					Name: "phpmyadmin-https-port",
				}},
			}, {
				Name:        []string{"--phpmyadmin-port"},
				Description: `Router port to be used for PHPMyAdmin (dba) container access`,
				Args: []model.Arg{{
					Name: "phpmyadmin-port",
				}},
			}, {
				Name:        []string{"--project-name"},
				Description: `Provide the project name of project to configure (normally the same as the last part of directory name)`,
				Args: []model.Arg{{
					Name: "project-name",
				}},
			}, {
				Name:        []string{"--project-tld"},
				Description: `Set the top-level domain to be used for projects, defaults to ddev.site`,
				Args: []model.Arg{{
					Name: "project-tld",
				}},
			}, {
				Name:        []string{"--project-type"},
				Description: `Provide the project type (one of backdrop, drupal10, drupal6, drupal7, drupal8, drupal9, laravel, magento, magento2, php, shopware6, typo3, wordpress). This is autodetected and this flag is necessary only to override the detection`,
				Args: []model.Arg{{
					Name: "project-type",
				}},
			}, {
				Name:        []string{"--projectname"},
				Description: `Provide the project name of project to configure (normally the same as the last part of directory name)`,
				Args: []model.Arg{{
					Name: "projectname",
				}},
			}, {
				Name:        []string{"--projecttype"},
				Description: `Provide the project type (one of backdrop, drupal10, drupal6, drupal7, drupal8, drupal9, laravel, magento, magento2, php, shopware6, typo3, wordpress). This is autodetected and this flag is necessary only to override the detection`,
				Args: []model.Arg{{
					Name: "projecttype",
				}},
			}, {
				Name:        []string{"--show-config-location"},
				Description: `Output the location of the config.yaml file if it exists, or error that it doesn't exist`,
			}, {
				Name:        []string{"--sitename"},
				Description: `Provide the project name of project to configure (normally the same as the last part of directory name) This is the same as project-name and is included only for backwards compatibility`,
				Args: []model.Arg{{
					Name: "sitename",
				}},
			}, {
				Name:        []string{"--timezone"},
				Description: `Specify timezone for containers and php, like Europe/London or America/Denver or GMT or UTC`,
				Args: []model.Arg{{
					Name: "timezone",
				}},
			}, {
				Name:        []string{"--upload-dir"},
				Description: `Sets the project's upload directory, the destination directory of the import-files command`,
				Args: []model.Arg{{
					Name: "upload-dir",
				}},
			}, {
				Name:        []string{"--use-dns-when-possible"},
				Description: `Use DNS for hostname resolution instead of /etc/hosts when possible`,
			}, {
				Name:        []string{"--web-environment"},
				Description: `Set the environment variables in the web container: --web-environment="TYPO3_CONTEXT=Development,SOMEENV=someval"`,
				Args: []model.Arg{{
					Name: "web-environment",
				}},
			}, {
				Name:        []string{"--web-environment-add"},
				Description: `Append environment variables to the web container: --web-environment="TYPO3_CONTEXT=Development,SOMEENV=someval"`,
				Args: []model.Arg{{
					Name: "web-environment-add",
				}},
			}, {
				Name:        []string{"--web-image"},
				Description: `Sets the web container image`,
				Args: []model.Arg{{
					Name: "web-image",
				}},
			}, {
				Name:        []string{"--web-image-default"},
				Description: `Sets the default web container image for this ddev version`,
			}, {
				Name:        []string{"--web-working-dir"},
				Description: `Overrides the default working directory for the web service`,
				Args: []model.Arg{{
					Name: "web-working-dir",
				}},
			}, {
				Name:        []string{"--web-working-dir-default"},
				Description: `Unsets a web service working directory override`,
			}, {
				Name:        []string{"--webimage-extra-packages"},
				Description: `A comma-delimited list of Debian packages that should be added to web container when the project is started`,
				Args: []model.Arg{{
					Name: "webimage-extra-packages",
				}},
			}, {
				Name:        []string{"--webserver-type"},
				Description: `Sets the project's desired webserver type: nginx-fpm or apache-fpm`,
				Args: []model.Arg{{
					Name: "webserver-type",
				}},
			}, {
				Name:        []string{"--working-dir-defaults"},
				Description: `Unsets all service working directory overrides`,
			}, {
				Name:        []string{"--xdebug-enabled"},
				Description: `Whether or not XDebug is enabled in the web container`,
			}},
			Subcommands: []model.Subcommand{{
				Name:        []string{"global"},
				Description: `Change global configuration`,
				Options: []model.Option{{
					Name:        []string{"--auto-restart-containers"},
					Description: `If true, automatically restart containers after a reboot or docker restart`,
				}, {
					Name:        []string{"--disable-http2"},
					Description: `Optionally disable http2 in ddev-router, 'ddev global --disable-http2' or "ddev global --disable-http2=false'`,
				}, {
					Name:        []string{"--fail-on-hook-fail"},
					Description: `If true, 'ddev start' will fail when a hook fails`,
				}, {
					Name:        []string{"--instrumentation-opt-in"},
					Description: `Instrumentation-opt-in=true`,
				}, {
					Name:        []string{"--internet-detection-timeout-ms"},
					Description: `Increase timeout when checking internet timeout, in milliseconds`,
					Args: []model.Arg{{
						Name: "internet-detection-timeout-ms",
					}},
				}, {
					Name:        []string{"--letsencrypt-email"},
					Description: `Email associated with Let's Encrypt, "ddev global --letsencrypt-email=me@example.com'`,
					Args: []model.Arg{{
						Name: "letsencrypt-email",
					}},
				}, {
					Name:        []string{"--mutagen-enabled"},
					Description: `If true, web container will use mutagen caching/asynchronous updates`,
				}, {
					Name:        []string{"--nfs-mount-enabled"},
					Description: `Enable NFS mounting on all projects globally`,
				}, {
					Name:        []string{"--no-bind-mounts"},
					Description: `If true, don't use bind-mounts - useful for environments like remote docker where bind-mounts are impossible`,
				}, {
					Name:        []string{"--omit-containers"},
					Description: `For example, --omit-containers=dba,ddev-ssh-agent`,
					Args: []model.Arg{{
						Name: "omit-containers",
					}},
				}, {
					Name:        []string{"--required-docker-compose-version"},
					Description: `Override default docker-compose version`,
					Args: []model.Arg{{
						Name: "required-docker-compose-version",
					}},
				}, {
					Name:        []string{"--router-bind-all-interfaces"},
					Description: `Router-bind-all-interfaces=true`,
				}, {
					Name:        []string{"--simple-formatting"},
					Description: `If true, use simple formatting and no color for tables`,
				}, {
					Name:        []string{"--table-style"},
					Description: `Table style for list and describe, see ~/.ddev/global_config.yaml for values`,
					Args: []model.Arg{{
						Name: "table-style",
					}},
				}, {
					Name:        []string{"--use-docker-compose-from-path"},
					Description: `If true, use docker-compose from path instead of private ~/.ddev/bin/docker-compose`,
				}, {
					Name:        []string{"--use-hardened-images"},
					Description: `If true, use more secure 'hardened' images for an actual internet deployment`,
				}, {
					Name:        []string{"--use-letsencrypt"},
					Description: `Enables experimental Let's Encrypt integration, 'ddev global --use-letsencrypt' or "ddev global --use-letsencrypt=false'`,
				}, {
					Name:        []string{"--web-environment"},
					Description: `Set the environment variables in the web container: --web-environment="TYPO3_CONTEXT=Development,SOMEENV=someval"`,
					Args: []model.Arg{{
						Name: "web-environment",
					}},
				}, {
					Name:        []string{"--web-environment-add"},
					Description: `Append environment variables to the web container: --web-environment="TYPO3_CONTEXT=Development,SOMEENV=someval"`,
					Args: []model.Arg{{
						Name: "web-environment-add",
					}},
				}},
			}},
		}, {
			Name:        []string{"d", "dbg", "debug"},
			Description: `A collection of debugging commands`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"capabilities"},
				Description: `Show capabilities of this version of ddev`,
			}, {
				Name:        []string{"check-db-match"},
				Description: `Verify that the database in the ddev-db server matches the configured type/version`,
			}, {
				Name:        []string{"compose-config"},
				Description: `Prints the docker-compose configuration of the current project`,
			}, {
				Name:        []string{"configyaml"},
				Description: `Prints the project config.*.yaml usage`,
			}, {
				Name:        []string{"dockercheck"},
				Description: `Diagnose DDEV docker/colima setup`,
			}, {
				Name:        []string{"download-images"},
				Description: `Download all images required by ddev`,
			}, {
				Name:        []string{"fix-commands"},
				Description: `Fix up custom/shell commands without running ddev start`,
			}, {
				Name:        []string{"get-volume-db-version"},
				Description: `Get the database type/version found in the ddev-dbserver's database volume, which may not be the same as the configured database type/version`,
			}, {
				Name:        []string{"migrate-database"},
				Description: `Migrate a mysql or mariadb database to a different dbtype:dbversion; works only with mysql and mariadb, not with postgres`,
			}, {
				Name:        []string{"mutagen"},
				Description: `Allows access to any mutagen command`,
			}, {
				Name:        []string{"nfsmount"},
				Description: `Checks to see if nfs mounting works for current project`,
			}, {
				Name:        []string{"refresh"},
				Description: `Refreshes docker cache for project`,
			}, {
				Name:        []string{"router-nginx-config"},
				Description: `Prints the nginx config of the router`,
			}, {
				Name:        []string{"test"},
				Description: `Run diagnostics on ddev using the embedded test_ddev.sh script`,
			}},
		}, {
			Name:        []string{"delete"},
			Description: `Remove all project information (including database) for an existing project`,
			Options: []model.Option{{
				Name:        []string{"--all", "-a"},
				Description: `Delete all projects`,
			}, {
				Name:        []string{"--clean-containers"},
				Description: `Clean up all ddev docker containers which are not required by this version of ddev`,
			}, {
				Name:        []string{"--omit-snapshot", "-O"},
				Description: `Omit/skip database snapshot`,
			}, {
				Name:        []string{"--yes", "-y"},
				Description: `Yes - skip confirmation prompt`,
			}},
			Subcommands: []model.Subcommand{{
				Name:        []string{"images"},
				Description: `Deletes drud/ddev-* docker images not in use by current ddev version`,
				Options: []model.Option{{
					Name:        []string{"--all", "-a"},
					Description: `If set, deletes all Docker images created by ddev`,
				}, {
					Name:        []string{"--yes", "-y"},
					Description: `Yes - skip confirmation prompt`,
				}},
			}},
		}, {
			Name:        []string{"desc", "st", "status", "describe"},
			Description: `Get a detailed description of a running ddev project`,
		}, {
			Name:        []string{".", "exec"},
			Description: `Execute a shell command in the container for a service. Uses the web service by default`,
			Options: []model.Option{{
				Name:        []string{"--dir", "-d"},
				Description: `Defines the execution directory within the container`,
				Args: []model.Arg{{
					Name: "dir",
				}},
			}, {
				Name:        []string{"--raw"},
				Description: `Use raw exec (do not interpret with bash inside container)`,
			}, {
				Name:        []string{"--service", "-s"},
				Description: `Defines the service to connect to. [e.g. web, db]`,
				Args: []model.Arg{{
					Name: "service",
				}},
			}},
		}, {
			Name:        []string{"export-db"},
			Description: `Dump a database to a file or to stdout`,
			Options: []model.Option{{
				Name:        []string{"--bzip2"},
				Description: `Use bzip2 compression`,
			}, {
				Name:        []string{"--file", "-f"},
				Description: `Provide the path to output the dump`,
				Args: []model.Arg{{
					Name: "file",
				}},
			}, {
				Name:        []string{"--gzip", "-z"},
				Description: `Use gzip compression`,
			}, {
				Name:        []string{"--target-db", "-d"},
				Description: `If provided, target-db is alternate database to export`,
				Args: []model.Arg{{
					Name: "target-db",
				}},
			}, {
				Name:        []string{"--xz"},
				Description: `Use xz compression`,
			}},
		}, {
			Name:        []string{"get"},
			Description: `Get/Download a 3rd party add-on (service, provider, etc.)`,
			Options: []model.Option{{
				Name:        []string{"--all"},
				Description: `List unofficial add-ons for 'ddev get' in addition to the official ones`,
			}, {
				Name:        []string{"--list"},
				Description: `List available add-ons for 'ddev get'`,
			}},
		}, {
			Name:        []string{"hostname"},
			Description: `Manage your hostfile entries`,
			Options: []model.Option{{
				Name:        []string{"--fire-bazooka"},
				Description: `Alias of --remove-inactive`,
			}, {
				Name:        []string{"--remove", "-r"},
				Description: `Remove the provided host name - ip correlation`,
			}, {
				Name:        []string{"--remove-inactive", "-R"},
				Description: `Remove host names of inactive projects`,
			}},
		}, {
			Name:        []string{"import-db"},
			Description: `Import a sql file into the project`,
			Options: []model.Option{{
				Name:        []string{"--extract-path"},
				Description: `If provided asset is an archive, provide the path to extract within the archive`,
				Args: []model.Arg{{
					Name: "extract-path",
				}},
			}, {
				Name:        []string{"--no-drop"},
				Description: `Set if you do NOT want to drop the db before importing`,
			}, {
				Name:        []string{"--progress", "-p"},
				Description: `Display a progress bar during import`,
			}, {
				Name:        []string{"--src", "-f"},
				Description: `Provide the path to a sql dump in .sql or tar/tar.gz/tgz/zip format`,
				Args: []model.Arg{{
					Name: "src",
				}},
			}, {
				Name:        []string{"--target-db", "-d"},
				Description: `If provided, target-db is alternate database to import into`,
				Args: []model.Arg{{
					Name: "target-db",
				}},
			}},
		}, {
			Name:        []string{"import-files"},
			Description: `Pull the uploaded files directory of an existing project to the default public upload directory of your project`,
			Options: []model.Option{{
				Name:        []string{"--extract-path"},
				Description: `If provided asset is an archive, optionally provide the path to extract within the archive`,
				Args: []model.Arg{{
					Name: "extract-path",
				}},
			}, {
				Name:        []string{"--src"},
				Description: `Provide the path to the source directory or tar/tar.gz/tgz/zip archive of files to import`,
				Args: []model.Arg{{
					Name: "src",
				}},
			}},
		}, {
			Name:        []string{"l", "ls", "list"},
			Description: `List projects`,
			Options: []model.Option{{
				Name:        []string{"--active-only", "-A"},
				Description: `If set, only currently active projects will be displayed`,
			}, {
				Name:        []string{"--continuous"},
				Description: `If set, project information will be emitted until the command is stopped`,
			}, {
				Name:        []string{"--continuous-sleep-interval", "-I"},
				Description: `Time in seconds between ddev list --continuous output lists`,
				Args: []model.Arg{{
					Name: "continuous-sleep-interval",
				}},
			}, {
				Name:        []string{"--wrap-table", "-W"},
				Description: `Display table with wrapped text if required`,
			}},
		}, {
			Name:        []string{"logs"},
			Description: `Get the logs from your running services`,
			Options: []model.Option{{
				Name:        []string{"--follow", "-f"},
				Description: `Follow the logs in real time`,
			}, {
				Name:        []string{"--service", "-s"},
				Description: `Defines the service to retrieve logs from. [e.g. web, db]`,
				Args: []model.Arg{{
					Name: "service",
				}},
			}, {
				Name:        []string{"--tail"},
				Description: `How many lines to show`,
				Args: []model.Arg{{
					Name: "tail",
				}},
			}, {
				Name:        []string{"--time", "-t"},
				Description: `Add timestamps to logs`,
			}},
		}, {
			Name:        []string{"mutagen"},
			Description: `Commands for mutagen status and sync, etc`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"monitor"},
				Description: `Monitor mutagen status`,
			}, {
				Name:        []string{"reset"},
				Description: `Reset mutagen for project`,
			}, {
				Name:        []string{"st", "status"},
				Description: `Shows mutagen sync status`,
				Options: []model.Option{{
					Name:        []string{"--verbose", "-l"},
					Description: `Extended/verbose output for mutagen status`,
				}},
			}, {
				Name:        []string{"sync"},
				Description: `Explicit sync for mutagen`,
				Options: []model.Option{{
					Name:        []string{"--verbose"},
					Description: `Extended/verbose output for mutagen status`,
				}},
			}},
		}, {
			Name:        []string{"sc", "stop-containers", "pause"},
			Description: `Uses 'docker stop' to pause/stop the containers belonging to a project`,
			Options: []model.Option{{
				Name:        []string{"--all", "-a"},
				Description: `Pause all running projects`,
			}},
		}, {
			Name:        []string{"powerdown", "poweroff"},
			Description: `Completely stop all projects and containers`,
		}, {
			Name:        []string{"pull"},
			Description: `Pull files and database using a configured provider plugin`,
		}, {
			Name:        []string{"push"},
			Description: `Push files and database using a configured provider plugin`,
		}, {
			Name:        []string{"restart"},
			Description: `Restart a project or several projects`,
			Options: []model.Option{{
				Name:        []string{"--all", "-a"},
				Description: `Restart all projects`,
			}},
		}, {
			Name:        []string{"service"},
			Description: `Add or remove, enable or disable extra services`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"disable"},
				Description: `Disable a 3rd party service`,
			}, {
				Name:        []string{"enable"},
				Description: `Enable a 3rd party service`,
			}},
		}, {
			Name:        []string{"share"},
			Description: `Share project on the internet via ngrok`,
			Options: []model.Option{{
				Name:        []string{"--subdomain"},
				Description: `Ngrok --subdomain argument, as in "ngrok --subdomain my-subdomain:, requires paid ngrok.com account"`,
				Args: []model.Arg{{
					Name: "subdomain",
				}},
			}},
		}, {
			Name:        []string{"snapshot"},
			Description: `Create a database snapshot for one or more projects`,
			Options: []model.Option{{
				Name:        []string{"--all", "-a"},
				Description: `Snapshot all projects. Will start the project if it is stopped or paused`,
			}, {
				Name:        []string{"--cleanup", "-C"},
				Description: `Cleanup snapshots`,
			}, {
				Name:        []string{"--list", "-l"},
				Description: `List snapshots`,
			}, {
				Name:        []string{"--name", "-n"},
				Description: `Provide a name for the snapshot`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"--yes", "-y"},
				Description: `Yes - skip confirmation prompt`,
			}},
		}, {
			Name:        []string{"ssh"},
			Description: `Starts a shell session in the container for a service. Uses web service by default`,
			Options: []model.Option{{
				Name:        []string{"--dir", "-d"},
				Description: `Defines the destination directory within the container`,
				Args: []model.Arg{{
					Name: "dir",
				}},
			}, {
				Name:        []string{"--service", "-s"},
				Description: `Defines the service to connect to. [e.g. web, db]`,
				Args: []model.Arg{{
					Name: "service",
				}},
			}},
		}, {
			Name:        []string{"add", "start"},
			Description: `Start a ddev project`,
			Options: []model.Option{{
				Name:        []string{"--all", "-a"},
				Description: `Start all projects`,
			}, {
				Name:        []string{"--select", "-s"},
				Description: `Interactively select a project to start`,
			}, {
				Name:        []string{"--skip-confirmation", "-y"},
				Description: `Skip any confirmation steps`,
			}},
		}, {
			Name:        []string{"remove", "rm", "stop"},
			Description: `Stop and remove the containers of a project. Does not lose or harm anything unless you add --remove-data`,
			Options: []model.Option{{
				Name:        []string{"--all", "-a"},
				Description: `Stop and remove all running or container-stopped projects and remove from global projects list`,
			}, {
				Name:        []string{"--omit-snapshot", "-O"},
				Description: `Omit/skip database snapshot`,
			}, {
				Name:        []string{"--remove-data", "-R"},
				Description: `Remove stored project data (MySQL, logs, etc.)`,
			}, {
				Name:        []string{"--select", "-s"},
				Description: `Interactively select a project to stop`,
			}, {
				Name:        []string{"--snapshot", "-S"},
				Description: `Create database snapshot`,
			}, {
				Name:        []string{"--stop-ssh-agent"},
				Description: `Stop the ddev-ssh-agent container`,
			}, {
				Name:        []string{"--unlist", "-U"},
				Description: `Remove the project from global project list, it won't show in ddev list until started again`,
			}},
		}, {
			Name:        []string{"version"},
			Description: `Print ddev version and component versions`,
		}, {
			Name:        []string{"help"},
			Description: `Help about any command`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"auth"},
				Description: `A collection of authentication commands`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"ssh"},
					Description: `Add ssh key authentication to the ddev-ssh-auth container`,
				}},
			}, {
				Name:        []string{"clean"},
				Description: `Removes items ddev has created`,
			}, {
				Name:        []string{"composer"},
				Description: `Executes a composer command within the web container`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"create"},
					Description: `Executes 'composer create-project' within the web container with the arguments and flags provided`,
				}, {
					Name:        []string{"create-project"},
					Description: ``,
				}},
			}, {
				Name:        []string{"config"},
				Description: `Create or modify a ddev project configuration in the current directory`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"global"},
					Description: `Change global configuration`,
				}},
			}, {
				Name:        []string{"d", "dbg", "debug"},
				Description: `A collection of debugging commands`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"capabilities"},
					Description: `Show capabilities of this version of ddev`,
				}, {
					Name:        []string{"check-db-match"},
					Description: `Verify that the database in the ddev-db server matches the configured type/version`,
				}, {
					Name:        []string{"compose-config"},
					Description: `Prints the docker-compose configuration of the current project`,
				}, {
					Name:        []string{"configyaml"},
					Description: `Prints the project config.*.yaml usage`,
				}, {
					Name:        []string{"dockercheck"},
					Description: `Diagnose DDEV docker/colima setup`,
				}, {
					Name:        []string{"download-images"},
					Description: `Download all images required by ddev`,
				}, {
					Name:        []string{"fix-commands"},
					Description: `Fix up custom/shell commands without running ddev start`,
				}, {
					Name:        []string{"get-volume-db-version"},
					Description: `Get the database type/version found in the ddev-dbserver's database volume, which may not be the same as the configured database type/version`,
				}, {
					Name:        []string{"migrate-database"},
					Description: `Migrate a mysql or mariadb database to a different dbtype:dbversion; works only with mysql and mariadb, not with postgres`,
				}, {
					Name:        []string{"mutagen"},
					Description: `Allows access to any mutagen command`,
				}, {
					Name:        []string{"nfsmount"},
					Description: `Checks to see if nfs mounting works for current project`,
				}, {
					Name:        []string{"refresh"},
					Description: `Refreshes docker cache for project`,
				}, {
					Name:        []string{"router-nginx-config"},
					Description: `Prints the nginx config of the router`,
				}, {
					Name:        []string{"test"},
					Description: `Run diagnostics on ddev using the embedded test_ddev.sh script`,
				}},
			}, {
				Name:        []string{"delete"},
				Description: `Remove all project information (including database) for an existing project`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"images"},
					Description: `Deletes drud/ddev-* docker images not in use by current ddev version`,
				}},
			}, {
				Name:        []string{"desc", "st", "status", "describe"},
				Description: `Get a detailed description of a running ddev project`,
			}, {
				Name:        []string{".", "exec"},
				Description: `Execute a shell command in the container for a service. Uses the web service by default`,
			}, {
				Name:        []string{"export-db"},
				Description: `Dump a database to a file or to stdout`,
			}, {
				Name:        []string{"get"},
				Description: `Get/Download a 3rd party add-on (service, provider, etc.)`,
			}, {
				Name:        []string{"hostname"},
				Description: `Manage your hostfile entries`,
			}, {
				Name:        []string{"import-db"},
				Description: `Import a sql file into the project`,
			}, {
				Name:        []string{"import-files"},
				Description: `Pull the uploaded files directory of an existing project to the default public upload directory of your project`,
			}, {
				Name:        []string{"l", "ls", "list"},
				Description: `List projects`,
			}, {
				Name:        []string{"logs"},
				Description: `Get the logs from your running services`,
			}, {
				Name:        []string{"mutagen"},
				Description: `Commands for mutagen status and sync, etc`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"monitor"},
					Description: `Monitor mutagen status`,
				}, {
					Name:        []string{"reset"},
					Description: `Reset mutagen for project`,
				}, {
					Name:        []string{"st", "status"},
					Description: `Shows mutagen sync status`,
				}, {
					Name:        []string{"sync"},
					Description: `Explicit sync for mutagen`,
				}},
			}, {
				Name:        []string{"sc", "stop-containers", "pause"},
				Description: `Uses 'docker stop' to pause/stop the containers belonging to a project`,
			}, {
				Name:        []string{"powerdown", "poweroff"},
				Description: `Completely stop all projects and containers`,
			}, {
				Name:        []string{"pull"},
				Description: `Pull files and database using a configured provider plugin`,
			}, {
				Name:        []string{"push"},
				Description: `Push files and database using a configured provider plugin`,
			}, {
				Name:        []string{"restart"},
				Description: `Restart a project or several projects`,
			}, {
				Name:        []string{"service"},
				Description: `Add or remove, enable or disable extra services`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"disable"},
					Description: `Disable a 3rd party service`,
				}, {
					Name:        []string{"enable"},
					Description: `Enable a 3rd party service`,
				}},
			}, {
				Name:        []string{"share"},
				Description: `Share project on the internet via ngrok`,
			}, {
				Name:        []string{"snapshot"},
				Description: `Create a database snapshot for one or more projects`,
			}, {
				Name:        []string{"ssh"},
				Description: `Starts a shell session in the container for a service. Uses web service by default`,
			}, {
				Name:        []string{"add", "start"},
				Description: `Start a ddev project`,
			}, {
				Name:        []string{"remove", "rm", "stop"},
				Description: `Stop and remove the containers of a project. Does not lose or harm anything unless you add --remove-data`,
			}, {
				Name:        []string{"version"},
				Description: `Print ddev version and component versions`,
			}},
		}},
	}
}