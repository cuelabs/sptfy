package spotifyclient

import (
	"net/url"

)

type SptfyRpcClient struct {
	SptfyHost *url.URL
}