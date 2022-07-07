module service

go 1.17

require discover v0.0.0

require config v0.0.0

require (
	github.com/go-kit/kit v0.12.0 // indirect
	github.com/go-kit/log v0.2.0 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
)

replace discover => ../discover

replace config => ../config
