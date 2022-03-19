SERVICE := am-stats-dataprep-api
NAMESPACE := aftermath-render
REGISTRY := ghcr.io/byvko-dev
# 
VERSION = $(shell git rev-parse --short HEAD)
TAG := ${REGISTRY}/${SERVICE}

echo:
	@echo ${TAG}

pull:
	git pull

build:
	docker build -t ${TAG}:${VERSION} -t ${TAG}:latest .
	docker image prune -f

push:
	docker push ${TAG}:latest

restart:
	kubectl rollout restart deployment/${SERVICE} -n ${NAMESPACE}
	