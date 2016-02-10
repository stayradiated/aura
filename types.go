package aura

type Links struct {
	Albums  []string `json:"albums,omitempty"`
	Tracks  []string `json:"tracks,omitempty"`
	Artists []string `json:"artists,omitempty"`
}

type Track struct {
	// REQUIRED
	ID     string    `json:"id"`     // A unique identifier
	Title  string    `json:"title"`  // The song's name
	Artist *[]string `json:"artist"` // The recording artists
	Type   string    `json:"type"`   // The string 'track'

	// Links
	Links Links `json:"links"` // Link to other collections

	// OPTIONAL
	Album         string    `json:"album,omitempty"`          // The name of the release the track appears on
	Track         int       `json:"track,omitempty"`          // The index of the track on its album
	TrackTotal    int       `json:"tracktotal,omitempty"`     // The number of tracks on the album
	Disc          int       `json:"disc,omitempty"`           // The index of the medium in the album
	DiscTotal     int       `json:"disctotal,omitempty"`      // The number of media in the album
	Year          int       `json:"year,omitempty"`           // The year the track was released
	Month         int       `json:"month,omitempty"`          // The release date’s month
	Day           int       `json:"day,omitempty"`            // The release date’s day of the month
	BPM           int       `json:"bpm,omitempty"`            // Tempo, in beats per minute
	Genre         *[]string `json:"genre,omitempty"`          // The track’s musical genres
	RecordingMBID string    `json:"recording_mbid,omitempty"` // A MusicBrainz recording id
	TrackMBID     string    `json:"track_mbid,omitempty"`     // A MusicBrainz track id
	Composer      *[]string `json:"composer,omitempty"`       // The names of the music’s composers
	AlbumArtist   *[]string `json:"albumartist,omitempty"`    // The artists for the release the track appears on
	Comments      string    `json:"comments,omitempty"`       // Free-form, user-specified information
	Format        string    `json:"format,omitempty"`         // The MIME type of the associated audio file
	Duration      float64   `json:"duration,omitempty"`       // The (approximate) length of the audio in seconds
	FrameRate     int       `json:"framerate,omitempty"`      // The number of frames per second in the audio
	FrameCount    int       `json:"framecount,omitempty"`     // The total number of frames in the audio.
	Channels      int       `json:"channels,omitempty"`       // The number of audio channels. (A frame consists of one sample per channel.
	BitRate       int       `json:"bitrate,omitempty"`        // The number of bits per second in the encoding
	BitDepth      int       `json:"bitdepth,omitempty"`       // The number of bits per sample
	Size          int       `json:"size,omitempty"`           // The size of the audio file in bytes
	Rating        int       `json:"rating,omitempty"`         // Track rating out of 100
}

type TrackList []Track

func (t TrackList) IDs() []string {
	IDs := make([]string, len(t))
	for i, track := range t {
		IDs[i] = track.ID
	}
	return IDs
}

type Album struct {
	// REQUIRED
	ID     string    `json:"id"`     // A unique identifier
	Title  string    `json:"title"`  // The album’s name
	Artist *[]string `json:"artist"` // The names of the artists responsible for the release
	Type   string    `json:"type"`   // The string 'album'

	// Links
	Links Links `json:"links"` // Link to other collections

	// OPTIONAL
	TrackTotal       int       `json:"tracktotal,omitempty"`         // The number of tracks on the album
	DiscTotal        int       `json:"disctotal,omitempty"`          // The number of media in the album
	Year             int       `json:"year,omitempty"`               // The year the album was released
	Month            int       `json:"month,omitempty"`              // The release date’s month
	Day              int       `json:"day,omitempty"`                // The release date’s day of the month
	Genre            *[]string `json:"genre,omitempty"`              // The album’s musical genres
	ReleaseMBID      string    `json:"release_mbid,omitempty"`       // A MusicBrainz release id
	ReleaseGroupMBID string    `json:"release_group_mbid,omitempty"` // A MusicBrainz release group id
}

type AlbumList []Album

func (a AlbumList) IDs() []string {
	IDs := make([]string, len(a))
	for i, album := range a {
		IDs[i] = album.ID
	}
	return IDs
}

type Artist struct {
	// REQUIRED
	ID   string `json:"id"`   // A unique identifier
	Name string `json:"name"` // The artist’s name
	Type string `json:"type"` // The string 'artist'

	// Links
	Links Links `json:"links"` // Link to other collections

	// OPTIONAL
	ArtistMBID string `json:"artist_mbid,omitempty"` // A MusicBrainz artist id
}

type ArtistList []Artist

func (a ArtistList) IDs() []string {
	IDs := make([]string, len(a))
	for i, artist := range a {
		IDs[i] = artist.ID
	}
	return IDs
}
