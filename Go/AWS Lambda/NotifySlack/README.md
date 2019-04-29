#### Refeference: 
[can-slack-incoming-webhooks-post-message-to-all-private-groups](https://stackoverflow.com/questions/31589744/can-slack-incoming-webhooks-post-message-to-all-private-groups)

[aws-lambda-go#for-developers-on-windows](https://github.com/aws/aws-lambda-go#for-developers-on-windows)

##### ＠Step 1
create slack app. see [here](https://api.slack.com/)

create aws account see [here](https://aws.amazon.com/tw/lambda/features/)

##### ＠Step 2
Implement function to send request, then zip it and upload it to aws lambda service.(maybe here could be more automative when func update?)

GO packing instructions in linux base:
<pre>
GOOS=linux go build -o main .
zip deployment.zip main
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
Test event to check lambda func work or not from aws console(web).

**Note**: Slack API there are **two** types of "incoming webhook" now.
<pre>
<b>default type</b>: create as part of a Slack app.
<b>legacy type</b>: needs to be installed as app from the Slack App Directory.
Only the legacy type supports multiple destinations.
</pre>
