# wowapi
Golang API client for World of Warcraft.

In order to try this API you will need an OAuth access token from one of your registered OAuth clients. You can register a new client at https://develop.battle.net/access/clients.

For more information on the API, see https://develop.battle.net/documentation/world-of-warcraft.

This API client is generated using Go templates and GitHub Actions.

This project is not affiliated with Blizzard Entertainment.

## Usage
```go
  profileClient, err := profile.NewClient(myClientId, myClientSecret, myRegion, nil)
  if err != nil {
    log.Fatalln(err)
  }
  res, err := profileClient.CharacterProfileSummary("tichondrius", "sighmir", "profile-us", "en_US")
  if err != nil {
    log.Fatalln(err)
  }
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    log.Fatalln(err)
  }
  fmt.Println(string(body))
```
For more information check the [examples](examples) directory.

## License
    wowapi - Golang API client for World of Warcraft
    Copyright (C) 2023  Guilherme Caulada

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Affero General Public License as published
    by the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.

    You should have received a copy of the GNU Affero General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.