# mgorepo

An extensible and generic mongo repository pattern implementation

## Requirements

- Golang >= 1.20
- Docker (for testing only)

## CLI

**Pending to add**

## Installation

```bash
go get github.com/Drafteame/mgorepo
```

## Usage

This package was made with the purpose of have a way to generify mongo actions using the repository pattern on a DDD
environment. It's not an ORM, it's just a way to have a generic way to interact with mongo.

### Main concepts

#### Repository

A repository is a struct that will be used to interact with the database, it should be created using the `NewRepository`
generic function and can be configured to be extended with custom methods using struct embedding.

The sign of the `NewRepository` function is:

```go
func NewRepository[
	M    Model,
	D    Dao,
	SF   SearchFilters,
	SORD SearchOrders,
	SO   SearchOptions[SORD, SF],
	UF   UpdateFields,
](
	db             Driver,
	collectionName string,
	filterBuilders []func(SF) (*bson.E, error),
	orderBuilders  []func(SORD) (*bson.E, error),
	updateBuilders []func(UF) (*bson.E, error),
) Repository[M, D, SF, SORD, SO, UF] {}
```

##### Generic constraints

- `M Model`: Is a struct that will be used to interact with the repository, it should be created using just native go 
  types and should contain all document fields of a collection. Fields should be public.
- `D Dao`: Is a struct that will be used to interact with the database, it should be created using mongo data types,
  should contain all document fields of a collection and be compliant with the generic interface `DaoFiller[M Model]`.
  Fields should be public.
  ID field of a `Dao` should be defined as `ID primitive.ObjectID` using tag `bson:"_id,omitempty"`.
- `SF SearchFilters`: Is a struct that will be used to filter the results of a search, it should be created using just
  pointers of native go types and each field should represent a type of filter that can be applied to the collection. Fields should
  be public.
- `SORD SearchOrders`: Is a struct that will be used to order the results of a search, it should be created using just
  native go types and each field should represent a type of order that can be applied to the collection. fields should
  be public.
- `SO SearchOptions[SORD, SF]`: Is an interface that will be used to configure the options of a search. This interface 
  contain 4 methods to bring information to a search:
  - `GetFilters() SF`: Should return a `SF` struct with the filters that should be applied to the search.
  - `GetOrders() SORD`: Should return a `SORD` struct with the orders that should be applied to the search.
  - `GetLimit() int64`: Should return an int64 with the limit of results that should be returned by the 
    search.
  - `GetSkip() int64`: Should return an int64 with the number of results that should be skipped by the search.
- `UF UpdateFields`: Is a struct that will be used to update a document, it should be created using just native go types
  and each field should represent a possible updated fields of a document. Fields should be public.

##### Parameters

- `db Driver`: Is an interface that implement 2 methods:
  - `Client() *mongo.Client`: Should return a mongo client that will be used to interact with the database (package used
    to interact with mongo is `go.mongodb.org/mongo-driver/mongo`).
  - `DbName() string`: Should return a string with the name of the database that will be used on the operations.
- `collectionName string`: Is a string with the name of the collection that will be used by the repository.
- `filterBuilders []func(SF) (*bson.E, error)`: Is a slice of functions that will be used to build the filters that will
  be applied to the search. Each function should receive a `SF` struct and return a `*bson.E` with the filter that will
  be applied to the search. If an error is returned, the operation will be aborted. Each function should build a filter 
  based in just one field of the `SF` struct, making as result that the slice of functions should have the same length
  as the number of fields of the `SF` struct.
- `orderBuilders []func(SORD) (*bson.E, error)`: Is a slice of functions that will be used to build the orders that will
  be applied to the search. Each function should receive a `SORD` struct and return a `*bson.E` with the order that will
  be applied to the search. If an error is returned, the operation will be aborted. Each function should build an order 
  based in just one field of the `SORD` struct, making as result that the slice of functions should have the same length
  as the number of fields of the `SORD` struct.
- `updateBuilders []func(UF) (*bson.E, error)`: Is a slice of functions that will be used to build the updates that will
  be applied to the document. Each function should receive a `UF` struct and return a `*bson.E` with the update that will
  be applied to the document. If an error is returned, the operation will be aborted. Each function should build an 
  update based in just one field of the `UF` struct, making as result that the slice of functions should have the same 
  length as the number of fields of the `UF` struct.

##### Example

- Basic usage
- Setup timestamps
- Clock interface for timestamps
- Filtering different that equals
- Option builders pattern
- Custom logger
- Custom repository methods
