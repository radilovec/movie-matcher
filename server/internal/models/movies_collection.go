package models

type MovieCollectionType string

const (
	Upcoming   MovieCollectionType = "upcoming"
	TopRated   MovieCollectionType = "top_rated"
	NowPlaying MovieCollectionType = "now_playing"
	Popular    MovieCollectionType = "popular"
)

func (m MovieCollectionType) IsValid() bool {
	switch m {
	case Upcoming, TopRated, NowPlaying, Popular:
		return true
	default:
		return false
	}
}
