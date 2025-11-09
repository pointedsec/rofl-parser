package model

// MetadataValidationError representa los errores de validación del bloque Metadata JSON
type MetadataValidationError struct {
	MissingFields []string `json:"missingFields,omitempty"`
	ExtraFields   []string `json:"extraFields,omitempty"`
}

// PlayerStatsValidationError representa los errores de validación de cada objeto StatsJSON
type PlayerStatsValidationError struct {
	PlayerIndex   int      `json:"playerIndex"`
	MissingFields []string `json:"missingFields,omitempty"`
	ExtraFields   []string `json:"extraFields,omitempty"`
}
