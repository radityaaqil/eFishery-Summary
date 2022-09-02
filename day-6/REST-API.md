# REST API

REST API :

- http/https
- Methods
- Status Codes
- JSON
- Authorization

Tools pemanggil API :
- Postman
- REST CLIENT VS CODE

## Definition
- Representatinal State Transfer (REST) is an architectural style that defines a set of constraints to be used for creating web services
REST API is a way of accessing web services in a simple and flexible way without having any processing. Commonly used nowadays standardized, stateless

- HTTPS --> secured version of HTTP by using asymmetric public key infrastructure (encrypted response)

- HTTP 1 --> one request at a time unlike HTTP 2 (similar to websocket)

- HTTP messages/body/payload. Two types of messages req by client to trigger an action on the server and res, the answer from the server
- HTTP headers let the client and the server pass additional information with an HTTP request or response. An HTTP Header consists of its case-insensitive name followed by a colon(:), then by its value

## REST METHODS
Methods :

- GET
- POST
- PUT
- PATCH
- DELETE

## Status Codes
### Status Code : 2xx - success ; 4xx - error caused by client ; 5xx - error caused by server
2xx (usually with GET method) :
- 200 OK
- 201 CREATED - Indicates that the request has succeeded and a new resource has been created as a result
- 202 ACCEPTED - Indicates that the request has been received but not completed yet. It is typically used in log running requests and batch processing

4xx :
- 400 - Bad Request - Request could not be understood by the server due to incorrect syntax. The client SHOULD NOT repeat the request without modification
- 401 - Unauthorized - Indicates that the req requires user authentication information
- 403 - Forbidden - Unauthorized request. The client does not have access rights to the content. Unlike 401, the client's identity is known to the server
- 405 - Method not allowed - Req is known but has been disabled and cannot be used for that resource

5xx :
- 500 - Internal Server Error - The Server encountered a unexpected condition that prevented it from fulfilling the request.
- 502 - Bad Gateway - The server got an invalid response while working as a gateway to get the response needed to handle the request.
- 504 - Gateway Timeout - The server is acting as a gateway and cannot get a response in time for a request.

## JSON
JSON - (Javascript Object Notation) - lightweight data-interchange format.

JSON Web Token - Header, Payload, Signature
- Header - Algorithm & token type
- Payload - Claims, usually additional user data
- Signature - To create the signature part you have to take the encoded header, the encoded payload, a secret, the algorithm specified in the header and sign that.

## ECHO FRAMEWORK
ECHO FRAMEWORK - Hi Performance, extensible, minimalist GO framework

## TESTING
Same folder as file to be tested. ended with _testing.go

- Unit Tests --> Integration Tests --> End to end tests
slower-->>>>
- End to end Tests --> Integration Tests --> Unit Testing
cheaper-->>>>

Unit tests --> every functions
E2E & Integration --> functionality tests

## LOGGING

import "log"

Log to terminal and to file

TASK ---> CRUD APP with log

GLOG & LOGRUS --> eFishery uses logrus