package types

type CurseforgeAPIResponse interface {
	SetRawResponseBody(resp string)
}

type RawResponse struct {
	CurseforgeAPIResponse
	RawBody string `json:"body"`
}

func (r *RawResponse) SetRawResponseBody(resp string) {
	r.RawBody = resp
}
