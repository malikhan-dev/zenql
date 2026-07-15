module github.com/malikhan-dev/zenql/integration-test

go 1.25

require github.com/malikhan-dev/zenql/contracts/v2 v2.0.4

replace github.com/malikhan-dev/zenql/contracts/v2 => ../contracts/v2


require github.com/malikhan-dev/zenql/expressions/Sifu v1.0.1

replace github.com/malikhan-dev/zenql/expressions/Sifu => ../expressions/Sifu


require github.com/malikhan-dev/zenql/collections/Thor/v2 v2.0.4

replace github.com/malikhan-dev/zenql/collections/Thor/v2 => ../collections/Thor/v2
