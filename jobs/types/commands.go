package types

import "github.com/deployment-io/jobs-runner-kit/enums/parameters_enums"

type Runner interface {
	Run(parameters map[parameters_enums.Key]interface{}, logger Logger) (map[parameters_enums.Key]interface{}, error)
}

type Command interface {
	Runner
}