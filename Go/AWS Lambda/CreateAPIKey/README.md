#### Refeference: 
[aws-sdk-go github](https://github.com/aws/aws-sdk-go)

##### ＠Note

use credentials on local enviroment. In contrast, use IAM on lambda.

Policy:
<pre>
AWSAppSyncPushToCloudWatchLogs
AWSAppSyncAdministrator
</pre>

Role:
<pre>
Allows Lambda functions to call AWS services on your behalf.
</pre>
