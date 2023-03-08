# Test a weird case 
POST http://localhost:{{port}}/fizzbuzz
{
	"int1": 4,
	"int2": 3,
	"limit": 15
}
HTTP 200
[Asserts]
jsonpath "$.result[0]" == "1" 
jsonpath "$.result[2]" == "" 
jsonpath "$.result[4]" == "5" 
jsonpath "$.result[11]" == "" 
jsonpath "$.result" count == 15


# Test a wrong case with negative int
POST http://localhost:{{port}}/fizzbuzz
{
	"int1": -2,
	"int2": 3,
	"limit": 12,
	"str1": "fizz",
	"str2": "buzz"
}
HTTP 500
[Asserts]
jsonpath "$.error" == "input integers must be strictly positive" 


# Test a typical case 
POST http://localhost:{{port}}/fizzbuzz
{
	"int1": 2,
	"int2": 3,
	"limit": 12,
	"str1": "fizz",
	"str2": "buzz"
}
HTTP 200
[Asserts]
jsonpath "$.result[0]" == "1" 
jsonpath "$.result[2]" == "buzz" 
jsonpath "$.result[5]" == "fizzbuzz" 
jsonpath "$.result[11]" == "fizzbuzz" 
jsonpath "$.result" count == 12
