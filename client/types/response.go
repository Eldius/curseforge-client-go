package types

type CurseforgeAPIResponse interface {
	SetRawResponseBody(string)
}

type RawResponse struct {
	RawBody string `json:"body"`
}

func (r *RawResponse) SetRawResponseBody(resp string) {
	r.RawBody = resp
}
