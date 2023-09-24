GET http://api/user/insert

Content-Type: application/json

HTTP 405
Content-Type: application/json; charset=utf-8

[Asserts]
jsonpath "$.Error" == "Method not allowed"