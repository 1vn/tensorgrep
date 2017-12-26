run:
	rerun -watch ./ -ignore vendor bin tmp -run sh -c 'go build -i -o ./bin/tensorgrep ./cmd/main.go && ./bin/tensorgrep -d=tmp/training'