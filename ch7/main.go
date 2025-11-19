package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	league := League{
		Teams: []Team{
			{name: "Hornets", playerNames: []string{"Sean", "Blizard", "John"}},
			{name: "Lakers", playerNames: []string{"LeBron", "Davis", "Reaves"}},
			{name: "Warriors", playerNames: []string{"Curry", "Thompson", "Green"}},
			{name: "Celtics", playerNames: []string{"Tatum", "Brown", "Holiday"}},
		},
		Wins: make(map[string]int),
	}

	league.MatchResult("Hornets", 105, "Lakers", 98)
	league.MatchResult("Warriors", 112, "Celtics", 110)
	league.MatchResult("Lakers", 95, "Warriors", 102)
	league.MatchResult("Hornets", 108, "Celtics", 103)
	league.MatchResult("Warriors", 115, "Hornets", 109)
	league.MatchResult("Celtics", 100, "Lakers", 97)

	rankings := league.Ranking()
	for i, team := range rankings {
		println(i+1, team.name, league.Wins[team.name], "wins")
	}

	RankPrinter(&league, os.Stdout)
}

// You have been asked to manage a basketball league and are going to write a program to help you.
// Define two types. The first one, called Team, has a field for the name of the team and a field for the player names.
// The second type is called League and has a field called Teams for the teams in the league and a field called Wins that maps a team’s name to its number of wins.

type Team struct {
	name        string
	playerNames []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

// Add two methods to League. The first method is called MatchResult.
// It takes four parameters: the name of the first team, its score in the game, the name of the second team, and its score in the game.
// This method should update the Wins field in League.
// Add a second method to League called Ranking that returns a slice of the team names in order of wins.
// Build your data structures and call these methods from the main function in your program using some sample data.

func (l *League) MatchResult(teamOne string, teamOneScore int, teamTwo string, teamTwoScore int) {
	var winner string
	if teamOneScore > teamTwoScore {
		winner = teamOne
	} else {
		winner = teamTwo
	}
	l.Wins[winner] += 1
}

func (l League) Ranking() []Team {
	var rankedTeams []Team
	rankedTeams = append(rankedTeams, l.Teams...)
	sort.Slice(rankedTeams, func(i, j int) bool {
		return l.Wins[rankedTeams[i].name] > l.Wins[rankedTeams[j].name]
	})
	return rankedTeams
}

func (l League) RankingStrings() []string {
	rankedTeams := l.Ranking()
	var teamNames []string
	for _, team := range rankedTeams {
		teamNames = append(teamNames, team.name)
	}
	return teamNames
}

// Define an interface called Ranker that has a single method called Ranking that returns a slice of strings.
// Write a function called RankPrinter with two parameters, the first of type Ranker and the second of type io.Writer[…]”
type Ranker interface {
	RankingStrings() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	rankings := r.RankingStrings()
	for i, team := range rankings {
		io.WriteString(w, fmt.Sprintf("%d. %s\n", i+1, team))
	}
}
