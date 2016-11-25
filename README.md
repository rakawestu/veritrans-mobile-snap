# Veritrans Mobile Snap

Example implementation of Merchant Server required for new [Veritrans Mobile SDK](https://github.com/veritrans/veritrans-android).

## Endpoints

### Normal Transactions
```
POST /charge
```

This endpoint will redirect user request to Snap endpoint `/transactions` with added `Server Key`.

The response will be returned to user.

### Installment

```
POST /installment/charge
```

This will redirect user request to Snap endpoint `/transactions` with added installment data and `Server Key`.

The response will be returned to user.

### Save Card Endpoints

MongoDB database is required to implement this endpoints.
Please provide the URL and collection name using these environment variables

```
MONGODB_URL = mongodb://$username:$password@$mongodb_url:$mongodb_port/$collection_name
MONGODB_NAME = $collection_name
```

#### Get Cards

```
GET /users/{user_id}/tokens
GET /installment/users/{user_id}/tokens
```

 This will return array of JSON object contains saved token and masked card.

#### Save Cards

```
POST /users/{user_id}/tokens
POST /installment/users/{user_id}/tokens
```

This will handle provided array of JSON object contains saved token and masked card to be saved into mongo DB.
## How to deploy to heroku using toolbelt

- Please create a heroku app in your heroku dashboard.
- Please initialise your heroku credentials.

```
heroku login
```
- Add your heroku app into this source code.

```
heroku git:remote -a {{heroku-app-name}}
```

- Please set these environment variables in heroku app settings.
  - SERVER_KEY : Veritrans Server Key
  - PRODUCTION : true to set environment to production mode (make sure your server key also for production)
  - MONGODB_URL : Mongo DB url
  - MONGODB_NAME : Mongo DB collection name
- Deploy changes to heroku.

```
github push heroku master
```
