package main

type InfoClimatResponse struct {
	RequestState int16            `json:"request_state"`
	RequestKey   string           `json:"request_key"`
	Message      string           `json:"message"`
	ModelRun     string           `json:"model_run"`
	Source       string           `json:"source"`
	Meteo        map[string]Meteo `json:"-"`
}

type Meteo struct {
	Temperature struct {
		TwoM       float32 `json:"2m"`
		Sol        float32 `json:"sol"`
		Five00HPa  float32 `json:"500hPa"`
		Eight50HPa float32 `json:"850hPa"`
	} `json:"temperature"`
	Pression struct {
		NiveauDeLaMer int `json:"niveau_de_la_mer"`
	} `json:"pression"`
	Pluie    float32 `json:"pluie"`
	Humidite struct {
		TwoM float32 `json:"2m"`
	} `json:"humidite"`
	VentMoyen struct {
		One0M float32 `json:"10m"`
	} `json:"vent_moyen"`
	RisqueNeige string `json:"risque_neige"`
	Nebulosite  struct {
		Haute   int `json:"haute"`
		Moyenne int `json:"moyenne"`
		Basse   int `json:"basse"`
		Totale  int `json:"totale"`
	} `json:"nebulosite"`
}
