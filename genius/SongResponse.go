package genius

// SongResponse is the response from a song request to the genius API
type SongResponse struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Response struct {
		Song struct {
			AnnotationCount int    `json:"annotation_count"`
			APIPath         string `json:"api_path"`
			Description     struct {
				Dom struct {
					Tag      string        `json:"tag"`
					Children []interface{} `json:"children"`
				} `json:"dom"`
			} `json:"description"`
			EmbedContent             string `json:"embed_content"`
			FeaturedVideo            bool   `json:"featured_video"`
			FullTitle                string `json:"full_title"`
			HeaderImageThumbnailURL  string `json:"header_image_thumbnail_url"`
			HeaderImageURL           string `json:"header_image_url"`
			ID                       int    `json:"id"`
			LyricsOwnerID            int    `json:"lyrics_owner_id"`
			LyricsState              string `json:"lyrics_state"`
			Path                     string `json:"path"`
			PyongsCount              int    `json:"pyongs_count"`
			RecordingLocation        string `json:"recording_location"`
			ReleaseDate              string `json:"release_date"`
			SongArtImageThumbnailURL string `json:"song_art_image_thumbnail_url"`
			SongArtImageURL          string `json:"song_art_image_url"`
			Stats                    struct {
				AcceptedAnnotations   int  `json:"accepted_annotations"`
				Contributors          int  `json:"contributors"`
				Hot                   bool `json:"hot"`
				IqEarners             int  `json:"iq_earners"`
				Transcribers          int  `json:"transcribers"`
				UnreviewedAnnotations int  `json:"unreviewed_annotations"`
				VerifiedAnnotations   int  `json:"verified_annotations"`
				Pageviews             int  `json:"pageviews"`
			} `json:"stats"`
			Title               string `json:"title"`
			TitleWithFeatured   string `json:"title_with_featured"`
			URL                 string `json:"url"`
			CurrentUserMetadata struct {
				Permissions         []string `json:"permissions"`
				ExcludedPermissions []string `json:"excluded_permissions"`
				Interactions        struct {
					Pyong     bool `json:"pyong"`
					Following bool `json:"following"`
				} `json:"interactions"`
				Relationships struct {
				} `json:"relationships"`
				IqByAction struct {
				} `json:"iq_by_action"`
			} `json:"current_user_metadata"`
			Album struct {
				APIPath     string `json:"api_path"`
				CoverArtURL string `json:"cover_art_url"`
				FullTitle   string `json:"full_title"`
				ID          int    `json:"id"`
				Name        string `json:"name"`
				URL         string `json:"url"`
				Artist      struct {
					APIPath        string `json:"api_path"`
					HeaderImageURL string `json:"header_image_url"`
					ID             int    `json:"id"`
					ImageURL       string `json:"image_url"`
					IsMemeVerified bool   `json:"is_meme_verified"`
					IsVerified     bool   `json:"is_verified"`
					Name           string `json:"name"`
					URL            string `json:"url"`
				} `json:"artist"`
			} `json:"album"`
			CustomPerformances    []interface{} `json:"custom_performances"`
			DescriptionAnnotation struct {
				Type           string `json:"_type"`
				AnnotatorID    int    `json:"annotator_id"`
				AnnotatorLogin string `json:"annotator_login"`
				APIPath        string `json:"api_path"`
				Classification string `json:"classification"`
				Fragment       string `json:"fragment"`
				ID             int    `json:"id"`
				IsDescription  bool   `json:"is_description"`
				Path           string `json:"path"`
				Range          struct {
					Content string `json:"content"`
				} `json:"range"`
				SongID               int           `json:"song_id"`
				URL                  string        `json:"url"`
				VerifiedAnnotatorIds []interface{} `json:"verified_annotator_ids"`
				Annotatable          struct {
					APIPath          string `json:"api_path"`
					ClientTimestamps struct {
						UpdatedByHumanAt int `json:"updated_by_human_at"`
						LyricsUpdatedAt  int `json:"lyrics_updated_at"`
					} `json:"client_timestamps"`
					Context   string `json:"context"`
					ID        int    `json:"id"`
					ImageURL  string `json:"image_url"`
					LinkTitle string `json:"link_title"`
					Title     string `json:"title"`
					Type      string `json:"type"`
					URL       string `json:"url"`
				} `json:"annotatable"`
				Annotations []struct {
					APIPath string `json:"api_path"`
					Body    struct {
						Dom struct {
							Tag      string        `json:"tag"`
							Children []interface{} `json:"children"`
						} `json:"dom"`
					} `json:"body"`
					CommentCount        int         `json:"comment_count"`
					Community           bool        `json:"community"`
					CustomPreview       interface{} `json:"custom_preview"`
					HasVoters           bool        `json:"has_voters"`
					ID                  int         `json:"id"`
					Pinned              bool        `json:"pinned"`
					ShareURL            string      `json:"share_url"`
					Source              interface{} `json:"source"`
					State               string      `json:"state"`
					URL                 string      `json:"url"`
					Verified            bool        `json:"verified"`
					VotesTotal          int         `json:"votes_total"`
					CurrentUserMetadata struct {
						Permissions         []string `json:"permissions"`
						ExcludedPermissions []string `json:"excluded_permissions"`
						Interactions        struct {
							Cosign bool        `json:"cosign"`
							Pyong  bool        `json:"pyong"`
							Vote   interface{} `json:"vote"`
						} `json:"interactions"`
						IqByAction struct {
						} `json:"iq_by_action"`
					} `json:"current_user_metadata"`
					Authors []struct {
						Attribution float64     `json:"attribution"`
						PinnedRole  interface{} `json:"pinned_role"`
						User        struct {
							APIPath string `json:"api_path"`
							Avatar  struct {
								Tiny struct {
									URL         string `json:"url"`
									BoundingBox struct {
										Width  int `json:"width"`
										Height int `json:"height"`
									} `json:"bounding_box"`
								} `json:"tiny"`
								Thumb struct {
									URL         string `json:"url"`
									BoundingBox struct {
										Width  int `json:"width"`
										Height int `json:"height"`
									} `json:"bounding_box"`
								} `json:"thumb"`
								Small struct {
									URL         string `json:"url"`
									BoundingBox struct {
										Width  int `json:"width"`
										Height int `json:"height"`
									} `json:"bounding_box"`
								} `json:"small"`
								Medium struct {
									URL         string `json:"url"`
									BoundingBox struct {
										Width  int `json:"width"`
										Height int `json:"height"`
									} `json:"bounding_box"`
								} `json:"medium"`
							} `json:"avatar"`
							HeaderImageURL              string `json:"header_image_url"`
							HumanReadableRoleForDisplay string `json:"human_readable_role_for_display"`
							ID                          int    `json:"id"`
							Iq                          int    `json:"iq"`
							Login                       string `json:"login"`
							Name                        string `json:"name"`
							RoleForDisplay              string `json:"role_for_display"`
							URL                         string `json:"url"`
							CurrentUserMetadata         struct {
								Permissions         []interface{} `json:"permissions"`
								ExcludedPermissions []string      `json:"excluded_permissions"`
								Interactions        struct {
									Following bool `json:"following"`
								} `json:"interactions"`
								Features []interface{} `json:"features"`
							} `json:"current_user_metadata"`
						} `json:"user"`
					} `json:"authors"`
					CosignedBy       []interface{} `json:"cosigned_by"`
					RejectionComment interface{}   `json:"rejection_comment"`
					VerifiedBy       interface{}   `json:"verified_by"`
				} `json:"annotations"`
			} `json:"description_annotation"`
			FeaturedArtists []interface{} `json:"featured_artists"`
			Media           []struct {
				Provider   string `json:"provider"`
				ProviderID string `json:"provider_id,omitempty"`
				Type       string `json:"type"`
				URL        string `json:"url"`
				NativeURI  string `json:"native_uri,omitempty"`
				Start      int    `json:"start,omitempty"`
			} `json:"media"`
			PrimaryArtist struct {
				APIPath        string `json:"api_path"`
				HeaderImageURL string `json:"header_image_url"`
				ID             int    `json:"id"`
				ImageURL       string `json:"image_url"`
				IsMemeVerified bool   `json:"is_meme_verified"`
				IsVerified     bool   `json:"is_verified"`
				Name           string `json:"name"`
				URL            string `json:"url"`
			} `json:"primary_artist"`
			ProducerArtists []struct {
				APIPath        string `json:"api_path"`
				HeaderImageURL string `json:"header_image_url"`
				ID             int    `json:"id"`
				ImageURL       string `json:"image_url"`
				IsMemeVerified bool   `json:"is_meme_verified"`
				IsVerified     bool   `json:"is_verified"`
				Name           string `json:"name"`
				URL            string `json:"url"`
			} `json:"producer_artists"`
			SongRelationships []struct {
				Type  string        `json:"type"`
				Songs []interface{} `json:"songs"`
			} `json:"song_relationships"`
			VerifiedAnnotationsBy []interface{} `json:"verified_annotations_by"`
			VerifiedContributors  []interface{} `json:"verified_contributors"`
			VerifiedLyricsBy      []interface{} `json:"verified_lyrics_by"`
			WriterArtists         []struct {
				APIPath        string `json:"api_path"`
				HeaderImageURL string `json:"header_image_url"`
				ID             int    `json:"id"`
				ImageURL       string `json:"image_url"`
				IsMemeVerified bool   `json:"is_meme_verified"`
				IsVerified     bool   `json:"is_verified"`
				Name           string `json:"name"`
				URL            string `json:"url"`
			} `json:"writer_artists"`
		} `json:"song"`
	} `json:"response"`
}
