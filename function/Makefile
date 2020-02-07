.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./pingpong/pingpong
	
build:
	GOOS=linux GOARCH=amd64 go build -o pingpong/pingpong ./pingpong


sam:
	sam deploy --stack-name apigw-pingpong --s3-bucket dihmuzikien-artifacts --capabilities CAPABILITY_IAM