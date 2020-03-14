# Get an Deezer Access token

1. Create a Deezer app and get amn App ID

`https://connect.deezer.com/oauth/auth.php?app_id=<APP_ID>&redirect_uri=<URL>&perms=offline_access,manage_library,email`

2. Get the return code and pass ti to the access_token endpoint

`https://connect.deezer.com/oauth/access_token.php?app_id=<APP_ID>&secret=<APP_SECRET>&code=<CODE>`

3. Get the access token from the response

`access_token=<ACCESS_TOEN>&expires=0`
