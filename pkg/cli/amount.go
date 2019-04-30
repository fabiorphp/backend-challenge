package cli

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"net/http"
)

func Amount(c *cli.Context) error {
	if len(c.Args()) > 1 {
		return ErrTooManyArgs
	}

	if len(c.Args()) != 1 {
		return ErrBasketID
	}

	id := c.Args()[0]

	url := fmt.Sprintf("%s/baskets/%s/amount", c.String("host"), id)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	if res.StatusCode == http.StatusNotFound {
		return ErrBasketNotFound
	}

	var data map[string]float64

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return err
	}

	fmt.Fprintf(c.App.Writer, "basket amount: $%.2f\n", data["amount"])

	return nil
}
