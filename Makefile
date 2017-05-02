.PHONY: dist build
install:
	@npm install

dev: install
	@npm run dev

build:
	@npm run build

image:
	@docker build -t 10.211.55.40/daocloud/dce-app-update-plugin:0.0.1 .
