module github.com/malikhan-dev/zenql/databases/v2

go 1.25

require (
	github.com/go-sql-driver/mysql v1.10.0
	github.com/lib/pq v1.12.3
	github.com/malikhan-dev/zenql/streams/v2 v2.0.5
)

require (
	filippo.io/edwards25519 v1.2.0 // indirect
	github.com/malikhan-dev/zenql/contracts/v2 v2.0.5 // indirect
)

replace github.com/malikhan-dev/zenql/contracts/v2 => ../../contracts/v2

replace github.com/malikhan-dev/zenql/streams/v2 => ../../streams/v2
