module main

go 1.17

require discover v0.0.0

require service v0.0.0

require endpoint v0.0.0

require transport v0.0.0

require config v0.0.0

require (
	github.com/go-kit/kit v0.12.0 // indirect
	github.com/go-kit/log v0.2.0 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/satori/go.uuid v1.2.0
)

replace discover => ./discover

replace service => ./service

replace endpoint => ./endpoint

replace transport => ./transport

replace config => ./config
