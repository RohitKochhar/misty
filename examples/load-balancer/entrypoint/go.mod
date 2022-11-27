module rohitsingh/mistyExamples/entrypoint

go 1.19

require (
	github.com/gorilla/mux v1.8.0
	github.com/rohitkochhar/reed-http-utills v0.1.2
	rohitsingh/misty-go v0.0.0
)

// This line should be uncommented before containerization
replace rohitsingh/misty-go => /clients

// This line should be uncommented before running a go mod tidy locally
// replace rohitsingh/misty-go => ../../../clients/misty-go
