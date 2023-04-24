// Code generated. DO NOT EDIT.
// To generate use: go generate
//go:generate go run ../../scripts/gen.go ../../api-spec/classic-game-data-apis.json api.go classic
package classic

import (
  "encoding/json"
  "fmt"
  "net/http"
  "net/url"
  "strings"
)

const (
  BlizzardAPIUrl = "https://%s.api.blizzard.com"
  BlizzardTokenUrl = "https://oauth.battle.net/token"
)

type AccessToken struct {
  AccessToken string `json:"access_token"`
  TokenType string `json:"token_type"`
  ExpiresIn int `json:"expires_in"`
  Scope string `json:"scope"`
}

type Client struct {
  clientId string
  clientSecret string
  accessToken string
  region string
  url string
  httpClient *http.Client
}

func stringWithDefault(s string, d string) string {
  if s != "" {
    return d
  }
  return s
}

func (c *Client) getAccessToken() error {
  params := url.Values{}
  params.Add("grant_type", `client_credentials`)
  body := strings.NewReader(params.Encode())

  req, err := http.NewRequest("POST", BlizzardTokenUrl, body)
  if err != nil {
    return err
  }
  req.SetBasicAuth(c.clientId, c.clientSecret)
  req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

  resp, err := c.httpClient.Do(req)
  if err != nil {
    return err
  }
  defer resp.Body.Close()

  token := &AccessToken{}
  if err := json.NewDecoder(resp.Body).Decode(token); err != nil {
    return err
  }
  c.accessToken = token.AccessToken
  return nil
}

func NewClient(clientId string, clientSecret string, region string, httpClient *http.Client) (*Client, error) {
  if httpClient == nil {
    httpClient = http.DefaultClient
  }
  client := &Client{
    clientId: clientId,
    clientSecret: clientSecret,
    httpClient: httpClient,
    region: region,
    url: fmt.Sprintf(BlizzardAPIUrl, region),
  }
  if err := client.getAccessToken(); err != nil {
    return nil, err
  }
  return client, nil
}

