POST http://api/user/insert
{
    "name": "Example name",
    "email": ""
}

HTTP 400
Content-Type: application/json; charset=utf-8

[Asserts]
jsonpath "$.Error" == "Email cannot be empty"
