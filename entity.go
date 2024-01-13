package urlmetadata

type UrlMetadataIconData struct {
	Type  string `json:"type"`
	Href  string `json:"href"`
	Color string `json:"color"`
}

type UrlMetadataIcon struct {
	Icon     UrlMetadataIconData `json:"icon"`
	IconMask UrlMetadataIconData `json:"iconMask"`
	IconAlt  UrlMetadataIconData `json:"iconAlt"`
}

type UrlMetadataOpenGraph struct {
	Url            string `json:"url"`
	Locale         string `json:"locale"`
	Title          string `json:"title"`
	Type           string `json:"type"`
	Description    string `json:"description"`
	Image          string `json:"image"`
	ImageType      string `json:"imageType"`
	ImageAlt       string `json:"imageAlt"`
	ImageWidth     string `json:"imageWidth"`
	ImageHeight    string `json:"imageHeight"`
	ImageSecureUrl string `json:"imageSecureUrl"`
	SiteName       string `json:"siteName"`
}

type UrlMetadataTwitter struct {
	Url           string         `json:"url"`
	Title         string         `json:"title"`
	Description   string         `json:"description"`
	Image         string         `json:"image"`
	ImageAlt      string         `json:"imageAlt"`
	SiteName      string         `json:"siteName"`
	Card          string         `json:"card"`
	Creator       string         `json:"creator"`
	CreatorId     string         `json:"creatorId"`
	AppIphone     UrlMetadataApp `json:"appIphone"`
	AppIpad       UrlMetadataApp `json:"appIpad"`
	AppGooglePlay UrlMetadataApp `json:"appGooglePlay"`
}

type UrlMetadataApp struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type UrlMetadataMsApp struct {
	TileImage string `json:"tileImage"`
	TileColor string `json:"tileColor"`
}

type UrlMetadataAppleApp struct {
	AppleMobileWebAppCapable        string `json:"appleMobileWebAppCapable"`
	AppleMobileWebAppStatusBarStyle string `json:"appleMobileWebAppStatusBarStyle"`
	AppleMobileWebAppTitle          string `json:"appleMobileWebAppTitle"`
	AppleMobileWebAppIcon           string `json:"appleMobileWebAppIcon"`
}

type UrlMetadata struct {
	Icon        UrlMetadataIcon      `json:"icon"`
	Title       string               `json:"title"`
	Description string               `json:"description"`
	OpenGraph   UrlMetadataOpenGraph `json:"openGraph"`
	Twitter     UrlMetadataTwitter   `json:"twitter"`
	Body        string               `json:"body"`
	ViewPort    string               `json:"viewPort"`
	ThemeColor  string               `json:"themeColor"`
	ColorScheme string               `json:"colorScheme"`
	AppleApp    UrlMetadataAppleApp  `json:"appleApp"`
	MsApp       UrlMetadataMsApp     `json:"msApp"`
	HostName    string               `json:"hostName"`
}
