POST http://api/user/insert

Content-Type: application/json

HTTP 400
Content-Type: application/json; charset=utf-8

[Asserts]
jsonpath "$.Error" == "EOF"
