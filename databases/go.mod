module github.com/malikhan-dev/zenql/databases

go 1.25

require (
	github.com/go-sql-driver/mysql v1.10.0
	github.com/lib/pq v1.12.3
	filippo.io/edwards25519 v1.2.0 // indirect
	github.com/malikhan-dev/zenql/contracts v2.0.0
	github.com/malikhan-dev/zenql/streams v2.0.0
)
replace github.com/malikhan-dev/zenql/contracts => ../contracts
replace github.com/malikhan-dev/zenql/streams => ../streams