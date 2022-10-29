# API Tester
A server written in Go using Gin and assert library to test endpoints.

So far this API has one endpoint 'api/test'.

The fields you need to try the endpoint are:
{"Uri": "Uri of the endpoint", "Status Code": number, "Body": "body you want to ckeck"}

Use the method POST to send data about the API you want to test.

#### Example:

{"Uri": "http://localhost:10000/costumer/2", "Status Code": 200, "Body": {"Id":"2", "name": "Carlos Marcano"}}

The server will make a GET request to "http://localhost:10000/costumer/2" and check if the response match with status code and body defined.

If there is no match, the test failed and will send message with the response received and expected.

If all match, then the test passed.
