package aura

type Features struct {
	Tracks  *TracksFeature
	Albums  *AlbumsFeature
	Artists *ArtistsFeature
}

type Include struct {
	Albums  bool
	Tracks  bool
	Artists bool
}

type Entities struct {
	Albums  AlbumList  `json:"albums,omitempty"`
	Tracks  TrackList  `json:"tracks,omitempty"`
	Artists ArtistList `json:"artists,omitempty"`
}

type EntitiesMap struct {
	Albums  map[string]Album  `json:"albums,omitempty"`
	Tracks  map[string]Track  `json:"tracks,omitempty"`
	Artists map[string]Artist `json:"artists,omitempty"`
}

func (e Entities) asMap() EntitiesMap {
	result := EntitiesMap{}
	if e.Tracks != nil {
		result.Tracks = make(map[string]Track)
		for _, track := range e.Tracks {
			result.Tracks[track.ID] = track
		}
	}
	if e.Albums != nil {
		result.Albums = make(map[string]Album)
		for _, album := range e.Albums {
			result.Albums[album.ID] = album
		}
	}
	if e.Artists != nil {
		result.Artists = make(map[string]Artist)
		for _, artist := range e.Artists {
			result.Artists[artist.ID] = artist
		}
	}
	return result
}

type Response struct {
	Result   []string    `json:"result"`
	Entities EntitiesMap `json:"entities"`
}
