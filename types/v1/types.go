package types

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
)

type HexNumber big.Int

func (i HexNumber) MarshalJSON() ([]byte, error) {
	hexNumber := fmt.Sprintf("\"0x%x\"", (*big.Int)(&i).Uint64())

	return []byte(hexNumber), nil
}

func (i *HexNumber) UnmarshalJSON(data []byte) error {
	var resultStr string
	err := json.Unmarshal(data, &resultStr)
	if err != nil {
		return err
	}

	hex := strings.Replace(resultStr, "0x", "", 1)

	value, success := new(big.Int).SetString(hex, 16)
	if !success {
		return fmt.Errorf("could not parse hex value %v", resultStr)
	}

	*i = HexNumber(*value)

	return nil
}
