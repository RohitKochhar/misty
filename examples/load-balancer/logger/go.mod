module rohitsingh/mistyExamples/logger

go 1.19

require rohitsingh/misty-go v0.0.0

require (
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/rohitkochhar/reed-http-utills v0.1.2 // indirect
)

// This line should be uncommented before containerization
replace rohitsingh/misty-go => /clients

// This line should be uncommented before running a go mod tidy locally
// replace rohitsingh/misty-go => ../../../clients/misty-go
