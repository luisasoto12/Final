package main

import (
	"fmt"
)

func main() {
	var nombre string
	var cantidadDesarrollador float64
	var cantidadContador float64
	var cantidadDiseñador float64
	var cantidadJefeProy float64
	var cantidadAnalistaDatos float64
	var cantidadHorasDesarrollador float64
	var cantidadHorasContador float64
	var cantidadHorasDiseñador float64
	var cantidadHorasJefeProy float64
	var cantidadHorasAnalistaDatos float64
	var costoHorDes float64
	var costoHorContador float64
	var costoHorDiseñador float64
	var costoHorJefeProy float64
	var costoHorAnalistaDatos float64
	var total float64
	var costoPlataforma float64
	var respuesta string
	var capitalizacion float64
	var oficina float64
	var IVA float64 = 0.19
	var RETEICA float64 = 0.01
	var RenF float64 = 0.11
	var Riesgos float64 = 0.16
	var costosPlataforma map[string]float64

	var costosInfraestructura = map[string]map[string]float64{
		"Azure": {
			"computadores": 2000000,
			"nube":         200000,
			"licencia":     100000,
		},
		"AWS": {
			"computadores": 2000000,
			"nube":         180000,
			"licencia":     120000,
		},
		"Otro": {
			"computadores": 2000000,
			"nube":         250000,
			"licencia":     80000,
		},
	}

	// Ingresar datos
	fmt.Print("Ingrese el nombre del Software a cotizar: ")
	fmt.Scanln(&nombre)

	fmt.Print("Ingrese la cantidad de desarrolladores: ")
	fmt.Scanln(&cantidadDesarrollador)

	fmt.Print("Ingrese la cantidad de contadores: ")
	fmt.Scanln(&cantidadContador)

	fmt.Print("Ingrese la cantidad de diseñadores: ")
	fmt.Scanln(&cantidadDiseñador)

	fmt.Print("Ingrese la cantidad de Jefes de proyecto: ")
	fmt.Scanln(&cantidadJefeProy)

	fmt.Print("Ingrese la cantidad de analista de datos: ")
	fmt.Scanln(&cantidadAnalistaDatos)

	fmt.Print("Ingrese la cantidad de horas a trabajar de los desarrolladores: ")
	fmt.Scanln(&cantidadHorasDesarrollador)

	fmt.Print("Ingrese la cantidad de horas a trabajar de los contadores: ")
	fmt.Scanln(&cantidadHorasContador)

	fmt.Print("Ingrese la cantidad de horas a trabajar de los diseñadores: ")
	fmt.Scanln(&cantidadHorasDiseñador)

	fmt.Print("Ingrese la cantidad de horas a trabajar de los jefes de proyecto: ")
	fmt.Scanln(&cantidadHorasJefeProy)

	fmt.Print("Ingrese la cantidad de horas a trabajar de los analista de datos: ")
	fmt.Scanln(&cantidadHorasAnalistaDatos)

	fmt.Print("Ingrese el costo de hora de los desarrolladores: ")
	fmt.Scanln(&costoHorDes)

	fmt.Print("Ingrese el costo de hora de los contadores: ")
	fmt.Scanln(&costoHorContador)

	fmt.Print("Ingrese el costo de hora de los diseñadores: ")
	fmt.Scanln(&costoHorDiseñador)

	fmt.Print("Ingrese el costo de hora de los jefes de proyecto: ")
	fmt.Scanln(&costoHorJefeProy)

	fmt.Print("Ingrese el costo de hora de los analista de datos: ")
	fmt.Scanln(&costoHorAnalistaDatos)

	// Calcular esfuerzo
	esfuerzo := (cantidadDesarrollador * cantidadHorasDesarrollador * costoHorDes) +
		(cantidadContador * cantidadHorasContador * costoHorContador) +
		(cantidadDiseñador * cantidadHorasDiseñador * costoHorDiseñador) +
		(cantidadJefeProy * cantidadHorasJefeProy * costoHorJefeProy) +
		(cantidadAnalistaDatos * cantidadHorasAnalistaDatos * costoHorAnalistaDatos)

	// Calcular total de la cotización del esfuerzo humano realizado

	total = esfuerzo

	// Seleccionar plataforma
	var plataforma string
	fmt.Println("Elige una opción para la plataforma:")
	for key := range costosInfraestructura {
		fmt.Printf("- %s\n", key)
	}

	fmt.Print("Opción seleccionada: ")
	fmt.Scanln(&plataforma)

	//Definimos el costo de la infraestructura en la variable plataforma.

	costosPlataforma = costosInfraestructura[plataforma]

	for key, costo := range costosPlataforma {
		total += costo
		costoPlataforma += costo
		fmt.Printf("Costo de %s: $%.2f\n", key, costo)
	}

	//Calcular el total de los costos (plataforma + esfuerzo humano)
	TotalCostos := esfuerzo + costoPlataforma

	// Imprimir los costos de la mano de obra y plataforma
	fmt.Println("--------------------------------")
	fmt.Printf("Para producir el Software %s, el costo en mano de obra es: $%.2f\n", nombre, esfuerzo)
	fmt.Println("--------------------------------")
	fmt.Printf("El costo total de la plataforma seleccionada es: $%.2f\n", costoPlataforma)
	fmt.Println("--------------------------------")
	fmt.Printf("\nEl total de la cotización para el Software %s con respecto a la infraestructura es: $%.2f\n", nombre, total)
	fmt.Println("--------------------------------")
	fmt.Println("El costo del programa de software es $", TotalCostos)
	fmt.Println("--------------------------------")

	// Gastos de oficina
	fmt.Println("¿Desea disponer de oficina?(si/no)")
	fmt.Scanln(&respuesta)

	if respuesta == "si" {
		oficina = 5000000.00
		fmt.Println("--------------------------------")
		fmt.Println("Se añadirá a la cuenta $5000000 que es lo que cuesta el canon de oficina")
		fmt.Println("--------------------------------")
	} else if respuesta == "no" {
		oficina = 2000000.00
		fmt.Println("Se añadirá a la cuenta $2000000 que es lo que cuesta el teletrabajo")
		fmt.Println("--------------------------------")
	} else {
		fmt.Println("No comprendo tu mensaje")
		fmt.Println("--------------------------------")
	}

	fmt.Println("¿Cuánto es su capital?: ")
	fmt.Scanln(&capitalizacion)

	// Capitalización
	var porcentaje float64
	if capitalizacion >= 0 && capitalizacion <= 100000000 {
		porcentaje = 0.20
	} else if capitalizacion > 100000000 && capitalizacion < 500000000 {
		porcentaje = 0.15
	} else if capitalizacion >= 500000000 {
		porcentaje = 0.1
	}

	//
	fmt.Println("--------------------------------")
	fmt.Println("El porcentaje que se debe capitalizar es ", porcentaje, "%, que es el porcentaje perteneciente a las ganancias que se incluye en el capital")
	fmt.Println("--------------------------------")

	// Impuestos
	fmt.Println("El valor de los gastos es $", oficina, " que incluye los gastos administrativos")
	fmt.Println("--------------------------------")

	Total_Costos_Gastos := TotalCostos + oficina

	aplicaRiesgos := (Total_Costos_Gastos * Riesgos)
	Total_Riesgos := (Total_Costos_Gastos + aplicaRiesgos)
	fmt.Println("El monto dispuesto para los riesgos es de $", aplicaRiesgos, " y sumándolo al total es $", Total_Riesgos)
	fmt.Println("--------------------------------")

	aplicaREF := Total_Riesgos * RenF
	aplicaRETEICA := aplicaREF * RETEICA
	TotalRetencion := TotalCostos + aplicaRETEICA + aplicaREF
	aplicaIVA := TotalRetencion * IVA
	Total_definitivo := TotalRetencion + aplicaIVA

	fmt.Println("La retención en la fuente es: $", aplicaREF)
	fmt.Println("--------------------------------")
	fmt.Println("El RETEICA en el total es: $", aplicaRETEICA)
	fmt.Println("--------------------------------")
	fmt.Println("El IVA en el total es: $", aplicaIVA)
	fmt.Println("--------------------------------")
	fmt.Println("El valor final del software es $", Total_definitivo)
	fmt.Println("--------------------------------")
}
