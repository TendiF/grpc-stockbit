// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: proto/movie.proto

package proto

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetMovieParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Search string `protobuf:"bytes,2,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetMovieParams) Reset() {
	*x = GetMovieParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_movie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovieParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieParams) ProtoMessage() {}

func (x *GetMovieParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_movie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieParams.ProtoReflect.Descriptor instead.
func (*GetMovieParams) Descriptor() ([]byte, []int) {
	return file_proto_movie_proto_rawDescGZIP(), []int{0}
}

func (x *GetMovieParams) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetMovieParams) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetMovieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalResults int64    `protobuf:"varint,1,opt,name=total_results,json=totalResults,proto3" json:"total_results,omitempty"`
	Data         []*Movie `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetMovieResponse) Reset() {
	*x = GetMovieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_movie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieResponse) ProtoMessage() {}

func (x *GetMovieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_movie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieResponse.ProtoReflect.Descriptor instead.
func (*GetMovieResponse) Descriptor() ([]byte, []int) {
	return file_proto_movie_proto_rawDescGZIP(), []int{1}
}

func (x *GetMovieResponse) GetTotalResults() int64 {
	if x != nil {
		return x.TotalResults
	}
	return 0
}

func (x *GetMovieResponse) GetData() []*Movie {
	if x != nil {
		return x.Data
	}
	return nil
}

type Movie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Year   string `protobuf:"bytes,2,opt,name=year,proto3" json:"year,omitempty"`
	ImdbID string `protobuf:"bytes,3,opt,name=imdbID,proto3" json:"imdbID,omitempty"`
	Poster string `protobuf:"bytes,4,opt,name=poster,proto3" json:"poster,omitempty"`
}

func (x *Movie) Reset() {
	*x = Movie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_movie_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Movie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movie) ProtoMessage() {}

func (x *Movie) ProtoReflect() protoreflect.Message {
	mi := &file_proto_movie_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Movie.ProtoReflect.Descriptor instead.
func (*Movie) Descriptor() ([]byte, []int) {
	return file_proto_movie_proto_rawDescGZIP(), []int{2}
}

func (x *Movie) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Movie) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *Movie) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

func (x *Movie) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

type GetDetailMovieParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetDetailMovieParams) Reset() {
	*x = GetDetailMovieParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_movie_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDetailMovieParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDetailMovieParams) ProtoMessage() {}

