// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["pytest"] = model.Subcommand{
		Name: []string{"pytest"},
		Args: []model.Arg{{
			Templates:   []model.Template{model.TemplateFilepaths, model.TemplateFolders},
			Name:        "File or Directory",
			Description: `The test file or directory you want to run pytest on. If omitted, pytest will run all       of the files of the form test_*.py or *_test.py in the current directory       and its subdirectories`,
			IsOptional:  true,
		}},
		Options: []model.Option{{
			Name:        []string{"--assert"},
			Description: `Control assertion debugging tools. 'plain' performs no assertion debugging. 'rewrite' (the default) rewrites assert statements in test modules on import to provide assert expression information`,
			Args: []model.Arg{{
				Name:        "Mode",
				Suggestions: []model.Suggestion{{Name: []string{`plain`}}, {Name: []string{`rewrite`}}},
			}},
		}, {
			Name:        []string{"--basetemp"},
			Description: `Base temporary directory for this test run.(warning: this directory is removed if it exists)`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFolders},
				Name:      "Directory",
			}},
		}, {
			Name:        []string{"-c"},
			Description: `Load configuration from "file" instead of trying to locate one of the implicit configuration files`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "File",
			}},
		}, {
			Name:        []string{"--cache-clear"},
			Description: `Remove all cache contents at start of test run`,
		}, {
			Name:        []string{"--cache-show"},
			Description: `Show cache contents, don't perform collection or tests. Optional argument: glob (default: '*')`,
			Args: []model.Arg{{
				Name:       "Glob",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"--capture"},
			Description: `Per-test capturing method`,
			Args: []model.Arg{{
				Name:        "Method",
				Description: `One of fd|sys|no|tee-sys`,
				Suggestions: []model.Suggestion{{Name: []string{`fd`}}, {Name: []string{`sys`}}, {Name: []string{`no`}}, {Name: []string{`tee-sys`}}},
			}},
		}, {
			Name:        []string{"--code-highlight"},
			Description: `Whether code should be highlighted (only if --color is also enabled)`,
			Args: []model.Arg{{
				Name:        "Highlight",
				Suggestions: []model.Suggestion{{Name: []string{`yes`}}, {Name: []string{`no`}}},
			}},
		}, {
			Name:        []string{"--co", "--collect-only"},
			Description: `Only collect tests, don't execute them`,
		}, {
			Name:        []string{"--collect-in-virtualenv"},
			Description: `Don't ignore tests in a local virtualenv directory`,
		}, {
			Name:        []string{"--color"},
			Description: `Color terminal output`,
			Args: []model.Arg{{
				Name:        "Color",
				Suggestions: []model.Suggestion{{Name: []string{`yes`}}, {Name: []string{`no`}}, {Name: []string{`auto`}}},
			}},
		}, {
			Name:        []string{"--confcutdir"},
			Description: `Only load conftest.py's relative to specified dir`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFolders},
				Name:      "Dir",
			}},
		}, {
			Name:        []string{"--continue-on-collection-errors"},
			Description: `Force test execution even if collection errors occur`,
		}, {
			Name:        []string{"--debug"},
			Description: `Store internal tracing debug information in this log file. This file is opened with 'w' and truncated as a result, care advised. Defaults to 'pytestdebug.log'`,
			Args: []model.Arg{{
				Name:       "Debug File Name",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"--durations"},
			Description: `Show N slowest setup/test durations (N=0 for all)`,
			Args: []model.Arg{{
				Name:        "N",
				Description: `(N=0 for all)`,
			}},
		}, {
			Name:        []string{"--durations-min"},
			Description: `Minimal duration in seconds for inclusion in slowest list`,
			Args: []model.Arg{{
				Name: "N",
			}},
		}, {
			Name:        []string{"--deselect"},
			Description: `Deselect item (via node id prefix) during collection (multi-allowed)`,
			Args: []model.Arg{{
				Name: "nodeid_prefix",
			}},
		}, {
			Name:        []string{"--disable-warnings", "--disable-pytest-warnings"},
			Description: `Disable warnings summary`,
		}, {
			Name:        []string{"--doctest-continue-on-failure"},
			Description: `For a given doctest, continue to run after the first failure`,
		}, {
			Name:        []string{"--doctest-ignore-import-errors"},
			Description: `Ignore doctest ImportErrors`,
		}, {
			Name:        []string{"--doctest-modules"},
			Description: `Run doctests in all .py modules`,
		}, {
			Name:        []string{"--doctest-report"},
			Description: `Choose another output format for diffs on doctest failure`,
			Args: []model.Arg{{
				Name:        "Output format",
				Description: `None,cdiff,ndiff,udiff,only_first_failure`,
				Suggestions: []model.Suggestion{{Name: []string{`none`}}, {Name: []string{`cdiff`}}, {Name: []string{`ndiff`}}, {Name: []string{`udiff`}}, {Name: []string{`only_first_failure`}}},
			}},
		}, {
			Name:        []string{"--doctest-glob"},
			Description: `Doctests file matching pattern, default: test*.txt`,
			Args: []model.Arg{{
				Name: "Pattern",
			}},
		}, {
			Name:        []string{"--exitfirst", "-x"},
			Description: `Exit instantly on first error or failed test`,
		}, {
			Name:        []string{"--failed-first", "--ff"},
			Description: `Run all tests, but run the last failures first`,
		}, {
			Name:        []string{"--fixtures"},
			Description: `Shows builtin and custom fixtures. Note that this command omits fixtures with leading _ unless the -v option is added`,
		}, {
			Name:        []string{"--fixtures-per-test"},
			Description: `Show fixtures per test`,
		}, {
			Name:        []string{"--full-trace"},
			Description: `Don't cut any tracebacks (default is to cut)`,
		}, {
			Name:        []string{"--help", "-h"},
			Description: `This shows help on command line and config-line options`,
		}, {
			Name:        []string{"--ignore"},
			Description: `Ignore path during collection (multi-allowed)`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "Path",
			}},
		}, {
			Name:        []string{"--ignore-glob"},
			Description: `Ignore path pattern during collection (multi-allowed)`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "Path",
			}},
		}, {
			Name:        []string{"--import-mode"},
			Description: `Prepend/append to sys.path when importing test modules and conftest files, default is to prepend`,
			Args: []model.Arg{{
				Name:        "Mode",
				Suggestions: []model.Suggestion{{Name: []string{`prepend`}}, {Name: []string{`append`}}, {Name: []string{`importlib`}}},
			}},
		}, {
			Name:        []string{"--junit-xml"},
			Description: `Create junit-xml style report file at given path`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths, model.TemplateFolders},
				Name:      "Path",
			}},
		}, {
			Name:        []string{"--junit-prefix"},
			Description: `Prepend prefix to classnames in junit-xml output`,
			Args: []model.Arg{{
				Name:        "Str",
				Description: `String to prepend`,
			}},
		}, {
			Name:        []string{"-k"},
			Description: `Only run tests which match the given substring expression. An expression is a python evaluatable expression where all names are substring-matched against test names and their parent classes. Example: -k 'test_method or test_other' matches all test functions and classes whose name contains 'test_method' or 'test_other', while -k 'not test_method' matches those that don't contain 'test_method' in their names. -k 'not test_method and not test_other' will eliminate the matches. Additionally keywords are matched to classes and functions containing extra names in their 'extra_keyword_matches' set, as well as functions which have names assigned directly to them. The matching is case- insensitive`,
			Args: []model.Arg{{
				Name:        "Expression",
				Description: `Ex: 'test_method or test_other'`,
			}},
		}, {
			Name:        []string{"--keep-duplicates"},
			Description: `Keep duplicate tests`,
		}, {
			Name:        []string{"--showlocals", "-l"},
			Description: `Show locals in tracebacks (disabled by default)`,
		}, {
			Name:        []string{"--last-failed-no-failures", "--lfnf"},
			Description: `Which tests to run with no previously (known) failures`,
			Args: []model.Arg{{
				Name:        "Tests",
				Suggestions: []model.Suggestion{{Name: []string{`all`}}, {Name: []string{`none`}}},
			}},
		}, {
			Name:        []string{"--last-failed", "--lf"},
			Description: `Rerun only the tests that failed at the last run (or all if none failed)`,
		}, {
			Name:        []string{"--log-auto-indent"},
			Description: `Auto-indent multiline messages passed to the logging module. Accepts true|on, false|off or an integer`,
			Args: []model.Arg{{
				Name:        "Log Auto Indent Setting",
				Suggestions: []model.Suggestion{{Name: []string{`true`}}, {Name: []string{`false`}}},
			}},
		}, {
			Name:        []string{"--log-cli-level"},
			Description: `Cli logging level`,
			Args: []model.Arg{{
				Name:        "Log CLI Level",
				Suggestions: []model.Suggestion{{Name: []string{`CRITICAL`}}, {Name: []string{`ERROR`}}, {Name: []string{`WARNING`}}, {Name: []string{`INFO`}}, {Name: []string{`DEBUG`}}},
			}},
		}, {
			Name:        []string{"--log-cli-format"},
			Description: `Log format as used by the logging module`,
			Args: []model.Arg{{
				Name: "Log CLI Format",
			}},
		}, {
			Name:        []string{"--log-cli-date-format"},
			Description: `Log date format as used by the logging module`,
			Args: []model.Arg{{
				Name: "Log CLI Date Format",
			}},
		}, {
			Name:        []string{"--log-date-format"},
			Description: `Log date format as used by the logging module`,
			Args: []model.Arg{{
				Name: "Log Date Format",
			}},
		}, {
			Name:        []string{"--log-format"},
			Description: `Log format as used by the logging module`,
			Args: []model.Arg{{
				Name: "Log Format",
			}},
		}, {
			Name:        []string{"--log-file"},
			Description: `Path to a file where logging will be written to`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "Log File Path",
			}},
		}, {
			Name:        []string{"--log-file-level"},
			Description: `Log file logging level`,
			Args: []model.Arg{{
				Name:        "Log File Level",
				Suggestions: []model.Suggestion{{Name: []string{`CRITICAL`}}, {Name: []string{`ERROR`}}, {Name: []string{`WARNING`}}, {Name: []string{`INFO`}}, {Name: []string{`DEBUG`}}},
			}},
		}, {
			Name:        []string{"--log-file-date-format"},
			Description: `Log date format as used by the logging module`,
			Args: []model.Arg{{
				Name: "Log File Date Format",
			}},
		}, {
			Name:        []string{"--log-file-format"},
			Description: `Log format as used by the logging module`,
			Args: []model.Arg{{
				Name: "Log File Format",
			}},
		}, {
			Name:        []string{"--log-level"},
			Description: `Level of messages to catch/display. Not set by default, so it depends on the root/parent log handler's effective level, where it is "WARNING" by default`,
			Args: []model.Arg{{
				Name:        "Level",
				Suggestions: []model.Suggestion{{Name: []string{`CRITICAL`}}, {Name: []string{`ERROR`}}, {Name: []string{`WARNING`}}, {Name: []string{`INFO`}}, {Name: []string{`DEBUG`}}},
			}},
		}, {
			Name:        []string{"-m"},
			Description: `Only run tests matching given mark expression`,
			Args: []model.Arg{{
				Name: "Mark Expression",
			}},
		}, {
			Name:        []string{"--markers"},
			Description: `Show markers (builtin, plugin and per-project ones)`,
		}, {
			Name:        []string{"--maxfail"},
			Description: `Exit after first num failures or errors`,
			Args: []model.Arg{{
				Name: "num",
			}},
		}, {
			Name:        []string{"--new-first", "--nf"},
			Description: `Run tests from new files first, then the rest of the tests sorted by file mtime`,
		}, {
			Name:        []string{"--noconftest"},
			Description: `Don't load any conftest.py files`,
		}, {
			Name:        []string{"--no-header"},
			Description: `Disable header`,
		}, {
			Name:        []string{"--no-summary"},
			Description: `Disable summary`,
		}, {
			Name:        []string{"--override-ini", "-o"},
			Description: `Override ini option with "option=value" style"`,
			Args: []model.Arg{{
				Name:        "Override INI",
				Description: `Ex: "-o xfail_strict=True -o cache_dir=cache`,
			}},
		}, {
			Name:        []string{"-p"},
			Description: `Early-load given plugin module name or entry point (multi-allowed)`,
			Args: []model.Arg{{
				Name: "Plugin name",
			}},
		}, {
			Name:        []string{"--pastebin"},
			Description: `Send failed|all info to bpaste.net pastebin service`,
			Args: []model.Arg{{
				Name:        "mode",
				Suggestions: []model.Suggestion{{Name: []string{`failed`}}, {Name: []string{`all`}}},
			}},
		}, {
			Name:        []string{"--pdb"},
			Description: `Start the interactive Python debugger on errors or KeyboardInterrupt`,
		}, {
			Name:        []string{"--pdbcls"},
			Description: `Specify a custom interactive Python debugger for use with --pdb`,
			Args: []model.Arg{{
				Name:        "modulename:classname",
				Description: `Ex: --pdbcls=IPython.terminal.debugger:TerminalPdb`,
			}},
		}, {
			Name:        []string{"--pyargs"},
			Description: `Try to interpret all arguments as python packages`,
		}, {
			Name:        []string{"--quiet", "-q"},
			Description: `Decrease verbosity`,
		}, {
			Name:        []string{"-r"},
			Description: `Show extra test summary info as specified by chars: (f)ailed, (E)rror, (s)kipped, (x)failed, (X)passed, (p)assed, (P)assed with output, (a)ll except passed (p/P), or (A)ll. (w)arnings are enabled by default (see --disable-warnings), 'N' can be used to reset the list. (default: 'fE')`,
			Args: []model.Arg{{
				Name:        "chars",
				Suggestions: []model.Suggestion{{Name: []string{`a`}}, {Name: []string{`A`}}, {Name: []string{`E`}}, {Name: []string{`f`}}, {Name: []string{`N`}}, {Name: []string{`p`}}, {Name: []string{`P`}}, {Name: []string{`s`}}, {Name: []string{`w`}}, {Name: []string{`x`}}, {Name: []string{`X`}}},
			}},
		}, {
			Name:        []string{"--rootdir"},
			Description: `Define root directory for tests. Can be relative path: 'root_dir', './root_dir', 'root_dir/another_dir/'; absolute path: '/home/user/root_dir'; path with variables:'$HOME/root_dir'`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFolders},
				Name:      "Root Dir",
			}},
		}, {
			Name:        []string{"--runxfail"},
			Description: `Report the results of xfail tests as if they were not marked`,
		}, {
			Name:        []string{"-s"},
			Description: `Shortcut for --capture=no`,
		}, {
			Name:        []string{"--setup-only"},
			Description: `Only setup fixtures, do not execute tests`,
		}, {
			Name:        []string{"--setup-show"},
			Description: `Show setup of fixtures while executing tests`,
		}, {
			Name:        []string{"--setup-plan"},
			Description: `Show what fixtures and tests would be executed but don't execute anything`,
		}, {
			Name:        []string{"--show-capture"},
			Description: `Controls how captured stdout/stderr/log is shown on failed tests`,
			Args: []model.Arg{{
				Name:        "Capture method",
				Suggestions: []model.Suggestion{{Name: []string{`no`}}, {Name: []string{`stdout`}}, {Name: []string{`stderr`}}, {Name: []string{`log`}}, {Name: []string{`all`}}},
			}},
		}, {
			Name:        []string{"--stepwise", "--sw"},
			Description: `Exit on test failure and continue from last failing test next time`,
		}, {
			Name:        []string{"--stepwise-skip", "--sw-skip"},
			Description: `Ignore the first failing test but stop on the next failing test`,
		}, {
			Name:        []string{"--strict"},
			Description: `Alias to --strict-markers`,
		}, {
			Name:        []string{"--strict-config"},
			Description: `Any warnings encountered while parsing the "pytest" section of the configuration file raise errors`,
		}, {
			Name:        []string{"--strict-markers"},
			Description: `Markers not registered in the "markers" section of the configuration file raise errors`,
		}, {
			Name:        []string{"--tb"},
			Description: `Traceback print mode`,
			Args: []model.Arg{{
				Name:        "Traceback print mode",
				Suggestions: []model.Suggestion{{Name: []string{`auto`}}, {Name: []string{`long`}}, {Name: []string{`short`}}, {Name: []string{`line`}}, {Name: []string{`native`}}, {Name: []string{`no`}}},
			}},
		}, {
			Name:        []string{"--trace"},
			Description: `Immediately break when running each test`,
		}, {
			Name:        []string{"--trace-config"},
			Description: `Trace considerations of conftest.py files`,
		}, {
			Name:        []string{"--verbose", "-v"},
			Description: `Increase verbosity`,
		}, {
			Name:        []string{"--verbosity"},
			Description: `Set verbosity. Default is 0`,
			Args: []model.Arg{{
				Name: "Verbosity level",
			}},
		}, {
			Name:        []string{"--version", "-V"},
			Description: `Display pytest version and information about plugins. When given twice, also display information about plugins`,
		}, {
			Name:        []string{"--pythonwarnings", "-W"},
			Description: `Set which warnings to report, see -W option of python itself`,
			Args: []model.Arg{{
				Name: "Warnings to report",
			}},
		}},
	}
}