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

```bash
export TWITTER_API_KEY='your_api_key'
export TWITTER_API_SECRET_KEY='your_api_secret_key'
export TWITTER_ACCESS_TOKEN='your_access_token'
export TWITTER_ACCESS_TOKEN_SECRET='your_access_token_secret' 
```



