/*
   Copyright 2019 Evan Hazlett <ejhazlett@gmail.com>

   Permission is hereby granted, free of charge, to any person obtaining a copy of
   this software and associated documentation files (the "Software"), to deal in the
   Software without restriction, including without limitation the rights to use, copy,
   modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
   and to permit persons to whom the Software is furnished to do so, subject to the
   following conditions:

   The above copyright notice and this permission notice shall be included in all copies
   or substantial portions of the Software.

   THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
   INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR
   PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE
   FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
   TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE
   USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package server

import (
	"io/ioutil"
	"runtime"
	"runtime/pprof"

	"github.com/ehazlett/atlas"
	ptypes "github.com/gogo/protobuf/types"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	api "github.com/ehazlett/atlas/api/services/nameserver/v1"
	"github.com/ehazlett/atlas/ds"
	"github.com/ehazlett/ttlcache"
)

var (
	empty = &ptypes.Empty{}
)

// Server is an Atlas server
type Server struct {
	cfg   *atlas.Config
	ds    ds.Datastore
	cache *ttlcache.TTLCache
}

// NewServer returns a new server
func NewServer(cfg *atlas.Config) (*Server, error) {
	ds, err := atlas.GetDatastore(cfg.Datastore)
	if err != nil {
		return nil, err
	}
	srv := &Server{
		cfg: cfg,
		ds:  ds,
	}
	if cfg.CacheTTL != 0 {
		c, err := ttlcache.NewTTLCache(cfg.CacheTTL)
		if err != nil {
			return nil, err
		}
		srv.cache = c
	}

	return srv, nil
}

// Register enables callers to register this service with an existing GRPC server
func (s *Server) Register(server *grpc.Server) error {
	logrus.Debug("registering nameserver with grpc")
	api.RegisterNameserverServer(server, s)
	return nil
}

// Start starts the embedded DNS server
func (s *Server) Start() error {
	return s.startDNSServer()
}

// Stop is used to stop and release resources
func (s *Server) Stop() error {
	return nil
}

// GenerateProfile generates a new Go profile
func (s *Server) GenerateProfile() (string, error) {
	tmpfile, err := ioutil.TempFile("", "atlas-profile-")
	if err != nil {
		return "", err
	}
	runtime.GC()
	if err := pprof.WriteHeapProfile(tmpfile); err != nil {
		return "", err
	}
	tmpfile.Close()
	return tmpfile.Name(), nil
}
