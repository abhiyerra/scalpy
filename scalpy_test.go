package scalpy

import (
	"testing"
)

func TestScalpGoodUrl(t *testing.T) {
	url := "https://github.com/abhiyerra/feedbackjs/issues/2"
	scalp := ScalpUrl(url)

	if scalp == nil {
		t.Fatal("Couldn't parse the url correctly")
	}

	if scalp.HostingService != Github {
		t.Fatal("Invalid hosting service.")
	}

	if scalp.Project != "abhiyerra" {
		t.Fatal("Invalid repo.")
	}

	if scalp.Repo != "feedbackjs" {
		t.Fatal("Invalid repo.")
	}

	if scalp.IssueId != "2" {
		t.Fatal("Invalid repo.")
	}
}

func TestScalpBadUrl(t *testing.T) {
	url := "https://abhiyerra.com/asdf"
	scalp := ScalpUrl(url)

	if scalp != nil {
		t.Fatal("Somehow parsed the url correctly")
	}
}

func TestScalpBadPath(t *testing.T) {
	url := "https://github.com/asdf"
	scalp := ScalpUrl(url)

	if scalp != nil {
		t.Fatal("Somehow parsed the url correctly")
	}
}

func TestScalpGithubIssue(t *testing.T) {
	url := "https://github.com/abhiyerra/feedbackjs/issues/2"
	scalp := ScalpUrl(url)

	if issue := scalp.GithubIssue(); issue == nil {
		t.Fatal("Couldn't get the correct github issue")
	}
}
