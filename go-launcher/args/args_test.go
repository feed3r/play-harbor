package args

import "testing"

func TestParseArgs_Valid(t *testing.T) {
	url, exe, err := ParseArgs([]string{"url", "game.exe"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if url != "url" || exe != "game.exe" {
		t.Errorf("unexpected values: %s %s", url, exe)
	}
}

func TestParseArgs_Invalid(t *testing.T) {
	_, _, err := ParseArgs([]string{"onlyone"})
	if err == nil {
		t.Error("expected error for missing args, got nil")
	}
}
