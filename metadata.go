package urlmetadata

import (
	"bytes"
	"net/http"

	"io"

	"golang.org/x/net/html"
)

type UrlMetadataParseOptions struct {
	Body bool
}

func Parse(url string, opts ...UrlMetadataParseOptions) (*UrlMetadata, error) {
	opt := UrlMetadataParseOptions{
		Body: false,
	}

	if len(opts) > 0 {
		opt = opts[0]
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyStr := ""

	if opt.Body {
		bodyB, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		bodyStr = string(bytes.Replace(bodyB, []byte("\r"), []byte("\r\n"), -1))
	}

	meta := extract(resp.Body)
	meta.Body = bodyStr

	return meta, nil
}

func extract(resp io.Reader) *UrlMetadata {
	z := html.NewTokenizer(resp)

	titleFound := false

	meta := new(UrlMetadata)

	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return meta
		case html.StartTagToken, html.SelfClosingTagToken:
			t := z.Token()
			if t.Data == `body` {
				return meta
			}
			if t.Data == "title" {
				titleFound = true
			}
			if t.Data == "link" {
				href, contentType, color, ok := extractLinkProperty(t, "icon")
				if ok {
					meta.Icon.Icon.Href = href
					meta.Icon.Icon.Type = contentType
					meta.Icon.Icon.Color = color
				}

				href, contentType, color, ok = extractLinkProperty(t, "mask-icon")
				if ok {
					meta.Icon.IconMask.Href = href
					meta.Icon.IconMask.Type = contentType
					meta.Icon.IconMask.Color = color
				}

				href, contentType, color, ok = extractLinkProperty(t, "alternate icon")
				if ok {
					meta.Icon.IconAlt.Href = href
					meta.Icon.IconAlt.Type = contentType
					meta.Icon.IconAlt.Color = color
				}
			}
			if t.Data == "meta" {
				// Root
				v, ok := extractMetaProperty(t, "description")
				if ok {
					meta.Description = v
				}

				v, ok = extractMetaProperty(t, "hostname")
				if ok {
					meta.HostName = v
				}

				v, ok = extractMetaProperty(t, "theme-color")
				if ok {
					meta.ThemeColor = v
				}

				v, ok = extractMetaProperty(t, "color-scheme")
				if ok {
					meta.ColorScheme = v
				}

				v, ok = extractMetaProperty(t, "viewport")
				if ok {
					meta.ViewPort = v
				}

				// OpenGraph
				v, ok = extractMetaProperty(t, "og:url")
				if ok {
					meta.OpenGraph.Url = v
				}

				v, ok = extractMetaProperty(t, "og:locale")
				if ok {
					meta.OpenGraph.Locale = v
				}

				v, ok = extractMetaProperty(t, "og:title")
				if ok {
					meta.OpenGraph.Title = v
				}

				v, ok = extractMetaProperty(t, "og:type")
				if ok {
					meta.OpenGraph.Type = v
				}

				v, ok = extractMetaProperty(t, "og:description")
				if ok {
					meta.OpenGraph.Description = v
				}

				v, ok = extractMetaProperty(t, "og:image")
				if ok {
					meta.OpenGraph.Image = v
				}

				v, ok = extractMetaProperty(t, "og:image:alt")
				if ok {
					meta.OpenGraph.ImageAlt = v
				}

				v, ok = extractMetaProperty(t, "og:image:type")
				if ok {
					meta.OpenGraph.ImageType = v
				}

				v, ok = extractMetaProperty(t, "og:image:width")
				if ok {
					meta.OpenGraph.ImageWidth = v
				}

				v, ok = extractMetaProperty(t, "og:image:height")
				if ok {
					meta.OpenGraph.ImageHeight = v
				}

				v, ok = extractMetaProperty(t, "og:image:secure_url")
				if ok {
					meta.OpenGraph.ImageSecureUrl = v
				}

				v, ok = extractMetaProperty(t, "og:site_name")
				if ok {
					meta.OpenGraph.SiteName = v
				}

				// Twitter
				v, ok = extractMetaProperty(t, "twitter:card")
				if ok {
					meta.Twitter.Card = v
				}

				v, ok = extractMetaProperty(t, "twitter:title")
				if ok {
					meta.Twitter.Title = v
				}

				v, ok = extractMetaProperty(t, "twitter:description")
				if ok {
					meta.Twitter.Description = v
				}

				v, ok = extractMetaProperty(t, "twitter:site")
				if ok {
					meta.Twitter.SiteName = v
				}

				v, ok = extractMetaProperty(t, "twitter:creator")
				if ok {
					meta.Twitter.Creator = v
				}

				v, ok = extractMetaProperty(t, "twitter:creator:id")
				if ok {
					meta.Twitter.CreatorId = v
				}

				v, ok = extractMetaProperty(t, "twitter:image")
				if ok {
					meta.Twitter.Image = v
				}

				v, ok = extractMetaProperty(t, "twitter:image:alt")
				if ok {
					meta.Twitter.ImageAlt = v
				}

				v, ok = extractMetaProperty(t, "twitter:app:name:iphone")
				if ok {
					meta.Twitter.AppIphone.Name = v
				}

				v, ok = extractMetaProperty(t, "twitter:app:id:iphone")
				if ok {
					meta.Twitter.AppIphone.Id = v
				}

				v, ok = extractMetaProperty(t, "twitter:app:url:iphone")
				if ok {
					meta.Twitter.AppIphone.Url = v
				}

				v, ok = extractMetaProperty(t, "twitter:app:name:ipad")
				if ok {
					meta.Twitter.AppIpad.Name = v
				}

				v, ok = extractMetaProperty(t, "twitter:app:id:ipad")
				if ok {
					meta.Twitter.AppIpad.Id = v
				}

				v, ok = extractMetaProperty(t, "twitter:app:url:ipad")
				if ok {
					meta.Twitter.AppIpad.Url = v
				}

				v, ok = extractMetaProperty(t, "twitter:app:name:googleplay")
				if ok {
					meta.Twitter.AppGooglePlay.Name = v
				}

				v, ok = extractMetaProperty(t, "twitter:app:id:googleplay")
				if ok {
					meta.Twitter.AppGooglePlay.Id = v
				}

				v, ok = extractMetaProperty(t, "twitter:app:url:googleplay")
				if ok {
					meta.Twitter.AppGooglePlay.Url = v
				}

				// Ms App
				v, ok = extractMetaProperty(t, "twitter:app:url:googleplay")
				if ok {
					meta.Twitter.AppGooglePlay.Url = v
				}

			}
		case html.TextToken:
			if titleFound {
				t := z.Token()
				meta.Title = t.Data
				titleFound = false
			}
		}
	}
}

func extractMetaProperty(t html.Token, prop string) (content string, ok bool) {
	for _, attr := range t.Attr {
		if (attr.Key == "property" && attr.Val == prop) || (attr.Key == "name" && attr.Val == prop) {
			ok = true
		}

		if attr.Key == "content" {
			content = attr.Val
		}
	}

	return
}

func extractLinkProperty(t html.Token, prop string) (href, contentType, color string, ok bool) {
	for _, attr := range t.Attr {
		if attr.Key == "rel" && attr.Val == prop {
			ok = true
		}

		if attr.Key == "href" {
			href = attr.Val
		}

		if attr.Key == "type" {
			contentType = attr.Val
		}

		if attr.Key == "color" {
			color = attr.Val
		}
	}

	return
}
