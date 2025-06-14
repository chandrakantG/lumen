# lumen

1. create database named `lumen`

2. executes schema.sql file to create tables and add data in tables

3. run application waith `make run` command

4. use below curl for get product revenue 
```
curl --location 'http://localhost:8081/api/v1/revenue/product' \
--header 'Content-Type: application/json' \
--data '{
    "start":"2025-5-11",
    "end":"2025-5-16"
}'
```