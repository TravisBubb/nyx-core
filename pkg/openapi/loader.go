package openapi

type Loader interface{
    Load([]byte) error
}
