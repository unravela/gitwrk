package gitwrk

import "testing"

func TestParseSimpleCommitMessage(t *testing.T) {
	d := parseSpent("Spent 1h40m")
	if len(d) != 1 {
		t.Error("The parse function must return only one duration")
	}

	if d[0].Minutes() != 100 {
		t.Error("The duration need to be 1h40m (100m)")
	}
}

func TestParseMultipleDurationsCommitMessage(t *testing.T) {
	d := parseSpent("Spent 1h40m, 30m, 3h20m")
	if len(d) != 3 {
		t.Error("The parse function must return only one duration")
	}

	if d[1].Minutes() != 30 {
		t.Error("The second duration need to be 30m")
	}

	if d[2].Minutes() != ((3 * 60) + 20) {
		t.Error("The second duration need to be 3h20m")
	}
}

func TestParseAsGitTrailerLine(t *testing.T) {
	// the git trailer line is convention described here: https://git-scm.com/docs/git-interpret-trailers
	d := parseSpent("spent: 1h")

	if len(d) != 1 || d[0].Hours() != 1 {
		t.Error("The tailer line cannot be parsed correctly")
	}
}
