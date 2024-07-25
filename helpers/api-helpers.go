package helpers

type Response struct {
	Results interface{} `json:"results"`
	Errors  interface{} `json:"errors"`
}

func CreateResponse(res interface{}, err interface{}) Response {
	return Response{
		Results: res,
		Errors:  err,
	}
}
