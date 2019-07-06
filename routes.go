package wrapper

import (
	"errors"
	"io/ioutil"

	"github.com/afex/hystrix-go/hystrix"
)

/*
	DoGet : synchronous Call to api

	Parameters:
	* api: Specifies client and is apart of hystrix stream name
	* route: Route makes hystrix stream name unique once appended
	* url: Endpoint
	* callback: Anything to do if it fails
*/
func (hw HystrixWrapper) DoGet(api, route, url string) ([]byte, error) {

	var bytes []byte

	err := hystrix.Do(api+"_"+route, func() error {

		resp, err := Client.clients[api].Get(url)

		if err != nil {
			return errors.New("Get call Failed")
		}

		bytes, err = ioutil.ReadAll(resp.Body)

		if err != nil {
			return errors.New("Get call didn't fail, but problem reading body")
		}

		return nil
	}, func(err error) error {

		return err
	})

	if err != nil {
		return nil, err
	}

	return bytes, nil
}
