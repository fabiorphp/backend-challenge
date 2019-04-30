package cli

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"net/http"
)

var (
	ErrProductCode     = errors.New("product code is required")
	ErrProductNotFound = errors.New("product not found")
)

func Add(c *cli.Context) error {
	if len(c.Args()) > 2 {
		return ErrTooManyArgs
	}

	if c.Args().Get(0) == "" {
		return ErrBasketID
	}

	if c.Args().Get(1) == "" {
		return ErrProductCode
	}

	url := fmt.Sprintf(
		"%s/baskets/%s/items",
		c.String("host"),
		c.Args().Get(0),
	)

	buf := new(bytes.Buffer)

	body := map[string]string{"code": c.Args().Get(1)}

	err := json.NewEncoder(buf).Encode(body)

	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, buf)

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

	if res.StatusCode == http.StatusBadRequest {
		return ErrProductNotFound
	}

	fmt.Fprint(c.App.Writer, "product added\n")

	return nil
}
