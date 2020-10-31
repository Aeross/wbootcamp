# wbootcamp
Golang Bootcamp challenge

To run the project (Go 1.15 required):

Clone the project, navigate to the wbootcamp directory and run the project with: 

``go run main.go``

the only dependency (gorilla mux) will be downloaded automatically.

To use another port:

``go run main.go -addr :PORT``

Two endpoints will be available: ``/hello`` and ``/rickmorty/{id}`` where ``{id}`` is a positive integer. The max amount of ids are **671**.

The first endpoint will return:

``"Hello World!"``

The second endpoint will return information about a Rick and Morty character based on the provided character id. 
There is a total of **671** characters according to the Rick and Morty API. Therefore, any number above **671** will return an error.

i.e: 

GET to ``http://localhost:4000/rickmorty/1``

returns:

```
{
    "id": 1,
    "name": "Rick Sanchez",
    "status": "Alive",
    "species": "Human",
    "location": {
        "name": "Earth (Replacement Dimension)",
        "url": "https://rickandmortyapi.com/api/location/20"
    },
    "origin": {
        "name": "Earth (C-137)",
        "url": "https://rickandmortyapi.com/api/location/1"
    }
}
```

GET to ``http://localhost:4000/rickmorty/672``

returns:

```
{
    "error": "Character not found"
}
```

To run the tests:

Assuming that you're in the wbootcamp directory. Execute:

``go test -v ./handlers``

You'll get a ``PASS`` output like:

```
=== RUN   TestHelloWorld
--- PASS: TestHelloWorld (0.00s)
=== RUN   TestGetCharacterByID
--- PASS: TestGetCharacterByID (0.48s)
PASS
ok      github.com/Aeross/wbootcamp/handlers    0.486s
```
modify the expected string to something different like:

``expected := "Random String!"``

To get a ``FAIL`` output like:

```
=== RUN   TestHelloWorld
    handlers_test.go:51: unexpected body response: got Hello World! want Random String!
--- FAIL: TestHelloWorld (0.00s)
=== RUN   TestGetCharacterByID
--- PASS: TestGetCharacterByID (0.49s)
FAIL
FAIL    github.com/Aeross/wbootcamp/handlers    0.498s
FAIL
```
