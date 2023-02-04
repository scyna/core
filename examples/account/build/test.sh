curl --location --request POST 'http://localhost:8080/account/create' \
--header 'Client-Id: account_test' \
--header 'Client-Secret: 123456' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "a5@gmail.com",
    "password": "123456",
    "name": "Nguyen Van A"
}'
