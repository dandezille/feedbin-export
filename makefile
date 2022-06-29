APP_NAME=feedbin-export

build:
	podman build -t $(APP_NAME) .

build-nc:
	podman build --no-cache -t $(APP_NAME) .

run:
	podman run -it --rm $(APP_NAME)

tag:
	podman tag $(APP_NAME) dandezille/$(APP_NAME)

publish: build-nc tag
	podman push dandezille/$(APP_NAME)
