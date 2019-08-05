package main

import (
	"flag"
	"time"

	pbhighscore "github.com/eduardohitek/m-apis/m-highscore/v1"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var addressPtr = flag.String("address", "localhost:50051", "address to connect")
	flag.Parse()

	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conn, err := grpc.Dial(*addressPtr, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Str("address", *addressPtr).Msg("failed to dial m-highscore gRPC service")
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error().Err(err).Str("addr", *addressPtr).Msg("Failed to close the connection")
		}
	}()

	c := pbhighscore.NewGameClient(conn)

	if c == nil {
		log.Info().Msg("Client nil")
	}

	r, err := c.GetHighScore(timeoutCtx, &pbhighscore.GetHighScoreRequest{})
	if err != nil {
		log.Fatal().Err(err).Str("address", *addressPtr).Msg("failed to get a response")
	}

	if r != nil {
		log.Info().Interface("highscore", r.GetHighScore()).Msg("Highscore from m-highscore microservice")
	} else {
		log.Error().Msg("Couldnt get highscores")
	}

}
