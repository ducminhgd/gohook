package googlechat

const (
	IMAGE_TYPE_SQUARE = "SQUARE"
	IMAGE_TYPE_CIRCLE = "CIRCLE"

	LOAD_INDICATOR_NONE    = "NONE"
	LOAD_INDICATOR_SPINNER = "SPINNER"

	INTERACTION_UNSPECIFIED = "INTERACTION_UNSPECIFIED"
	INTERACTION_OPEN_DIALOG = "OPEN_DIALOG"

	OPEN_AS_FULLSIZE = "FULL_SIZE"
	OPEN_AS_OVERLAY  = "OVERLAY"

	ON_CLOSE_NOTHING = "NOTHING"
	ON_CLOSE_RELOAD  = "RELOAD"

	CONTROL_TYPE_SWITCH    = "SWITCH"
	CONTROL_TYPE_CHECK_BOX = "CHECK_BOX"
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

type ActionParameter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Action struct {
	Function      string            `json:"function"`
	Parameters    []ActionParameter `json:"parameters"`
	LoadIndicator string            `json:"loadIndicator"`
	PersistValues bool              `json:"persistValues"`
	Interaction   string            `json:"interaction"`
}

type OnClick struct {
	Action                Action   `json:"action"`
	OpenLink              OpenLink `json:"openLink"`
	OpenDynamicLinkAction Action   `json:"openDynamicLinkAction"`
	Card                  Card     `json:"card"`
}

type OpenLink struct {
	URL     string `json:"url"`
	OpenAs  string `json:"openAs"`
	OnClose string `json:"onClose"`
}

type Section struct {
	Header                    string `json:"header"`
	Widgets                   Widget `json:"widgets"`
	Collapsible               bool   `json:"collapsible"`
	UncollapsibleWidgetsCount int32  `json:"uncollapsibleWidgetsCount"`
}

type DecoratedText struct {
	StartIcon   Icon    `json:"startIcon"`
	TopLabel    string  `json:"topLabel"`
	Text        string  `json:"text"`
	WrapText    bool    `json:"wrapText"`
	BottomLabel string  `json:"bottomLabel"`
	OnClick     OnClick `json:"onClick"`

	// Union field control can be only one of the following:
	Button        Button        `json:"button"`
	SwitchControl SwitchControl `json:"switchControl"`
	EndIcon       Icon          `json:"endIcon"`
	// End of list of possible types for union field control.
}

type Button struct {
	Text     string  `json:"text"`
	Icon     Icon    `json:"icon"`
	Color    Color   `json:"color"`
	OnClick  OnClick `json:"onClick"`
	Disabled bool    `json:"disabled"`
	AltText  string  `json:"altText"`
}

type SwitchControl struct {
	Name           string `json:"name"`
	Value          string `json:"value"`
	Selected       bool   `json:"selected"`
	OnChangeAction Action `json:"onChangeAction"`
	ControlType    string `json:"controlType"`
}

type Color struct {
	Red   float32 `json:"red"`
	Green float32 `json:"green"`
	Blue  float32 `json:"blue"`
	Alpha float32 `json:"alpha"`
}

type Icon struct {
	AltText   string `json:"altText"`
	ImageType string `json:"imageType"`

	// Union field control can be only one of the following:
	KnownIcon string `json:"knownIcon"`
	IconURL   string `json:"iconUrl"`
	// End of list of possible types for union field control.
}

type Widget struct {
	CardID string `json:"cardId"`
	Card   Card   `json:"card"`
}

type Card struct {
	Header      CardHeader   `json:"header"`
	Sections    []Section    `json:"sections"`
	CardActions []CardAction `json:"cardActions"`
	Name        string       `json:"name"`
	// FixedFooter    CardFixedFooter `json:"fixedFooter"`
	DisplayStyle   string     `json:"displayStyle"`
	PeekCardHeader CardHeader `json:"peekCardHeader"`
}

type CardAction struct {
	Name    string  `json:"name"`
	OnClick OnClick `json:"onClick"`
}

// type CardFixedFooter struct {
// 	PrimaryButton   Button `json:"primaryButton"`
// 	SecondaryButton Button `json:"secondaryButton"`
// }
