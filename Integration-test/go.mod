module github.com/malikhan-dev/zenql/integration-test

go 1.27.3

replace github.com/malikhan-dev/zenql/contracts/v2 => ../contracts/v2

replace github.com/malikhan-dev/zenql/expressions/Sifu => ../expressions/Sifu

require (
	github.com/malikhan-dev/zenql/collections/Thor/v2 v2.0.5
	github.com/malikhan-dev/zenql/contracts/v2 v2.0.7
	github.com/malikhan-dev/zenql/expressions/Sifu v1.0.4
)

replace github.com/malikhan-dev/zenql/collections/Thor/v2 => ../collections/Thor/v2
