package model

type Ciudadano struct {
	ID                     int     `json:"id"`
	ESTRATO_SOCIOECONOMICO float64 `json:"estrato_socioeconomico"`
	SEGURIDAD_NOCTURNA     float64 `json:"seguridad_nocturna"`
	GRUPOS_EDAD            float64 `json:"grupos_edad"`
	CONFIANZA_POLICIA      float64 `json:"confianza_policia"`
	PRONTO_DELITO          float64 `json:"pronto_delito"`
	CLUSTER                int     `json:"cluster"`
}
