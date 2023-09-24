POST http://api/user/insert
{
    "name": "perfectly valid name",
    "email": "email@example.com"
}

HTTP 201
Content-Type: application/json; charset=utf-8

[Asserts]
jsonpath "$.id" exists
jsonpath "$.name" == "perfectly valid name"
jsonpath "$.email" == "email@example.com"