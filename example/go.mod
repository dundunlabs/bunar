module github.com/dundunlabs/bunar/example

go 1.19

replace github.com/dundunlabs/bunar => ../

require (
	github.com/dundunlabs/bunar v0.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.4.0
	github.com/uptrace/bun v1.1.10
	github.com/uptrace/bun/dialect/pgdialect v1.1.10
	github.com/uptrace/bun/driver/pgdriver v1.1.10
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/tmthrgd/go-hex v0.0.0-20190904060850-447a3041c3bc // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	golang.org/x/crypto v0.5.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	mellium.im/sasl v0.3.1 // indirect
)
