// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["lp"] = model.Subcommand{
		Name:        []string{"lp"},
		Description: `Print files`,
		Args: []model.Arg{{
			Templates:   []model.Template{model.TemplateFilepaths},
			Description: `Filepath you want to print`,
		}},
		Options: []model.Option{{
			Name:        []string{"-E"},
			Description: `Force encryption when connecting to the server`,
		}, {
			Name:        []string{"-U"},
			Description: `Specifies the username to use when connecting to the server`,
			Args: []model.Arg{{
				Name: "Username",
			}},
		}, {
			Name:        []string{"-c"},
			Description: `This  option is provided for backwards-compatibility only. On systems that support it, this option forces the print file to be copied to the spool directory before printing.  In CUPS, print files are always sent to the scheduler via IPP which has the  same effect`,
		}, {
			Name:        []string{"-d"},
			Description: `Print files to the named printer`,
			Args: []model.Arg{{
				Name: "Destination printer name",
			}},
		}, {
			Name:        []string{"-h"},
			Description: `Chooses an alternate server`,
			Args: []model.Arg{{
				Name: "hostname[:port]",
			}},
		}, {
			Name:        []string{"-i"},
			Description: `Specifies an existing job to modify`,
			Args: []model.Arg{{
				Name: "job-id",
			}},
		}, {
			Name:        []string{"-m"},
			Description: `Sends an email when the job is completed`,
		}, {
			Name:        []string{"-n"},
			Description: `Sets the number of copies to print`,
			Args: []model.Arg{{
				Name: "copies",
			}},
		}, {
			Name:        []string{"-q"},
			Description: `Sets the job priority from 1 (lowest) to 100 (highest). The default priority is 50`,
			Args: []model.Arg{{
				Name:        "priority",
				Suggestions: []model.Suggestion{{Name: []string{`1`}}, {Name: []string{`2`}}, {Name: []string{`3`}}, {Name: []string{`4`}}, {Name: []string{`5`}}, {Name: []string{`6`}}, {Name: []string{`7`}}, {Name: []string{`8`}}, {Name: []string{`9`}}, {Name: []string{`10`}}, {Name: []string{`11`}}, {Name: []string{`12`}}, {Name: []string{`13`}}, {Name: []string{`14`}}, {Name: []string{`15`}}, {Name: []string{`16`}}, {Name: []string{`17`}}, {Name: []string{`18`}}, {Name: []string{`19`}}, {Name: []string{`20`}}, {Name: []string{`21`}}, {Name: []string{`22`}}, {Name: []string{`23`}}, {Name: []string{`24`}}, {Name: []string{`25`}}, {Name: []string{`26`}}, {Name: []string{`27`}}, {Name: []string{`28`}}, {Name: []string{`29`}}, {Name: []string{`30`}}, {Name: []string{`31`}}, {Name: []string{`32`}}, {Name: []string{`33`}}, {Name: []string{`34`}}, {Name: []string{`35`}}, {Name: []string{`36`}}, {Name: []string{`37`}}, {Name: []string{`38`}}, {Name: []string{`39`}}, {Name: []string{`40`}}, {Name: []string{`41`}}, {Name: []string{`42`}}, {Name: []string{`43`}}, {Name: []string{`44`}}, {Name: []string{`45`}}, {Name: []string{`46`}}, {Name: []string{`47`}}, {Name: []string{`48`}}, {Name: []string{`49`}}, {Name: []string{`50`}}, {Name: []string{`51`}}, {Name: []string{`52`}}, {Name: []string{`53`}}, {Name: []string{`54`}}, {Name: []string{`55`}}, {Name: []string{`56`}}, {Name: []string{`57`}}, {Name: []string{`58`}}, {Name: []string{`59`}}, {Name: []string{`60`}}, {Name: []string{`61`}}, {Name: []string{`62`}}, {Name: []string{`63`}}, {Name: []string{`64`}}, {Name: []string{`65`}}, {Name: []string{`66`}}, {Name: []string{`67`}}, {Name: []string{`68`}}, {Name: []string{`69`}}, {Name: []string{`70`}}, {Name: []string{`71`}}, {Name: []string{`72`}}, {Name: []string{`73`}}, {Name: []string{`74`}}, {Name: []string{`75`}}, {Name: []string{`76`}}, {Name: []string{`77`}}, {Name: []string{`78`}}, {Name: []string{`79`}}, {Name: []string{`80`}}, {Name: []string{`81`}}, {Name: []string{`82`}}, {Name: []string{`83`}}, {Name: []string{`84`}}, {Name: []string{`85`}}, {Name: []string{`86`}}, {Name: []string{`87`}}, {Name: []string{`88`}}, {Name: []string{`89`}}, {Name: []string{`90`}}, {Name: []string{`91`}}, {Name: []string{`92`}}, {Name: []string{`93`}}, {Name: []string{`94`}}, {Name: []string{`95`}}, {Name: []string{`96`}}, {Name: []string{`97`}}, {Name: []string{`98`}}, {Name: []string{`99`}}, {Name: []string{`100`}}},
			}},
		}, {
			Name:        []string{"-s"},
			Description: `Do not report the resulting job IDs (silent mode)`,
		}, {
			Name:        []string{"-t"},
			Description: `Sets the job name`,
			Args: []model.Arg{{
				Name: "name",
			}},
		}, {
			Name:        []string{"-H"},
			Description: `Specifies when the job should be printed.  A value of immediate will print the file immediately, a value of hold will hold the job indefinitely, and a UTC time value (HH:MM) will hold the job until the specified UTC (not local) time. Use a value of resume with the -i option to resume a held job.  Use a value of restart with the -i option to restart a completed job`,
			Args: []model.Arg{{
				Name: "pages",
			}},
		}, {
			Name:        []string{"-P"},
			Description: `Specifies which pages to print in the document. The list can contain a list of numbers and ranges (#-#) separated by commas, e.g., "1,3-5,16". The page numbers refer to the output pages and not the document's original  pages - options like  "number-up" can affect the numbering of the pages`,
			Args: []model.Arg{{
				Name:        "hh:mm",
				Suggestions: []model.Suggestion{{Name: []string{`hold`}}, {Name: []string{`immediate`}}, {Name: []string{`restart`}}, {Name: []string{`resume`}}},
			}},
		}, {
			Name:        []string{"-o"},
			Description: `Sets one or more job options`,
			Args: []model.Arg{{
				Name: "name=value",
			}},
		}, {
			Name:        []string{"--help"},
			Description: `Show help for lp`,
		}},
	}
}