package cli

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"net/http"
)

func Create(c *cli.Context) error {
	url := fmt.Sprintf("%s/baskets", c.String("host"))

	req, err := http.NewRequest(http.MethodPost, url, nil)

	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	var data map[string]int64

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return err
	}

	fmt.Fprintf(c.App.Writer, "basket id: %d\n", data["id"])

	return nil
}
