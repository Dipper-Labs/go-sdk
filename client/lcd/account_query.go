package lcd

import (
	"encoding/json"
	"fmt"
)

func (c *client) QueryAccount(address string) (accountInfo AccountBody, err error) {
	path := fmt.Sprintf(UriQueryAccount, address)

	if _, body, err := c.httpClient.Get(path, nil); err != nil {
		return accountInfo, err
	} else {
		if err := json.Unmarshal(body, &accountInfo); err != nil {
			return accountInfo, err
		} else {
			return accountInfo, nil
		}
	}
}
