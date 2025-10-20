module golangexample

go 1.22

toolchain go1.24.2

require (
	jsongoparser v0.0.0
	lispgoparser v0.0.0
	peggoparser v0.0.0
)

replace jsongoparser => ../canopy/json-go

replace lispgoparser => ../canopy/lisp-go

replace peggoparser => ../canopy/peg-go
