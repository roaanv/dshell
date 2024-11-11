image:
	docker build --platform linux/arm64,linux/amd64 -t roaanv/dshell:latest .
	docker tag roaanv/dshell:latest roaanv/dshell:0.0.2
