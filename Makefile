IMAGE_TAG_BASE ?= cnosdb/init-config
DIST ?= dist
init-config-tidy:
	cd initconfig && go mod tidy
init-config-build: init-config-tidy
	cd initconfig && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ../$(DIST)/initconfig main.go
init-config-image:
	cp initconfig/Dockerfile $(DIST)/Dockerfile
	cd $(DIST)&&docker build -t $(IMAGE_TAG_BASE):latest .
init-config: init-config-build init-config-image

helm-package:
	cd $(DIST)&&helm package ../charts/cnosdb

build: init-config helm-package

