package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	roflparser "rofl-parser"
)

func main() {
	replaysDir := "./replays"
	targetDir := "./target"

	// Crear el directorio target si no existe
	if err := os.MkdirAll(targetDir, os.ModePerm); err != nil {
		log.Fatalf("No se pudo crear el directorio target: %v", err)
	}

	// Procesar todos los archivos .rofl en ./replays
	err := filepath.WalkDir(replaysDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Ext(path) != ".rofl" {
			return nil
		}

		fmt.Printf("Procesando archivo: %s\n", path)
		roflData, err := roflparser.New(path, true)
		if err != nil {
			log.Printf("Error leyendo ROFL %s: %v", path, err)
			return nil
		}

		// Procesar StatsJSON para que sea un array JSON real
		var statsArr []map[string]interface{}
		if roflData.Metadata.StatsJSON != "" {
			if err := json.Unmarshal([]byte(roflData.Metadata.StatsJSON), &statsArr); err == nil {
				roflData.Metadata.Stats = statsArr
			}
		}

		// Guardar el Metadata como JSON en ./target/<nombre>.json
		baseName := filepath.Base(path)
		jsonName := baseName + ".json"
		jsonPath := filepath.Join(targetDir, jsonName)

		// Marshal completo, incluyendo el campo Stats como array
		jsonBytes, err := json.MarshalIndent(roflData.Metadata, "", "  ")
		if err != nil {
			log.Printf("Error serializando Metadata de %s: %v", path, err)
			return nil
		}

		if err := os.WriteFile(jsonPath, jsonBytes, 0644); err != nil {
			log.Printf("Error guardando JSON de %s: %v", path, err)
			return nil
		}

		fmt.Printf("Metadata guardado en: %s\n", jsonPath)
		return nil
	})

	if err != nil {
		log.Fatalf("Error procesando archivos: %v", err)
	}
}
