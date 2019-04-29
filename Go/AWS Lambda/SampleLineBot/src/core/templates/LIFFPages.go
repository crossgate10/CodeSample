package tmpl

var SmapleBotLifftmpl = `
<head>
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>SampleBot</title>
	<style type="text/css">
		html{
            font-size: 11px;
        }
	</style>
</head>
<body>
<img id="userPic" />
<script src="https://d.line-scdn.net/liff/1.0/sdk.js"></script>
<script>
	window.onload = function(e) {
    	liff.init(function (data) {
			liff.getProfile()
            .then(profile => {
				document.getElementById("userPic").src = profile.pictureUrl;
            })
            .catch((err) => {
              alert('error', err);
            });
    	});
	};
</script>
</body>
`