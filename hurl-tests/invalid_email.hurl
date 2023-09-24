POST http://api/user/insert
{
    "name": "Example name",
    "email": "invalidemail"
}

HTTP 400
Content-Type: application/json; charset=utf-8

[Asserts]
jsonpath "$.Error" == "Email must have an @ character"
