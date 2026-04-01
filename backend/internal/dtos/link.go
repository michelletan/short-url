package dtos

type CreateLinkRequest struct {
    URL string `json:"url"`
}

type CreateLinkResponse struct {
    URL string `json:"url"`
}

type GetLinksResponse struct {
	URLs []string `json:"urls"`
}