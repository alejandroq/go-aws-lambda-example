S3_BUCKET := "hello-world-go-xyz"

all: build

build: 
	GOOS=linux go build -o handler src/main.go
	zip dist/handler.zip handler
	rm handler

package:
	# aws cloudformation deploy --template-fil packaged.yml --stack-name go-lambda-test-1 --capabilities CAPABILITY_IAM
	aws cloudformation package --template-file template.yml --s3-bucket ${S3_BUCKET} --output-template-file packaged.yml