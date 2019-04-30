package cli

import (
	"fmt"
	"github.com/urfave/cli"
	"net/http"
)

func Delete(c *cli.Context) error {
	if len(c.Args()) > 1 {
		return ErrTooManyArgs
	}

	if len(c.Args()) != 1 {
		return ErrBasketID
	}

	id := c.Args()[0]

	url := fmt.Sprintf("%s/baskets/%s", c.String("host"), id)

	req, err := http.NewRequest(http.MethodDelete, url, nil)

	if err != nil {
		return err
	}

	_, err = http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	fmt.Fprint(c.App.Writer, "basket deleted\n")

	return nil
}
