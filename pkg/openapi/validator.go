package openapi

import "github.com/pb33f/libopenapi"

type Validator interface {
	Validate([]byte) error
}

type OpenAPIv3Validator struct{}

func (v *OpenAPIv3Validator) Validate(b []byte) error {
	document, err := libopenapi.NewDocument(b)
	if err != nil {
		return err
	}

    _, errors := document.BuildV3Model()
    if len(errors) > 0{
        return errors[0]
    }

	return nil
}
