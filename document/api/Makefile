.PHONY: compile
compile:
	redocly bundle ./src/index.yaml -o ./api.yaml --ext yaml

.PHONY: lint
lint:
	redocly lint ./src/index.yaml

.PHONY: lint-compile
lint-compile:
	redocly lint ./api.yaml
	
.PHONY: setup
setup:
	npm install @redocly/cli -g
