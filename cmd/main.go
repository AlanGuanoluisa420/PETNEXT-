
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// MODELOS DEL SISTEMA

type Cat struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Active bool   `json:"active"`
}

type Feeding struct {
	ID       int       `json:"id"`
	CatID    int       `json:"cat_id"`
	Food     string    `json:"food"`
	Quantity float64   `json:"quantity_grams"`
	Date     time.Time `json:"date"`
}

type Cleaning struct {
	ID     int       `json:"id"`
	Litter string    `json:"litter_box"`
	CatID  int       `json:"cat_id"`
	Date   time.Time `json:"date"`
	Notes  string    `json:"notes"`
}

// ALMACENAMIENTO DE INFORMACIÓN

var cats []Cat
var feedings []Feeding
var cleanings []Cleaning

var reader = bufio.NewReader(os.Stdin)

func loadData(file string, target interface{}) {
	data, err := os.ReadFile(file)

	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println("Error leyendo", file, ":", err)
		}
		return
	}

	if err := json.Unmarshal(data, target); err != nil {
		fmt.Println("Error procesando JSON:", err)
	}
}

func saveData(file string, data interface{}) {
	content, err := json.MarshalIndent(data, "", "  ")

	if err != nil {
		fmt.Println("Error convirtiendo datos:", err)
		return
	}

	err = os.WriteFile(file, content, 0644)

	if err != nil {
		fmt.Println("Error guardando información:", err)
	}
}

// FUNCIONES DE ENTRADA

func input(message string) string {
	fmt.Print(message)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func inputInt(message string) int {
	for {
		value := input(message)
		number, err := strconv.Atoi(value)

		if err == nil {
			return number
		}

		fmt.Println("Ingrese un número válido.")
	}
}

func inputFloat(message string) float64 {
	for {
		value := input(message)
		number, err := strconv.ParseFloat(value, 64)

		if err == nil && number > 0 {
			return number
		}

		fmt.Println("Ingrese una cantidad mayor que cero.")
	}
}

// GENERACIÓN DE IDENTIFICADORES

func nextCatID() int {
	maxID := 0

	for _, cat := range cats {
		if cat.ID > maxID {
			maxID = cat.ID
		}
	}

	return maxID + 1
}

func nextFeedingID() int {
	maxID := 0

	for _, record := range feedings {
		if record.ID > maxID {
			maxID = record.ID
		}
	}

	return maxID + 1
}

func nextCleaningID() int {
	maxID := 0

	for _, record := range cleanings {
		if record.ID > maxID {
			maxID = record.ID
		}
	}

	return maxID + 1
}

func findCat(id int) bool {
	for _, cat := range cats {
		if cat.ID == id && cat.Active {
			return true
		}
	}

	return false
}

// MODULO 1: GESTIÓN DE MASCOTAS

func registerCat() {
	name := input("Nombre de la mascota: ")

	if name == "" {
		fmt.Println("El nombre no puede estar vacío.")
		return
	}

	age := inputInt("Edad en meses: ")

	if age < 0 {
		fmt.Println("La edad no puede ser negativa.")
		return
	}

	cat := Cat{
		ID:     nextCatID(),
		Name:   name,
		Age:    age,
		Active: true,
	}

	cats = append(cats, cat)
	saveData("data/cats.json", cats)

	fmt.Println("Mascota registrada correctamente en PETNEXT.")
}

func listCats() {
	fmt.Println("\n--- MASCOTAS REGISTRADAS ---")

	for _, cat := range cats {
		if cat.Active {
			fmt.Printf(
				"ID: %d | Nombre: %s | Edad: %d meses\n",
				cat.ID, cat.Name, cat.Age,
			)
		}
	}
}

// MODULO 2: GESTIÓN DE ALIMENTACIÓN

func registerFeeding() {
	listCats()

	id := inputInt("ID de la mascota: ")

	if !findCat(id) {
		fmt.Println("La mascota no existe o está inactiva.")
		return
	}

	food := input("Tipo de alimento: ")

	if food == "" {
		fmt.Println("Debe indicar el tipo de alimento.")
		return
	}

	quantity := inputFloat("Cantidad en gramos: ")

	record := Feeding{
		ID:       nextFeedingID(),
		CatID:    id,
		Food:     food,
		Quantity: quantity,
		Date:     time.Now(),
	}

	feedings = append(feedings, record)
	saveData("data/feeding.json", feedings)

	fmt.Println("Alimentación registrada correctamente.")
}

// MODULO 3: LIMPIEZA DE ARENEROS

func registerCleaning() {
	listCats()

	id := inputInt("ID de la mascota: ")

	if !findCat(id) {
		fmt.Println("La mascota no existe o está inactiva.")
		return
	}

	litter := input("Identificación del arenero: ")

	if litter == "" {
		fmt.Println("Debe identificar el arenero.")
		return
	}

	notes := input("Observaciones: ")

	record := Cleaning{
		ID:     nextCleaningID(),
		CatID:  id,
		Litter: litter,
		Date:   time.Now(),
		Notes:  notes,
	}

	cleanings = append(cleanings, record)
	saveData("data/cleaning.json", cleanings)

	fmt.Println("Limpieza registrada correctamente.")
}

// MODULO 4: HISTORIAL DE ACTIVIDADES

func showHistory() {
	fmt.Println("\n--- HISTORIAL DE ALIMENTACIÓN ---")

	for _, record := range feedings {
		fmt.Printf(
			"ID: %d | Mascota: %d | Alimento: %s | Cantidad: %.2f g | Fecha: %s\n",
			record.ID,
			record.CatID,
			record.Food,
			record.Quantity,
			record.Date.Format("02/01/2006 15:04"),
		)
	}

	fmt.Println("\n--- HISTORIAL DE LIMPIEZA ---")

	for _, record := range cleanings {
		fmt.Printf(
			"ID: %d | Arenero: %s | Mascota: %d | Fecha: %s | Observaciones: %s\n",
			record.ID,
			record.Litter,
			record.CatID,
			record.Date.Format("02/01/2006 15:04"),
			record.Notes,
		)
	}
}

// MENÚ PRINCIPAL

func main() {
	err := os.MkdirAll("data", 0755)

	if err != nil {
		fmt.Println("No se pudo crear la carpeta de datos:", err)
		return
	}

	loadData("data/cats.json", &cats)
	loadData("data/feeding.json", &feedings)
	loadData("data/cleaning.json", &cleanings)

	for {
		fmt.Println("\n======================================")
		fmt.Println("              PETNEXT")
		fmt.Println("       PET MANAGEMENT SYSTEM")
		fmt.Println("======================================")
		fmt.Println("1. Registrar mascota")
		fmt.Println("2. Consultar mascotas")
		fmt.Println("3. Registrar alimentación")
		fmt.Println("4. Registrar limpieza de arenero")
		fmt.Println("5. Consultar historial")
		fmt.Println("6. Salir")
		fmt.Println("======================================")

		option := inputInt("Seleccione una opción: ")

		switch option {
		case 1:
			registerCat()

		case 2:
			listCats()

		case 3:
			registerFeeding()

		case 4:
			registerCleaning()

		case 5:
			showHistory()

		case 6:
			fmt.Println("Gracias por utilizar PETNEXT.")
			return

		default:
			fmt.Println("Opción no válida.")
		}
	}
}
