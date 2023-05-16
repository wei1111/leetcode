fmt:
	find . -name "*.go" | egrep -v "(log|output|kitex_gen)" | xargs goimports -w
	find . -name "*.go" | egrep -v "(log|output|kitex_gen)" | xargs gofmt -w