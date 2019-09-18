package howto

func HowTo() string {
	return `
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
}
