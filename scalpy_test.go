package scalpy

import (
	"testing"
)

func TestScalp_good(t *testing.T) {
	url := "https://github.com/abhiyerra/feedbackjs/issues/2"
	scalp := ScalpUrl(url)

	if scalp == nil {
		t.Fatal("Couldn't parse the url correctly")
	}

	if scalp.HostingService != Github {
		t.Fatal("Invalid hosting service.")
	}

	if scalp.Repo != "abhiyerra/feedbackjs" {
		t.Fatal("Invalid repo.")
	}

	if scalp.IssueId != "2" {
		t.Fatal("Invalid repo.")
	}
}

func TestScalp_badHost(t *testing.T) {
	url := "https://abhiyerra.com/asdf"
	scalp := ScalpUrl(url)

	if scalp != nil {
		t.Fatal("Somehow parsed the url correctly")
	}
}

func TestScalp_badPath(t *testing.T) {
	url := "https://github.com/asdf"
	scalp := ScalpUrl(url)

	if scalp != nil {
		t.Fatal("Somehow parsed the url correctly")
	}
}
