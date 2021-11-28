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

func (s *Server) GetDetailMovie(ctx context.Context, params *proto.GetDetailMovieParams) (*proto.GetMovieDetailResponse, error) {
	response := proto.GetMovieDetailResponse{}

	resp, err := http.Get("http://www.omdbapi.com/?apikey=faf7e5bb&i=" + params.Id)

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)

	res := types.Movie{}
	error := json.Unmarshal([]byte(sb), &res)

	if error != nil {
		log.Fatalln(error)
	}

	response.Actors = res.Actors
	response.Awards = res.Awards
	response.BoxOffice = res.BoxOffice
	response.Country = res.Country
	response.DVD = res.DVD
	response.Director = res.Director
	response.Genre = res.Genre
	response.ImdbID = res.ImdbID
	response.ImdbRating = res.ImdbRating
	response.ImdbVotes = res.ImdbVotes
	response.Language = res.Language
	response.Metascore = res.Metascore
	response.Plot = res.Plot
	response.Poster = res.Poster
	response.Production = res.Production
	response.Rated = res.Rated
	response.Released = res.Released
	response.Response = res.Response
	response.Runtime = res.Runtime
	response.Title = res.Title
	response.Type = res.Type
	response.Website = res.Website
	response.Writer = res.Writer
	response.Year = res.Year

	for _, v := range res.Ratings {
		response.Ratings = append(response.Ratings, &proto.Rating{
			Source: v.Source,
			Value:  v.Value,
		})
	}

	return &response, nil
}
