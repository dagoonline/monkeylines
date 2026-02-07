package main

import (
	"fmt"
	"math/rand"
)

type exchange struct {
	Insult   string
	Comeback string
}

// A theme groups insult/comeback templates that are topically related.
// Each template uses %s placeholders filled from the theme's word lists,
// producing a large number of unique but coherent pairs.
type theme struct {
	insults   []string
	comebacks []string
	nouns     []string
}

var themes = []theme{
	{
		insults: []string{
			"You code like a %s!",
			"I've seen a %s write better code!",
			"A %s could push cleaner commits!",
			"My grandmother's %s has better coding skills!",
		},
		comebacks: []string{
			"How appropriate. You fight like a %s.",
			"And I thought you smelled like a %s.",
			"I'm glad you attended your %s family reunion.",
			"At least a %s can learn. You can't.",
		},
		nouns: []string{
			"dairy farmer", "cow", "rubber duck", "sea sponge",
			"three-headed monkey", "bilge rat", "barnacle",
			"parrot", "jellyfish", "cabin boy",
		},
	},
	{
		insults: []string{
			"Your code is as ugly as a %s!",
			"Your functions smell worse than a %s!",
			"Your variables are messier than a %s!",
			"I've seen cleaner logic in a %s!",
		},
		comebacks: []string{
			"Yet yours couldn't even pass a %s.",
			"Funny, your %s said the same about your code.",
			"At least my code doesn't crash like your %s.",
			"Still prettier than your %s.",
		},
		nouns: []string{
			"merge conflict", "core dump", "stack trace",
			"spaghetti monster", "legacy codebase", "regex pattern",
			"minified bundle", "memory dump", "corrupted database",
		},
	},
	{
		insults: []string{
			"You debug slower than a %s!",
			"Your debugging skills are worse than a %s!",
			"Even a %s finds bugs faster than you!",
			"A %s could step through code better!",
		},
		comebacks: []string{
			"At least I know when to %s.",
			"I'd rather %s than watch you try.",
			"First you'd better learn to %s.",
			"Too bad you can't even %s.",
		},
		nouns: []string{
			"break out of an infinite loop", "exit vim",
			"read a stack trace", "use a debugger",
			"check the logs", "write a unit test",
			"set a breakpoint", "read the docs",
		},
	},
	{
		insults: []string{
			"Your deploys are as reliable as a %s!",
			"Your CI pipeline is held together with %s!",
			"Your production server runs on %s!",
			"Your uptime is shorter than a %s!",
		},
		comebacks: []string{
			"Your last deploy broke more things than a %s.",
			"At least my %s doesn't page me at 3am.",
			"Better than your %s that never even builds.",
			"Says the one whose %s is always on fire.",
		},
		nouns: []string{
			"chocolate teapot", "mass of duct tape", "house of cards",
			"rubber band and a prayer", "coffee-stained napkin sketch",
			"server held together with hope", "cron job from 2003",
		},
	},
	{
		insults: []string{
			"Your git history reads like a %s!",
			"Your commit messages are worse than %s!",
			"Your branches are more tangled than %s!",
			"Your pull requests look like %s!",
		},
		comebacks: []string{
			"Your last commit message was just '%s'.",
			"At least I don't %s on every push.",
			"Better than your strategy of '%s'.",
			"Says the one who thinks '%s' is version control.",
		},
		nouns: []string{
			"fix stuff", "force push to main",
			"commit directly to production", "WIP WIP WIP final FINAL",
			"rebase and pray", "copy the whole folder",
			"undo undo undo", "asdfasdf",
		},
	},
	{
		insults: []string{
			"You handle errors like a %s handles a sword!",
			"Your exception handling is as graceful as a %s!",
			"You catch bugs like a %s catches cannonballs!",
			"Your error messages are as helpful as a %s!",
		},
		comebacks: []string{
			"Your approach is even worse: %s.",
			"At least I don't %s like you do.",
			"Better than your technique of '%s'.",
			"Says the one whose strategy is '%s'.",
		},
		nouns: []string{
			"catch and ignore everything",
			"blame the user",
			"restart and hope for the best",
			"comment out the failing test",
			"wrap it all in a try-catch and move on",
			"print 'this should never happen'",
			"ship it anyway",
		},
	},
	{
		insults: []string{
			"Your architecture looks like it was designed by a %s!",
			"I've seen better design patterns in a %s!",
			"Your system design reminds me of a %s!",
			"A %s has more structure than your codebase!",
		},
		comebacks: []string{
			"Your architecture collapses faster than a %s.",
			"At least my code doesn't look like a %s.",
			"Better than your %s of a system.",
			"Says the one building on a %s.",
		},
		nouns: []string{
			"shipwreck", "sandcastle at high tide",
			"house built on quicksand", "tower of Babel",
			"Jenga tower in an earthquake", "soggy cardboard box",
			"pirate ship full of holes", "haunted spaghetti factory",
		},
	},
	{
		insults: []string{
			"Your tests cover less than a %s!",
			"Your test suite is as thorough as a %s!",
			"I've seen more coverage from a %s!",
			"Your QA process is just %s!",
		},
		comebacks: []string{
			"You don't even know what a %s is.",
			"At least I have tests. You just %s.",
			"Better than your approach of '%s'.",
			"Says the one who thinks %s is optional.",
		},
		nouns: []string{
			"blindfolded monkey poking a keyboard",
			"crossing your fingers and deploying",
			"testing in production",
			"asking the intern if it works",
			"refreshing the page once",
			"code review", "integration testing",
		},
	},
}

func randomChoice[T any](slice []T) T {
	return slice[rand.Intn(len(slice))]
}

func generateExchange() exchange {
	t := randomChoice(themes)
	noun1 := randomChoice(t.nouns)
	noun2 := randomChoice(t.nouns)
	return exchange{
		Insult:   fmt.Sprintf(randomChoice(t.insults), noun1),
		Comeback: fmt.Sprintf(randomChoice(t.comebacks), noun2),
	}
}

func generateMessage() string {
	ex := generateExchange()
	if rand.Intn(2) == 0 {
		return ex.Insult
	}
	return ex.Comeback
}
