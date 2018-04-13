.PHONY: default init build install release dep get_deps clean build_amd64 build_386 upx

# https://golang.org/doc/install/source#environment
GOOS := $(shell go env | awk -F= '$$1=="GOOS" {print $$2}' | awk -F '"' '{print $$2}') # 此处 awk 需使用两个 $
GOARCH := $(shell go env | awk -F= '$$1=="GOARCH" {print $$2}' | awk -F '"' '{print $$2}')
OSS = darwin dragonfly freebsd linux netbsd openbsd plan9 solaris windows
PKG =
# ifeq ($(strip $(GOOS)), windows)
# 	GOARCH := $(strip $(GOARCH)).exe
# endif

default:
	@echo "JWT info: please choose a target for 'make'"
	@echo "available target: init build install release dep get_deps clean build_amd64 build_386 upx"

init: get-deps dep clean
	go install -ldflags "-s -w"

build:
	@ go build -ldflags "-s -w" -o dist/jwt_$(strip $(GOOS))_$(strip $(if \
    $(findstring windows,$(GOOS)),\
    $(strip $(GOARCH)).exe,\
    $(strip $(GOARCH))\
	))

install:
	go install -ldflags "-s -w"

release: amd64 386

dep:
	dep ensure

get_deps:
	go get -u github.com/golang/dep/cmd/dep

clean:
	go clean -i
	rm -rf dist/jwt* jbls

build_amd64:
	@ $(foreach OS,\
	$(OSS),\
	$(shell CGO_ENABLED=0 GOOS=$(OS) GOARCH=amd64 go build -ldflags "-s -w" -o dist/jwt_$(OS)_amd64$(if $(findstring windows,$(OS)),.exe)))
	@ echo done

build_386:
	@ $(foreach OS,\
	$(OSS),\
	$(shell CGO_ENABLED=0 GOOS=$(OS) GOARCH=386 go build -ldflags "-s -w" -o dist/jwt_$(OS)_386$(if $(findstring windows,$(OS)),.exe)))
	@ echo done

# 压缩。需要安装 https://github.com/upx/upx
upx:
	upx $(if $(PKG),$(PKG),dist/jwt*)