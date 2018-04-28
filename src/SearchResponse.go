package main

// SearchResponse is the response from a search request to the genius API
type SearchResponse struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Response struct {
		Hits []struct {
			Highlights []interface{} `json:"highlights"`
			Index      string        `json:"index"`
			Type       string        `json:"type"`
			Result     struct {
				AnnotationCount          int    `json:"annotation_count"`
				APIPath                  string `json:"api_path"`
				FullTitle                string `json:"full_title"`
				HeaderImageThumbnailURL  string `json:"header_image_thumbnail_url"`
				HeaderImageURL           string `json:"header_image_url"`
				ID                       int    `json:"id"`
				LyricsOwnerID            int    `json:"lyrics_owner_id"`
				LyricsState              string `json:"lyrics_state"`
				Path                     string `json:"path"`
				PyongsCount              int    `json:"pyongs_count"`
				SongArtImageThumbnailURL string `json:"song_art_image_thumbnail_url"`
				Stats                    struct {
					Hot                   bool `json:"hot"`
					UnreviewedAnnotations int  `json:"unreviewed_annotations"`
					Pageviews             int  `json:"pageviews"`
				} `json:"stats"`
				Title             string `json:"title"`
				TitleWithFeatured string `json:"title_with_featured"`
				URL               string `json:"url"`
				PrimaryArtist     struct {
					APIPath        string `json:"api_path"`
					HeaderImageURL string `json:"header_image_url"`
					ID             int    `json:"id"`
					ImageURL       string `json:"image_url"`
					IsMemeVerified bool   `json:"is_meme_verified"`
					IsVerified     bool   `json:"is_verified"`
					Name           string `json:"name"`
					URL            string `json:"url"`
				} `json:"primary_artist"`
			} `json:"result"`
		} `json:"hits"`
	} `json:"response"`
}
