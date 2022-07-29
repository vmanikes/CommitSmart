# Introduction
I used Golang to build this banking application

## Code Structure
The structure of this code is motivated from Clean Architecture by Uncle Bob

## API Docs
API docs are stored in `docs/swagger.yaml` and you can use any online swagger viewer to convert yaml to UI

## Database
- I have used postgres to have persistence to my application
- To keep things simple I have created only 2 preexisting users in the DB with 0 balance. They are generated during the
migration application phase
- The database migrations are stored in `migrations` directory

## How to run
- I am using docker compose to start the dependencies
- Make sure you have `Make` installed
- Run `make install` to install the tool that is used to create and apply migration
- Run `make start` to run the app and the app should be exposed on port 8080
- Run `make stop` to clear all the docker dependencies

## Privileges took to complete this project in favor of time
- Not handling auth
- Skipping pagination when getting transactions
- Skipping Unit and E2e testing. I know that this might cause a huge issue for my candidacy but this will take me more than
3 hours to write all the test covering all the cases. If you can give me more time to write tests I will add them to the app.
Regardless of they not being present, I tested the app with the best of my knowledge to address all the cases
- Ideally the user / account information will be behind an auth token and passed in HTTP Header of the request. But to keep
  things simple I am passing the account number as a parameter. I know its gonna be sent in Plain Text in the request but I
  did not want to use POST when retrieving information

## Library Usage
- `github.com/go-playground/validator/v10`: Using this to validate request models
- `github.com/shopspring/decimal`: Using this for monetary values as floating point numbers have precision issues
- `github.com/flannel-dev-lab/cyclops/v2`: My web framework
- `github.com/caarlos0/env/v6`: Helps mapping env vars to struct

## Sample Curl
### Add deposit
```shell
curl --location --request POST 'http://localhost:8080/api/v1/deposit' \
--header 'Content-Type: application/json' \
--data-raw '{
  "amount": 1000.05,
  "account_number": 1234
}'
```

### Withdraw
```shell
curl --location --request POST 'http://localhost:8080/api/v1/withdraw' \
--header 'Content-Type: application/json' \
--data-raw '{
  "amount": 100,
  "account_number": 1234,
  "user_id": "e1ef9440-1b5c-4d59-859f-fa2c431c1b94"
}'
```

### Transfer
```shell
curl --location --request POST 'http://localhost:8080/api/v1/transfer' \
--header 'Content-Type: application/json' \
--data-raw '{
  "amount": 100,
  "current_account_number": 1234,
  "destination_account_number": 5678
}'
```

### Transactions
```shell
curl --location --request GET 'http://localhost:8080/api/v1/users/e1ef9440-1b5c-4d59-859f-fa2c431c1b94/transactions?filter_type=deposit&filter_value=2022-07-28'
```