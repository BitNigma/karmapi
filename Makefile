.PHONE: build 

build : 
		go build -o taro ./cmd

.DEFAULT_GOAL := build
