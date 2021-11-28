package movie

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/TendiF/grpc-stockbit/proto"
	"github.com/TendiF/grpc-stockbit/types"
)

type Server struct {
	proto.UnimplementedMovieServiceServer
}

func (s *Server) GetMovies(ctx context.Context, params *proto.GetMovieParams) (*proto.GetMovieResponse, error) {
	// http://www.omdbapi.com/?apikey=faf7e5bb&s=Batman&page=2

	resp, err := http.Get("http://www.omdbapi.com/?apikey=faf7e5bb&s=" + params.Search + "&page=" + strconv.Itoa(int(params.Page)))

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)

	res := types.ImdbResponse{}
	error := json.Unmarshal([]byte(sb), &res)

	if error != nil {
		return nil, error
	}

	movieResponse := proto.GetMovieResponse{}
	i, _ := strconv.ParseInt(res.TotalResults, 10, 64)
	for _, v := range res.Search {
		movieResponse.Data = append(movieResponse.Data, &proto.Movie{
			Title:  v.Title,
			ImdbID: v.ImdbID,
			Year:   v.Year,
			Poster: v.Poster,
		})
	}
	movieResponse.TotalResults = i

	return &movieResponse, nil
}

func (s *Server) GetDetailMovie(ctx context.Context, params *proto.GetDetailMovieParams) (*proto.GetDetailMovieParams, error) {

	log.Print("Params: ", params)

	return params, nil
}
