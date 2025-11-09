# rofl-parser

Librería en Go para parsear archivos `.rofl` de League of Legends y extraer su metadata en formato JSON.

## ¿Qué es esto?

`rofl-parser` es una librería que permite leer archivos de repeticiones (`.rofl`) de League of Legends, validar su formato, extraer el bloque de metadata JSON y analizar los datos de la partida y de los jugadores.  
Incluye validación de campos, detección de campos faltantes y extra, y conversión directa a estructuras Go y JSON.

## Características

- **Validación de Magic Number**: Solo procesa archivos `.rofl` válidos.
- **Extracción automática del bloque JSON**: Busca el bloque que comienza con `{"gameLength":` en el archivo binario.
- **Parseo de Metadata**: Convierte el bloque JSON a la estructura `MetadataJson`.
- **Parseo de StatsJSON**: Convierte el campo `statsJson` en un array de estadísticas de jugadores.
- **Detección de campos faltantes y extra**: Informa qué campos no están presentes y cuáles sobran en el JSON.
- **Compatible con Go Modules**.
- **Parseo desde archivos o streams**: Permite procesar archivos `.rofl` tanto desde disco como desde un `io.Reader` (por ejemplo, archivos subidos por el usuario vía API).

## Instalación

1. Clona el repositorio:
   ```sh
   git clone https://github.com/pointedsec/rofl-parser.git
   ```

2. Importa la librería en tu proyecto Go:
   ```go
   import "github.com/pointedsec/rofl-parser"
   ```

## Uso básico

```go
package main

import (
    "fmt"
    roflparser "github.com/pointedsec/rofl-parser"
)

func main() {
    // Parsear desde archivo en disco
    rofl, err := roflparser.New("ruta/al/archivo.rofl", true)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Printf("Versión: %s\n", rofl.Metadata.GameVersion)
    fmt.Printf("Duración: %d\n", rofl.Metadata.GameLength)
    fmt.Printf("Stats jugadores: %+v\n", rofl.Metadata.Stats)
}
```

### Parsear desde un `io.Reader` (por ejemplo, archivo subido por API)

```go
package main

import (
    "fmt"
    "os"
    roflparser "github.com/pointedsec/rofl-parser"
)

func main() {
    file, err := os.Open("ruta/al/archivo.rofl")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    rofl, err := roflparser.NewFromReader(file, true)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Printf("Versión: %s\n", rofl.Metadata.GameVersion)
}
```

## Estructuras principales

- `Rofl`: Estructura principal del archivo.
- `MetadataJson`: Metadata de la partida.
- `PlayerStatsJson`: Estadísticas de cada jugador.

## Ejemplo de salida

Al procesar un archivo `.rofl`, la librería muestra:
- Campos faltantes y extra en el JSON.
- Metadata parseada.
- Estadísticas de jugadores en formato array.

## Requisitos

- Go 1.18 o superior.
- Archivo `.rofl` válido.

## Licencia

MIT

---

**Desarrollado con amor por pointedsec.**