curl --location --request POST 'http://localhost:8080/ex/hello/hello' \
--header 'Client-Id: hello_test' \
--header 'Client-Secret: 123456' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Alice"
}'

curl --location --request POST 'http://localhost:8080/ex/hello/add' \
--header 'Client-Id: hello_test' \
--header 'Client-Secret: 123456' \
--header 'Content-Type: application/json' \
--data-raw '{
    "a": 23,
    "b": 5
}'

curl --location --request POST 'http://localhost:8080/ex/hello/add' \
--header 'Client-Id: hello_test' \
--header 'Client-Secret: 123456' \
--header 'Content-Type: application/json' \
--data-raw '{
    "a": 85,
    "b": 60
}'

