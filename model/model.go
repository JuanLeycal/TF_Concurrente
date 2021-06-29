package model

type Ciudadano struct {
	ID                     int     `json:"id"`
	ESTRATO_SOCIOECONOMICO float32 `json:"estrato_socioeconomico"`
	SEGURIDAD_NOCTURNA     float32 `json:"seguridad_nocturna"`
	GRUPOS_EDAD            float32 `json:"grupos_edad"`
	CONFIANZA_POLICIA      float32 `json:"confianza_policia"`
	PRONTO_DELITO          float32 `json:"pronto_delito"`
	CLUSTER                int     `json:"cluster"`
}
