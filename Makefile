IMAGE_TAG_BASE ?= registry.cn-hangzhou.aliyuncs.com/cnosdb/cnosdb-init-config
DIST ?= dist
setup-buildx:
	export cnosdb_build_count=`docker buildx ls |grep -v NAME/NODE|grep -c cnosdb-container`; \
	if [ $$cnosdb_build_count -eq 0 ]; then \
		docker buildx create --driver=docker-container --name=cnosdb-container; \
	fi 
init-config-tidy:
	cd initconfig && go mod tidy
init-config-build: init-config-tidy
	cd initconfig && GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o ../$(DIST)/initconfig_arm64 main.go
	cd initconfig && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ../$(DIST)/initconfig_amd64 main.go
init-config-image: setup-buildx
	cp initconfig/Dockerfile $(DIST)/Dockerfile
	cd $(DIST)&&docker buildx build -f ./Dockerfile --builder cnosdb-container --platform linux/amd64,linux/arm64 -t $(IMAGE_TAG_BASE):latest . --push

init-config: init-config-build init-config-image

helm-package:
	cd $(DIST)&&helm package ../charts/cnosdb

build: init-config helm-package

