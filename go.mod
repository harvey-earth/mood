module github.com/harvey-earth/mood

go 1.22.4

require (
	github.com/go-sql-driver/mysql v1.8.1
	github.com/gorilla/mux v1.8.1
	github.com/mattn/go-sqlite3 v1.14.22
	go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux v0.55.0
	go.opentelemetry.io/otel v1.30.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.30.0
	go.opentelemetry.io/otel/sdk v1.30.0
	go.opentelemetry.io/otel/trace v1.30.0
)

require (
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	go.opentelemetry.io/otel/metric v1.30.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/lib/pq v1.10.9
)
