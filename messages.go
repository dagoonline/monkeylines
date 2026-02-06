package main

import (
	"fmt"
	"math/rand"
)

var insultTemplates = []string{
	"You code like a %s!",
	"Your commits smell like %s!",
	"You fight bugs like a %s!",
	"I've seen %s write better code!",
	"Your code is as ugly as a %s!",
	"You debug slower than a %s!",
	"My %s could merge better than you!",
	"Your functions are as useless as a %s!",
	"You refactor like a %s!",
	"Soon your code will be as dead as a %s!",
	"Your variables smell like a %s!",
	"I've taught %s to code better!",
	"Every bug you fix spawns %s!",
	"Your pull requests look like %s!",
	"You handle exceptions like a %s!",
}

// comebackArgType declares what word forms a comeback template expects,
// so adding new templates only requires specifying the format — no fragile
// string matching needed.
type comebackArgType int

const (
	argNoun       comebackArgType = iota // single noun
	argVerbNoun                          // verb, noun
	argGerundNoun                        // gerund, noun
	argPastNoun                          // past verb, noun
	argNounVerb                          // noun, verb
	argNounGerund                        // noun, gerund
)

type comebackTemplate struct {
	format string
	args   comebackArgType
}

var comebackTemplates = []comebackTemplate{
	{"How appropriate. You %s like a %s.", argVerbNoun},
	{"First you'd better stop %s like a %s.", argGerundNoun},
	{"I'm glad you attended your %s reunion.", argNoun},
	{"And I thought you smelled like a %s.", argNoun},
	{"At least I know how to %s a %s.", argVerbNoun},
	{"You make me want to %s my %s.", argVerbNoun},
	{"I'd %s, but I don't want to dirty my %s.", argVerbNoun},
	{"Then you better stop %s your %s.", argGerundNoun},
	{"Even so, my %s can still %s.", argNounVerb},
	{"Too bad nobody will %s your %s.", argVerbNoun},
	{"I've %s worse %s than you.", argPastNoun},
	{"Yet you still can't %s a simple %s.", argVerbNoun},
	{"Funny, your %s said the same about your %s.", argNounGerund},
	{"That explains the %s in your repository.", argNoun},
	{"I'll %s that into your %s.", argVerbNoun},
}

var insultNouns = []string{
	"dairy farmer", "rubber duck", "segfault", "null pointer",
	"deprecated function", "memory leak", "infinite loop", "stack overflow",
	"merge conflict", "legacy codebase", "untested module", "spaghetti monster",
	"floating point", "race condition", "dead code", "code monkey",
	"keyboard warrior", "copy-paster", "tab user", "vim user",
}

var comebackNouns = []string{
	"cow", "compiler", "debugger", "garbage collector",
	"exception handler", "code reviewer", "unit test", "documentation",
	"git history", "production server", "staging environment", "bug tracker",
	"coffee machine", "rubber duck", "stack trace", "error log",
	"keyboard", "monitor", "semicolon", "curly brace",
}

var verbs = []string{
	"debug", "compile", "refactor", "deploy", "merge", "commit",
	"push", "pull", "branch", "rebase", "squash", "cherry-pick",
	"rollback", "hotfix", "optimize", "minify", "lint", "test",
}

var gerunds = []string{
	"debugging", "compiling", "refactoring", "deploying", "merging", "committing",
	"pushing", "pulling", "branching", "rebasing", "squashing", "cherry-picking",
	"rolling back", "hotfixing", "optimizing", "minifying", "linting", "testing",
}

var pastVerbs = []string{
	"debugged", "compiled", "refactored", "deployed", "merged", "committed",
	"pushed", "pulled", "branched", "rebased", "squashed", "cherry-picked",
	"rolled back", "hotfixed", "optimized", "minified", "linted", "tested",
}

func randomChoice[T any](slice []T) T {
	return slice[rand.Intn(len(slice))]
}

func generateMessage() string {
	if rand.Intn(2) == 0 {
		tpl := randomChoice(insultTemplates)
		return fmt.Sprintf(tpl, randomChoice(insultNouns))
	}

	cb := randomChoice(comebackTemplates)
	noun := randomChoice(comebackNouns)

	switch cb.args {
	case argNoun:
		return fmt.Sprintf(cb.format, noun)
	case argVerbNoun:
		return fmt.Sprintf(cb.format, randomChoice(verbs), noun)
	case argGerundNoun:
		return fmt.Sprintf(cb.format, randomChoice(gerunds), noun)
	case argPastNoun:
		return fmt.Sprintf(cb.format, randomChoice(pastVerbs), noun)
	case argNounVerb:
		return fmt.Sprintf(cb.format, noun, randomChoice(verbs))
	case argNounGerund:
		return fmt.Sprintf(cb.format, noun, randomChoice(gerunds))
	default:
		return fmt.Sprintf(cb.format, randomChoice(verbs), noun)
	}
}
