package wrapper

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
)

/*
	DoGet : synchronous Get Call to api

	Parameters:
	* api: Specifies client and is apart of hystrix stream name
	* route: Route makes hystrix stream name unique once appended
	* url: Endpoint
*/
func (hw HystrixWrapper) DoGet(api, route, url string) ([]byte, error) {

	var byteResponse []byte

	err := hystrix.Do(api+"_"+route, func() error {

		resp, err := Client.clients[api].Get(url)

		if err != nil {
			return err
		}

		byteResponse, err = ioutil.ReadAll(resp.Body)

		if err != nil {
			return err
		}

		return nil

	}, func(err error) error {

		return err
	})

	if err != nil {
		return nil, err
	}

	return byteResponse, nil
}

/*
	GoGet : asynchronous Get Call to api

	Parameters:
	* api: Specifies client and is apart of hystrix stream name
	* route: Route makes hystrix stream name unique once appended
	* url: Endpoint
	* bytesChan: Bytes will be received in this channel
	* errChan: Error will be received in this channel
*/
func (hw HystrixWrapper) GoGet(api, route, url string, bytesChan chan []byte, errChan chan error) {

	hystrix.Go(api+"_"+route, func() error {

		resp, err := Client.clients[api].Get(url)

		if err != nil {
			errChan <- err

			return err
		}

		bytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			errChan <- err

			return err
		}

		bytesChan <- bytes

		return nil

	}, func(err error) error {

		errChan <- err

		return err
	})
}

/*
	DoPost : synchronous Post Call to api

	Parameters:
	* api: Specifies client and is apart of hystrix stream name
	* route: Route makes hystrix stream name unique once appended
	* url: Endpoint
	* body: bytes of the post body
*/
func (hw HystrixWrapper) DoRequestWithBody(api, route, url, method string, body []byte) ([]byte, error) {

	var byteResponse []byte

	err := hystrix.Do(api+"_"+route, func() error {

		ioReader := bytes.NewReader(body)

		request, err := http.NewRequest(method, url, ioReader)

		if err != nil {
			return err
		}

		resp, err := Client.clients[api].Do(request)

		if err != nil {
			return err
		}

		byteResponse, err = ioutil.ReadAll(resp.Body)

		if err != nil {
			return err
		}

		return nil

	}, func(err error) error {

		return err
	})

	if err != nil {
		return nil, err
	}

	return byteResponse, nil
}

/*
	GoPost : asynchronous Post Call to api

	Parameters:
	* api: Specifies client and is apart of hystrix stream name
	* route: Route makes hystrix stream name unique once appended
	* url: Endpoint
	* body: bytes of the post body
	* bytesChan: Bytes will be received in this channel
	* errChan: Error will be received in this channel
*/
func (hw HystrixWrapper) GoRequestWithBody(api, route, url, method string, body []byte, bytesChan chan []byte, errChan chan error) {

	hystrix.Go(api+"_"+route, func() error {

		ioReader := bytes.NewReader(body)

		request, err := http.NewRequest(method, url, ioReader)

		if err != nil {
			return err
		}

		resp, err := Client.clients[api].Do(request)

		if err != nil {
			errChan <- err

			return err
		}

		bytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			errChan <- err

			return err
		}

		bytesChan <- bytes

		return nil

	}, func(err error) error {

		errChan <- err

		return err
	})
}
