package models

type MovieCollectionType string

const (
	Upcoming   MovieCollectionType = "upcoming"
	TopRated   MovieCollectionType = "top_rated"
	NowPlaying MovieCollectionType = "now_playing"
	Popular    MovieCollectionType = "popular"
	Unique     MovieCollectionType = "unique"
)
