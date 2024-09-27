package types

type CurseforgeAPIResponse interface {
	SetResponse(string)
}

type RawResponse struct {
	RawBody string `json:"body"`
}

func (r RawResponse) SetResponse(resp string) {
	r.RawBody = resp
}
