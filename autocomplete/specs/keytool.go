// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["keytool"] = model.Subcommand{
		Name:        []string{"keytool"},
		Description: `Key and Certificate Management Tool`,
		Options: []model.Option{{
			Name:        []string{"-h", "-?", "--help", "-help"},
			Description: `Show help message`,
		}, {
			Name:        []string{"-conf"},
			Description: `Specify pre-configured options file`,
			Args: []model.Arg{{
				Name: "url",
			}},
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"-certreq"},
			Description: `Generates a certificate request`,
			Options: []model.Option{{
				Name:        []string{"-h", "-?", "--help", "-help"},
				Description: `Show help message`,
			}, {
				Name:        []string{"-v"},
				Description: `Verbose output`,
			}, {
				Name:        []string{"-alias"},
				Description: `Alias name of the entry to process`,
				Args: []model.Arg{{
					Name: "alias",
				}},
			}, {
				Name:        []string{"-keystore"},
				Description: `Keystore name`,
				Args: []model.Arg{{
					Name: "keystore",
				}},
			}, {
				Name:        []string{"-storepass"},
				Description: `Keystore password`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-storetype"},
				Description: `Keystore type`,
				Args: []model.Arg{{
					Name: "type",
				}},
			}, {
				Name:        []string{"-providername"},
				Description: `Provider name`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"-addprovider"},
				Description: `Add security provider by name (e.g. SunPKCS11)`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"-providerclass"},
				Description: `Add security provider by fully-qualified class name`,
				Args: []model.Arg{{
					Name: "class",
				}},
			}, {
				Name:        []string{"-providerarg"},
				Description: `Configure argument for -addprovider or -providerclass`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-providerpath"},
				Description: `Provider classpath`,
				Args: []model.Arg{{
					Name: "list",
				}},
			}, {
				Name:        []string{"-protected"},
				Description: `Password through protected mechanism`,
			}, {
				Name:        []string{"-sigalg"},
				Description: `Signature algorithm name`,
				Args: []model.Arg{{
					Name: "alg",
				}},
			}, {
				Name:        []string{"-file"},
				Description: `Output file name`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "file",
				}},
			}, {
				Name:        []string{"-keypass"},
				Description: `Key password`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-dname"},
				Description: `Distinguished name`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"-ext"},
				Description: `X.509 extension`,
				Args: []model.Arg{{
					Name: "value",
				}},
			}},
		}, {
			Name:        []string{"-changealias"},
			Description: `Changes an entry's alias`,
			Options: []model.Option{{
				Name:        []string{"-h", "-?", "--help", "-help"},
				Description: `Show help message`,
			}, {
				Name:        []string{"-v"},
				Description: `Verbose output`,
			}, {
				Name:        []string{"-alias"},
				Description: `Alias name of the entry to process`,
				Args: []model.Arg{{
					Name: "alias",
				}},
			}, {
				Name:        []string{"-keystore"},
				Description: `Keystore name`,
				Args: []model.Arg{{
					Name: "keystore",
				}},
			}, {
				Name:        []string{"-storepass"},
				Description: `Keystore password`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-storetype"},
				Description: `Keystore type`,
				Args: []model.Arg{{
					Name: "type",
				}},
			}, {
				Name:        []string{"-providername"},
				Description: `Provider name`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"-addprovider"},
				Description: `Add security provider by name (e.g. SunPKCS11)`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"-providerclass"},
				Description: `Add security provider by fully-qualified class name`,
				Args: []model.Arg{{
					Name: "class",
				}},
			}, {
				Name:        []string{"-providerarg"},
				Description: `Configure argument for -addprovider or -providerclass`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-providerpath"},
				Description: `Provider classpath`,
				Args: []model.Arg{{
					Name: "list",
				}},
			}, {
				Name:        []string{"-protected"},
				Description: `Password through protected mechanism`,
			}, {
				Name:        []string{"-destalias"},
				Description: `Destination alias`,
				Args: []model.Arg{{
					Name: "alias",
				}},
			}, {
				Name:        []string{"-keypass"},
				Description: `Key password`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-cacerts"},
				Description: `Access the cacerts keystore`,
			}},
		}, {
			Name:        []string{"-delete"},
			Description: `Deletes an entry`,
			Options: []model.Option{{
				Name:        []string{"-h", "-?", "--help", "-help"},
				Description: `Show help message`,
			}, {
				Name:        []string{"-v"},
				Description: `Verbose output`,
			}, {
				Name:        []string{"-alias"},
				Description: `Alias name of the entry to process`,
				Args: []model.Arg{{
					Name: "alias",
				}},
			}, {
				Name:        []string{"-keystore"},
				Description: `Keystore name`,
				Args: []model.Arg{{
					Name: "keystore",
				}},
			}, {
				Name:        []string{"-storepass"},
				Description: `Keystore password`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-storetype"},
				Description: `Keystore type`,
				Args: []model.Arg{{
					Name: "type",
				}},
			}, {
				Name:        []string{"-providername"},
				Description: `Provider name`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"-addprovider"},
				Description: `Add security provider by name (e.g. SunPKCS11)`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"-providerclass"},
				Description: `Add security provider by fully-qualified class name`,
				Args: []model.Arg{{
					Name: "class",
				}},
			}, {
				Name:        []string{"-providerarg"},
				Description: `Configure argument for -addprovider or -providerclass`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-providerpath"},
				Description: `Provider classpath`,
				Args: []model.Arg{{
					Name: "list",
				}},
			}, {
				Name:        []string{"-protected"},
				Description: `Password through protected mechanism`,
			}, {
				Name:        []string{"-cacerts"},
				Description: `Access the cacerts keystore`,
			}},
		}, {
			Name:        []string{"-exportcert"},
			Description: `Exports certificate`,
			Options: []model.Option{{
				Name:        []string{"-h", "-?", "--help", "-help"},
				Description: `Show help message`,
			}, {
				Name:        []string{"-v"},
				Description: `Verbose output`,
			}, {
				Name:        []string{"-alias"},
				Description: `Alias name of the entry to process`,
				Args: []model.Arg{{
					Name: "alias",
				}},
			}, {
				Name:        []string{"-keystore"},
				Description: `Keystore name`,
				Args: []model.Arg{{
					Name: "keystore",
				}},
			}, {
				Name:        []string{"-storepass"},
				Description: `Keystore password`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-storetype"},
				Description: `Keystore type`,
				Args: []model.Arg{{
					Name: "type",
				}},
			}, {
				Name:        []string{"-providername"},
				Description: `Provider name`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"-addprovider"},
				Description: `Add security provider by name (e.g. SunPKCS11)`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"-providerclass"},
				Description: `Add security provider by fully-qualified class name`,
				Args: []model.Arg{{
					Name: "class",
				}},
			}, {
				Name:        []string{"-providerarg"},
				Description: `Configure argument for -addprovider or -providerclass`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-providerpath"},
				Description: `Provider classpath`,
				Args: []model.Arg{{
					Name: "list",
				}},
			}, {
				Name:        []string{"-protected"},
				Description: `Password through protected mechanism`,
			}, {
				Name:        []string{"-rfc"},
				Description: `Output in RFC style`,
			}, {
				Name:        []string{"-file"},
				Description: `Output file name`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "file",
				}},
			}, {
				Name:        []string{"-cacerts"},
				Description: `Access the cacerts keystore`,
			}},
		}, {
			Name:        []string{"-genkeypair"},
			Description: `Generate a key pair`,
			Options: []model.Option{{
				Name:        []string{"-h", "-?", "--help", "-help"},
				Description: `Show help message`,
			}, {
				Name:        []string{"-v"},
				Description: `Verbose output`,
			}, {
				Name:        []string{"-alias"},
				Description: `Alias name of the entry to process`,
				Args: []model.Arg{{
					Name: "alias",
				}},
			}, {
				Name:        []string{"-keystore"},
				Description: `Keystore name`,
				Args: []model.Arg{{
					Name: "keystore",
				}},
			}, {
				Name:        []string{"-storepass"},
				Description: `Keystore password`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-storetype"},
				Description: `Keystore type`,
				Args: []model.Arg{{
					Name: "type",
				}},
			}, {
				Name:        []string{"-providername"},
				Description: `Provider name`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"-addprovider"},
				Description: `Add security provider by name (e.g. SunPKCS11)`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"-providerclass"},
				Description: `Add security provider by fully-qualified class name`,
				Args: []model.Arg{{
					Name: "class",
				}},
			}, {
				Name:        []string{"-providerarg"},
				Description: `Configure argument for -addprovider or -providerclass`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-providerpath"},
				Description: `Provider classpath`,
				Args: []model.Arg{{
					Name: "list",
				}},
			}, {
				Name:        []string{"-protected"},
				Description: `Password through protected mechanism`,
			}, {
				Name:        []string{"-keyalg"},
				Description: `Key algorithm name`,
				Args: []model.Arg{{
					Name: "alg",
				}},
			}, {
				Name:        []string{"-keysize"},
				Description: `Key bit size`,
				Args: []model.Arg{{
					Name: "size",
				}},
			}, {
				Name:        []string{"-groupname"},
				Description: `Group name. For example, an Elliptic Curve name`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"-sigalg"},
				Description: `Signature algorithm name`,
				Args: []model.Arg{{
					Name: "alg",
				}},
			}, {
				Name:        []string{"-destalias"},
				Description: `Destination alias`,
				Args: []model.Arg{{
					Name: "alias",
				}},
			}, {
				Name:        []string{"-dname"},
				Description: `Distinguished name`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"-startdate"},
				Description: `Certificate validity start date/time`,
				Args: []model.Arg{{
					Name: "date",
				}},
			}, {
				Name:        []string{"-ext"},
				Description: `X.509 extension name`,
				Args: []model.Arg{{
					Name: "value",
				}},
			}, {
				Name:        []string{"-validity"},
				Description: `Validity number of days`,
				Args: []model.Arg{{
					Name: "days",
				}},
			}, {
				Name:        []string{"-keypass"},
				Description: `Key password`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}},
		}, {
			Name:        []string{"-genseckey"},
			Description: `Generates a secret key`,
			Options: []model.Option{{
				Name:        []string{"-h", "-?", "--help", "-help"},
				Description: `Show help message`,
			}, {
				Name:        []string{"-v"},
				Description: `Verbose output`,
			}, {
				Name:        []string{"-alias"},
				Description: `Alias name of the entry to process`,
				Args: []model.Arg{{
					Name: "alias",
				}},
			}, {
				Name:        []string{"-keystore"},
				Description: `Keystore name`,
				Args: []model.Arg{{
					Name: "keystore",
				}},
			}, {
				Name:        []string{"-storepass"},
				Description: `Keystore password`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-storetype"},
				Description: `Keystore type`,
				Args: []model.Arg{{
					Name: "type",
				}},
			}, {
				Name:        []string{"-providername"},
				Description: `Provider name`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"-addprovider"},
				Description: `Add security provider by name (e.g. SunPKCS11)`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"-providerclass"},
				Description: `Add security provider by fully-qualified class name`,
				Args: []model.Arg{{
					Name: "class",
				}},
			}, {
				Name:        []string{"-providerarg"},
				Description: `Configure argument for -addprovider or -providerclass`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-providerpath"},
				Description: `Provider classpath`,
				Args: []model.Arg{{
					Name: "list",
				}},
			}, {
				Name:        []string{"-protected"},
				Description: `Password through protected mechanism`,
			}, {
				Name:        []string{"-keypass"},
				Description: `Key password`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-keyalg"},
				Description: `Key algorithm name`,
				Args: []model.Arg{{
					Name: "alg",
				}},
			}, {
				Name:        []string{"-keysize"},
				Description: `Key bit size`,
				Args: []model.Arg{{
					Name: "size",
				}},
			}},
		}, {
			Name:        []string{"-gencert"},
			Description: `Generates certificate from a certificate request`,
			Options: []model.Option{{
				Name:        []string{"-h", "-?", "--help", "-help"},
				Description: `Show help message`,
			}, {
				Name:        []string{"-v"},
				Description: `Verbose output`,
			}, {
				Name:        []string{"-alias"},
				Description: `Alias name of the entry to process`,
				Args: []model.Arg{{
					Name: "alias",
				}},
			}, {
				Name:        []string{"-keystore"},
				Description: `Keystore name`,
				Args: []model.Arg{{
					Name: "keystore",
				}},
			}, {
				Name:        []string{"-storepass"},
				Description: `Keystore password`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-storetype"},
				Description: `Keystore type`,
				Args: []model.Arg{{
					Name: "type",
				}},
			}, {
				Name:        []string{"-providername"},
				Description: `Provider name`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"-addprovider"},
				Description: `Add security provider by name (e.g. SunPKCS11)`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"-providerclass"},
				Description: `Add security provider by fully-qualified class name`,
				Args: []model.Arg{{
					Name: "class",
				}},
			}, {
				Name:        []string{"-providerarg"},
				Description: `Configure argument for -addprovider or -providerclass`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-providerpath"},
				Description: `Provider classpath`,
				Args: []model.Arg{{
					Name: "list",
				}},
			}, {
				Name:        []string{"-protected"},
				Description: `Password through protected mechanism`,
			}, {
				Name:        []string{"-rfc"},
				Description: `Output in RFC style`,
			}, {
				Name:        []string{"-infile"},
				Description: `Input file name`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "file",
				}},
			}, {
				Name:        []string{"-outfile"},
				Description: `Output file name`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "file",
				}},
			}, {
				Name:        []string{"-sigalg"},
				Description: `Signature algorithm name`,
				Args: []model.Arg{{
					Name: "alg",
				}},
			}, {
				Name:        []string{"-dname"},
				Description: `Distinguished name`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"-startdate"},
				Description: `Certificate validity start date/time`,
				Args: []model.Arg{{
					Name: "date",
				}},
			}, {
				Name:        []string{"-ext"},
				Description: `X.509 extension name`,
				Args: []model.Arg{{
					Name: "value",
				}},
			}, {
				Name:        []string{"-validity"},
				Description: `Validity number of days`,
				Args: []model.Arg{{
					Name: "days",
				}},
			}, {
				Name:        []string{"-keypass"},
				Description: `Key password`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}},
		}, {
			Name:        []string{"-importcert"},
			Description: `Imports a certificate or a certificate chain`,
			Options: []model.Option{{
				Name:        []string{"-h", "-?", "--help", "-help"},
				Description: `Show help message`,
			}, {
				Name:        []string{"-v"},
				Description: `Verbose output`,
			}, {
				Name:        []string{"-alias"},
				Description: `Alias name of the entry to process`,
				Args: []model.Arg{{
					Name: "alias",
				}},
			}, {
				Name:        []string{"-keystore"},
				Description: `Keystore name`,
				Args: []model.Arg{{
					Name: "keystore",
				}},
			}, {
				Name:        []string{"-storepass"},
				Description: `Keystore password`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-storetype"},
				Description: `Keystore type`,
				Args: []model.Arg{{
					Name: "type",
				}},
			}, {
				Name:        []string{"-providername"},
				Description: `Provider name`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"-addprovider"},
				Description: `Add security provider by name (e.g. SunPKCS11)`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"-providerclass"},
				Description: `Add security provider by fully-qualified class name`,
				Args: []model.Arg{{
					Name: "class",
				}},
			}, {
				Name:        []string{"-providerarg"},
				Description: `Configure argument for -addprovider or -providerclass`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-providerpath"},
				Description: `Provider classpath`,
				Args: []model.Arg{{
					Name: "list",
				}},
			}, {
				Name:        []string{"-protected"},
				Description: `Password through protected mechanism`,
			}, {
				Name:        []string{"-noprompt"},
				Description: `Do not prompt`,
			}, {
				Name:        []string{"-trustcacerts"},
				Description: `Trust certificates from cacerts`,
			}, {
				Name:        []string{"-file"},
				Description: `Output file name`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "file",
				}},
			}, {
				Name:        []string{"-keypass"},
				Description: `Key password`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-cacerts"},
				Description: `Access the cacerts keystore`,
			}},
		}, {
			Name:        []string{"-importpass"},
			Description: `Imports a password`,
			Options: []model.Option{{
				Name:        []string{"-h", "-?", "--help", "-help"},
				Description: `Show help message`,
			}, {
				Name:        []string{"-v"},
				Description: `Verbose output`,
			}, {
				Name:        []string{"-alias"},
				Description: `Alias name of the entry to process`,
				Args: []model.Arg{{
					Name: "alias",
				}},
			}, {
				Name:        []string{"-keystore"},
				Description: `Keystore name`,
				Args: []model.Arg{{
					Name: "keystore",
				}},
			}, {
				Name:        []string{"-storepass"},
				Description: `Keystore password`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-storetype"},
				Description: `Keystore type`,
				Args: []model.Arg{{
					Name: "type",
				}},
			}, {
				Name:        []string{"-providername"},
				Description: `Provider name`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"-addprovider"},
				Description: `Add security provider by name (e.g. SunPKCS11)`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"-providerclass"},
				Description: `Add security provider by fully-qualified class name`,
				Args: []model.Arg{{
					Name: "class",
				}},
			}, {
				Name:        []string{"-providerarg"},
				Description: `Configure argument for -addprovider or -providerclass`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-providerpath"},
				Description: `Provider classpath`,
				Args: []model.Arg{{
					Name: "list",
				}},
			}, {
				Name:        []string{"-protected"},
				Description: `Password through protected mechanism`,
			}, {
				Name:        []string{"-keypass"},
				Description: `Key password`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-keyalg"},
				Description: `Key algorithm name`,
				Args: []model.Arg{{
					Name: "alg",
				}},
			}, {
				Name:        []string{"-keysize"},
				Description: `Key bit size`,
				Args: []model.Arg{{
					Name: "size",
				}},
			}},
		}, {
			Name:        []string{"-importkeystore"},
			Description: `Imports one or all entries from another keystore`,
			Options: []model.Option{{
				Name:        []string{"-h", "-?", "--help", "-help"},
				Description: `Show help message`,
			}, {
				Name:        []string{"-v"},
				Description: `Verbose output`,
			}, {
				Name:        []string{"-srckeystore"},
				Description: `Source keystore name`,
				Args: []model.Arg{{
					Name: "keystore",
				}},
			}, {
				Name:        []string{"-destkeystore"},
				Description: `Destination keystore name`,
				Args: []model.Arg{{
					Name: "keystore",
				}},
			}, {
				Name:        []string{"-srcstoretype"},
				Description: `Source keystore type`,
				Args: []model.Arg{{
					Name: "type",
				}},
			}, {
				Name:        []string{"-deststoretype"},
				Description: `Destination keystore type`,
				Args: []model.Arg{{
					Name: "type",
				}},
			}, {
				Name:        []string{"-srcstorepass"},
				Description: `Source keystore password`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-deststorepass"},
				Description: `Destination keystore password`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-srcprotected"},
				Description: `Source keystore password protected`,
			}, {
				Name:        []string{"-destprotected"},
				Description: `Destination keystore password protected`,
			}, {
				Name:        []string{"-srcprovidername"},
				Description: `Source keystore provider name`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"-destprovidername"},
				Description: `Destination keystore provider name`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"-srcalias"},
				Description: `Source alias`,
				Args: []model.Arg{{
					Name: "alias",
				}},
			}, {
				Name:        []string{"-destalias"},
				Description: `Destination alias`,
				Args: []model.Arg{{
					Name: "alias",
				}},
			}, {
				Name:        []string{"-srckeypass"},
				Description: `Source key password`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-destkeypass"},
				Description: `Destination key password`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-noprompt"},
				Description: `Do not prompt`,
			}, {
				Name:        []string{"-addprovider"},
				Description: `Add security provider by name (e.g. SunPKCS11)`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"-providerclass"},
				Description: `Add security provider by fully-qualified class name`,
				Args: []model.Arg{{
					Name: "class",
				}},
			}, {
				Name:        []string{"-providerarg"},
				Description: `Configure argument for -addprovider or -providerclass`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-providerpath"},
				Description: `Provider classpath`,
				Args: []model.Arg{{
					Name: "list",
				}},
			}},
		}, {
			Name:        []string{"-keypasswd"},
			Description: `Changes the key password of an entry`,
			Options: []model.Option{{
				Name:        []string{"-h", "-?", "--help", "-help"},
				Description: `Show help message`,
			}, {
				Name:        []string{"-v"},
				Description: `Verbose output`,
			}, {
				Name:        []string{"-alias"},
				Description: `Alias name of the entry to process`,
				Args: []model.Arg{{
					Name: "alias",
				}},
			}, {
				Name:        []string{"-keystore"},
				Description: `Keystore name`,
				Args: []model.Arg{{
					Name: "keystore",
				}},
			}, {
				Name:        []string{"-storepass"},
				Description: `Keystore password`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-storetype"},
				Description: `Keystore type`,
				Args: []model.Arg{{
					Name: "type",
				}},
			}, {
				Name:        []string{"-providername"},
				Description: `Provider name`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"-addprovider"},
				Description: `Add security provider by name (e.g. SunPKCS11)`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"-providerclass"},
				Description: `Add security provider by fully-qualified class name`,
				Args: []model.Arg{{
					Name: "class",
				}},
			}, {
				Name:        []string{"-providerarg"},
				Description: `Configure argument for -addprovider or -providerclass`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-providerpath"},
				Description: `Provider classpath`,
				Args: []model.Arg{{
					Name: "list",
				}},
			}, {
				Name:        []string{"-keypass"},
				Description: `Key password`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-new"},
				Description: `New password`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}},
		}, {
			Name:        []string{"-list"},
			Description: `Lists entries in a keystore`,
			Options: []model.Option{{
				Name:        []string{"-h", "-?", "--help", "-help"},
				Description: `Show help message`,
			}, {
				Name:        []string{"-v"},
				Description: `Verbose output`,
			}, {
				Name:        []string{"-alias"},
				Description: `Alias name of the entry to process`,
				Args: []model.Arg{{
					Name: "alias",
				}},
			}, {
				Name:        []string{"-keystore"},
				Description: `Keystore name`,
				Args: []model.Arg{{
					Name: "keystore",
				}},
			}, {
				Name:        []string{"-storepass"},
				Description: `Keystore password`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-storetype"},
				Description: `Keystore type`,
				Args: []model.Arg{{
					Name: "type",
				}},
			}, {
				Name:        []string{"-providername"},
				Description: `Provider name`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"-addprovider"},
				Description: `Add security provider by name (e.g. SunPKCS11)`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"-providerclass"},
				Description: `Add security provider by fully-qualified class name`,
				Args: []model.Arg{{
					Name: "class",
				}},
			}, {
				Name:        []string{"-providerarg"},
				Description: `Configure argument for -addprovider or -providerclass`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-providerpath"},
				Description: `Provider classpath`,
				Args: []model.Arg{{
					Name: "list",
				}},
			}, {
				Name:        []string{"-protected"},
				Description: `Password through protected mechanism`,
			}, {
				Name:        []string{"-rfc"},
				Description: `Output in RFC style`,
			}, {
				Name:        []string{"-cacerts"},
				Description: `Access the cacerts keystore`,
			}},
		}, {
			Name:        []string{"-printcert"},
			Description: `Prints the content of a certificate`,
			Options: []model.Option{{
				Name:        []string{"-h", "-?", "--help", "-help"},
				Description: `Show help message`,
			}, {
				Name:        []string{"-v"},
				Description: `Verbose output`,
			}, {
				Name:        []string{"-rfc"},
				Description: `Output in RFC style`,
			}, {
				Name:        []string{"-file"},
				Description: `Input file name`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "file",
				}},
			}, {
				Name:        []string{"-sslserver"},
				Description: `SSL server host and port`,
				Args: []model.Arg{{
					Name: "server:[port]",
				}},
			}, {
				Name:        []string{"-jarfile"},
				Description: `Signed jar file`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "file",
				}},
			}},
		}, {
			Name:        []string{"-printcertreq"},
			Description: `Prints the content of a certificate request`,
			Options: []model.Option{{
				Name:        []string{"-h", "-?", "--help", "-help"},
				Description: `Show help message`,
			}, {
				Name:        []string{"-v"},
				Description: `Verbose output`,
			}, {
				Name:        []string{"-file"},
				Description: `Input file name`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "file",
				}},
			}},
		}, {
			Name:        []string{"-printcrl"},
			Description: `Prints the content of a CRL file`,
			Options: []model.Option{{
				Name:        []string{"-h", "-?", "--help", "-help"},
				Description: `Show help message`,
			}, {
				Name:        []string{"-v"},
				Description: `Verbose output`,
			}, {
				Name:        []string{"-file"},
				Description: `Input file name`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "file",
				}},
			}},
		}, {
			Name:        []string{"-storepasswd"},
			Description: `Changes the store password of a keystore`,
			Options: []model.Option{{
				Name:        []string{"-h", "-?", "--help", "-help"},
				Description: `Show help message`,
			}, {
				Name:        []string{"-v"},
				Description: `Verbose output`,
			}, {
				Name:        []string{"-alias"},
				Description: `Alias name of the entry to process`,
				Args: []model.Arg{{
					Name: "alias",
				}},
			}, {
				Name:        []string{"-keystore"},
				Description: `Keystore name`,
				Args: []model.Arg{{
					Name: "keystore",
				}},
			}, {
				Name:        []string{"-storepass"},
				Description: `Keystore password`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-storetype"},
				Description: `Keystore type`,
				Args: []model.Arg{{
					Name: "type",
				}},
			}, {
				Name:        []string{"-providername"},
				Description: `Provider name`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"-addprovider"},
				Description: `Add security provider by name (e.g. SunPKCS11)`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"-providerclass"},
				Description: `Add security provider by fully-qualified class name`,
				Args: []model.Arg{{
					Name: "class",
				}},
			}, {
				Name:        []string{"-providerarg"},
				Description: `Configure argument for -addprovider or -providerclass`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-providerpath"},
				Description: `Provider classpath`,
				Args: []model.Arg{{
					Name: "list",
				}},
			}, {
				Name:        []string{"-new"},
				Description: `New password`,
				Args: []model.Arg{{
					Name: "arg",
				}},
			}, {
				Name:        []string{"-cacerts"},
				Description: `Access the cacerts keystore`,
			}},
		}},
	}
}