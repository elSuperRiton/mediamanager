CURR_DIR = $(PWD)

all: plugins

PHONY: plugins
plugins:
	cd $(CURR_DIR) && cd ./testplugin && go build -buildmode=plugin -o ../testplugin.so && \
	cd $(CURR_DIR) && cd ./testwrongplugin && go build -buildmode=plugin -o ../testwrongplugin.so