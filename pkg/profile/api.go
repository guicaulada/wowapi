// Code generated. DO NOT EDIT.
// To generate use: go generate
//go:generate go run ../../scripts/gen.go ../../api-spec/profile-apis.json api.go profile
package profile

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"path/filepath"
)

const (
	BlizzardAPIUrl = "https://us.api.blizzard.com"
)

type Client struct {
	token string
}

func stringWithDefault(s string, d string) string {
	if s != "" {
		return d
	}
	return s
}

func getResponseBody(req *http.Request) (string, error) {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// Resource: Account Profile API
// Name: Account Profile Summary
// Description: Returns a profile summary for an account. Because this endpoint provides data about the current logged-in user's World of Warcraft account, it requires an access token with the wow. profile scope acquired via the Authorization Code Flow.
// Path: /profile/user/wow
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) AccountProfileSummary(namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/user/wow", )), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Account Profile API
// Name: Protected Character Profile Summary
// Description: Returns a protected profile summary for a character. Because this endpoint provides data about the current logged-in user's World of Warcraft account, it requires an access token with the wow. profile scope acquired via the Authorization Code Flow.
// Path: /profile/user/wow/protected-character/{realmId}-{characterId}
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   realmId: The ID of the character's realm. (int) (required) (default: 1)
//   characterId: The ID of the character. (int) (required) (default: 12345)
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) ProtectedCharacterProfileSummary(realmId int, characterId int, namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/user/wow/protected-character/%v-%v", realmId, characterId)), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Account Profile API
// Name: Account Collections Index
// Description: Returns an index of collection types for an account. Because this endpoint provides data about the current logged-in user's World of Warcraft account, it requires an access token with the wow. profile scope acquired via the Authorization Code Flow.
// Path: /profile/user/wow/collections
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) AccountCollectionsIndex(namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/user/wow/collections", )), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Account Profile API
// Name: Account Mounts Collection Summary
// Description: Returns a summary of the mounts an account has obtained. Because this endpoint provides data about the current logged-in user's World of Warcraft account, it requires an access token with the wow. profile scope acquired via the Authorization Code Flow.
// Path: /profile/user/wow/collections/mounts
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) AccountMountsCollectionSummary(namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/user/wow/collections/mounts", )), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Account Profile API
// Name: Account Pets Collection Summary
// Description: Returns a summary of the battle pets an account has obtained. Because this endpoint provides data about the current logged-in user's World of Warcraft account, it requires an access token with the wow. profile scope acquired via the Authorization Code Flow.
// Path: /profile/user/wow/collections/pets
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) AccountPetsCollectionSummary(namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/user/wow/collections/pets", )), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Account Profile API
// Name: Account Toys Collection Summary
// Description: Returns a summary of the toys an account has obtained. Because this endpoint provides data about the current logged-in user's World of Warcraft account, it requires an access token with the wow. profile scope acquired via the Authorization Code Flow.
// Path: /profile/user/wow/collections/toys
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) AccountToysCollectionSummary(namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/user/wow/collections/toys", )), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Account Profile API
// Name: Account Heirlooms Collection Summary
// Description: Returns a summary of the heirlooms an account has obtained. Because this endpoint provides data about the current logged-in user's World of Warcraft account, it requires an access token with the wow. profile scope acquired via the Authorization Code Flow.
// Path: /profile/user/wow/collections/heirlooms
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) AccountHeirloomsCollectionSummary(namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/user/wow/collections/heirlooms", )), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Character Achievements API
// Name: Character Achievements Summary
// Description: Returns a summary of the achievements a character has completed.
// Path: /profile/wow/character/{realmSlug}/{characterName}/achievements
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   realmSlug: The slug of the realm. (string) (required) (default: tichondrius)
//   characterName: The lowercase name of the character. (string) (required) (default: charactername)
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CharacterAchievementsSummary(realmSlug string, characterName string, namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/wow/character/%v/%v/achievements", realmSlug, characterName)), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Character Achievements API
// Name: Character Achievement Statistics
// Description: Returns a character's statistics as they pertain to achievements.
// Path: /profile/wow/character/{realmSlug}/{characterName}/achievements/statistics
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   realmSlug: The slug of the realm. (string) (required) (default: tichondrius)
//   characterName: The lowercase name of the character. (string) (required) (default: charactername)
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CharacterAchievementStatistics(realmSlug string, characterName string, namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/wow/character/%v/%v/achievements/statistics", realmSlug, characterName)), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Character Appearance API
// Name: Character Appearance Summary
// Description: Returns a summary of a character's appearance settings.
// Path: /profile/wow/character/{realmSlug}/{characterName}/appearance
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   realmSlug: The slug of the realm. (string) (required) (default: tichondrius)
//   characterName: The lowercase name of the character. (string) (required) (default: charactername)
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CharacterAppearanceSummary(realmSlug string, characterName string, namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/wow/character/%v/%v/appearance", realmSlug, characterName)), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Character Collections API
// Name: Character Collections Index
// Description: Returns an index of collection types for a character.
// Path: /profile/wow/character/{realmSlug}/{characterName}/collections
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   realmSlug: The slug of the realm. (string) (required) (default: tichondrius)
//   characterName: The lowercase name of the character. (string) (required) (default: charactername)
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CharacterCollectionsIndex(realmSlug string, characterName string, namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/wow/character/%v/%v/collections", realmSlug, characterName)), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Character Collections API
// Name: Character Mounts Collection Summary
// Description: Returns a summary of the mounts a character has obtained.
// Path: /profile/wow/character/{realmSlug}/{characterName}/collections/mounts
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   realmSlug: The slug of the realm. (string) (required) (default: tichondrius)
//   characterName: The lowercase name of the character. (string) (required) (default: charactername)
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CharacterMountsCollectionSummary(realmSlug string, characterName string, namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/wow/character/%v/%v/collections/mounts", realmSlug, characterName)), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Character Collections API
// Name: Character Pets Collection Summary
// Description: Returns a summary of the battle pets a character has obtained.
// Path: /profile/wow/character/{realmSlug}/{characterName}/collections/pets
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   realmSlug: The slug of the realm. (string) (required) (default: tichondrius)
//   characterName: The lowercase name of the character. (string) (required) (default: charactername)
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CharacterPetsCollectionSummary(realmSlug string, characterName string, namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/wow/character/%v/%v/collections/pets", realmSlug, characterName)), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Character Collections API
// Name: Character Toys Collection Summary
// Description: Returns a summary of the toys a character has obtained.
// Path: /profile/wow/character/{realmSlug}/{characterName}/collections/toys
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   realmSlug: The slug of the realm. (string) (required) (default: tichondrius)
//   characterName: The lowercase name of the character. (string) (required) (default: charactername)
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CharacterToysCollectionSummary(realmSlug string, characterName string, namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/wow/character/%v/%v/collections/toys", realmSlug, characterName)), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Character Collections API
// Name: Character Heirlooms Collection Summary
// Description: Returns a summary of the heirlooms a character has obtained.
// Path: /profile/wow/character/{realmSlug}/{characterName}/collections/heirlooms
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   realmSlug: The slug of the realm. (string) (required) (default: tichondrius)
//   characterName: The lowercase name of the character. (string) (required) (default: charactername)
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CharacterHeirloomsCollectionSummary(realmSlug string, characterName string, namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/wow/character/%v/%v/collections/heirlooms", realmSlug, characterName)), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Character Encounters API
// Name: Character Encounters Summary
// Description: Returns a summary of a character's encounters.
// Path: /profile/wow/character/{realmSlug}/{characterName}/encounters
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   realmSlug: The slug of the realm. (string) (required) (default: tichondrius)
//   characterName: The lowercase name of the character. (string) (required) (default: charactername)
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CharacterEncountersSummary(realmSlug string, characterName string, namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/wow/character/%v/%v/encounters", realmSlug, characterName)), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Character Encounters API
// Name: Character Dungeons
// Description: Returns a summary of a character's completed dungeons.
// Path: /profile/wow/character/{realmSlug}/{characterName}/encounters/dungeons
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   realmSlug: The slug of the realm. (string) (required) (default: tichondrius)
//   characterName: The lowercase name of the character. (string) (required) (default: charactername)
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CharacterDungeons(realmSlug string, characterName string, namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/wow/character/%v/%v/encounters/dungeons", realmSlug, characterName)), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Character Encounters API
// Name: Character Raids
// Description: Returns a summary of a character's completed raids.
// Path: /profile/wow/character/{realmSlug}/{characterName}/encounters/raids
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   realmSlug: The slug of the realm. (string) (required) (default: tichondrius)
//   characterName: The lowercase name of the character. (string) (required) (default: charactername)
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CharacterRaids(realmSlug string, characterName string, namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/wow/character/%v/%v/encounters/raids", realmSlug, characterName)), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Character Equipment API
// Name: Character Equipment Summary
// Description: Returns a summary of the items equipped by a character.
// Path: /profile/wow/character/{realmSlug}/{characterName}/equipment
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   realmSlug: The slug of the realm. (string) (required) (default: tichondrius)
//   characterName: The lowercase name of the character. (string) (required) (default: charactername)
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CharacterEquipmentSummary(realmSlug string, characterName string, namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/wow/character/%v/%v/equipment", realmSlug, characterName)), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Character Hunter Pets API
// Name: Character Hunter Pets Summary
// Description: If the character is a hunter, returns a summary of the character's hunter pets. Otherwise, returns an HTTP 404 Not Found error.
// Path: /profile/wow/character/{realmSlug}/{characterName}/hunter-pets
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   realmSlug: The slug of the realm. (string) (required) (default: tichondrius)
//   characterName: The lowercase name of the character. (string) (required) (default: charactername)
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CharacterHunterPetsSummary(realmSlug string, characterName string, namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/wow/character/%v/%v/hunter-pets", realmSlug, characterName)), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Character Media API
// Name: Character Media Summary
// Description: Returns a summary of the media assets available for a character (such as an avatar render).
// Path: /profile/wow/character/{realmSlug}/{characterName}/character-media
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   realmSlug: The slug of the realm. (string) (required) (default: tichondrius)
//   characterName: The lowercase name of the character. (string) (required) (default: charactername)
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CharacterMediaSummary(realmSlug string, characterName string, namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/wow/character/%v/%v/character-media", realmSlug, characterName)), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Character Mythic Keystone Profile API
// Name: Character Mythic Keystone Profile Index
// Description: Returns the Mythic Keystone profile index for a character.
// Path: /profile/wow/character/{realmSlug}/{characterName}/mythic-keystone-profile
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   realmSlug: The slug of the realm. (string) (required) (default: tichondrius)
//   characterName: The lowercase name of the character. (string) (required) (default: charactername)
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CharacterMythicKeystoneProfileIndex(realmSlug string, characterName string, namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/wow/character/%v/%v/mythic-keystone-profile", realmSlug, characterName)), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Character Mythic Keystone Profile API
// Name: Character Mythic Keystone Season Details
// Description: Returns the Mythic Keystone season details for a character. Returns a 404 Not Found for characters that have not yet completed a Mythic Keystone dungeon for the specified season.
// Path: /profile/wow/character/{realmSlug}/{characterName}/mythic-keystone-profile/season/{seasonId}
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   realmSlug: The slug of the realm. (string) (required) (default: tichondrius)
//   characterName: The lowercase name of the character. (string) (required) (default: charactername)
//   seasonId: The ID of the Mythic Keystone season. (string) (required) (default: 1)
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CharacterMythicKeystoneSeasonDetails(realmSlug string, characterName string, seasonId string, namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/wow/character/%v/%v/mythic-keystone-profile/season/%v", realmSlug, characterName, seasonId)), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Character Professions API
// Name: Character Professions Summary
// Description: Returns a summary of professions for a character.
// Path: /profile/wow/character/{realmSlug}/{characterName}/professions
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   realmSlug: The slug of the realm. (string) (required) (default: tichondrius)
//   characterName: The lowercase name of the character. (string) (required) (default: charactername)
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CharacterProfessionsSummary(realmSlug string, characterName string, namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/wow/character/%v/%v/professions", realmSlug, characterName)), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Character Profile API
// Name: Character Profile Summary
// Description: Returns a profile summary for a character.
// Path: /profile/wow/character/{realmSlug}/{characterName}
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   realmSlug: The slug of the realm. (string) (required) (default: tichondrius)
//   characterName: The lowercase name of the character. (string) (required) (default: charactername)
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CharacterProfileSummary(realmSlug string, characterName string, namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/wow/character/%v/%v", realmSlug, characterName)), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Character Profile API
// Name: Character Profile Status
// Description: Returns the status and a unique ID for a character. A client should delete information about a character from their application if any of the following conditions occur:an HTTP 404 Not Found error is returnedthe is_valid value is falsethe returned character ID doesn't match the previously recorded value for the characterThe following example illustrates how to use this endpoint:A client requests and stores information about a character, including its unique character ID and the timestamp of the request. After 30 days, the client makes a request to the status endpoint to verify if the character information is still valid. If character cannot be found, is not valid, or the characters IDs do not match, the client removes the information from their application. If the character is valid and the character IDs match, the client retains the data for another 30 days.
// Path: /profile/wow/character/{realmSlug}/{characterName}/status
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   realmSlug: The slug of the realm. (string) (required) (default: tichondrius)
//   characterName: The lowercase name of the character. (string) (required) (default: charactername)
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CharacterProfileStatus(realmSlug string, characterName string, namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/wow/character/%v/%v/status", realmSlug, characterName)), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Character PvP API
// Name: Character PvP Bracket Statistics
// Description: Returns the PvP bracket statistics for a character.
// Path: /profile/wow/character/{realmSlug}/{characterName}/pvp-bracket/{pvpBracket}
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   realmSlug: The slug of the realm. (string) (required) (default: tichondrius)
//   characterName: The lowercase name of the character. (string) (required) (default: charactername)
//   pvpBracket: The PvP bracket type. (string) (required) (default: 3v3)
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CharacterPvPBracketStatistics(realmSlug string, characterName string, pvpBracket string, namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/wow/character/%v/%v/pvp-bracket/%v", realmSlug, characterName, pvpBracket)), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Character PvP API
// Name: Character PvP Summary
// Description: Returns a PvP summary for a character.
// Path: /profile/wow/character/{realmSlug}/{characterName}/pvp-summary
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   realmSlug: The slug of the realm. (string) (required) (default: tichondrius)
//   characterName: The lowercase name of the character. (string) (required) (default: charactername)
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CharacterPvPSummary(realmSlug string, characterName string, namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/wow/character/%v/%v/pvp-summary", realmSlug, characterName)), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Character Quests API
// Name: Character Quests
// Description: Returns a character's active quests as well as a link to the character's completed quests.
// Path: /profile/wow/character/{realmSlug}/{characterName}/quests
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   realmSlug: The slug of the realm. (string) (required) (default: tichondrius)
//   characterName: The lowercase name of the character. (string) (required) (default: charactername)
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CharacterQuests(realmSlug string, characterName string, namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/wow/character/%v/%v/quests", realmSlug, characterName)), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Character Quests API
// Name: Character Completed Quests
// Description: Returns a list of quests that a character has completed.
// Path: /profile/wow/character/{realmSlug}/{characterName}/quests/completed
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   realmSlug: The slug of the realm. (string) (required) (default: tichondrius)
//   characterName: The lowercase name of the character. (string) (required) (default: charactername)
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CharacterCompletedQuests(realmSlug string, characterName string, namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/wow/character/%v/%v/quests/completed", realmSlug, characterName)), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Character Reputations API
// Name: Character Reputations Summary
// Description: Returns a summary of a character's reputations.
// Path: /profile/wow/character/{realmSlug}/{characterName}/reputations
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   realmSlug: The slug of the realm. (string) (required) (default: tichondrius)
//   characterName: The lowercase name of the character. (string) (required) (default: charactername)
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CharacterReputationsSummary(realmSlug string, characterName string, namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/wow/character/%v/%v/reputations", realmSlug, characterName)), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Character Soulbinds API
// Name: Character Soulbinds
// Description: Returns a character's soulbinds.
// Path: /profile/wow/character/{realmSlug}/{characterName}/soulbinds
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   realmSlug: The slug of the realm. (string) (required) (default: tichondrius)
//   characterName: The lowercase name of the character. (string) (required) (default: charactername)
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CharacterSoulbinds(realmSlug string, characterName string, namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/wow/character/%v/%v/soulbinds", realmSlug, characterName)), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Character Specializations API
// Name: Character Specializations Summary
// Description: Returns a summary of a character's specializations.
// Path: /profile/wow/character/{realmSlug}/{characterName}/specializations
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   realmSlug: The slug of the realm. (string) (required) (default: tichondrius)
//   characterName: The lowercase name of the character. (string) (required) (default: charactername)
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CharacterSpecializationsSummary(realmSlug string, characterName string, namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/wow/character/%v/%v/specializations", realmSlug, characterName)), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Character Statistics API
// Name: Character Statistics Summary
// Description: Returns a statistics summary for a character.
// Path: /profile/wow/character/{realmSlug}/{characterName}/statistics
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   realmSlug: The slug of the realm. (string) (required) (default: tichondrius)
//   characterName: The lowercase name of the character. (string) (required) (default: charactername)
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CharacterStatisticsSummary(realmSlug string, characterName string, namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/wow/character/%v/%v/statistics", realmSlug, characterName)), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Character Titles API
// Name: Character Titles Summary
// Description: Returns a summary of titles a character has obtained.
// Path: /profile/wow/character/{realmSlug}/{characterName}/titles
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   realmSlug: The slug of the realm. (string) (required) (default: tichondrius)
//   characterName: The lowercase name of the character. (string) (required) (default: charactername)
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) CharacterTitlesSummary(realmSlug string, characterName string, namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/profile/wow/character/%v/%v/titles", realmSlug, characterName)), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Guild API
// Name: Guild
// Description: Returns a single guild by its name and realm.
// Path: /data/wow/guild/{realmSlug}/{nameSlug}
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   realmSlug: The slug of the realm. (string) (required) (default: tichondrius)
//   nameSlug: The slug of the guild. (string) (required) (default: guild-slug)
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) Guild(realmSlug string, nameSlug string, namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/data/wow/guild/%v/%v", realmSlug, nameSlug)), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Guild API
// Name: Guild Activity
// Description: Returns a single guild's activity by name and realm.
// Path: /data/wow/guild/{realmSlug}/{nameSlug}/activity
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   realmSlug: The slug of the realm. (string) (required) (default: tichondrius)
//   nameSlug: The slug of the guild. (string) (required) (default: guild-slug)
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) GuildActivity(realmSlug string, nameSlug string, namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/data/wow/guild/%v/%v/activity", realmSlug, nameSlug)), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Guild API
// Name: Guild Achievements
// Description: Returns a single guild's achievements by name and realm.
// Path: /data/wow/guild/{realmSlug}/{nameSlug}/achievements
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   realmSlug: The slug of the realm. (string) (required) (default: tichondrius)
//   nameSlug: The slug of the guild. (string) (required) (default: guild-slug)
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) GuildAchievements(realmSlug string, nameSlug string, namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/data/wow/guild/%v/%v/achievements", realmSlug, nameSlug)), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}

// Resource: Guild API
// Name: Guild Roster
// Description: Returns a single guild's roster by its name and realm.
// Path: /data/wow/guild/{realmSlug}/{nameSlug}/roster
// HttpMethod: GET
// CnRegion: false
// Parameters: 
//   realmSlug: The slug of the realm. (string) (required) (default: tichondrius)
//   nameSlug: The slug of the guild. (string) (required) (default: guild-slug)
//   namespace: The namespace to use to locate this document. (string) (required) (default: profile-us)
//   locale: The locale to reflect in localized data. (string)  (default: en_US)
func (c *Client) GuildRoster(realmSlug string, nameSlug string, namespace string, locale string) (string, error) {
	req, err := http.NewRequest("GET", filepath.Join(BlizzardAPIUrl, fmt.Sprintf("/data/wow/guild/%v/%v/roster", realmSlug, nameSlug)), nil)
	if err != nil {
		return "", err
	}
	
	q := req.URL.Query() 
	q.Add("namespace", stringWithDefault(fmt.Sprint(namespace), "profile-us"))
	q.Add("locale", stringWithDefault(fmt.Sprint(locale), "en_US"))
	q.Add("access-token", c.token)
	req.URL.RawQuery = q.Encode()
	
	return getResponseBody(req)
}
