package main

import (
	"fmt"
	"strings"
)

type dog struct {
	weight float64
}

const dogFoodPerOneKg = 2

func (d dog) calcFoodPerMonth() float64 {
	return dogFoodPerOneKg * d.weight
}

func (d dog) String() string {
	return "dog"
}

func (d dog) returnkWeight() float64 {
	return d.weight
}

type cat struct {
	weight float64
}

const catFoodPerOneKg = 7

func (c cat) calcFoodPerMonth() float64 {
	return catFoodPerOneKg * c.weight
}

func (c cat) String() string {
	return "cat"
}

func (c cat) returnkWeight() float64 {
	return c.weight
}

type cow struct {
	weight float64
}

const cowFoodPerOneKg = 25

func (c cow) calcFoodPerMonth() float64 {
	return cowFoodPerOneKg * c.weight
}

func (c cow) String() string {
	return "cow"
}

func (c cow) returnkWeight() float64 {
	return c.weight
}

type foodPerMonth interface {
	calcFoodPerMonth() float64
}

type checkAnimalWeight interface {
	returnkWeight() float64
}

type animalTotal interface {
	foodPerMonth
	fmt.Stringer
	checkAnimalWeight
}

func printFarmData(animals []animalTotal) float64 {
	var totalFood float64

	for _, animal := range animals {
		fmt.Printf("%v with weight %v kg will eat %v kg food per month.\n", strings.Title(animal.String()), animal.returnkWeight(), animal.calcFoodPerMonth())
		totalFood += animal.calcFoodPerMonth()
	}

	return totalFood
}

func main() {
	animals := []animalTotal{
		dog{weight: 10},
		cat{weight: 5},
		cow{weight: 100},
		cat{weight: 3},
		dog{weight: 13.7},
	}

	totalFoodRequired := printFarmData(animals)
	fmt.Printf("Total food required: %v kg.", totalFoodRequired)
}
