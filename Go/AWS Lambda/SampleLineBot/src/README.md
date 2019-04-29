#### Refeference: 
[awslabs/aws-lambda-go-api-proxy](https://github.com/awslabs/aws-lambda-go-api-proxy)

[gin-gonic](https://github.com/gin-gonic/gin)

##### ＠Prerequisites
create aws account. see [here](https://aws.amazon.com/tw/lambda/features/)

create Line develop account. see [here](https://developers.line.biz/en/)

##### ＠Step 1
create API Gateway resources & methods, then click 'Deploy API'. see [here](https://aws.amazon.com/tw/api-gateway/)

##### ＠Step 2
Implement function to handle linebot request & web request, then zip it and upload it to aws lambda service.

GO packing instructions in mac OS/linux base:
<pre>
GOOS=linux go build -o main .
zip NameYouLike.zip main
</pre>

GO packing instructions in windows base(use git bash terminal):
<pre>
clone first: go.exe get -u github.com/aws/aws-lambda-go/cmd/build-lambda-zip
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
go build -o main main.go
build-lambda-zip.exe -o main.zip main
</pre>

##### ＠Step 3
set varaiable(channel secret & channel secret token) in Lambda

happy Coding With Linebot in Lambda & API Gateway.
