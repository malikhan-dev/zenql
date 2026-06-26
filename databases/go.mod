module github.com/malikhan-dev/zenql/databases

go 1.25


require (
	github.com/go-sql-driver/mysql v1.10.0
	github.com/lib/pq v1.12.3
)

require filippo.io/edwards25519 v1.2.0 // indirect



require github.com/malikhan-dev/zenql/contracts v1.8.3

replace github.com/malikhan-dev/zenql/contracts => ../contracts

require github.com/malikhan-dev/zenql/streams v1.8.3

replace github.com/malikhan-dev/zenql/streams => ../streams