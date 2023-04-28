// Code generated. DO NOT EDIT.
// To generate use: go generate
//go:generate go run ../../scripts/gen.go ../../api-spec/retail-game-data-apis.json api.go retail
package retail

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
  locale string
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

func NewClient(clientId string, clientSecret string, region string, locale string, httpClient *http.Client) (*Client, error) {
  if httpClient == nil {
    httpClient = http.DefaultClient
  }
  client := &Client{
    clientId: clientId,
    clientSecret: clientSecret,
    httpClient: httpClient,
    region: region,
    locale: locale,
    url: fmt.Sprintf(BlizzardAPIUrl, region),
  }
  if err := client.getAccessToken(); err != nil {
    return nil, err
  }
  return client, nil
}

// Resource: Achievement API
//
// Name: Achievement Categories Index
//
// Description: Returns an index of achievement categories.
//
// Path: /data/wow/achievement-category/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) AchievementCategoriesIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/achievement-category/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Achievement API
//
// Name: Achievement Category
//
// Description: Returns an achievement category by ID.
//
// Path: /data/wow/achievement-category/{achievementCategoryId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   achievementCategoryId: The ID of the achievement category. (int) (required) (default: 81)
func (c *Client) AchievementCategory(achievementCategoryId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/achievement-category/%v", achievementCategoryId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Achievement API
//
// Name: Achievements Index
//
// Description: Returns an index of achievements.
//
// Path: /data/wow/achievement/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) AchievementsIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/achievement/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Achievement API
//
// Name: Achievement
//
// Description: Returns an achievement by ID.
//
// Path: /data/wow/achievement/{achievementId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   achievementId: The ID of the achievement. (int) (required) (default: 6)
func (c *Client) Achievement(achievementId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/achievement/%v", achievementId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Achievement API
//
// Name: Achievement Media
//
// Description: Returns media for an achievement by ID.
//
// Path: /data/wow/media/achievement/{achievementId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   achievementId: The ID of the achievement. (int) (required) (default: 6)
func (c *Client) AchievementMedia(achievementId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/media/achievement/%v", achievementId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Auction House API
//
// Name: Auctions
//
// Description: Returns all active auctions for a connected realm. See the Connected Realm API for information about retrieving a list of connected realm IDs. Auction house data updates at a set interval. The value was initially set at 1 hour; however, it might change over time without notice. Depending on the number of active auctions on the specified connected realm, the response from this endpoint may be rather large, sometimes exceeding 10 MB.
//
// Path: /data/wow/connected-realm/{connectedRealmId}/auctions
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   connectedRealmId: The ID of the connected realm. (int) (required) (default: 1146)
func (c *Client) Auctions(connectedRealmId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/connected-realm/%v/auctions", connectedRealmId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("dynamic-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Auction House API
//
// Name: Commodities
//
// Description: Returns all active auctions for commodity items for the entire game region. Auction house data updates at a set interval. The value was initially set at 1 hour; however, it might change over time without notice. Depending on the number of active auctions on the specified connected realm, the response from this endpoint may be rather large, sometimes exceeding 10 MB.
//
// Path: /data/wow/auctions/commodities
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) Commodities() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/auctions/commodities", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("dynamic-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Azerite Essence API
//
// Name: Azerite Essences Index
//
// Description: Returns an index of azerite essences.
//
// Path: /data/wow/azerite-essence/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) AzeriteEssencesIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/azerite-essence/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Azerite Essence API
//
// Name: Azerite Essence
//
// Description: Returns an azerite essence by ID.
//
// Path: /data/wow/azerite-essence/{azeriteEssenceId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   azeriteEssenceId: The ID of the azerite essence. (string) (required) (default: 2)
func (c *Client) AzeriteEssence(azeriteEssenceId string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/azerite-essence/%v", azeriteEssenceId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Azerite Essence API
//
// Name: Azerite Essence Search
//
// Description: Performs a search of azerite essences. The fields below are provided for example. For more detail see the Search Guide.
//
// Path: /data/wow/search/azerite-essence
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   orderby: The field to sort the result set by. (string)  (default: name)
//   page: The result page number. (int)  (default: 1)
//   search: Search query string. For more detail see the Search Guide: https://develop.battle.net/documentation/world-of-warcraft/guides/search
func (c *Client) AzeriteEssenceSearch(orderby string, page int, search string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/search/azerite-essence", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("orderby", stringWithDefault(fmt.Sprint(orderby), "name"))
  q.Add("_page", stringWithDefault(fmt.Sprint(page), "1"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = fmt.Sprintf("%s&%s", q.Encode(), search)
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Azerite Essence API
//
// Name: Azerite Essence Media
//
// Description: Returns media for an azerite essence by ID.
//
// Path: /data/wow/media/azerite-essence/{azeriteEssenceId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   azeriteEssenceId: The ID of the azerite essence. (int) (required) (default: 2)
func (c *Client) AzeriteEssenceMedia(azeriteEssenceId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/media/azerite-essence/%v", azeriteEssenceId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
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
func (c *Client) ConnectedRealmsIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/connected-realm/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("dynamic-%s", c.region))
  q.Add("locale", c.locale)
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
//   connectedRealmId: The ID of the connected realm. (int) (required) (default: 11)
func (c *Client) ConnectedRealm(connectedRealmId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/connected-realm/%v", connectedRealmId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("dynamic-%s", c.region))
  q.Add("locale", c.locale)
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
//   orderby: The field to sort the result set by. (string)  (default: id)
//   page: The result page number. (int)  (default: 1)
//   search: Search query string. For more detail see the Search Guide: https://develop.battle.net/documentation/world-of-warcraft/guides/search
func (c *Client) ConnectedRealmsSearch(orderby string, page int, search string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/search/connected-realm", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("dynamic-%s", c.region))
  q.Add("orderby", stringWithDefault(fmt.Sprint(orderby), "id"))
  q.Add("_page", stringWithDefault(fmt.Sprint(page), "1"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = fmt.Sprintf("%s&%s", q.Encode(), search)
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Covenant API
//
// Name: Covenant Index
//
// Description: Returns an index of covenants.
//
// Path: /data/wow/covenant/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) CovenantIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/covenant/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Covenant API
//
// Name: Covenant
//
// Description: Returns a covenant by ID.
//
// Path: /data/wow/covenant/{covenantId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   covenantId: The ID of the covenant. (int) (required) (default: 1)
func (c *Client) Covenant(covenantId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/covenant/%v", covenantId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Covenant API
//
// Name: Covenant Media
//
// Description: Returns media for a covenant by ID.
//
// Path: /data/wow/media/covenant/{covenantId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   covenantId: The ID of the covenant. (int) (required) (default: 1)
func (c *Client) CovenantMedia(covenantId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/media/covenant/%v", covenantId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Covenant API
//
// Name: Soulbind Index
//
// Description: Returns an index of soulbinds.
//
// Path: /data/wow/covenant/soulbind/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) SoulbindIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/covenant/soulbind/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Covenant API
//
// Name: Soulbind
//
// Description: Returns a soulbind by ID.
//
// Path: /data/wow/covenant/soulbind/{soulbindId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   soulbindId: The ID of the soulbind. (int) (required) (default: 1)
func (c *Client) Soulbind(soulbindId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/covenant/soulbind/%v", soulbindId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Covenant API
//
// Name: Conduit Index
//
// Description: Returns an index of conduits.
//
// Path: /data/wow/covenant/conduit/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) ConduitIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/covenant/conduit/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Covenant API
//
// Name: Conduit
//
// Description: Returns a conduit by ID.
//
// Path: /data/wow/covenant/conduit/{conduitId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   conduitId: The ID of the conduit. (int) (required) (default: 1)
func (c *Client) Conduit(conduitId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/covenant/conduit/%v", conduitId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
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
func (c *Client) CreatureFamiliesIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/creature-family/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
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
func (c *Client) CreatureFamily(creatureFamilyId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/creature-family/%v", creatureFamilyId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
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
func (c *Client) CreatureTypesIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/creature-type/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
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
func (c *Client) CreatureType(creatureTypeId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/creature-type/%v", creatureTypeId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
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
//   creatureId: The ID of the creature. (int) (required) (default: 42722)
func (c *Client) Creature(creatureId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/creature/%v", creatureId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
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
//   orderby: The field to sort the result set by. (string)  (default: id)
//   page: The result page number. (int)  (default: 1)
//   search: Search query string. For more detail see the Search Guide: https://develop.battle.net/documentation/world-of-warcraft/guides/search
func (c *Client) CreatureSearch(orderby string, page int, search string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/search/creature", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
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
//   creatureDisplayId: The ID of the creature display. (int) (required) (default: 30221)
func (c *Client) CreatureDisplayMedia(creatureDisplayId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/media/creature-display/%v", creatureDisplayId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
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
func (c *Client) CreatureFamilyMedia(creatureFamilyId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/media/creature-family/%v", creatureFamilyId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
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
func (c *Client) GuildCrestComponentsIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/guild-crest/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
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
func (c *Client) GuildCrestBorderMedia(borderId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/media/guild-crest/border/%v", borderId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
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
func (c *Client) GuildCrestEmblemMedia(emblemId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/media/guild-crest/emblem/%v", emblemId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Heirloom API
//
// Name: Heirloom Index
//
// Description: Returns an index of heirlooms.
//
// Path: /data/wow/heirloom/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) HeirloomIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/heirloom/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Heirloom API
//
// Name: Heirloom
//
// Description: Returns an heirloom by id.
//
// Path: /data/wow/heirloom/{heirloomId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   heirloomId: The ID of the heirloom. (int) (required) (default: 1)
func (c *Client) Heirloom(heirloomId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/heirloom/%v", heirloomId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
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
func (c *Client) ItemClassesIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/item-class/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
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
func (c *Client) ItemClass(itemClassId string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/item-class/%v", itemClassId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Item API
//
// Name: Item Sets Index
//
// Description: Returns an index of item sets.
//
// Path: /data/wow/item-set/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) ItemSetsIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/item-set/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Item API
//
// Name: Item Set
//
// Description: Returns an item set by ID.
//
// Path: /data/wow/item-set/{itemSetId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   itemSetId: The ID of the item set. (int) (required) (default: 1)
func (c *Client) ItemSet(itemSetId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/item-set/%v", itemSetId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
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
func (c *Client) ItemSubclass(itemClassId string, itemSubclassId string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/item-class/%v/item-subclass/%v", itemClassId, itemSubclassId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
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
func (c *Client) Item(itemId string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/item/%v", itemId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
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
func (c *Client) ItemMedia(itemId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/media/item/%v", itemId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
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
//   orderby: The field to sort the result set by. (string)  (default: id)
//   page: The result page number. (int)  (default: 1)
//   search: Search query string. For more detail see the Search Guide: https://develop.battle.net/documentation/world-of-warcraft/guides/search
func (c *Client) ItemSearch(orderby string, page int, search string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/search/item", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("orderby", stringWithDefault(fmt.Sprint(orderby), "id"))
  q.Add("_page", stringWithDefault(fmt.Sprint(page), "1"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = fmt.Sprintf("%s&%s", q.Encode(), search)
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Journal API
//
// Name: Journal Expansions Index
//
// Description: Returns an index of journal expansions.
//
// Path: /data/wow/journal-expansion/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) JournalExpansionsIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/journal-expansion/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Journal API
//
// Name: Journal Expansion
//
// Description: Returns a journal expansion by ID.
//
// Path: /data/wow/journal-expansion/{journalExpansionId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   journalExpansionId: The ID of the journal expansion. (int) (required) (default: 68)
func (c *Client) JournalExpansion(journalExpansionId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/journal-expansion/%v", journalExpansionId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Journal API
//
// Name: Journal Encounters Index
//
// Description: Returns an index of journal encounters.
//
// Path: /data/wow/journal-encounter/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) JournalEncountersIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/journal-encounter/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Journal API
//
// Name: Journal Encounter
//
// Description: Returns a journal encounter by ID.
//
// Path: /data/wow/journal-encounter/{journalEncounterId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   journalEncounterId: The ID of the journal encounter. (int) (required) (default: 89)
func (c *Client) JournalEncounter(journalEncounterId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/journal-encounter/%v", journalEncounterId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Journal API
//
// Name: Journal Encounter Search
//
// Description: Performs a search of journal encounters. The fields below are provided for example. For more detail see the Search Guide.
//
// Path: /data/wow/search/journal-encounter
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   orderby: The field to sort the result set by. (string)  (default: id)
//   page: The result page number. (int)  (default: 1)
//   search: Search query string. For more detail see the Search Guide: https://develop.battle.net/documentation/world-of-warcraft/guides/search
func (c *Client) JournalEncounterSearch(orderby string, page int, search string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/search/journal-encounter", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("orderby", stringWithDefault(fmt.Sprint(orderby), "id"))
  q.Add("_page", stringWithDefault(fmt.Sprint(page), "1"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = fmt.Sprintf("%s&%s", q.Encode(), search)
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Journal API
//
// Name: Journal Instances Index
//
// Description: Returns an index of journal instances.
//
// Path: /data/wow/journal-instance/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) JournalInstancesIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/journal-instance/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Journal API
//
// Name: Journal Instance
//
// Description: Returns a journal instance.
//
// Path: /data/wow/journal-instance/{journalInstanceId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   journalInstanceId: The ID of the journal instance. (int) (required) (default: 63)
func (c *Client) JournalInstance(journalInstanceId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/journal-instance/%v", journalInstanceId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Journal API
//
// Name: Journal Instance Media
//
// Description: Returns media for a journal instance by ID.
//
// Path: /data/wow/media/journal-instance/{journalInstanceId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   journalInstanceId: The ID of the journal instance. (int) (required) (default: 63)
func (c *Client) JournalInstanceMedia(journalInstanceId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/media/journal-instance/%v", journalInstanceId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
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
//   orderby: The field to sort the result set by. (string)  (default: id)
//   page: The result page number. (int)  (default: 1)
//   search: Search query string. For more detail see the Search Guide: https://develop.battle.net/documentation/world-of-warcraft/guides/search
func (c *Client) MediaSearch(orderby string, page int, search string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/search/media", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("orderby", stringWithDefault(fmt.Sprint(orderby), "id"))
  q.Add("_page", stringWithDefault(fmt.Sprint(page), "1"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = fmt.Sprintf("%s&%s", q.Encode(), search)
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Modified Crafting API
//
// Name: Modified Crafting Index
//
// Description: Returns the parent index for Modified Crafting.
//
// Path: /data/wow/modified-crafting/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) ModifiedCraftingIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/modified-crafting/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Modified Crafting API
//
// Name: Modified Crafting Category Index
//
// Description: Returns the index of Modified Crafting categories.
//
// Path: /data/wow/modified-crafting/category/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) ModifiedCraftingCategoryIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/modified-crafting/category/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Modified Crafting API
//
// Name: Modified Crafting Category
//
// Description: Returns a Modified Crafting category by ID.
//
// Path: /data/wow/modified-crafting/category/{categoryId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   categoryId: The ID of the Modified Crafting category. (int) (required) (default: 1)
func (c *Client) ModifiedCraftingCategory(categoryId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/modified-crafting/category/%v", categoryId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Modified Crafting API
//
// Name: Modified Crafting Reagent Slot Type Index
//
// Description: Returns the index of Modified Crafting reagent slot types.
//
// Path: /data/wow/modified-crafting/reagent-slot-type/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) ModifiedCraftingReagentSlotTypeIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/modified-crafting/reagent-slot-type/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Modified Crafting API
//
// Name: Modified Crafting Reagent Slot Type
//
// Description: Returns a Modified Crafting reagent slot type by ID.
//
// Path: /data/wow/modified-crafting/reagent-slot-type/{slotTypeId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   slotTypeId: The ID of the Modified Crafting reagent slot type. (int) (required) (default: 16)
func (c *Client) ModifiedCraftingReagentSlotType(slotTypeId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/modified-crafting/reagent-slot-type/%v", slotTypeId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Mount API
//
// Name: Mounts Index
//
// Description: Returns an index of mounts.
//
// Path: /data/wow/mount/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) MountsIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/mount/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Mount API
//
// Name: Mount
//
// Description: Returns a mount by ID.
//
// Path: /data/wow/mount/{mountId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   mountId: The ID of the mount. (int) (required) (default: 6)
func (c *Client) Mount(mountId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/mount/%v", mountId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Mount API
//
// Name: Mount Search
//
// Description: Performs a search of mounts. The fields below are provided for example. For more detail see the Search Guide.
//
// Path: /data/wow/search/mount
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   orderby: The field to sort the result set by. (string)  (default: id)
//   page: The result page number. (int)  (default: 1)
//   search: Search query string. For more detail see the Search Guide: https://develop.battle.net/documentation/world-of-warcraft/guides/search
func (c *Client) MountSearch(orderby string, page int, search string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/search/mount", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("orderby", stringWithDefault(fmt.Sprint(orderby), "id"))
  q.Add("_page", stringWithDefault(fmt.Sprint(page), "1"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = fmt.Sprintf("%s&%s", q.Encode(), search)
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Mythic Keystone Affix API
//
// Name: Mythic Keystone Affixes Index
//
// Description: Returns an index of mythic keystone affixes.
//
// Path: /data/wow/keystone-affix/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) MythicKeystoneAffixesIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/keystone-affix/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Mythic Keystone Affix API
//
// Name: Mythic Keystone Affix
//
// Description: Returns a mythic keystone affix by ID.
//
// Path: /data/wow/keystone-affix/{keystoneAffixId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   keystoneAffixId: The ID of the mythic keystone affix. (int) (required) (default: 1)
func (c *Client) MythicKeystoneAffix(keystoneAffixId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/keystone-affix/%v", keystoneAffixId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Mythic Keystone Affix API
//
// Name: Mythic Keystone Affix Media
//
// Description: Returns media for a mythic keystone affix by ID.
//
// Path: /data/wow/media/keystone-affix/{keystoneAffixId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   keystoneAffixId: The ID of the mythic keystone affix. (int) (required) (default: 1)
func (c *Client) MythicKeystoneAffixMedia(keystoneAffixId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/media/keystone-affix/%v", keystoneAffixId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Mythic Keystone Dungeon API
//
// Name: Mythic Keystone Dungeons Index
//
// Description: Returns an index of Mythic Keystone dungeons.
//
// Path: /data/wow/mythic-keystone/dungeon/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) MythicKeystoneDungeonsIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/mythic-keystone/dungeon/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("dynamic-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Mythic Keystone Dungeon API
//
// Name: Mythic Keystone Dungeon
//
// Description: Returns a Mythic Keystone dungeon by ID.
//
// Path: /data/wow/mythic-keystone/dungeon/{dungeonId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   dungeonId: The ID of the dungeon. (int) (required) (default: 197)
func (c *Client) MythicKeystoneDungeon(dungeonId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/mythic-keystone/dungeon/%v", dungeonId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("dynamic-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Mythic Keystone Dungeon API
//
// Name: Mythic Keystone Index
//
// Description: Returns an index of links to other documents related to Mythic Keystone dungeons.
//
// Path: /data/wow/mythic-keystone/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) MythicKeystoneIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/mythic-keystone/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("dynamic-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Mythic Keystone Dungeon API
//
// Name: Mythic Keystone Periods Index
//
// Description: Returns an index of Mythic Keystone periods.
//
// Path: /data/wow/mythic-keystone/period/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) MythicKeystonePeriodsIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/mythic-keystone/period/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("dynamic-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Mythic Keystone Dungeon API
//
// Name: Mythic Keystone Period
//
// Description: Returns a Mythic Keystone period by ID.
//
// Path: /data/wow/mythic-keystone/period/{periodId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   periodId: The ID of the Mythic Keystone season period. (int) (required) (default: 880)
func (c *Client) MythicKeystonePeriod(periodId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/mythic-keystone/period/%v", periodId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("dynamic-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Mythic Keystone Dungeon API
//
// Name: Mythic Keystone Seasons Index
//
// Description: Returns an index of Mythic Keystone seasons.
//
// Path: /data/wow/mythic-keystone/season/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) MythicKeystoneSeasonsIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/mythic-keystone/season/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("dynamic-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Mythic Keystone Dungeon API
//
// Name: Mythic Keystone Season
//
// Description: Returns a Mythic Keystone season by ID.
//
// Path: /data/wow/mythic-keystone/season/{seasonId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   seasonId: The ID of the Mythic Keystone season. (int) (required) (default: 8)
func (c *Client) MythicKeystoneSeason(seasonId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/mythic-keystone/season/%v", seasonId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("dynamic-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Mythic Keystone Leaderboard API
//
// Name: Mythic Keystone Leaderboards Index
//
// Description: Returns an index of Mythic Keystone Leaderboard dungeon instances for a connected realm.
//
// Path: /data/wow/connected-realm/{connectedRealmId}/mythic-leaderboard/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   connectedRealmId: The ID of the connected realm. (int) (required) (default: 11)
func (c *Client) MythicKeystoneLeaderboardsIndex(connectedRealmId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/connected-realm/%v/mythic-leaderboard/index", connectedRealmId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("dynamic-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Mythic Keystone Leaderboard API
//
// Name: Mythic Keystone Leaderboard
//
// Description: Returns a weekly Mythic Keystone Leaderboard by period.
//
// Path: /data/wow/connected-realm/{connectedRealmId}/mythic-leaderboard/{dungeonId}/period/{period}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   connectedRealmId: The ID of the connected realm. (int) (required) (default: 11)
//   dungeonId: The ID of the dungeon. (int) (required) (default: 197)
//   period: The unique identifier for the leaderboard period. (int) (required) (default: 641)
func (c *Client) MythicKeystoneLeaderboard(connectedRealmId int, dungeonId int, period int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/connected-realm/%v/mythic-leaderboard/%v/period/%v", connectedRealmId, dungeonId, period))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("dynamic-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Mythic Raid Leaderboard API
//
// Name: Mythic Raid Leaderboard
//
// Description: Returns the leaderboard for a given raid and faction.
//
// Path: /data/wow/leaderboard/hall-of-fame/{raid}/{faction}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   raid: The raid for a leaderboard. (string) (required) (default: uldir)
//   faction: Player faction (`alliance` or `horde`). (string) (required) (default: alliance)
func (c *Client) MythicRaidLeaderboard(raid string, faction string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/leaderboard/hall-of-fame/%v/%v", raid, faction))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("dynamic-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Pet API
//
// Name: Pets Index
//
// Description: Returns an index of battle pets.
//
// Path: /data/wow/pet/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) PetsIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/pet/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Pet API
//
// Name: Pet
//
// Description: Returns a battle pets by ID.
//
// Path: /data/wow/pet/{petId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   petId: The ID of the pet. (int) (required) (default: 39)
func (c *Client) Pet(petId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/pet/%v", petId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Pet API
//
// Name: Pet Media
//
// Description: Returns media for a battle pet by ID.
//
// Path: /data/wow/media/pet/{petId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   petId: The ID of the pet. (int) (required) (default: 39)
func (c *Client) PetMedia(petId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/media/pet/%v", petId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Pet API
//
// Name: Pet Abilities Index
//
// Description: Returns an index of pet abilities.
//
// Path: /data/wow/pet-ability/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) PetAbilitiesIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/pet-ability/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Pet API
//
// Name: Pet Ability
//
// Description: Returns a pet ability by ID.
//
// Path: /data/wow/pet-ability/{petAbilityId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   petAbilityId: The ID of the pet ability. (int) (required) (default: 110)
func (c *Client) PetAbility(petAbilityId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/pet-ability/%v", petAbilityId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Pet API
//
// Name: Pet Ability Media
//
// Description: Returns media for a pet ability by ID.
//
// Path: /data/wow/media/pet-ability/{petAbilityId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   petAbilityId: The ID of the pet ability. (int) (required) (default: 110)
func (c *Client) PetAbilityMedia(petAbilityId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/media/pet-ability/%v", petAbilityId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
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
func (c *Client) PlayableClassesIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/playable-class/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
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
func (c *Client) PlayableClass(classId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/playable-class/%v", classId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
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
func (c *Client) PlayableClassMedia(playableClassId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/media/playable-class/%v", playableClassId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Playable Class API
//
// Name: PvP Talent Slots
//
// Description: Returns the PvP talent slots for a playable class by ID.
//
// Path: /data/wow/playable-class/{classId}/pvp-talent-slots
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   classId: The ID of the playable class. (int) (required) (default: 7)
func (c *Client) PvPTalentSlots(classId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/playable-class/%v/pvp-talent-slots", classId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
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
func (c *Client) PlayableRacesIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/playable-race/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
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
func (c *Client) PlayableRace(playableRaceId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/playable-race/%v", playableRaceId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Playable Specialization API
//
// Name: Playable Specializations Index
//
// Description: Returns an index of playable specializations.
//
// Path: /data/wow/playable-specialization/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) PlayableSpecializationsIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/playable-specialization/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Playable Specialization API
//
// Name: Playable Specialization
//
// Description: Returns a playable specialization by ID.
//
// Path: /data/wow/playable-specialization/{specId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   specId: The ID of the playable specialization. (int) (required) (default: 262)
func (c *Client) PlayableSpecialization(specId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/playable-specialization/%v", specId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Playable Specialization API
//
// Name: Playable Specialization Media
//
// Description: Returns media for a playable specialization by ID.
//
// Path: /data/wow/media/playable-specialization/{specId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   specId: The ID of the playable specialization. (int) (required) (default: 262)
func (c *Client) PlayableSpecializationMedia(specId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/media/playable-specialization/%v", specId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
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
func (c *Client) PowerTypesIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/power-type/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
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
func (c *Client) PowerType(powerTypeId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/power-type/%v", powerTypeId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Profession API
//
// Name: Professions Index
//
// Description: Returns an index of professions.
//
// Path: /data/wow/profession/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) ProfessionsIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/profession/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Profession API
//
// Name: Profession
//
// Description: Returns a profession by ID.
//
// Path: /data/wow/profession/{professionId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   professionId: The ID of the profession. (int) (required) (default: 164)
func (c *Client) Profession(professionId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/profession/%v", professionId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Profession API
//
// Name: Profession Media
//
// Description: Returns media for a profession by ID.
//
// Path: /data/wow/media/profession/{professionId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   professionId: The ID of the profession. (int) (required) (default: 164)
func (c *Client) ProfessionMedia(professionId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/media/profession/%v", professionId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Profession API
//
// Name: Profession Skill Tier
//
// Description: Returns a skill tier for a profession by ID.
//
// Path: /data/wow/profession/{professionId}/skill-tier/{skillTierId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   professionId: The ID of the profession. (int) (required) (default: 164)
//   skillTierId: The ID of the skill tier. (int) (required) (default: 2477)
func (c *Client) ProfessionSkillTier(professionId int, skillTierId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/profession/%v/skill-tier/%v", professionId, skillTierId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Profession API
//
// Name: Recipe
//
// Description: Returns a recipe by ID.
//
// Path: /data/wow/recipe/{recipeId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   recipeId: The ID of the recipe. (int) (required) (default: 1631)
func (c *Client) Recipe(recipeId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/recipe/%v", recipeId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Profession API
//
// Name: Recipe Media
//
// Description: Returns media for a recipe by ID.
//
// Path: /data/wow/media/recipe/{recipeId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   recipeId: The ID of the recipe. (int) (required) (default: 1631)
func (c *Client) RecipeMedia(recipeId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/media/recipe/%v", recipeId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
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
func (c *Client) PvPSeasonsIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/pvp-season/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("dynamic-%s", c.region))
  q.Add("locale", c.locale)
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
//   pvpSeasonId: The ID of the PvP season. (int) (required) (default: 33)
func (c *Client) PvPSeason(pvpSeasonId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/pvp-season/%v", pvpSeasonId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("dynamic-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: PvP Season API
//
// Name: PvP Leaderboards Index
//
// Description: Returns an index of PvP leaderboards for a PvP season.
//
// Path: /data/wow/pvp-season/{pvpSeasonId}/pvp-leaderboard/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   pvpSeasonId: The ID of the PvP season. (int) (required) (default: 33)
func (c *Client) PvPLeaderboardsIndex(pvpSeasonId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/pvp-season/%v/pvp-leaderboard/index", pvpSeasonId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("dynamic-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: PvP Season API
//
// Name: PvP Leaderboard
//
// Description: Returns the PvP leaderboard of a specific PvP bracket for a PvP season.
//
// Path: /data/wow/pvp-season/{pvpSeasonId}/pvp-leaderboard/{pvpBracket}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   pvpSeasonId: The ID of the PvP season. (int) (required) (default: 33)
//   pvpBracket: The PvP bracket type. (string) (required) (default: 3v3)
func (c *Client) PvPLeaderboard(pvpSeasonId int, pvpBracket string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/pvp-season/%v/pvp-leaderboard/%v", pvpSeasonId, pvpBracket))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("dynamic-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: PvP Season API
//
// Name: PvP Rewards Index
//
// Description: Returns an index of PvP rewards for a PvP season.
//
// Path: /data/wow/pvp-season/{pvpSeasonId}/pvp-reward/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   pvpSeasonId: The ID of the PvP season. (int) (required) (default: 33)
func (c *Client) PvPRewardsIndex(pvpSeasonId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/pvp-season/%v/pvp-reward/index", pvpSeasonId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("dynamic-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: PvP Tier API
//
// Name: PvP Tier Media
//
// Description: Returns media for a PvP tier by ID.
//
// Path: /data/wow/media/pvp-tier/{pvpTierId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   pvpTierId: The ID of the PvP tier. (int) (required) (default: 1)
func (c *Client) PvPTierMedia(pvpTierId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/media/pvp-tier/%v", pvpTierId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: PvP Tier API
//
// Name: PvP Tiers Index
//
// Description: Returns an index of PvP tiers.
//
// Path: /data/wow/pvp-tier/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) PvPTiersIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/pvp-tier/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: PvP Tier API
//
// Name: PvP Tier
//
// Description: Returns a PvP tier by ID.
//
// Path: /data/wow/pvp-tier/{pvpTierId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   pvpTierId: The ID of the PvP tier. (int) (required) (default: 1)
func (c *Client) PvPTier(pvpTierId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/pvp-tier/%v", pvpTierId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Quest API
//
// Name: Quests Index
//
// Description: Returns the parent index for quests.
//
// Path: /data/wow/quest/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) QuestsIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/quest/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Quest API
//
// Name: Quest
//
// Description: Returns a quest by ID.
//
// Path: /data/wow/quest/{questId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   questId: The ID of the quest. (int) (required) (default: 2)
func (c *Client) Quest(questId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/quest/%v", questId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Quest API
//
// Name: Quest Categories Index
//
// Description: Returns an index of quest categories (such as quests for a specific class, profession, or storyline).
//
// Path: /data/wow/quest/category/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) QuestCategoriesIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/quest/category/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Quest API
//
// Name: Quest Category
//
// Description: Returns a quest category by ID.
//
// Path: /data/wow/quest/category/{questCategoryId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   questCategoryId: The ID of the quest category. (string) (required) (default: 1)
func (c *Client) QuestCategory(questCategoryId string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/quest/category/%v", questCategoryId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Quest API
//
// Name: Quest Areas Index
//
// Description: Returns an index of quest areas.
//
// Path: /data/wow/quest/area/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) QuestAreasIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/quest/area/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Quest API
//
// Name: Quest Area
//
// Description: Returns a quest area by ID.
//
// Path: /data/wow/quest/area/{questAreaId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   questAreaId: The ID of the quest area. (string) (required) (default: 1)
func (c *Client) QuestArea(questAreaId string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/quest/area/%v", questAreaId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Quest API
//
// Name: Quest Types Index
//
// Description: Returns an index of quest types (such as PvP quests, raid quests, or account quests).
//
// Path: /data/wow/quest/type/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) QuestTypesIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/quest/type/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Quest API
//
// Name: Quest Type
//
// Description: Returns a quest type by ID.
//
// Path: /data/wow/quest/type/{questTypeId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   questTypeId: The ID of the quest type. (string) (required) (default: 1)
func (c *Client) QuestType(questTypeId string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/quest/type/%v", questTypeId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
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
func (c *Client) RealmsIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/realm/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("dynamic-%s", c.region))
  q.Add("locale", c.locale)
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
//   realmSlug: The slug of the realm. (string) (required) (default: tichondrius)
func (c *Client) Realm(realmSlug string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/realm/%v", realmSlug))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("dynamic-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Realm API
//
// Name: Realm Search
//
// Description: Performs a search of realms. The fields below are provided for example. For more detail see the Search Guide.
//
// Path: /data/wow/search/realm
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   orderby: The field to sort the result set by. (string)  (default: id)
//   page: The result page number. (int)  (default: 1)
//   search: Search query string. For more detail see the Search Guide: https://develop.battle.net/documentation/world-of-warcraft/guides/search
func (c *Client) RealmSearch(orderby string, page int, search string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/search/realm", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("dynamic-%s", c.region))
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
func (c *Client) RegionsIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/region/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("dynamic-%s", c.region))
  q.Add("locale", c.locale)
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
//   regionId: The ID of the region. (int) (required) (default: 1)
func (c *Client) Region(regionId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/region/%v", regionId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("dynamic-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Reputations API
//
// Name: Reputation Factions Index
//
// Description: Returns an index of reputation factions.
//
// Path: /data/wow/reputation-faction/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) ReputationFactionsIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/reputation-faction/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Reputations API
//
// Name: Reputation Faction
//
// Description: Returns a single reputation faction by ID.
//
// Path: /data/wow/reputation-faction/{reputationFactionId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   reputationFactionId: The ID of the reputation faction. (int) (required) (default: 21)
func (c *Client) ReputationFaction(reputationFactionId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/reputation-faction/%v", reputationFactionId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Reputations API
//
// Name: Reputation Tiers Index
//
// Description: Returns an index of reputation tiers.
//
// Path: /data/wow/reputation-tiers/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) ReputationTiersIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/reputation-tiers/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Reputations API
//
// Name: Reputation Tiers
//
// Description: Returns a single set of reputation tiers by ID.
//
// Path: /data/wow/reputation-tiers/{reputationTiersId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   reputationTiersId: The ID of the set of reputation tiers. (int) (required) (default: 2)
func (c *Client) ReputationTiers(reputationTiersId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/reputation-tiers/%v", reputationTiersId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Spell API
//
// Name: Spell
//
// Description: Returns a spell by ID.
//
// Path: /data/wow/spell/{spellId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   spellId: The ID of the spell. (int) (required) (default: 196607)
func (c *Client) Spell(spellId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/spell/%v", spellId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Spell API
//
// Name: Spell Media
//
// Description: Returns media for a spell by ID.
//
// Path: /data/wow/media/spell/{spellId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   spellId: The ID of the spell. (int) (required) (default: 196607)
func (c *Client) SpellMedia(spellId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/media/spell/%v", spellId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Spell API
//
// Name: Spell Search
//
// Description: Performs a search of spells. The fields below are provided for example. For more detail see the Search Guide.
//
// Path: /data/wow/search/spell
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   orderby: The field to sort the result set by. (string)  (default: id)
//   page: The result page number. (int)  (default: 1)
//   search: Search query string. For more detail see the Search Guide: https://develop.battle.net/documentation/world-of-warcraft/guides/search
func (c *Client) SpellSearch(orderby string, page int, search string) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/search/spell", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("orderby", stringWithDefault(fmt.Sprint(orderby), "id"))
  q.Add("_page", stringWithDefault(fmt.Sprint(page), "1"))
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = fmt.Sprintf("%s&%s", q.Encode(), search)
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Talent API
//
// Name: Talent Tree Index
//
// Description: Returns an index of talent trees.
//
// Path: /data/wow/talent-tree/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) TalentTreeIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/talent-tree/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Talent API
//
// Name: Talent Tree
//
// Description: Returns a talent tree by specialization ID.
//
// Path: /data/wow/talent-tree/{talentTreeId}/playable-specialization/{specId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   talentTreeId: The ID of the talent-tree. (int) (required) (default: 786)
//   specId: The ID of the playable-specialization. (int) (required) (default: 262)
func (c *Client) TalentTree(talentTreeId int, specId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/talent-tree/%v/playable-specialization/%v", talentTreeId, specId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Talent API
//
// Name: Talent Tree Nodes
//
// Description: Returns all talent tree nodes as well as links to associated playable specializations given a talent tree id. This is useful to generate loadout export codes.
//
// Path: /data/wow/talent-tree/{talentTreeId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   talentTreeId: The ID of the talent-tree. (int) (required) (default: 786)
func (c *Client) TalentTreeNodes(talentTreeId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/talent-tree/%v", talentTreeId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Talent API
//
// Name: Talents Index
//
// Description: Returns an index of talents.
//
// Path: /data/wow/talent/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) TalentsIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/talent/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Talent API
//
// Name: Talent
//
// Description: Returns a talent by ID.
//
// Path: /data/wow/talent/{talentId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   talentId: The ID of the talent. (int) (required) (default: 117163)
func (c *Client) Talent(talentId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/talent/%v", talentId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Talent API
//
// Name: PvP Talents Index
//
// Description: Returns an index of PvP talents.
//
// Path: /data/wow/pvp-talent/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) PvPTalentsIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/pvp-talent/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Talent API
//
// Name: PvP Talent
//
// Description: Returns a PvP talent by ID.
//
// Path: /data/wow/pvp-talent/{pvpTalentId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   pvpTalentId: The ID of the PvP talent. (int) (required) (default: 40)
func (c *Client) PvPTalent(pvpTalentId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/pvp-talent/%v", pvpTalentId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Tech Talent API
//
// Name: Tech Talent Tree Index
//
// Description: Returns an index of tech talent trees.
//
// Path: /data/wow/tech-talent-tree/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) TechTalentTreeIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/tech-talent-tree/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Tech Talent API
//
// Name: Tech Talent Tree
//
// Description: Returns a tech talent tree by ID.
//
// Path: /data/wow/tech-talent-tree/{techTalentTreeId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   techTalentTreeId: The ID of the tech talent tree. (int) (required) (default: 275)
func (c *Client) TechTalentTree(techTalentTreeId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/tech-talent-tree/%v", techTalentTreeId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Tech Talent API
//
// Name: Tech Talent Index
//
// Description: Returns an index of tech talents.
//
// Path: /data/wow/tech-talent/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) TechTalentIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/tech-talent/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Tech Talent API
//
// Name: Tech Talent
//
// Description: Returns a tech talent by ID.
//
// Path: /data/wow/tech-talent/{techTalentId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   techTalentId: The ID of the tech talent. (int) (required) (default: 863)
func (c *Client) TechTalent(techTalentId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/tech-talent/%v", techTalentId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Tech Talent API
//
// Name: Tech Talent Media
//
// Description: Returns media for a tech talent by ID.
//
// Path: /data/wow/media/tech-talent/{techTalentId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   techTalentId: The ID of the tech talent. (int) (required) (default: 863)
func (c *Client) TechTalentMedia(techTalentId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/media/tech-talent/%v", techTalentId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Title API
//
// Name: Titles Index
//
// Description: Returns an index of titles.
//
// Path: /data/wow/title/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) TitlesIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/title/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Title API
//
// Name: Title
//
// Description: Returns a title by ID.
//
// Path: /data/wow/title/{titleId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   titleId: The ID of the title. (int) (required) (default: 1)
func (c *Client) Title(titleId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/title/%v", titleId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Toy API
//
// Name: Toy Index
//
// Description: Returns an index of toys.
//
// Path: /data/wow/toy/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) ToyIndex() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/toy/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: Toy API
//
// Name: Toy
//
// Description: Returns a toy by id.
//
// Path: /data/wow/toy/{toyId}
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
//   toyId: The ID of the toy. (int) (required) (default: 30)
func (c *Client) Toy(toyId int) (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/toy/%v", toyId))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("static-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}

// Resource: WoW Token API
//
// Name: WoW Token Index (US, EU, KR, TW)
//
// Description: Returns the WoW Token index.
//
// Path: /data/wow/token/index
//
// HttpMethod: GET
//
// CnRegion: false
//
// Parameters: 
func (c *Client) WoWTokenIndexUSEUKRTW() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/token/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("dynamic-%s", c.region))
  q.Add("locale", c.locale)
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
func (c *Client) WoWTokenIndexCN() (*http.Response, error) {
  path, err := url.JoinPath(c.url, fmt.Sprintf("/data/wow/token/index", ))
  if err != nil {
    return nil, err
  }
  
  req, err := http.NewRequest("GET", path, nil)
  if err != nil {
    return nil, err
  }
  
  q := req.URL.Query() 
  q.Add("namespace", fmt.Sprintf("dynamic-%s", c.region))
  q.Add("locale", c.locale)
  q.Add("access-token", c.accessToken)
  req.URL.RawQuery = q.Encode()
  
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
  return c.httpClient.Do(req)
}
