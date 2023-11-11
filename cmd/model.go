package cmd

type root struct {
	pathZoity *string     `json:"path_zoity"`
	services  []*service  `json:"services"`
	sequences []*sequence `json:"sequences"`
}

type service struct {
	id      *string `json:"id"`
	name    *string `json:"name"`
	command *string `json:"command"`
	path    *string `json:"path"`
}

type sequence struct {
	name     *string   `json:"name"`
	services []*string `json:"services"`
}
