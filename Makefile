image:
	docker build --platform linux/arm64,linux/amd64 -t roaanv/dshell:latest .
	docker tag roaanv/dshell:latest docker.spacehuis.net/dshell:latest
