POST http://api/user/insert
{
    "name": "",
    "email": "email@example.com"
}

HTTP 400
Content-Type: application/json; charset=utf-8

[Asserts]
jsonpath "$.Error" == "Name cannot be empty"
