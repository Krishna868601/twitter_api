# twitter_api


## Introduction

This project demonstrates how to post and delete on twitter using the Go programming language. The program performs two key tasks:

1. **Posting a new tweet** using the Twitter API.
2. **Deleting an existing tweet** by providing its Tweet ID.

This project provides hands-on experience with the following:
- Using OAuth 1.0a for authentication with the Twitter API.
- Making HTTP POST and DELETE requests to the Twitter API.
- Handling API responses and errors.

## Setup Instructions

### 1. Set Up a Twitter Developer Account

To interact with the Twitter API, you need a Twitter Developer account. If you don't have one yet, follow these steps:

1. Visit the [Twitter Developer Platform](https://developer.twitter.com/).
2. Apply for a developer account. This process usually involves filling out some details about how you plan to use the API.
3. Once your account is approved, log in to the developer platform.

### 2. Create a New Project and App

1. Navigate to **Projects & Apps** → **Overview**.
2. Create a new project and name it.
3. Create an app within the project.

### 3. Generate API Keys and Tokens

1. After creating your app, navigate to **Keys and Tokens**.
2. Generate the following credentials:
   - **API Key**
   - **API Secret Key**
   - **Access Token**
   - **Access Token Secret**

3. Save these credentials. You'll need them to authenticate requests.

### 4. Set Environment Variables

For security reasons, it's best not to hardcode your API credentials. Instead, store them in environment variables. In your terminal, run the following commands:

```powershell
$env:TWITTER_API_KEY='your_api_key'
$env:TWITTER_API_SECRET_KEY='your_api_secret_key'
$env:TWITTER_ACCESS_TOKEN='your_access_token'
$env:TWITTER_ACCESS_TOKEN_SECRET='your_access_token_secret'

```


### Posting a New Tweet

The code posts a new tweet using the Twitter API's `POST /2/tweets` endpoint. To avoid issues with duplicate tweets, the program appends a timestamp to the tweet message, ensuring each tweet is unique.

#### Example API Request (Tweet Post)

```json
POST https://api.twitter.com/2/tweets
Authorization: OAuth 1.0a User Context
Content-Type: application/json
```

{
  "text": "Hello from Twitter API! 2024-10-10 12:30:45"
}
```

Execute the below command to run main.go 

```powershell
go run main.go
```
To delete the tweet run below command by passing tweet id using -delete flag
```powershell
go run main.go -delete -tweetID <your_tweet_id>
```

Console Output:
```powershell
PS C:\Users\krish\OneDrive\Documents\Code\go\src\github.com\Krishna868601\twitter_api> go run main.go
Tweet posted successfully! Tweet ID: 1844441015450481058
PS C:\Users\krish\OneDrive\Documents\Code\go\src\github.com\Krishna868601\twitter_api> go run main.go
Tweet posted successfully! Tweet ID: 1844441289619619944
PS C:\Users\krish\OneDrive\Documents\Code\go\src\github.com\Krishna868601\twitter_api> go run main.go
Tweet posted successfully! Tweet ID: 1844441309873860978
PS C:\Users\krish\OneDrive\Documents\Code\go\src\github.com\Krishna868601\twitter_api> go run main.go -delete -tweetID 1844441015450481058
Error while deleting tweet: 200 OK map[data:map[deleted:true]]
PS C:\Users\krish\OneDrive\Documents\Code\go\src\github.com\Krishna868601\twitter_api> go run main.go
Tweet posted successfully! Tweet ID: 1844442379148792319
PS C:\Users\krish\OneDrive\Documents\Code\go\src\github.com\Krishna868601\twitter_api> go run main.go
Tweet posted successfully! Tweet ID: 1844442403756753316
PS C:\Users\krish\OneDrive\Documents\Code\go\src\github.com\Krishna868601\twitter_api> go run main.go -delete -tweetID 1844442379148792319
Tweet deleted successfully!
```

### Error Handling


If the user attempts to post the same tweet content multiple times, the Twitter API will reject the request with a 403 Forbidden error. To resolve this, the program ensures the tweet message includes a timestamp to generate unique content.



