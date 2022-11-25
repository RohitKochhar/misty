module rohitsingh/mistyExamples/entrypoint

go 1.19

require (
	github.com/gorilla/mux v1.8.0
	rohitsingh/misty-go v0.0.0
	rohitsingh/misty-utils v0.0.0
)

replace rohitsingh/misty-go => /clients

replace rohitsingh/misty-utils => /misty-utils
