package spotifyclient

import (
	"net/url"

	"github.com/cuelabs/sptfy/internal/environment"
	"github.com/cuelabs/sptfy/pkg/user"
	"github.com/pkg/errors"
	/*
	pb "github.com/cuelabs/sptfy/rpc/sptfyapi"
	"github.com/cuelabs/sptfy/internal/auth"
	"fmt"
	*/
)

type SptfyRpcClient struct {
	SptfyHost *url.URL
}

func (s *SptfyRpcClient) RetrieveInfo(e *environment.Environment) (*user.SptfyUser, error) {
	return nil, errors.New("ERROR: Not implemented")
}

func (s *SptfyRpcClient) RetrieveAuth(e *environment.Environment) error {
	return errors.New("ERROR: Not implemented")
}