package api

type GetAgreementResponse struct {
	HasAgreedNotice int64 `json:"hasAgreedNotice"`
	HasAgreedHonest int64 `json:"hasAgreedHonest"`
}
