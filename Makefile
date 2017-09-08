

ARCH := 386 amd64
OS := linux darwin windows



status:
	dep status

install:
	dep ensure

update:
	dep ensure -update
	
build:
	go generate ./...
	go test ./...

package: build
	gox -os="$(OS)" -arch="$(ARCH)" -output "dist/{{.OS}}_{{.Arch}}/{{.Dir}}"