// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["kafkactl"] = model.Subcommand{
		Name:        []string{"kafkactl"},
		Description: `Command-line interface for Apache Kafka`,
		Options: []model.Option{{
			Name:        []string{"--config-file", "-C"},
			Description: `Config file. one of: [$HOME/.config/kafkactl $HOME/.kafkactl $SNAP_REAL_HOME/.config/kafkactl $SNAP_DATA/kafkactl /etc/kafkactl]`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "config-file",
			}},
			IsPersistent: true,
		}, {
			Name:         []string{"--verbose", "-V"},
			Description:  `Verbose output`,
			IsPersistent: true,
		}, {
			Name:         []string{"--help", "-h"},
			Description:  `Display help`,
			IsPersistent: true,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"edit", "alter"},
			Description: `Alter topics, partitions`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"partition"},
				Description: `Alter a partition`,
				Options: []model.Option{{
					Name:        []string{"--replicas", "-r"},
					Description: `Set replicas for a partition`,
					Args: []model.Arg{{
						Name: "replicas",
					}},
				}, {
					Name:        []string{"--validate-only", "-v"},
					Description: `Validate only`,
				}},
			}, {
				Name:        []string{"topic"},
				Description: `Alter a topic`,
				Options: []model.Option{{
					Name:        []string{"--config", "-c"},
					Description: `Configs in format "key=value"`,
					Args: []model.Arg{{
						Name: "config",
					}},
				}, {
					Name:        []string{"--partitions", "-p"},
					Description: `Number of partitions`,
					Args: []model.Arg{{
						Name: "partitions",
					}},
				}, {
					Name:        []string{"--replication-factor", "-r"},
					Description: `Replication factor`,
					Args: []model.Arg{{
						Name: "replication-factor",
					}},
				}, {
					Name:        []string{"--validate-only", "-v"},
					Description: `Validate only`,
				}},
			}},
		}, {
			Name:        []string{"attach"},
			Description: `Run kafkactl pod in kubernetes and attach to it`,
		}, {
			Name:        []string{"clone"},
			Description: `Clone topics, consumerGroups`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"cg", "consumer-group"},
				Description: `Clone existing consumerGroup with all offsets`,
			}, {
				Name:        []string{"topic"},
				Description: `Clone existing topic (number of partitions, replication factor, config entries) to new one`,
			}},
		}, {
			Name:        []string{"completion"},
			Description: `Generate shell auto-completion file`,
		}, {
			Name:        []string{"config"},
			Description: `Show and edit configurations`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"currentContext", "current-context"},
				Description: `Show current context`,
			}, {
				Name:        []string{"getContexts", "get-contexts"},
				Description: `List configured contexts`,
				Options: []model.Option{{
					Name:        []string{"--output", "-o"},
					Description: `Output format. One of: compact`,
					Args: []model.Arg{{
						Name: "output",
					}},
				}},
			}, {
				Name:        []string{"useContext", "use-context"},
				Description: `Switch active context`,
			}, {
				Name:        []string{"view"},
				Description: `Show contents of config file`,
			}},
		}, {
			Name:        []string{"consume"},
			Description: `Consume messages from a topic`,
			Options: []model.Option{{
				Name:        []string{"--exit", "-e"},
				Description: `Stop consuming when latest offset is reached`,
			}, {
				Name:        []string{"--from-beginning", "-b"},
				Description: `Set offset for consumer to the oldest offset`,
			}, {
				Name:        []string{"--group", "-g"},
				Description: `Consumer group to join`,
				Args: []model.Arg{{
					Name: "group",
				}},
			}, {
				Name:        []string{"--key-encoding"},
				Description: `Key encoding (auto-detected by default). One of: none|hex|base64`,
				Args: []model.Arg{{
					Name: "key-encoding",
				}},
			}, {
				Name:        []string{"--key-proto-type"},
				Description: `Key protobuf message type`,
				Args: []model.Arg{{
					Name: "key-proto-type",
				}},
			}, {
				Name:        []string{"--max-messages"},
				Description: `Stop consuming after n messages have been read`,
				Args: []model.Arg{{
					Name: "max-messages",
				}},
			}, {
				Name:        []string{"--offset"},
				Description: `Offsets in format "partition=offset (for partitions not specified, other parameters apply)"`,
				Args: []model.Arg{{
					Name: "offset",
				}},
			}, {
				Name:        []string{"--output", "-o"},
				Description: `Output format. One of: json|yaml`,
				Args: []model.Arg{{
					Name: "output",
				}},
			}, {
				Name:        []string{"--partitions", "-p"},
				Description: `Partitions to consume. The default is to consume from all partitions`,
				Args: []model.Arg{{
					Name: "partitions",
				}},
			}, {
				Name:        []string{"--print-headers"},
				Description: `Print message headers`,
			}, {
				Name:        []string{"--print-keys", "-k"},
				Description: `Print message keys`,
			}, {
				Name:        []string{"--print-schema", "-a"},
				Description: `Print details about avro schema used for decoding`,
			}, {
				Name:        []string{"--print-timestamps", "-t"},
				Description: `Print message timestamps`,
			}, {
				Name:        []string{"--proto-file"},
				Description: `Additional protobuf description file for searching message description`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "proto-file",
				}},
			}, {
				Name:        []string{"--proto-import-path"},
				Description: `Additional path to search files listed in proto 'import' directive`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFolders},
					Name:      "proto-import-path",
				}},
			}, {
				Name:        []string{"--protoset-file"},
				Description: `Additional compiled protobuf description file for searching message description`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "protoset-file",
				}},
			}, {
				Name:        []string{"--separator", "-s"},
				Description: `Separator to split key and value`,
				Args: []model.Arg{{
					Name: "separator",
				}},
			}, {
				Name:        []string{"--tail"},
				Description: `Show only the last n messages on the topic`,
				Args: []model.Arg{{
					Name: "tail",
				}},
			}, {
				Name:        []string{"--value-encoding"},
				Description: `Value encoding (auto-detected by default). One of: none|hex|base64`,
				Args: []model.Arg{{
					Name: "value-encoding",
				}},
			}, {
				Name:        []string{"--value-proto-type"},
				Description: `Value protobuf message type`,
				Args: []model.Arg{{
					Name: "value-proto-type",
				}},
			}},
		}, {
			Name:        []string{"create"},
			Description: `Create topics, consumerGroups, acls`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"acl", "access-control-list"},
				Description: `Create an acl`,
				Options: []model.Option{{
					Name:        []string{"--allow", "-a"},
					Description: `Acl of permissionType 'allow' (choose this or 'deny')`,
				}, {
					Name:        []string{"--cluster", "-c"},
					Description: `Create acl for the cluster`,
				}, {
					Name:        []string{"--deny", "-d"},
					Description: `Acl of permissionType 'deny' (choose this or 'allow')`,
				}, {
					Name:        []string{"--group", "-g"},
					Description: `Create acl for a consumer group`,
					Args: []model.Arg{{
						Name: "group",
					}},
				}, {
					Name:        []string{"--host"},
					Description: `Hosts to allow`,
					Args: []model.Arg{{
						Name: "host",
					}},
				}, {
					Name:        []string{"--operation", "-o"},
					Description: `Operations of acl`,
					Args: []model.Arg{{
						Name: "operation",
					}},
				}, {
					Name:        []string{"--pattern"},
					Description: `Pattern type. one of (match, prefixed, literal)`,
					Args: []model.Arg{{
						Name: "pattern",
					}},
				}, {
					Name:        []string{"--principal", "-p"},
					Description: `Principal to be authenticated`,
					Args: []model.Arg{{
						Name: "principal",
					}},
				}, {
					Name:        []string{"--topic", "-t"},
					Description: `Create acl for a topic`,
					Args: []model.Arg{{
						Name: "topic",
					}},
				}, {
					Name:        []string{"--validate-only", "-v"},
					Description: `Validate only`,
				}},
			}, {
				Name:        []string{"cg", "consumer-group"},
				Description: `Create a consumerGroup`,
				Options: []model.Option{{
					Name:        []string{"--newest"},
					Description: `Set the offset to newest offset (for all partitions or the specified partition)`,
				}, {
					Name:        []string{"--offset"},
					Description: `Set offset to this value. offset with value -1 is ignored`,
					Args: []model.Arg{{
						Name: "offset",
					}},
				}, {
					Name:        []string{"--oldest"},
					Description: `Set the offset to oldest offset (for all partitions or the specified partition)`,
				}, {
					Name:        []string{"--partition", "-p"},
					Description: `Partition to create group for. -1 stands for all partitions`,
					Args: []model.Arg{{
						Name: "partition",
					}},
				}, {
					Name:        []string{"--topic", "-t"},
					Description: `One or more topics to create group for`,
					Args: []model.Arg{{
						Name: "topic",
					}},
				}},
			}, {
				Name:        []string{"topic"},
				Description: `Create a topic`,
				Options: []model.Option{{
					Name:        []string{"--config", "-c"},
					Description: `Configs in format "key=value"`,
					Args: []model.Arg{{
						Name: "config",
					}},
				}, {
					Name:        []string{"--partitions", "-p"},
					Description: `Number of partitions`,
					Args: []model.Arg{{
						Name: "partitions",
					}},
				}, {
					Name:        []string{"--replication-factor", "-r"},
					Description: `Replication factor`,
					Args: []model.Arg{{
						Name: "replication-factor",
					}},
				}, {
					Name:        []string{"--validate-only", "-v"},
					Description: `Validate only`,
				}},
			}},
		}, {
			Name:        []string{"delete"},
			Description: `Delete topics, consumerGroups, consumer-group-offset, acls`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"acl", "access-control-list"},
				Description: `Delete an acl`,
				Options: []model.Option{{
					Name:        []string{"--allow", "-a"},
					Description: `Acl of permissionType 'allow'`,
				}, {
					Name:        []string{"--cluster", "-c"},
					Description: `Delete acl for the cluster`,
				}, {
					Name:        []string{"--deny", "-d"},
					Description: `Acl of permissionType 'deny'`,
				}, {
					Name:        []string{"--groups", "-g"},
					Description: `Delete acl for a consumer group`,
				}, {
					Name:        []string{"--operation", "-o"},
					Description: `Operation of acl`,
					Args: []model.Arg{{
						Name: "operation",
					}},
				}, {
					Name:        []string{"--pattern"},
					Description: `Pattern type. one of (any, match, prefixed, literal)`,
					Args: []model.Arg{{
						Name: "pattern",
					}},
				}, {
					Name:        []string{"--topics", "-t"},
					Description: `Delete acl for a topic`,
				}, {
					Name:        []string{"--validate-only", "-v"},
					Description: `Validate only`,
				}},
			}, {
				Name:        []string{"consumer-groups", "consumer-group"},
				Description: `Delete a consumer-group`,
			}, {
				Name:        []string{"cgo", "offset", "consumer-group-offset"},
				Description: `Delete a consumer-group-offset`,
				Options: []model.Option{{
					Name:        []string{"--partition", "-p"},
					Description: `Delete offset for this partition. -1 stands for all partitions`,
					Args: []model.Arg{{
						Name: "partition",
					}},
				}, {
					Name:        []string{"--topic", "-t"},
					Description: `Delete offset for this topic`,
					Args: []model.Arg{{
						Name: "topic",
					}},
				}},
			}, {
				Name:        []string{"topic"},
				Description: `Delete a topic`,
			}},
		}, {
			Name:        []string{"describe"},
			Description: `Describe topics, consumerGroups, brokers`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"broker"},
				Description: `Describe a broker`,
				Options: []model.Option{{
					Name:        []string{"--output", "-o"},
					Description: `Output format. One of: json|yaml|wide`,
					Args: []model.Arg{{
						Name: "output",
					}},
				}},
			}, {
				Name:        []string{"cg", "consumer-group"},
				Description: `Describe a consumerGroup`,
				Options: []model.Option{{
					Name:        []string{"--only-with-lag", "-l"},
					Description: `Show only partitions that have a lag`,
				}, {
					Name:        []string{"--output", "-o"},
					Description: `Output format. One of: json|yaml|wide`,
					Args: []model.Arg{{
						Name: "output",
					}},
				}, {
					Name:        []string{"--print-members", "-m"},
					Description: `Print group members`,
				}, {
					Name:        []string{"--print-topics", "-T"},
					Description: `Print topic details`,
				}, {
					Name:        []string{"--topic", "-t"},
					Description: `Show group details for given topic only`,
					Args: []model.Arg{{
						Name: "topic",
					}},
				}},
			}, {
				Name:        []string{"topic"},
				Description: `Describe a topic`,
				Options: []model.Option{{
					Name:        []string{"--output", "-o"},
					Description: `Output format. One of: json|yaml|wide`,
					Args: []model.Arg{{
						Name: "output",
					}},
				}, {
					Name:        []string{"--print-configs", "-c"},
					Description: `Print configs`,
				}, {
					Name:        []string{"--skip-empty", "-s"},
					Description: `Show only partitions that have a messages`,
				}},
			}},
		}, {
			Name:        []string{"list", "get"},
			Description: `Get info about topics, consumerGroups, acls, brokers`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"acl", "access-control-list"},
				Description: `List available acls`,
				Options: []model.Option{{
					Name:        []string{"--allow", "-a"},
					Description: `Acl of permissionType 'allow'`,
				}, {
					Name:        []string{"--cluster", "-c"},
					Description: `List acl for the cluster`,
				}, {
					Name:        []string{"--deny", "-d"},
					Description: `Acl of permissionType 'deny'`,
				}, {
					Name:        []string{"--groups", "-g"},
					Description: `List acl for consumer groups`,
				}, {
					Name:        []string{"--operation"},
					Description: `Operation of acl`,
					Args: []model.Arg{{
						Name: "operation",
					}},
				}, {
					Name:        []string{"--output", "-o"},
					Description: `Output format. One of: json|yaml`,
					Args: []model.Arg{{
						Name: "output",
					}},
				}, {
					Name:        []string{"--pattern"},
					Description: `Pattern type. one of (any, match, prefixed, literal)`,
					Args: []model.Arg{{
						Name: "pattern",
					}},
				}, {
					Name:        []string{"--topics", "-t"},
					Description: `List acl for topics`,
				}},
			}, {
				Name:        []string{"brokers"},
				Description: `List brokers`,
				Options: []model.Option{{
					Name:        []string{"--output", "-o"},
					Description: `Output format. One of: json|yaml|compact`,
					Args: []model.Arg{{
						Name: "output",
					}},
				}},
			}, {
				Name:        []string{"cg", "consumer-groups"},
				Description: `List available consumerGroups`,
				Options: []model.Option{{
					Name:        []string{"--output", "-o"},
					Description: `Output format. One of: json|yaml|wide|compact`,
					Args: []model.Arg{{
						Name: "output",
					}},
				}, {
					Name:        []string{"--topic", "-t"},
					Description: `Show groups for given topic only`,
					Args: []model.Arg{{
						Name: "topic",
					}},
				}},
			}, {
				Name:        []string{"topics"},
				Description: `List available topics`,
				Options: []model.Option{{
					Name:        []string{"--output", "-o"},
					Description: `Output format. One of: json|yaml|wide|compact`,
					Args: []model.Arg{{
						Name: "output",
					}},
				}},
			}},
		}, {
			Name:        []string{"produce"},
			Description: `Produce messages to a topic`,
			Options: []model.Option{{
				Name:        []string{"--file", "-f"},
				Description: `File to read input from`,
				Args: []model.Arg{{
					Name: "file",
				}},
			}, {
				Name:        []string{"--header", "-H"},
				Description: `Headers in format "key:value"`,
				Args: []model.Arg{{
					Name: "header",
				}},
			}, {
				Name:        []string{"--key", "-k"},
				Description: `Key to use for all messages`,
				Args: []model.Arg{{
					Name: "key",
				}},
			}, {
				Name:        []string{"--key-encoding"},
				Description: `Key encoding (none by default). One of: none|hex|base64`,
				Args: []model.Arg{{
					Name: "key-encoding",
				}},
			}, {
				Name:        []string{"--key-proto-type"},
				Description: `Key protobuf message type`,
				Args: []model.Arg{{
					Name: "key-proto-type",
				}},
			}, {
				Name:        []string{"--key-schema-version", "-K"},
				Description: `Avro schema version that should be used for key serialization (default is latest)`,
				Args: []model.Arg{{
					Name: "key-schema-version",
				}},
			}, {
				Name:        []string{"--lineSeparator", "-L"},
				Description: `Separator to split multiple messages from stdin or file`,
				Args: []model.Arg{{
					Name: "lineSeparator",
				}},
			}, {
				Name:        []string{"--max-message-bytes"},
				Description: `The maximum permitted size of a message (defaults to 1000000)`,
				Args: []model.Arg{{
					Name: "max-message-bytes",
				}},
			}, {
				Name:        []string{"--null-value"},
				Description: `Produce a null value (can be used instead of providing a value with --value)`,
			}, {
				Name:        []string{"--partition", "-p"},
				Description: `Partition to produce to`,
				Args: []model.Arg{{
					Name: "partition",
				}},
			}, {
				Name:        []string{"--partitioner", "-P"},
				Description: `The partitioning scheme to use. Can be "murmur2", "hash", "hash-ref" "manual", or "random". (default is murmur2)`,
				Args: []model.Arg{{
					Name: "partitioner",
				}},
			}, {
				Name:        []string{"--proto-file"},
				Description: `Additional protobuf description file for searching message description`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "proto-file",
				}},
			}, {
				Name:        []string{"--proto-import-path"},
				Description: `Additional path to search files listed in proto 'import' directive`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFolders},
					Name:      "proto-import-path",
				}},
			}, {
				Name:        []string{"--protoset-file"},
				Description: `Additional compiled protobuf description file for searching message description`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "protoset-file",
				}},
			}, {
				Name:        []string{"--rate", "-r"},
				Description: `Amount of messages per second to produce on the topic`,
				Args: []model.Arg{{
					Name: "rate",
				}},
			}, {
				Name:        []string{"--required-acks"},
				Description: `Required acks. One of "NoResponse", "WaitForLocal", "WaitForAll". (default is WaitForLocal)`,
				Args: []model.Arg{{
					Name: "required-acks",
				}},
			}, {
				Name:        []string{"--separator", "-S"},
				Description: `Separator to split key and value from stdin or file`,
				Args: []model.Arg{{
					Name: "separator",
				}},
			}, {
				Name:        []string{"--silent", "-s"},
				Description: `Do not write to standard output`,
			}, {
				Name:        []string{"--value", "-v"},
				Description: `Value to produce`,
				Args: []model.Arg{{
					Name: "value",
				}},
			}, {
				Name:        []string{"--value-encoding"},
				Description: `Value encoding (none by default). One of: none|hex|base64`,
				Args: []model.Arg{{
					Name: "value-encoding",
				}},
			}, {
				Name:        []string{"--value-proto-type"},
				Description: `Value protobuf message type`,
				Args: []model.Arg{{
					Name: "value-proto-type",
				}},
			}, {
				Name:        []string{"--value-schema-version", "-i"},
				Description: `Avro schema version that should be used for value serialization (default is latest)`,
				Args: []model.Arg{{
					Name: "value-schema-version",
				}},
			}},
		}, {
			Name:        []string{"reset"},
			Description: `Reset consumerGroupsOffset`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"cgo", "offset", "consumer-group-offset"},
				Description: `Reset a consumer group offset`,
				Options: []model.Option{{
					Name:        []string{"--all-topics"},
					Description: `Do the operation for all topics in the consumer group`,
				}, {
					Name:        []string{"--execute", "-e"},
					Description: `Execute the reset (as default only the results are displayed for validation)`,
				}, {
					Name:        []string{"--newest"},
					Description: `Set the offset to newest offset (for all partitions or the specified partition)`,
				}, {
					Name:        []string{"--offset"},
					Description: `Set offset to this value. offset with value -1 is ignored`,
					Args: []model.Arg{{
						Name: "offset",
					}},
				}, {
					Name:        []string{"--oldest"},
					Description: `Set the offset to oldest offset (for all partitions or the specified partition)`,
				}, {
					Name:        []string{"--output", "-o"},
					Description: `Output format. One of: json|yaml`,
					Args: []model.Arg{{
						Name: "output",
					}},
				}, {
					Name:        []string{"--partition", "-p"},
					Description: `Partition to apply the offset. -1 stands for all partitions`,
					Args: []model.Arg{{
						Name: "partition",
					}},
				}, {
					Name:        []string{"--topic", "-t"},
					Description: `One ore more topics to change offset for`,
					Args: []model.Arg{{
						Name: "topic",
					}},
				}},
			}},
		}, {
			Name:        []string{"version"},
			Description: `Print the version of kafkactl`,
		}, {
			Name:        []string{"help"},
			Description: `Help about any command`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"edit", "alter"},
				Description: `Alter topics, partitions`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"partition"},
					Description: `Alter a partition`,
				}, {
					Name:        []string{"topic"},
					Description: `Alter a topic`,
				}},
			}, {
				Name:        []string{"attach"},
				Description: `Run kafkactl pod in kubernetes and attach to it`,
			}, {
				Name:        []string{"clone"},
				Description: `Clone topics, consumerGroups`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"cg", "consumer-group"},
					Description: `Clone existing consumerGroup with all offsets`,
				}, {
					Name:        []string{"topic"},
					Description: `Clone existing topic (number of partitions, replication factor, config entries) to new one`,
				}},
			}, {
				Name:        []string{"completion"},
				Description: `Generate shell auto-completion file`,
			}, {
				Name:        []string{"config"},
				Description: `Show and edit configurations`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"currentContext", "current-context"},
					Description: `Show current context`,
				}, {
					Name:        []string{"getContexts", "get-contexts"},
					Description: `List configured contexts`,
				}, {
					Name:        []string{"useContext", "use-context"},
					Description: `Switch active context`,
				}, {
					Name:        []string{"view"},
					Description: `Show contents of config file`,
				}},
			}, {
				Name:        []string{"consume"},
				Description: `Consume messages from a topic`,
			}, {
				Name:        []string{"create"},
				Description: `Create topics, consumerGroups, acls`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"acl", "access-control-list"},
					Description: `Create an acl`,
				}, {
					Name:        []string{"cg", "consumer-group"},
					Description: `Create a consumerGroup`,
				}, {
					Name:        []string{"topic"},
					Description: `Create a topic`,
				}},
			}, {
				Name:        []string{"delete"},
				Description: `Delete topics, consumerGroups, consumer-group-offset, acls`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"acl", "access-control-list"},
					Description: `Delete an acl`,
				}, {
					Name:        []string{"consumer-groups", "consumer-group"},
					Description: `Delete a consumer-group`,
				}, {
					Name:        []string{"cgo", "offset", "consumer-group-offset"},
					Description: `Delete a consumer-group-offset`,
				}, {
					Name:        []string{"topic"},
					Description: `Delete a topic`,
				}},
			}, {
				Name:        []string{"describe"},
				Description: `Describe topics, consumerGroups, brokers`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"broker"},
					Description: `Describe a broker`,
				}, {
					Name:        []string{"cg", "consumer-group"},
					Description: `Describe a consumerGroup`,
				}, {
					Name:        []string{"topic"},
					Description: `Describe a topic`,
				}},
			}, {
				Name:        []string{"list", "get"},
				Description: `Get info about topics, consumerGroups, acls, brokers`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"acl", "access-control-list"},
					Description: `List available acls`,
				}, {
					Name:        []string{"brokers"},
					Description: `List brokers`,
				}, {
					Name:        []string{"cg", "consumer-groups"},
					Description: `List available consumerGroups`,
				}, {
					Name:        []string{"topics"},
					Description: `List available topics`,
				}},
			}, {
				Name:        []string{"produce"},
				Description: `Produce messages to a topic`,
			}, {
				Name:        []string{"reset"},
				Description: `Reset consumerGroupsOffset`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"cgo", "offset", "consumer-group-offset"},
					Description: `Reset a consumer group offset`,
				}},
			}, {
				Name:        []string{"version"},
				Description: `Print the version of kafkactl`,
			}},
		}},
	}
}