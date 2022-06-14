Run: `docker-compose up -d`

FxProxy: 
- http://localhost:8080/company (Valid)
- http://localhost:8080/company/abc123asd (Valid)
- http://localhost:8080/company/123 (Valid)
- http://localhost:8080/companys/123 (Invalid)


### Solution

I created two go apps, one for the proxy and another one for downstream service. Clean architecture / Repository pattern has been followed. 

### To improve

- Adding Metrics measurement 
- Making it more configurable
- CLI functionalities could be added
- Structured logging could be added
- TBA