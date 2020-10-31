package apiseeds

type SongResponse struct {
	Result struct {
		Artist struct {
			Name string `json:"name"`
		} `json:"artist"`
		Track struct {
			Name string `json:"name"`
			Text string `json:"text"`
			Lang struct {
				Code string `json:"code"`
				Name string `json:"name"`
			} `json:"lang"`
		} `json:"track"`
		copyright struct {
			Notice string `json:"notice"`
			Artist string `json:"artist"`
			Text   string `json:"text  "`
		} `json:"copyright"`
		Probability int     `json:"probability"`
		Similarity  float32 `json:"similarity "`
	} `json:"result"`
}
