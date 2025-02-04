set -e
for tag in none dmesg libc.membrk libc.memgrind
do
	echo "-tags=$tag"
	GOOS=darwin GOARCH=amd64 go build -tags=$tag -v ./...
	GOOS=darwin GOARCH=amd64 go test -tags=$tag -c -o /dev/null
	GOOS=darwin GOARCH=arm64 go build -tags=$tag -v ./...
	GOOS=darwin GOARCH=arm64 go test -tags=$tag -c -o /dev/null
	GOOS=freebsd GOARCH=386 go build -tags=$tag -v ./...
	GOOS=freebsd GOARCH=386 go test -tags=$tag -c -o /dev/null
	GOOS=freebsd GOARCH=amd64 go build -tags=$tag -v ./...
	GOOS=freebsd GOARCH=amd64 go test -tags=$tag -c -o /dev/null
	GOOS=freebsd GOARCH=arm go build -tags=$tag -v ./...
	GOOS=freebsd GOARCH=arm go test -tags=$tag -c -o /dev/null
	GOOS=linux GOARCH=386 go build -tags=$tag -v ./...
	GOOS=linux GOARCH=386 go test -tags=$tag -c -o /dev/null
	GOOS=linux GOARCH=amd64 go build -tags=$tag -v ./...
	GOOS=linux GOARCH=amd64 go test -tags=$tag -c -o /dev/null
	GOOS=linux GOARCH=arm go build -tags=$tag -v ./...
	GOOS=linux GOARCH=arm go test -tags=$tag -c -o /dev/null
	GOOS=linux GOARCH=arm64 go build -tags=$tag -v ./...
	GOOS=linux GOARCH=arm64 go test -tags=$tag -c -o /dev/null
	GOOS=linux GOARCH=ppc64le go test -tags=$tag -c -o /dev/null
	GOOS=linux GOARCH=riscv64 go build -tags=$tag -v ./...
	GOOS=linux GOARCH=s390x go build -tags=$tag -v ./...
	GOOS=linux GOARCH=s390x go test -tags=$tag -c -o /dev/null
	GOOS=netbsd GOARCH=amd64 go build -tags=$tag -v ./...
	GOOS=netbsd GOARCH=amd64 go test -tags=$tag -c -o /dev/null
	GOOS=openbsd GOARCH=386 go build -tags=$tag -v ./...
	GOOS=openbsd GOARCH=386 go test -tags=$tag -c -o /dev/null
	GOOS=openbsd GOARCH=amd64 go build -tags=$tag -v ./...
	GOOS=openbsd GOARCH=amd64 go test -tags=$tag -c -o /dev/null
	GOOS=openbsd GOARCH=arm64 go build -tags=$tag -v ./...
	GOOS=openbsd GOARCH=arm64 go test -tags=$tag -c -o /dev/null
	GOOS=windows GOARCH=386 go build -tags=$tag -v ./...
	GOOS=windows GOARCH=386 go test -tags=$tag -c -o /dev/null
	GOOS=windows GOARCH=amd64 go build -tags=$tag -v ./...
	GOOS=windows GOARCH=amd64 go test -tags=$tag -c -o /dev/null
	GOOS=windows GOARCH=arm64 go build -tags=$tag -v ./...
	GOOS=windows GOARCH=arm64 go test -tags=$tag -c -o /dev/null
done
