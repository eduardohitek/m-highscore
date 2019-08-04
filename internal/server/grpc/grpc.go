package grpc

import (
	"context"

	"net"

	pbhighscore "github.com/eduardohitek/m-apis/m-highscore/v1"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// Grpc asdfasfd
type Grpc struct {
	address string
	srv     *grpc.Server
}

// HighScore sdasfsd
var HighScore = 999999999999.0

// NewServer asdfasd
func NewServer(address string) *Grpc {
	return &Grpc{
		address: address,
	}
}

// SetHighScore sets the highscore to the HighScore variable
func (g *Grpc) SetHighScore(ctx context.Context, input *pbhighscore.SetHighScoreRequest) (*pbhighscore.SetHighScoreResponse, error) {
	log.Info().Msg("SetHighScore in m-highscore is called")
	HighScore = input.HighScore
	return &pbhighscore.SetHighScoreResponse{
		Set: true,
	}, nil
}

// GetHighScore returns the highscore from the HighScore variable
func (g *Grpc) GetHighScore(ctx context.Context, input *pbhighscore.GetHighScoreRequest) (*pbhighscore.GetHighScoreResponse, error) {
	// HighScore = 44444.0
	log.Info().Msg("GetHighScore in ms-highscore called")
	return &pbhighscore.GetHighScoreResponse{
		HighScore: HighScore,
	}, nil

}

// ListenAndServe asdff
func (g *Grpc) ListenAndServe() error {
	lis, err := net.Listen("tcp", g.address)
	if err != nil {
		return errors.Wrap(err, "failed to open port")
	}

	serverOpts := []grpc.ServerOption{}
	g.srv = grpc.NewServer(serverOpts...)

	pbhighscore.RegisterGameServer(g.srv, g)

	log.Info().Str("addrs", g.address).Msg("starting gRPC Server for m-highscore microservice")

	if err := g.srv.Serve(lis); err != nil {
		return errors.Wrap(err, "failed to start gRP Server for m-highscore microservice")
	}
	return nil
}
