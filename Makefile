ifeq (,$(wildcard .env))
    $(error file .env not exsist.)
endif

include .env
export

dev-build: 
	docker build -t $(IMAGE_DEV) -f Dockerfile.dev .

dev-run: dev-build
	docker run -p 8080:8080 -e WEATHER_KEY=$(WEATHER_KEY) $(IMAGE_DEV)

prod-build: 
	docker build -t $(IMAGE_PROD) -f Dockerfile.prod .

prod-run: prod-build
	docker run -p 8080:8080 -e WEATHER_KEY=$(WEATHER_KEY) $(IMAGE_PROD)

test: dev-build
	docker run $(IMAGE_DEV) go test -race ./...

cover: dev-build
	docker run $(IMAGE_DEV) go test -cover ./...