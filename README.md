Run: `docker-compose up -d`

### Endpoints to test

- GET http://localhost:8080/company (Valid)
- POST http://localhost:8080/company (Invalid)
- PUT http://localhost:8080/company (Invalid)
- GET http://localhost:8080/company/abc123asd (Valid)
- GET http://localhost:8080/company/123 (Valid)
- GET http://localhost:8080/companys/123 (Invalid)


### Solution

- I created two applications with Golang, one for the proxy and another one for downstream service. 
- Clean architecture / Repository pattern has been followed. So the project can be updated/extended pretty easily. 
- json file has been considered as a database. Hence the file handling has been done from repository directory. If we want to use other data sources, i.e. SQL/NoSQL DB, Redis etc. we can simply add new files for that in the directory.
- All `http methods` have been handled. So it's possible to request with `GET/POST/PUT...` http methods. We just need to insert the rule/condition in the json file (if we use any persistent DB, we can do the same).

### To improve

- Adding Metrics measurement 
- Making it more configurable
- CLI functionalities could be added
- Structured logging could be added
- TBA