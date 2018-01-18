all: build

build: 
	GOOS=linux go build -o main.go
	zip handler.zip ./main
	# aws lambda create-function         \
	# 	--region us-east-1             \ 
	# 	--function-name GoHelloWorld  \
	# 	--memory 128				   \
	# 	--runtime go1.x                \
	# 	--zip-file $$(pwd)/handler.zip \
	# 	--handler ./main