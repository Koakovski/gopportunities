package handler

import (
	"fmt"
)

func errParamIsRequired(paramName, paramType string) error {
	return fmt.Errorf("param: '%s' ( type %s ) is required", paramName, paramType)
}

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   uint64 `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}

	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}

	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}

	if r.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}

	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}

	if r.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}

	return nil
}

// UpdateOpeningRequest

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   uint64 `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	var hasData bool

	if r.Role != "" {
		hasData = true
	}

	if r.Company != "" {
		hasData = true
	}

	if r.Location != "" {
		hasData = true
	}

	if r.Remote != nil {
		hasData = true
	}

	if r.Link != "" {
		hasData = true
	}

	if r.Salary > 0 {
		hasData = true
	}

	if hasData {
		return nil
	}

	return fmt.Errorf("at least one field must be provided to update")
}
