package howto;

import "testing"

func TestHowTo(t *testing.T) {
	want := `
	cd ~/git/dbraley
	mkdir project-name
	cd project-name
	touch how_to.go
	code how_to.go
	touch how_to_test.go
	code how_to_test.go
	go test
	go mod init github.com/dbraley/project-name
	`
    if got := HowTo(); got != want {
        t.Errorf("HowTo() = %q, want %q", got, want)
    }
}