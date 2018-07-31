.PHONY: all

all: build deploy

build:
	GOOS=linux go build -o main main.go
	zip deployment.zip main

create: deployment.zip
	aws lambda create-function --region=us-east-1 --function-name=DiscoverMovies --zip-file=fileb://./deployment.zip --runtime=go1.x  --role=arn:aws:iam::045871928384:role/discovermovies  --handler=main
	touch .created .deployed

deploy: deployment.zip .created
	aws lambda update-function-code --region=us-east-1 --function-name=DiscoverMovies --zip-file=fileb://./deployment.zip
	touch .deployed

clean:
	rm -f *.zip main .deployed .created
