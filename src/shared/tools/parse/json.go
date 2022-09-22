package parse

import "encoding/json"

func Unmarshal(source, dest any) error {
	b, _ := json.Marshal(source)
	if err := json.Unmarshal(b, &dest); err != nil {
		return err
	}
	return nil
}
