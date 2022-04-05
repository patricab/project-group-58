module distributor

go 1.17

require Driver-go v0.0.0
require Network-go v0.0.0
require network v0.0.0
require fsm v0.0.0
require backup v0.0.0

replace Driver-go => ../Driver-go
replace Network-go => ../Network-go
replace network => ../network
replace fsm => ../fsm
replace backup  => ../backup
