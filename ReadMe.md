# Url Metadata

Url Metadata parser for Go.

## How to

1. Install

    ```go
    go get -u github.com/ariefsn/url-metadata
    ```

2. Import

    ```go
    import (
      "fmt"
      "github.com/ariefsn/urlmetadata"
    )
    ```

3. Use it

    ```go
      url := "https://github.com/ariefsn"

      opt := urlmetadata.UrlMetadataParseOptions{
        Body: false, // set true to parse the body string 
      }

      meta, err := urlmetadata.Parse(url, opt)
      if err != nil {
        fmt.Println(err)
        return
      }
    ```

4. Result

    ```json
      {
        "icon": {
          "icon": {
            "type": "image/svg+xml",
            "href": "https://github.githubassets.com/favicons/favicon.svg",
            "color": ""
          },
          "iconMask": {
            "type": "",
            "href": "https://github.githubassets.com/assets/pinned-octocat-093da3e6fa40.svg",
            "color": "#000000"
          },
          "iconAlt": {
            "type": "image/png",
            "href": "https://github.githubassets.com/favicons/favicon.png",
            "color": ""
          }
        },
        "title": "ariefsn (Arief Setiyo Nugroho) · GitHub",
        "description": "ariefsn has 43 repositories available. Follow their code on GitHub.",
        "openGraph": {
          "url": "https://github.com/ariefsn",
          "locale": "",
          "title": "ariefsn - Overview",
          "type": "profile",
          "description": "ariefsn has 43 repositories available. Follow their code on GitHub.",
          "image": "https://avatars.githubusercontent.com/u/33569737?v=4?s=400",
          "imageType": "",
          "imageAlt": "ariefsn has 43 repositories available. Follow their code on GitHub.",
          "imageWidth": "",
          "imageHeight": "",
          "imageSecureUrl": "",
          "siteName": "GitHub"
        },
        "twitter": {
          "url": "",
          "title": "ariefsn - Overview",
          "description": "ariefsn has 43 repositories available. Follow their code on GitHub.",
          "image": "",
          "imageAlt": "",
          "siteName": "@github",
          "card": "summary",
          "creator": "",
          "creatorId": "",
          "appIphone": {
            "id": "",
            "name": "",
            "url": ""
          },
          "appIpad": {
            "id": "",
            "name": "",
            "url": ""
          },
          "appGooglePlay": {
            "id": "",
            "name": "",
            "url": ""
          }
        },
        "body": "",
        "viewPort": "width=device-width",
        "themeColor": "#1e2327",
        "colorScheme": "light dark",
        "appleApp": {
          "appleMobileWebAppCapable": "",
          "appleMobileWebAppStatusBarStyle": "",
          "appleMobileWebAppTitle": "",
          "appleMobileWebAppIcon": ""
        },
        "msApp": {
          "tileImage": "",
          "tileColor": ""
        },
        "hostName": "github.com"
      }
    ```
