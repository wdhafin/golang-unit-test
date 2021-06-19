test:
	go test -v -cover -coverprofile=cover.out ./...

push:
	go fmt ./... && git push origin HEAD

push-force:
	go fmt ./... && git push -f origin HEAD