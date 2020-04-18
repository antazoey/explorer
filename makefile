build:
	mkdir -p functions
	cd api && \
	go get ./... && \
	GO111MODULE=on go build -o "../functions/chain-explorer-api"
