gen:
	mkdir -p force-common-js
	protoc --go_out=./ --js_out=./force-common-js ./common.proto
