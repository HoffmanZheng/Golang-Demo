module main

go 1.17

require discover v0.0.0

require service v0.0.0

require endpoint v0.0.0

require transport v0.0.0

require config v0.0.0

require github.com/satori/go.uuid v1.2.0 // indirect

replace discover => ./discover

replace service => ./service

replace endpoint => ./endpoint

replace transport => ./transport

replace config => ./config
