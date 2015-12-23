package main

import (
	"testing"

	"github.com/cloudnativego/gogo-engine"
)

func TestAddMatchShowsUpInRepository(t *testing.T) {
	match := gogo.NewMatch(19, "bob", "alfred")

	repo := newInMemoryRepository()
	err := repo.addMatch(match)
	if err != nil {
		t.Error("Got an error adding a match to repository, should not have.")
	}

	matches := repo.getMatches()
	if len(matches) != 1 {
		t.Errorf("Expected to have 1 match in the repository, got %d", len(matches))
	}

	if matches[0].PlayerBlack != "bob" {
		t.Errorf("Player 1's name should have been bob, got %s", matches[0].PlayerBlack)
	}
	if matches[0].PlayerWhite != "alfred" {
		t.Errorf("Player 2's name should have been alfred, got %s", matches[0].PlayerWhite)
	}
}

func TestGetMatchRetrievesProperMatch(t *testing.T) {
	match := gogo.NewMatch(19, "bob", "alfred")

	repo := newInMemoryRepository()
	err := repo.addMatch(match)
	if err != nil {
		t.Error("Got an error adding a match to repository, should not have.")
	}

	target, err := repo.getMatch(match.ID)
	if err != nil {
		t.Errorf("Got an error when retrieving match from repo instead of success. Err: %s", err.Error())
	}

	if target.GridSize != 19 {
		t.Errorf("Got the wrong gridsize. Expected 19, got %d", target.GridSize)
	}
}

func TestNewRepositoryIsEmpty(t *testing.T) {
	repo := newInMemoryRepository()

	matches := repo.getMatches()
	if len(matches) != 0 {
		t.Errorf("Expected to have 0 matches in newly created repository, got %d", len(matches))
	}
}