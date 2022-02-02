APP_NAME=feedbin-export

build:
	docker build -t $(APP_NAME) .

build-nc:
	docker build --no-cache -t $(APP_NAME) .

run:
	docker run -it --rm $(APP_NAME)

tag:
	docker tag $(APP_NAME) dandezille/$(APP_NAME)

publish: build-nc tag
	docker push dandezille/$(APP_NAME)
