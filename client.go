package wrapper

import (
	"errors"
	"net/http"
	"time"

	"github.com/afex/hystrix-go/hystrix"
)

// HystrixWrapper : Struct that holds all client information
type HystrixWrapper struct {
	clients     map[string]*http.Client
	contextType string
}

// Client : Holds all the clients that addressable by api string
var Client *HystrixWrapper

func init() {
	clients := make(map[string]*http.Client)

	Client = &HystrixWrapper{clients, "application/json"}

	Client.contextType = "application/json"
}

// InitializeClientWithTimeout : Takes in Options to configure specifics of a client
func (hw HystrixWrapper) InitializeClientWithTimeout(api string, options *http.Transport, timeout time.Duration) error {

	if api == "" {
		return errors.New("String api cannot be empty string, please define the api that will be hit by this client")
	}

	hw.clients[api] = &http.Client{Transport: options, Timeout: timeout}

	return nil
}

// InitializeClientWithoutTimeout : Takes in Options to configure specifics of a client without a timeout
func (hw HystrixWrapper) InitializeClientWithoutTimeout(api string, options *http.Transport) error {

	if api == "" {
		return errors.New("String api cannot be empty string, please define the api that will be hit by this client")
	}

	hw.clients[api] = &http.Client{Transport: options}

	return nil
}

// HystrixRoute : Set up a Hystrix Configuration for an api's route
func (hw HystrixWrapper) HystrixRoute(api, route string, timeout, maxConcurrentRequests, errorPercentThreshold int) {

	hystrix.ConfigureCommand(api+"_"+route, hystrix.CommandConfig{
		Timeout:               timeout,
		MaxConcurrentRequests: maxConcurrentRequests,
		ErrorPercentThreshold: errorPercentThreshold,
	})
}
