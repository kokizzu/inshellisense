// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["rustup"] = model.Subcommand{
		Name:        []string{"rustup"},
		Description: `The Rust toolchain installer`,
		Args: []model.Arg{{
			Name:        "toolchain",
			Description: `Release channel (e.g. +stable) or custom toolchain to set override`,
			IsOptional:  true,
		}},
		Options: []model.Option{{
			Name:        []string{"-h", "--help"},
			Description: `Prints help information`,
		}, {
			Name:        []string{"-V", "--version"},
			Description: `Prints version information`,
		}, {
			Name:        []string{"-q", "--quiet"},
			Description: `Disable progress output`,
		}, {
			Name:        []string{"-v", "--verbose"},
			Description: `Enable verbose output`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"show"},
			Description: `Show the active and installed toolchains or profiles`,
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Prints help information`,
			}},
			Subcommands: []model.Subcommand{{
				Name:        []string{"active-toolchain"},
				Description: `Show the active toolchain`,
			}, {
				Name:        []string{"home"},
				Description: `Display the computed value of RUSTUP_HOME`,
			}, {
				Name:        []string{"profile"},
				Description: `Show the current profile`,
			}, {
				Name:        []string{"keys"},
				Description: `Display the known PGP keys`,
			}, {
				Name:        []string{"help"},
				Description: `Prints this message or the help of the given subcommand(s)`,
			}},
		}, {
			Name:        []string{"update"},
			Description: `Update Rust toolchains and rustup`,
			Args: []model.Arg{{
				Name:        "toolchain",
				Description: `Toolchain name. For more information see "rustup help toolchain"`,
				Generator:   nil, // TODO: port over generator
				IsOptional:  true,
				IsVariadic:  true,
			}},
			Options: []model.Option{{
				Name:        []string{"--force"},
				Description: `Force an update, even if some components are missing`,
			}, {
				Name:        []string{"--force-non-host"},
				Description: `Install toolchains that require an emulator. See https://github.com/rust-lang/rustup/wiki/Non-host-toolchains`,
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Prints help information`,
			}, {
				Name:        []string{"--no-self-update"},
				Description: `Don't perform self update when running the "rustup update" command`,
			}},
		}, {
			Name:        []string{"check"},
			Description: `Check for updates to Rust toolchains and rustup`,
			Args: []model.Arg{{
				Name:        "toolchain",
				Description: `Toolchain name. For more information see "rustup help toolchain"`,
				IsOptional:  true,
			}},
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Prints help information`,
			}},
		}, {
			Name:        []string{"default"},
			Description: `Sets the default toolchain to the one specified. If the toolchain is not already installed then it is installed first`,
			Args: []model.Arg{{
				Name:        "toolchain",
				Description: `Toolchain name`,
				Generator:   nil, // TODO: port over generator
				IsOptional:  true,
			}},
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Prints help information`,
			}},
		}, {
			Name:        []string{"target"},
			Description: `Modify a toolchain's supported targets`,
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Prints help information`,
			}},
			Subcommands: []model.Subcommand{{
				Name:        []string{"list"},
				Description: `List installed and available targets`,
				Options: []model.Option{{
					Name:        []string{"-h", "--help"},
					Description: `Prints help information`,
				}, {
					Name:        []string{"-V", "--version"},
					Description: `Prints version information`,
				}},
			}, {
				Name:        []string{"add"},
				Description: `Add a target to a Rust toolchain`,
				Args: []model.Arg{{
					Name:        "target",
					Description: `Target triple`,
					Generator:   nil, // TODO: port over generator
				}},
				Options: []model.Option{{
					Name:        []string{"-h", "--help"},
					Description: `Prints help information`,
				}, {
					Name:        []string{"-V", "--version"},
					Description: `Prints version information`,
				}, {
					Name: []string{"--toolchain"},
					Args: []model.Arg{{
						Name:        "toolchain",
						Description: `Toolchain name`,
						Generator:   nil, // TODO: port over generator
					}},
				}},
			}, {
				Name:        []string{"remove"},
				Description: `Remove a target from a Rust toolchain`,
				Args: []model.Arg{{
					Name:        "target",
					Description: `Target triple`,
					Generator:   nil, // TODO: port over generator
				}},
				Options: []model.Option{{
					Name:        []string{"-h", "--help"},
					Description: `Prints help information`,
				}, {
					Name:        []string{"-V", "--version"},
					Description: `Prints version information`,
				}, {
					Name: []string{"--toolchain"},
					Args: []model.Arg{{
						Name:        "toolchain",
						Description: `Toolchain name`,
						Generator:   nil, // TODO: port over generator
					}},
				}},
			}, {
				Name:        []string{"help"},
				Description: `Prints this message or the help of the given subcommand(s)`,
				Args: []model.Arg{{
					Name:       "subcommand",
					IsOptional: true,
				}},
				Options: []model.Option{{
					Name:        []string{"-h", "--help"},
					Description: `Prints help information`,
				}, {
					Name:        []string{"-V", "--version"},
					Description: `Prints version information`,
				}},
			}},
		}, {
			Name:        []string{"toolchain"},
			Description: `Modify or query the installed toolchains`,
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Prints help information`,
			}},
			Subcommands: []model.Subcommand{{
				Name:        []string{"list"},
				Description: `List installed toolchains`,
				Options: []model.Option{{
					Name:        []string{"-h", "--help"},
					Description: `Prints help information`,
				}, {
					Name:        []string{"-V", "--version"},
					Description: `Prints version information`,
				}},
			}, {
				Name:        []string{"install"},
				Description: `Install or update a given toolchain`,
				Args: []model.Arg{{
					Name:        "toolchain",
					Description: `Toolchain name`,
					Generator:   nil, // TODO: port over generator
					IsVariadic:  true,
				}},
				Options: []model.Option{{
					Name:        []string{"-h", "--help"},
					Description: `Prints help information`,
				}, {
					Name:        []string{"-t", "--target"},
					Description: `Add specific targets on installation`,
					Args: []model.Arg{{
						Name:       "targets",
						Generator:  nil, // TODO: port over generator
						IsVariadic: true,
					}},
				}, {
					Name:        []string{"-c", "--component"},
					Description: `Add specific components on installation`,
					Args: []model.Arg{{
						Name:       "components",
						IsVariadic: true,
					}},
				}, {
					Name: []string{"--profile"},
					Args: []model.Arg{{
						Name:        "profile",
						Suggestions: []model.Suggestion{{Name: []string{`minimal`}}, {Name: []string{`default`}}, {Name: []string{`complete`}}},
					}},
				}, {
					Name:        []string{"--allow-downgrade"},
					Description: `Allow rustup to downgrade the toolchain to satisfy your component choice`,
				}, {
					Name:        []string{"--force"},
					Description: `Force an update, even if some components are missing`,
				}, {
					Name:        []string{"--no-self-update"},
					Description: `Don't perform self update when running the"rustup toolchain install" command`,
				}},
			}, {
				Name:        []string{"uninstall"},
				Description: `Uninstall a toolchain`,
				Args: []model.Arg{{
					Name:        "toolchain",
					Description: `Toolchain name`,
					Generator:   nil, // TODO: port over generator
					IsVariadic:  true,
				}},
				Options: []model.Option{{
					Name:        []string{"-h", "--help"},
					Description: `Prints help information`,
				}},
			}, {
				Name:        []string{"link"},
				Description: `Create a custom toolchain by symlinking to a directory`,
				Args: []model.Arg{{
					Name:        "toolchain",
					Description: `Custom toolchain name`,
					IsVariadic:  true,
				}, {
					Name:        "path",
					Description: `Path to the directory`,
				}},
				Options: []model.Option{{
					Name:        []string{"-h", "--help"},
					Description: `Prints help information`,
				}},
			}, {
				Name:        []string{"help"},
				Description: `Prints this message or the help of the given subcommand(s)`,
				Args: []model.Arg{{
					Name:       "subcommand",
					IsOptional: true,
				}},
				Options: []model.Option{{
					Name:        []string{"-h", "--help"},
					Description: `Prints help information`,
				}, {
					Name:        []string{"-V", "--version"},
					Description: `Prints version information`,
				}},
			}},
		}, {
			Name:        []string{"component"},
			Description: `Modify a toolchain's installed components`,
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Prints help information`,
			}},
			Subcommands: []model.Subcommand{{
				Name:        []string{"list"},
				Description: `List installed and available components`,
				Options: []model.Option{{
					Name:        []string{"-h", "--help"},
					Description: `Prints help information`,
				}, {
					Name:        []string{"-V", "--version"},
					Description: `Prints version information`,
				}},
			}, {
				Name:        []string{"add"},
				Description: `Add a component to a Rust toolchain`,
				Options: []model.Option{{
					Name:        []string{"-h", "--help"},
					Description: `Prints help information`,
				}, {
					Name:        []string{"-V", "--version"},
					Description: `Prints version information`,
				}},
			}, {
				Name:        []string{"remove"},
				Description: `Remove a component from a Rust toolchain`,
				Options: []model.Option{{
					Name:        []string{"-h", "--help"},
					Description: `Prints help information`,
				}, {
					Name:        []string{"-V", "--version"},
					Description: `Prints version information`,
				}},
			}, {
				Name:        []string{"help"},
				Description: `Prints this message or the help of the given subcommand(s)`,
				Args: []model.Arg{{
					Name:       "subcommand(s)",
					IsOptional: true,
				}},
			}},
		}, {
			Name:        []string{"override"},
			Description: `Modify directory toolchain overrides`,
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Prints help information`,
			}},
			Subcommands: []model.Subcommand{{
				Name:        []string{"list"},
				Description: `List directory toolchain overrides`,
				Options: []model.Option{{
					Name:        []string{"-h", "--help"},
					Description: `Prints help information`,
				}, {
					Name:        []string{"-V", "--version"},
					Description: `Prints version information`,
				}},
			}, {
				Name:        []string{"set"},
				Description: `Set the override toolchain for a directory`,
				Args: []model.Arg{{
					Name:        "toolchain",
					Description: `Toolchain name`,
					Generator:   nil, // TODO: port over generator
				}},
				Options: []model.Option{{
					Name:        []string{"-h", "--help"},
					Description: `Prints help information`,
				}, {
					Name:        []string{"-V", "--version"},
					Description: `Prints version information`,
				}},
			}, {
				Name:        []string{"unset"},
				Description: `Remove the override toolchain for a directory`,
				Options: []model.Option{{
					Name:        []string{"-h", "--help"},
					Description: `Prints help information`,
				}, {
					Name:        []string{"-V", "--version"},
					Description: `Prints version information`,
				}},
			}, {
				Name:        []string{"help"},
				Description: `Prints this message or the help of the given subcommand(s)`,
				Args: []model.Arg{{
					Name:       "subcommand(s)",
					IsOptional: true,
				}},
			}},
		}, {
			Name:        []string{"run"},
			Description: `Run a command with an environment configured for a given toolchain`,
			Args: []model.Arg{{
				Name:        "toolchain",
				Description: `Toolchain name`,
				Generator:   nil, // TODO: port over generator
			}, {
				Name:       "command",
				IsCommand:  true,
				IsVariadic: true,
			}},
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Prints help information`,
			}, {
				Name:        []string{"--install"},
				Description: `Install the requested toolchain if needed`,
			}},
		}, {
			Name:        []string{"which"},
			Description: `Display which binary will be run for a given command`,
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Prints help information`,
			}, {
				Name:        []string{"--toolchain"},
				Description: `Toolchain name`,
				Args: []model.Arg{{
					Name:      "toolchain",
					Generator: nil, // TODO: port over generator
				}},
			}},
		}, {
			Name:        []string{"doc"},
			Description: `Open the documentation for the current toolchain`,
			Args: []model.Arg{{
				Name:        "topic",
				Description: `Topic such as 'core', 'fn', 'usize', 'eprintln!', 'core::arch', 'alloc::format!', 'std::fs', 'std::fs::read_dir', 'std::io::Bytes', 'std::iter::Sum', 'std::io::error::Result' etc`,
				Generator:   nil, // TODO: port over generator
			}},
			Options: []model.Option{{
				Name:        []string{"--alloc"},
				Description: `The Rust core allocation and collections library`,
			}, {
				Name:        []string{"--book"},
				Description: `The Rust Programming Language book`,
			}, {
				Name:        []string{"--cargo"},
				Description: `The Cargo Book`,
			}, {
				Name:        []string{"--core"},
				Description: `The Rust Core Library`,
			}, {
				Name:        []string{"--edition-guide"},
				Description: `The Rust Edition Guide`,
			}, {
				Name:        []string{"--embedded-book"},
				Description: `The Embedded Rust Book`,
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Prints help information`,
			}, {
				Name:        []string{"--nomicon"},
				Description: `The Dark Arts of Advanced and Unsafe Rust Programming`,
			}, {
				Name:        []string{"--path"},
				Description: `Only print the path to the documentation`,
			}, {
				Name:        []string{"--proc_macro"},
				Description: `A support library for macro authors when defining new macros`,
			}, {
				Name:        []string{"--reference"},
				Description: `The Rust Reference`,
			}, {
				Name:        []string{"--rust-by-example"},
				Description: `A collection of runnable examples that illustrate various Rust concepts and standard libraries`,
			}, {
				Name:        []string{"--rustc"},
				Description: `The compiler for the Rust programming language`,
			}, {
				Name:        []string{"--rustdoc"},
				Description: `Generate documentation for Rust projects`,
			}, {
				Name:        []string{"--std"},
				Description: `Standard library API documentation`,
			}, {
				Name:        []string{"--test"},
				Description: `Support code for rustc's built in unit-test and micro-benchmarking framework`,
			}, {
				Name:        []string{"--unstable-book"},
				Description: `--unstable-book`,
			}, {
				Name:        []string{"--toolchain"},
				Description: `Toolchain name`,
				Args: []model.Arg{{
					Name:      "toolchain",
					Generator: nil, // TODO: port over generator
				}},
			}},
		}, {
			Name:        []string{"man"},
			Description: `View the man page for a given command`,
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Prints help information`,
			}, {
				Name:        []string{"--toolchain"},
				Description: `Toolchain name`,
				Args: []model.Arg{{
					Name:      "toolchain",
					Generator: nil, // TODO: port over generator
				}},
			}},
		}, {
			Name:        []string{"self"},
			Description: `Modify the rustup installation`,
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Prints help information`,
			}},
			Subcommands: []model.Subcommand{{
				Name:        []string{"update"},
				Description: `Download and install updates to rustup`,
				Options: []model.Option{{
					Name:        []string{"-h", "--help"},
					Description: `Prints help information`,
				}, {
					Name:        []string{"-V", "--version"},
					Description: `Prints version information`,
				}},
			}, {
				Name:        []string{"uninstall"},
				Description: `Uninstall rustup`,
				Options: []model.Option{{
					Name:        []string{"-h", "--help"},
					Description: `Prints help information`,
				}, {
					Name:        []string{"-V", "--version"},
					Description: `Prints version information`,
				}},
			}, {
				Name:        []string{"upgrade-data"},
				Description: `Upgrade the internal data format`,
				Options: []model.Option{{
					Name:        []string{"-h", "--help"},
					Description: `Prints help information`,
				}, {
					Name:        []string{"-V", "--version"},
					Description: `Prints version information`,
				}},
			}, {
				Name:        []string{"help"},
				Description: `Prints this message or the help of the given subcommand(s)`,
				Args: []model.Arg{{
					Name:       "subcommand(s)",
					IsOptional: true,
				}},
			}},
		}, {
			Name:        []string{"set"},
			Description: `Alter rustup settings`,
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Prints help information`,
			}},
			Subcommands: []model.Subcommand{{
				Name:        []string{"auto-self-update"},
				Description: `The rustup auto self update mode`,
				Args: []model.Arg{{
					Name:        "auto-self-update-mode",
					Suggestions: []model.Suggestion{{Name: []string{`enable`}}, {Name: []string{`disable`}}, {Name: []string{`check-only`}}},
					IsOptional:  true,
				}},
				Options: []model.Option{{
					Name:        []string{"-h", "--help"},
					Description: `Prints help information`,
				}, {
					Name:        []string{"-V", "--version"},
					Description: `Prints version information`,
				}},
			}, {
				Name: []string{"default-host"},
				Args: []model.Arg{{
					Name:      "host-triple",
					Generator: nil, // TODO: port over generator
				}},
				Options: []model.Option{{
					Name:        []string{"-h", "--help"},
					Description: `Prints help information`,
				}, {
					Name:        []string{"-V", "--version"},
					Description: `Prints version information`,
				}},
			}},
		}, {
			Name:        []string{"completions"},
			Description: `Generate tab-completion scripts for your shell`,
			Args: []model.Arg{{
				Name:        "shell",
				Suggestions: []model.Suggestion{{Name: []string{`zsh`}}, {Name: []string{`bash`}}, {Name: []string{`fish`}}, {Name: []string{`powershell`}}, {Name: []string{`elvish`}}},
			}, {
				Name:        "command",
				Suggestions: []model.Suggestion{{Name: []string{`rustup`}}, {Name: []string{`cargo`}}},
			}},
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Prints help information`,
			}},
		}, {
			Name:        []string{"help"},
			Description: `Prints this message or the help of the given subcommand(s)`,
			Args: []model.Arg{{
				Name:       "subcommand",
				IsOptional: true,
				IsVariadic: true,
			}},
		}},
	}
}