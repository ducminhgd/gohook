package googlechat

const (
	IMAGE_TYPE_SQUARE = "SQUARE"
	IMAGE_TYPE_CIRCLE = "CIRCLE"
)

type CardHeader struct {
	Title        string `json:"title"`
	Subtitle     string `json:"subtitle"`
	ImageType    string `json:"imageType"`
	ImageUrl     string `json:"imageUrl"`
	ImageAltText string `json:"imageAltText"`
}

type TextParagraph struct {
	Text string `json:"text"`
}

type Image struct {
	ImageURL string  `json:"imageUrl"`
	OnClick  OnClick `json:"onClick"`
	AltText  string  `json:"altText"`
}

type OnClick struct {
	//https://developers.google.com/chat/api/reference/rest/v1/cards#OnClick_1
}
