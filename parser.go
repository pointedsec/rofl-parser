package roflparser

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"

	"github.com/pointedsec/rofl-parser/model"
)

// NewFromReader permite parsear el archivo desde un io.Reader
func NewFromReader(reader io.Reader, verbose bool) (*model.Rofl, error) {
	// Lee todo el contenido en memoria
	allBytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("error leyendo datos: %w", err)
	}
	return parseRoflBytes(allBytes, verbose)
}

// New abre y parsea un archivo .rofl desde la ruta dada
func New(path string, verbose bool) (*model.Rofl, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error abriendo archivo: %w", err)
	}
	defer file.Close()

	allBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error leyendo archivo completo: %w", err)
	}
	return parseRoflBytes(allBytes, verbose)
}

// parseRoflBytes contiene la lógica principal del parseo
func parseRoflBytes(allBytes []byte, verbose bool) (*model.Rofl, error) {
	r := &model.Rofl{}
	buf := bytes.NewReader(allBytes)

	// --- Leer Magic y Signature ---
	if err := binary.Read(buf, binary.LittleEndian, &r.Magic); err != nil {
		return nil, fmt.Errorf("error leyendo magic: %w", err)
	}
	if !bytes.HasPrefix(r.Magic[:], []byte("RIOT")) {
		return nil, fmt.Errorf("magic number inválido: %v", r.Magic)
	}
	if verbose {
		fmt.Println("Magic OK:", string(r.Magic[:4]))
	}

	if err := binary.Read(buf, binary.LittleEndian, &r.Signature); err != nil {
		return nil, fmt.Errorf("error leyendo signature: %w", err)
	}
	if verbose {
		fmt.Printf("Signature: %x\n", r.Signature[:16])
	}

	// --- Leer longitudes y offsets ---
	if err := binary.Read(buf, binary.LittleEndian, &r.Lengths); err != nil {
		return nil, fmt.Errorf("error leyendo lengths: %w", err)
	}

	// Busca el inicio del JSON por la secuencia específica
	jsonStartSeq := []byte(`{"gameLength":`)
	start := bytes.Index(allBytes, jsonStartSeq)
	if start == -1 {
		return nil, fmt.Errorf("no se encontró el inicio del bloque JSON con '\"gameLength\":'")
	}

	// Busca el cierre del JSON correspondiente
	end := start
	bracketCount := 0
	foundStart := false
	for i := start; i < len(allBytes); i++ {
		if allBytes[i] == '{' {
			bracketCount++
			foundStart = true
		} else if allBytes[i] == '}' {
			bracketCount--
			if foundStart && bracketCount == 0 {
				end = i
				break
			}
		}
	}
	if end <= start {
		return nil, fmt.Errorf("no se encontró el final del bloque JSON")
	}
	metaBytes := allBytes[start : end+1]

	// --- Validación y análisis del JSON ---
	var metaMap map[string]interface{}
	if err := json.Unmarshal(metaBytes, &metaMap); err != nil {
		return nil, fmt.Errorf("error parseando metadata JSON: %w", err)
	}

	expectedFields := getJSONFields(reflect.TypeOf(model.MetadataJson{}))
	missingFields := []string{}
	extraFields := []string{}

	for _, field := range expectedFields {
		if _, ok := metaMap[field]; !ok {
			missingFields = append(missingFields, field)
		}
	}
	for key := range metaMap {
		found := false
		for _, field := range expectedFields {
			if key == field {
				found = true
				break
			}
		}
		if !found {
			extraFields = append(extraFields, key)
		}
	}

	fmt.Printf("Campos faltantes en MetadataJson: %d\n", len(missingFields))
	if len(missingFields) > 0 {
		fmt.Printf("Faltan: %v\n", missingFields)
	}
	if len(extraFields) > 0 {
		fmt.Printf("Campos extra en JSON: %v\n", extraFields)
	}

	if err := json.Unmarshal(metaBytes, &r.Metadata); err != nil {
		return nil, fmt.Errorf("error parseando MetadataJson: %w", err)
	}

	if r.Metadata.StatsJSON != "" {
		var statsArr []map[string]interface{}
		if err := json.Unmarshal([]byte(r.Metadata.StatsJSON), &statsArr); err != nil {
			return nil, fmt.Errorf("error parseando stats JSON: %w", err)
		}
		expectedStatsFields := getJSONFields(reflect.TypeOf(model.PlayerStatsJson{}))
		for idx, stats := range statsArr {
			missingStats := []string{}
			extraStats := []string{}
			for _, field := range expectedStatsFields {
				if _, ok := stats[field]; !ok {
					missingStats = append(missingStats, field)
				}
			}
			for key := range stats {
				found := false
				for _, field := range expectedStatsFields {
					if key == field {
						found = true
						break
					}
				}
				if !found {
					extraStats = append(extraStats, key)
				}
			}
			fmt.Printf("StatsJSON jugador %d: faltan %d campos, extras: %d\n", idx, len(missingStats), len(extraStats))
			if len(missingStats) > 0 {
				fmt.Printf("Faltan: %v\n", missingStats)
			}
			if len(extraStats) > 0 {
				fmt.Printf("Extras: %v\n", extraStats)
			}
		}
		if err := json.Unmarshal([]byte(r.Metadata.StatsJSON), &r.Metadata.Stats); err != nil {
			fmt.Printf("Advertencia: no se pudo parsear StatsJSON a PlayerStatsJson: %v\n", err)
		}
	}

	if verbose {
		fmt.Printf("Metadata cargada: Version=%s, GameLength=%d\n", r.Metadata.GameVersion, r.Metadata.GameLength)
	}

	return r, nil
}

// getJSONFields devuelve los nombres de los campos JSON de una estructura
func getJSONFields(t reflect.Type) []string {
	fields := []string{}
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("json")
		if tag != "" && tag != "-" {
			fields = append(fields, tag)
		}
	}
	return fields
}
