package grpc

import (
	"context"

	pbhighscore "github.com/eduardohitek/m-apis/m-highscore/v1"
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