func (x *GetDetailMovieParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_movie_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDetailMovieParams.ProtoReflect.Descriptor instead.
func (*GetDetailMovieParams) Descriptor() ([]byte, []int) {
	return file_proto_movie_proto_rawDescGZIP(), []int{3}
}

func (x *GetDetailMovieParams) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetMovieDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title      string    `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Year       string    `protobuf:"bytes,2,opt,name=year,proto3" json:"year,omitempty"`
	Rated      string    `protobuf:"bytes,3,opt,name=rated,proto3" json:"rated,omitempty"`
	Released   string    `protobuf:"bytes,4,opt,name=released,proto3" json:"released,omitempty"`
	Runtime    string    `protobuf:"bytes,5,opt,name=runtime,proto3" json:"runtime,omitempty"`
	Genre      string    `protobuf:"bytes,6,opt,name=genre,proto3" json:"genre,omitempty"`
	Director   string    `protobuf:"bytes,7,opt,name=director,proto3" json:"director,omitempty"`
	Writer     string    `protobuf:"bytes,8,opt,name=writer,proto3" json:"writer,omitempty"`
	Actors     string    `protobuf:"bytes,9,opt,name=actors,proto3" json:"actors,omitempty"`
	Plot       string    `protobuf:"bytes,10,opt,name=plot,proto3" json:"plot,omitempty"`
	Language   string    `protobuf:"bytes,11,opt,name=language,proto3" json:"language,omitempty"`
	Country    string    `protobuf:"bytes,12,opt,name=country,proto3" json:"country,omitempty"`
	Awards     string    `protobuf:"bytes,13,opt,name=awards,proto3" json:"awards,omitempty"`
	Ratings    []*Rating `protobuf:"bytes,14,rep,name=ratings,proto3" json:"ratings,omitempty"`
	Metascore  string    `protobuf:"bytes,15,opt,name=metascore,proto3" json:"metascore,omitempty"`
	ImdbRating string    `protobuf:"bytes,16,opt,name=imdbRating,proto3" json:"imdbRating,omitempty"`
	ImdbVotes  string    `protobuf:"bytes,17,opt,name=imdbVotes,proto3" json:"imdbVotes,omitempty"`
	ImdbID     string    `protobuf:"bytes,18,opt,name=imdbID,proto3" json:"imdbID,omitempty"`
	DVD        string    `protobuf:"bytes,19,opt,name=DVD,proto3" json:"DVD,omitempty"`
	BoxOffice  string    `protobuf:"bytes,20,opt,name=boxOffice,proto3" json:"boxOffice,omitempty"`
	Production string    `protobuf:"bytes,21,opt,name=production,proto3" json:"production,omitempty"`
	Website    string    `protobuf:"bytes,22,opt,name=website,proto3" json:"website,omitempty"`
	Response   string    `protobuf:"bytes,23,opt,name=response,proto3" json:"response,omitempty"`
	Type       string    `protobuf:"bytes,24,opt,name=type,proto3" json:"type,omitempty"`
	Poster     string    `protobuf:"bytes,25,opt,name=poster,proto3" json:"poster,omitempty"`
}

func (x *GetMovieDetailResponse) Reset() {
	*x = GetMovieDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_movie_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovieDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieDetailResponse) ProtoMessage() {}

func (x *GetMovieDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_movie_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieDetailResponse.ProtoReflect.Descriptor instead.
func (*GetMovieDetailResponse) Descriptor() ([]byte, []int) {
	return file_proto_movie_proto_rawDescGZIP(), []int{4}
}

func (x *GetMovieDetailResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetMovieDetailResponse) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *GetMovieDetailResponse) GetRated() string {
	if x != nil {
		return x.Rated
	}
	return ""
}

func (x *GetMovieDetailResponse) GetReleased() string {
	if x != nil {
		return x.Released
	}
	return ""
}

func (x *GetMovieDetailResponse) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *GetMovieDetailResponse) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *GetMovieDetailResponse) GetDirector() string {
	if x != nil {
		return x.Director
	}
	return ""
}

func (x *GetMovieDetailResponse) GetWriter() string {
	if x != nil {
		return x.Writer
	}
	return ""
}

func (x *GetMovieDetailResponse) GetActors() string {
	if x != nil {
		return x.Actors
	}
	return ""
}

func (x *GetMovieDetailResponse) GetPlot() string {
	if x != nil {
		return x.Plot
	}
	return ""
}

func (x *GetMovieDetailResponse) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *GetMovieDetailResponse) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *GetMovieDetailResponse) GetAwards() string {
	if x != nil {
		return x.Awards
	}
	return ""
}

func (x *GetMovieDetailResponse) GetRatings() []*Rating {
	if x != nil {
		return x.Ratings
	}
	return nil
}

func (x *GetMovieDetailResponse) GetMetascore() string {
	if x != nil {
		return x.Metascore
	}
	return ""
}

func (x *GetMovieDetailResponse) GetImdbRating() string {
	if x != nil {
		return x.ImdbRating
	}
	return ""
}

func (x *GetMovieDetailResponse) GetImdbVotes() string {
	if x != nil {
		return x.ImdbVotes
	}
	return ""
}

func (x *GetMovieDetailResponse) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

func (x *GetMovieDetailResponse) GetDVD() string {
	if x != nil {
		return x.DVD
	}
	return ""
}

func (x *GetMovieDetailResponse) GetBoxOffice() string {
	if x != nil {
		return x.BoxOffice
	}
	return ""
}

func (x *GetMovieDetailResponse) GetProduction() string {
	if x != nil {
		return x.Production
	}
	return ""
}

func (x *GetMovieDetailResponse) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *GetMovieDetailResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *GetMovieDetailResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GetMovieDetailResponse) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

type Rating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source string `protobuf:"bytes,1,opt,name=Source,proto3" json:"Source,omitempty"`
	Value  string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *Rating) Reset() {
	*x = Rating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_movie_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rating) ProtoMessage() {}

func (x *Rating) ProtoReflect() protoreflect.Message {
	mi := &file_proto_movie_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rating.ProtoReflect.Descriptor instead.
func (*Rating) Descriptor() ([]byte, []int) {
	return file_proto_movie_proto_rawDescGZIP(), []int{5}
}

func (x *Rating) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Rating) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_proto_movie_proto protoreflect.FileDescriptor

var file_proto_movie_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x59, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12,
	0x20, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x61, 0x0a, 0x05, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x79, 0x65, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f,
	0x73, 0x74, 0x65, 0x72, 0x22, 0x26, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa1, 0x05, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x61,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x72, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x65,
	0x6e, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x6c, 0x6f, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x6c, 0x6f, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x77, 0x61,
	0x72, 0x64, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x77, 0x61, 0x72, 0x64,
	0x73, 0x12, 0x27, 0x0a, 0x07, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0e, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x07, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x65,
	0x74, 0x61, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d,
	0x65, 0x74, 0x61, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6d, 0x64, 0x62,
	0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6d,
	0x64, 0x62, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6d, 0x64, 0x62,
	0x56, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x64,
	0x62, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x64, 0x62, 0x49, 0x44,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x12, 0x10,
	0x0a, 0x03, 0x44, 0x56, 0x44, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x44, 0x56, 0x44,
	0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6f, 0x78, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6f, 0x78, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x18, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74,
	0x65, 0x72, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72,
	0x22, 0x36, 0x0a, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xce, 0x01, 0x0a, 0x0c, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x6a, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12,
	0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x1d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x2d, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_movie_proto_rawDescOnce sync.Once
	file_proto_movie_proto_rawDescData = file_proto_movie_proto_rawDesc
)

func file_proto_movie_proto_rawDescGZIP() []byte {
	file_proto_movie_proto_rawDescOnce.Do(func() {
		file_proto_movie_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_movie_proto_rawDescData)
	})
	return file_proto_movie_proto_rawDescData
}

var file_proto_movie_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_movie_proto_goTypes = []interface{}{
	(*GetMovieParams)(nil),         // 0: proto.GetMovieParams
	(*GetMovieResponse)(nil),       // 1: proto.GetMovieResponse
	(*Movie)(nil),                  // 2: proto.Movie
	(*GetDetailMovieParams)(nil),   // 3: proto.GetDetailMovieParams
	(*GetMovieDetailResponse)(nil), // 4: proto.GetMovieDetailResponse
	(*Rating)(nil),                 // 5: proto.Rating
}
var file_proto_movie_proto_depIdxs = []int32{
	2, // 0: proto.GetMovieResponse.data:type_name -> proto.Movie
	5, // 1: proto.GetMovieDetailResponse.ratings:type_name -> proto.Rating
	0, // 2: proto.MovieService.GetMovies:input_type -> proto.GetMovieParams
	3, // 3: proto.MovieService.GetDetailMovie:input_type -> proto.GetDetailMovieParams
	1, // 4: proto.MovieService.GetMovies:output_type -> proto.GetMovieResponse
	4, // 5: proto.MovieService.GetDetailMovie:output_type -> proto.GetMovieDetailResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_movie_proto_init() }
func file_proto_movie_proto_init() {
	if File_proto_movie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_movie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovieParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_movie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovieResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_movie_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Movie); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_movie_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDetailMovieParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_movie_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovieDetailResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_movie_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rating); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_movie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_movie_proto_goTypes,
		DependencyIndexes: file_proto_movie_proto_depIdxs,
		MessageInfos:      file_proto_movie_proto_msgTypes,
	}.Build()
	File_proto_movie_proto = out.File
	file_proto_movie_proto_rawDesc = nil
	file_proto_movie_proto_goTypes = nil
	file_proto_movie_proto_depIdxs = nil
}
