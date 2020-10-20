package lcd

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func (c *client) QueryContractLog(txId []byte) (contractLog ContractLog, err error) {
	path := fmt.Sprintf(UriQueryContractLogs, hex.EncodeToString(txId))

	if _, body, err := c.httpClient.Get(path, nil); err != nil {
		return contractLog, err
	} else {
		if err := json.Unmarshal(body, &contractLog); err != nil {
			return contractLog, err
		} else {
			return contractLog, nil
		}
	}
}
