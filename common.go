package gephi_http_client

import (
	"encoding/json"
)

func marshal(enc *json.Encoder, operation string, o []interface{}) error {
	var err error
	m := make(map[string]interface{})

	for _, e := range o {
		m[operation] = e
		err = enc.Encode(m)
		if err != nil {
			return err
		}
		delete(m, operation)
	}

	return nil
}
