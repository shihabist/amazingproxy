Run: `docker-compose up -d`

Run Tests:
- `go test -v ./...`
### Endpoints to test

| Http method | URL | Validity |
|-------------|-----|----------|
| GET | http://localhost:8080/company | (Valid) |
| POST | http://localhost:8080/company | (Valid) |
| PUT | http://localhost:8080/company | (Invalid) |
| GET | http://localhost:8080/company/abc123asd | (Valid) |
| GET | http://localhost:8080/company/123 | (Valid) |
| GET | http://localhost:8080/companys/123 | (Invalid) |


### Solution

- I created two applications with Golang, one for the proxy and another one for downstream service. 
- Clean architecture / Repository pattern has been followed. So the project can be updated/extended pretty easily. 
- json file has been considered as a `database`. Hence the file handling has been done from repository directory. If we want to use other data sources, i.e. SQL/NoSQL DB, Redis etc. we can simply add new files for that in the directory.
- We can also use other feature-rich routing packages easily, just by adding a file for that which will implement the route interface{}.
- All `http methods` have been handled. So it's possible to request with `GET/POST/PUT...` http methods. We just need to insert the rule/condition in the json file (if we use any persistent DB, we can do the same).

### Further improvement

- Adding Metrics measurement 
- Making it more configurable
- CLI functionalities could be added
- Structured logging, request tracing could be added
- gRPC could be used for service-service communication
- We could use some load balancing algorithms (i.e. Round-robin) in case there are multiple instances of same service
- More tests should be added. Also each package of the codebase should have tests
- Code coverage, benchmarking data should have been added
- If there were lots of (say millions) `allowedUri` we could use `goroutines` while traversing over the list and matching them with `incomingUri`
- We would see code performance in a better way if we could test it with `pprof` and `trace`
- Necessary comments could be added 