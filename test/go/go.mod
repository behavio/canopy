module canopy/test

go 1.22.0

require (
	choicesgoparser v0.0.0
	extensionsgoparser v0.0.0
	nodeactionsgoparser v0.0.0
)

replace (
	choicesgoparser => ../grammars/choices-go
	extensionsgoparser => ../grammars/extensions-go
	nodeactionsgoparser => ../grammars/node_actions-go
)
