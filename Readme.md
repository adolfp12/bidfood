How to run
- clone this repository
- `go mod tidy`
- `go run ./cmd/server/main.go`


No.1
Testing endpoint :
- List All Product `curl --location 'http://localhost:8080/products' --header 'X-API-Key: bidfoodkey'`
- Add a New Product `curl --location 'http://localhost:8080/products' --header 'Content-Type: text/plain' --header 'X-API-Key: bidfoodkey' --data '{"name": "Lemon","desc": "Green"}''`
- Retrieve a product by ID `curl --location 'http://localhost:8080/products/1' --header 'X-API-Key: bidfoodkey'`
- Update an Existing Product `curl --location --request PUT 'http://localhost:8080/products/2' --header 'Content-Type: application/json' --header 'X-API-Key: bidfoodkey' --data '{"name": "Apple Green","desc": "Green"}'`
- Delete a Product `curl --location --request DELETE 'http://localhost:8080/products/2' --header 'X-API-Key: bidfoodkey'`


Bonus endpoint:
- Pagination `curl --location 'http://localhost:8080/products?page=1' --header 'X-API-Key: bidfoodkey'`
- Filter. This filter only on Name of Product `curl --location 'http://localhost:8080/products?filter=Apple' --header 'X-API-Key: bidfoodkey'`
- Channel. I use channel on gracefull shutdown. Please find on cmd/server/main.go

Additional :
- I create middlewware to logging http request

No. 2
You find my answer on ./answers/DESIGN.md

No. 3
You find my answer on ./answers/CODEREVIEW.md ; ./cmd/code_review/main.go ; ./cmd/code_review/main_test.go

No. 4
You can find on ./cmd/rate_limiter/main.go