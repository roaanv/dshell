image:
	docker build --platform linux/arm64,linux/amd64 -t rvos-docker-shell:latest .
	docker tag rvos-docker-shell:latest docker.spacehuis.net/rvos-docker-shell:latest
