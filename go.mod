module main

go 1.17

require fsm v0.0.0
require Driver-go v0.0.0
require Network-go v0.0.0
require network v0.0.0
require distributor v0.0.0
require backup v0.0.0

replace fsm  => ./fsm
replace backup  => ./backup
replace Driver-go => ./Driver-go
replace Network-go => ./Network-go
replace network => ./network
replace distributor => ./distributor