// Resource: Auction House API
//
// Name: Auction House Index
//
// Description: Returns an index of auction houses for a connected realm. See the Connected Realm API for information about retrieving a list of connected realm IDs.
//
// Path: /data/wow/connected-realm/{connectedRealmId}/auctions/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   connectedRealmId: The ID of the connected realm. (int) (required) (default: 4372)
//   namespace: The namespace to use to locate this document. (string) (required) (default: dynamic-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) AuctionHouseIndex(connectedRealmId int, namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/connected-realm/%v/auctions/index", connectedRealmId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "dynamic-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Auction House API
//
// Name: Auctions
//
// Description: Returns all active auctions for a specific auction house on a connected realm. See the Connected Realm API for information about retrieving a list of connected realm IDs. Auction house data updates at a set interval. The value was initially set at 1 hour; however, it might change over time without notice. Depending on the number of active auctions on the specified connected realm, the response from this endpoint may be rather large, sometimes exceeding 10 MB.
//
// Path: /data/wow/connected-realm/{connectedRealmId}/auctions/{auctionHouseId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   connectedRealmId: The ID of the connected realm. (int) (required) (default: 4372)
//   auctionHouseId: The ID of the auction house. (int) (required) (default: 2)
//   namespace: The namespace to use to locate this document. (string) (required) (default: dynamic-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) Auctions(connectedRealmId int, auctionHouseId int, namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/connected-realm/%v/auctions/%v", connectedRealmId, auctionHouseId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "dynamic-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Connected Realm API
//
// Name: Connected Realms Index
//
// Description: Returns an index of connected realms.
//
// Path: /data/wow/connected-realm/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   namespace: The namespace to use to locate this document. (string) (required) (default: dynamic-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) ConnectedRealmsIndex(namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/connected-realm/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "dynamic-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Connected Realm API
//
// Name: Connected Realm
//
// Description: Returns a connected realm by ID.
//
// Path: /data/wow/connected-realm/{connectedRealmId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   connectedRealmId: The ID of the connected realm. (int) (required) (default: 4388)
//   namespace: The namespace to use to locate this document. (string) (required) (default: dynamic-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) ConnectedRealm(connectedRealmId int, namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/connected-realm/%v", connectedRealmId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "dynamic-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Connected Realm API
//
// Name: Connected Realms Search
//
// Description: Performs a search of connected realms. The fields below are provided for example. For more detail see the Search Guide.
//
// Path: /data/wow/search/connected-realm
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   namespace: The namespace to use to locate this document. (string) (required) (default: dynamic-classic-us)
//   orderby: The field to sort the result set by. (string)  (default: id)
//   page: The result page number. (int)  (default: 1)
//   search: Search query string. For more detail see the Search Guide: https://develop.battle.net/documentation/world-of-warcraft/guides/search
func (c *Client) ConnectedRealmsSearch(namespace string, orderby string, page int, search string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/search/connected-realm", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "dynamic-classic-us"))
  q.Add("orderby", stringWithDefault(fmt.Sprint(orderby), "id"))
  q.Add("_page", stringWithDefault(fmt.Sprint(page), "1"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = fmt.Sprintf("%s&%s", q.Encode(), search)
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Creature API
//
// Name: Creature Families Index
//
// Description: Returns an index of creature families.
//
// Path: /data/wow/creature-family/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   namespace: The namespace to use to locate this document. (string) (required) (default: static-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CreatureFamiliesIndex(namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/creature-family/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "static-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Creature API
//
// Name: Creature Family
//
// Description: Returns a creature family by ID.
//
// Path: /data/wow/creature-family/{creatureFamilyId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   creatureFamilyId: The ID of the creature family. (int) (required) (default: 1)
//   namespace: The namespace to use to locate this document. (string) (required) (default: static-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CreatureFamily(creatureFamilyId int, namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/creature-family/%v", creatureFamilyId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "static-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Creature API
//
// Name: Creature Types Index
//
// Description: Returns an index of creature types.
//
// Path: /data/wow/creature-type/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   namespace: The namespace to use to locate this document. (string) (required) (default: static-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CreatureTypesIndex(namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/creature-type/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "static-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Creature API
//
// Name: Creature Type
//
// Description: Returns a creature type by ID.
//
// Path: /data/wow/creature-type/{creatureTypeId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   creatureTypeId: The ID of the creature type. (int) (required) (default: 1)
//   namespace: The namespace to use to locate this document. (string) (required) (default: static-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CreatureType(creatureTypeId int, namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/creature-type/%v", creatureTypeId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "static-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Creature API
//
// Name: Creature
//
// Description: Returns a creature by ID.
//
// Path: /data/wow/creature/{creatureId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   creatureId: The ID of the creature. (int) (required) (default: 30)
//   namespace: The namespace to use to locate this document. (string) (required) (default: static-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) Creature(creatureId int, namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/creature/%v", creatureId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "static-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Creature API
//
// Name: Creature Search
//
// Description: Performs a search of creatures. The fields below are provided for example. For more detail see the Search Guide.
//
// Path: /data/wow/search/creature
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   namespace: The namespace to use to locate this document. (string) (required) (default: static-us)
//   orderby: The field to sort the result set by. (string)  (default: id)
//   page: The result page number. (int)  (default: 1)
//   search: Search query string. For more detail see the Search Guide: https://develop.battle.net/documentation/world-of-warcraft/guides/search
func (c *Client) CreatureSearch(namespace string, orderby string, page int, search string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/search/creature", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "static-us"))
  q.Add("orderby", stringWithDefault(fmt.Sprint(orderby), "id"))
  q.Add("_page", stringWithDefault(fmt.Sprint(page), "1"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = fmt.Sprintf("%s&%s", q.Encode(), search)
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Creature API
//
// Name: Creature Display Media
//
// Description: Returns media for a creature display by ID.
//
// Path: /data/wow/media/creature-display/{creatureDisplayId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   creatureDisplayId: The ID of the creature display. (int) (required) (default: 30)
//   namespace: The namespace to use to locate this document. (string) (required) (default: static-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CreatureDisplayMedia(creatureDisplayId int, namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/media/creature-display/%v", creatureDisplayId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "static-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Creature API
//
// Name: Creature Family Media
//
// Description: Returns media for a creature family by ID.
//
// Path: /data/wow/media/creature-family/{creatureFamilyId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   creatureFamilyId: The ID of the creature family. (int) (required) (default: 1)
//   namespace: The namespace to use to locate this document. (string) (required) (default: static-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CreatureFamilyMedia(creatureFamilyId int, namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/media/creature-family/%v", creatureFamilyId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "static-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Guild Crest API
//
// Name: Guild Crest Components Index
//
// Description: Returns an index of guild crest media.
//
// Path: /data/wow/guild-crest/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   namespace: The namespace to use to locate this document. (string) (required) (default: static-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) GuildCrestComponentsIndex(namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/guild-crest/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "static-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Guild Crest API
//
// Name: Guild Crest Border Media
//
// Description: Returns media for a guild crest border by ID.
//
// Path: /data/wow/media/guild-crest/border/{borderId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   borderId: The ID of the guild crest border. (int) (required) 
//   namespace: The namespace to use to locate this document. (string) (required) (default: static-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) GuildCrestBorderMedia(borderId int, namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/media/guild-crest/border/%v", borderId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "static-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Guild Crest API
//
// Name: Guild Crest Emblem Media
//
// Description: Returns media for a guild crest emblem by ID.
//
// Path: /data/wow/media/guild-crest/emblem/{emblemId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   emblemId: The ID of the guild crest emblem. (int) (required) 
//   namespace: The namespace to use to locate this document. (string) (required) (default: static-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) GuildCrestEmblemMedia(emblemId int, namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/media/guild-crest/emblem/%v", emblemId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "static-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Item API
//
// Name: Item Classes Index
//
// Description: Returns an index of item classes.
//
// Path: /data/wow/item-class/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   namespace: The namespace to use to locate this document. (string) (required) (default: static-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) ItemClassesIndex(namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/item-class/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "static-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Item API
//
// Name: Item Class
//
// Description: Returns an item class by ID.
//
// Path: /data/wow/item-class/{itemClassId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   itemClassId: The ID of the item class. (string) (required) (default: 0)
//   namespace: The namespace to use to locate this document. (string) (required) (default: static-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) ItemClass(itemClassId string, namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/item-class/%v", itemClassId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "static-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Item API
//
// Name: Item Subclass
//
// Description: Returns an item subclass by ID.
//
// Path: /data/wow/item-class/{itemClassId}/item-subclass/{itemSubclassId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   itemClassId: The ID of the item class. (string) (required) (default: 0)
//   itemSubclassId: The ID of the item subclass. (string) (required) (default: 0)
//   namespace: The namespace to use to locate this document. (string) (required) (default: static-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) ItemSubclass(itemClassId string, itemSubclassId string, namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/item-class/%v/item-subclass/%v", itemClassId, itemSubclassId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "static-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Item API
//
// Name: Item
//
// Description: Returns an item by ID.
//
// Path: /data/wow/item/{itemId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   itemId: The ID of the item. (string) (required) (default: 19019)
//   namespace: The namespace to use to locate this document. (string) (required) (default: static-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) Item(itemId string, namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/item/%v", itemId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "static-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Item API
//
// Name: Item Media
//
// Description: Returns media for an item by ID.
//
// Path: /data/wow/media/item/{itemId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   itemId: The ID of the item. (int) (required) (default: 19019)
//   namespace: The namespace to use to locate this document. (string) (required) (default: static-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) ItemMedia(itemId int, namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/media/item/%v", itemId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "static-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Item API
//
// Name: Item Search
//
// Description: Performs a search of items. The fields below are provided for example. For more detail see the Search Guide.
//
// Path: /data/wow/search/item
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   namespace: The namespace to use to locate this document. (string) (required) (default: static-us)
//   orderby: The field to sort the result set by. (string)  (default: id)
//   page: The result page number. (int)  (default: 1)
//   search: Search query string. For more detail see the Search Guide: https://develop.battle.net/documentation/world-of-warcraft/guides/search
func (c *Client) ItemSearch(namespace string, orderby string, page int, search string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/search/item", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "static-us"))
  q.Add("orderby", stringWithDefault(fmt.Sprint(orderby), "id"))
  q.Add("_page", stringWithDefault(fmt.Sprint(page), "1"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = fmt.Sprintf("%s&%s", q.Encode(), search)
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Media Search API
//
// Name: Media Search
//
// Description: Performs a search of all types of media documents. The fields below are provided for example. For more detail see the Search Guide.
//
// Path: /data/wow/search/media
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   namespace: The namespace to use to locate this document. (string) (required) (default: static-us)
//   orderby: The field to sort the result set by. (string)  (default: id)
//   page: The result page number. (int)  (default: 1)
//   search: Search query string. For more detail see the Search Guide: https://develop.battle.net/documentation/world-of-warcraft/guides/search
func (c *Client) MediaSearch(namespace string, orderby string, page int, search string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/search/media", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "static-us"))
  q.Add("orderby", stringWithDefault(fmt.Sprint(orderby), "id"))
  q.Add("_page", stringWithDefault(fmt.Sprint(page), "1"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = fmt.Sprintf("%s&%s", q.Encode(), search)
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Playable Class API
//
// Name: Playable Classes Index
//
// Description: Returns an index of playable classes.
//
// Path: /data/wow/playable-class/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   namespace: The namespace to use to locate this document. (string) (required) (default: static-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) PlayableClassesIndex(namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/playable-class/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "static-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Playable Class API
//
// Name: Playable Class
//
// Description: Returns a playable class by ID.
//
// Path: /data/wow/playable-class/{classId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   classId: The ID of the playable class. (int) (required) (default: 7)
//   namespace: The namespace to use to locate this document. (string) (required) (default: static-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) PlayableClass(classId int, namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/playable-class/%v", classId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "static-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Playable Class API
//
// Name: Playable Class Media
//
// Description: Returns media for a playable class by ID.
//
// Path: /data/wow/media/playable-class/{playableClassId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   playableClassId: The ID of the playable class. (int) (required) (default: 7)
//   namespace: The namespace to use to locate this document. (string) (required) (default: static-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) PlayableClassMedia(playableClassId int, namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/media/playable-class/%v", playableClassId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "static-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Playable Race API
//
// Name: Playable Races Index
//
// Description: Returns an index of playable races.
//
// Path: /data/wow/playable-race/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   namespace: The namespace to use to locate this document. (string) (required) (default: static-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) PlayableRacesIndex(namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/playable-race/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "static-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Playable Race API
//
// Name: Playable Race
//
// Description: Returns a playable race by ID.
//
// Path: /data/wow/playable-race/{playableRaceId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   playableRaceId: The ID of the playable race. (int) (required) (default: 2)
//   namespace: The namespace to use to locate this document. (string) (required) (default: static-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) PlayableRace(playableRaceId int, namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/playable-race/%v", playableRaceId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "static-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Power Type API
//
// Name: Power Types Index
//
// Description: Returns an index of power types.
//
// Path: /data/wow/power-type/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   namespace: The namespace to use to locate this document. (string) (required) (default: static-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) PowerTypesIndex(namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/power-type/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "static-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Power Type API
//
// Name: Power Type
//
// Description: Returns a power type by ID.
//
// Path: /data/wow/power-type/{powerTypeId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   powerTypeId: The ID of the power type. (int) (required) 
//   namespace: The namespace to use to locate this document. (string) (required) (default: static-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) PowerType(powerTypeId int, namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/power-type/%v", powerTypeId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "static-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: PvP Season API
//
// Name: PvP Seasons Index
//
// Description: Returns an index of PvP seasons.
//
// Path: /data/wow/pvp-season/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   namespace: The namespace to use to locate this document. (string) (required) (default: dynamic-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) PvPSeasonsIndex(namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/pvp-season/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "dynamic-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: PvP Season API
//
// Name: PvP Season
//
// Description: Returns a PvP season by ID.
//
// Path: /data/wow/pvp-season/{pvpSeasonId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   pvpSeasonId: The ID of the PvP season. (int) (required) (default: 1)
//   namespace: The namespace to use to locate this document. (string) (required) (default: dynamic-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) PvPSeason(pvpSeasonId int, namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/pvp-season/%v", pvpSeasonId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "dynamic-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: PvP Season API
//
// Name: PvP Region Index
//
// Description: Returns an index of PvP Regions.
//
// Path: /data/wow/pvp-region/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   namespace: The namespace to use to locate this document. (string) (required) (default: dynamic-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) PvPRegionIndex(namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/pvp-region/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "dynamic-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: PvP Season API
//
// Name: PvP Regional Season Index
//
// Description: Returns an index of PvP Seasons in a PvP region.
//
// Path: /data/wow/pvp-region/{pvpRegionId}/pvp-season/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   pvpRegionId: The ID of the PvP region. (int) (required) (default: 1)
//   namespace: The namespace to use to locate this document. (string) (required) (default: dynamic-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) PvPRegionalSeasonIndex(pvpRegionId int, namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/pvp-region/%v/pvp-season/index", pvpRegionId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "dynamic-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: PvP Season API
//
// Name: PvP Regional Season
//
// Description: Returns a PvP season by region ID and season ID.
//
// Path: /data/wow/pvp-region/{pvpRegionId}/pvp-season/{pvpSeasonId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   pvpRegionId: The ID of the PvP region. (int) (required) (default: 1)
//   pvpSeasonId: The ID of the PvP season. (int) (required) (default: 1)
//   namespace: The namespace to use to locate this document. (string) (required) (default: dynamic-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) PvPRegionalSeason(pvpRegionId int, pvpSeasonId int, namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/pvp-region/%v/pvp-season/%v", pvpRegionId, pvpSeasonId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "dynamic-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: PvP Season API
//
// Name: PvP Leaderboards Index
//
// Description: Returns an index of PvP leaderboards for a PvP season in a given PvP region.
//
// Path: /data/wow/pvp-region/{pvpRegionId}/pvp-season/{pvpSeasonId}/pvp-leaderboard/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   pvpRegionId: The ID of the PvP region. (int) (required) (default: 1)
//   pvpSeasonId: The ID of the PvP season. (int) (required) (default: 1)
//   namespace: The namespace to use to locate this document. (string) (required) (default: dynamic-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) PvPLeaderboardsIndex(pvpRegionId int, pvpSeasonId int, namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/pvp-region/%v/pvp-season/%v/pvp-leaderboard/index", pvpRegionId, pvpSeasonId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "dynamic-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: PvP Season API
//
// Name: PvP Leaderboard
//
// Description: Returns the PvP leaderboard of a specific PvP bracket for a PvP season in a given PvP region.
//
// Path: /data/wow/pvp-region/{pvpRegionId}/pvp-season/{pvpSeasonId}/pvp-leaderboard/{pvpBracket}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   pvpRegionId: The ID of the PvP region. (int) (required) (default: 1)
//   pvpSeasonId: The ID of the PvP season. (int) (required) (default: 1)
//   pvpBracket: The PvP bracket type. (string) (required) (default: 3v3)
//   namespace: The namespace to use to locate this document. (string) (required) (default: dynamic-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) PvPLeaderboard(pvpRegionId int, pvpSeasonId int, pvpBracket string, namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/pvp-region/%v/pvp-season/%v/pvp-leaderboard/%v", pvpRegionId, pvpSeasonId, pvpBracket))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "dynamic-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: PvP Season API
//
// Name: PvP Rewards Index
//
// Description: Returns an index of PvP rewards for a PvP season in a given PvP region.
//
// Path: /data/wow/pvp-region/{pvpRegionId}/pvp-season/{pvpSeasonId}/pvp-reward/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   pvpRegionId: The ID of the PvP region. (int) (required) (default: 1)
//   pvpSeasonId: The ID of the PvP season. (int) (required) (default: 1)
//   namespace: The namespace to use to locate this document. (string) (required) (default: dynamic-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) PvPRewardsIndex(pvpRegionId int, pvpSeasonId int, namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/pvp-region/%v/pvp-season/%v/pvp-reward/index", pvpRegionId, pvpSeasonId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "dynamic-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Realm API
//
// Name: Realms Index
//
// Description: Returns an index of realms.
//
// Path: /data/wow/realm/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   namespace: The namespace to use to locate this document. (string) (required) (default: dynamic-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) RealmsIndex(namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/realm/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "dynamic-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Realm API
//
// Name: Realm
//
// Description: Returns a single realm by slug or ID.
//
// Path: /data/wow/realm/{realmSlug}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   realmSlug: The slug of the realm. (string) (required) (default: westfall)
//   namespace: The namespace to use to locate this document. (string) (required) (default: dynamic-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) Realm(realmSlug string, namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/realm/%v", realmSlug))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "dynamic-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Realm API
//
// Name: Realm Search
//
// Description: Performs a search of realms. The fields below are examples only. For more detail see the Search Guide.
//
// Path: /data/wow/search/realm
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   namespace: The namespace to use to locate this document. (string) (required) (default: dynamic-classic-us)
//   orderby: The field to sort the result set by. (string)  (default: id)
//   page: The result page number. (int)  (default: 1)
//   search: Search query string. For more detail see the Search Guide: https://develop.battle.net/documentation/world-of-warcraft/guides/search
func (c *Client) RealmSearch(namespace string, orderby string, page int, search string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/search/realm", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "dynamic-classic-us"))
  q.Add("orderby", stringWithDefault(fmt.Sprint(orderby), "id"))
  q.Add("_page", stringWithDefault(fmt.Sprint(page), "1"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = fmt.Sprintf("%s&%s", q.Encode(), search)
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Region API
//
// Name: Regions Index
//
// Description: Returns an index of regions.
//
// Path: /data/wow/region/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   namespace: The namespace to use to locate this document. (string) (required) (default: dynamic-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) RegionsIndex(namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/region/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "dynamic-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Region API
//
// Name: Region
//
// Description: Returns a region by ID.
//
// Path: /data/wow/region/{regionId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   regionId: The ID of the region. (int) (required) (default: 41)
//   namespace: The namespace to use to locate this document. (string) (required) (default: dynamic-classic-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) Region(regionId int, namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/region/%v", regionId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "dynamic-classic-us"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: WoW Token API
//
// Name: WoW Token Index (CN)
//
// Description: Returns the WoW Token index.
//
// Path: /data/wow/token/index
//
// HttpMethod: GET
//
// CnRegion: true
//
// Parameters: 
//   namespace: The namespace to use to locate this document. (string) (required) (default: dynamic-classic-cn)
//   locale: The locale to reflect in localized data. (string)  (default: zh_CN)
func (c *Client) WoWTokenIndexCN(namespace string, locale string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/token/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "dynamic-classic-cn"))
  q.Add("locale", stringWithDefault(fmt.Sprint(locale), "zh_CN"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}
