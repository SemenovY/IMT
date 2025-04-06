package main

import (
	"errors"
	"fmt"
	"math"
)

const HeightToIMT = 100
const IMTPower = 2

func main() {
	fmt.Println("__ Калькулятор индекса массы тела __")
	for {
		userHeight, userKg := getUserInput()
		IMT, err := calculateIMT(userHeight, userKg)
		if err != nil {
			fmt.Println(err)
			continue
		}
		outputResult(IMT)
		isRepeatCalculation := checkRepeatCalculation()
		if !isRepeatCalculation {
			break
		}
	}
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Println(result)
	switch {
	case imt < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case imt < 18.5:
		fmt.Println("У вас норма")
	case imt < 25:
		fmt.Println("У вас избыточная масса")
	case imt < 30:
		fmt.Println("У вас 1-я степень ожирения")
	case imt < 35:
		fmt.Println("У вас 2-я степень ожирения")
	default:
		fmt.Println("У вас 3-я степень ожирения")
	}
}

func calculateIMT(userHeight float64, userKg float64) (float64, error) {
	if userHeight <= 0 || userKg <= 0 {
		return 0, errors.New("NO_PARAMS_ERROR")
	}
	IMT := userKg / math.Pow(userHeight/HeightToIMT, IMTPower)
	return IMT, nil
}

func getUserInput() (float64, float64) {
	var userHeight float64 // height in santimeters
	var userKg float64     // weight in kilograms
	fmt.Println("Введите свой рост в сантиментрах: ")
	fmt.Scan(&userHeight)
	fmt.Println("Введите свой вес: ")
	fmt.Scan(&userKg)
	return userHeight, userKg
}

func checkRepeatCalculation() bool {
	var userChoice string
	fmt.Println("Вы хотите сделать ещё расчёт (y/n): ")
	fmt.Scan(&userChoice)
	if userChoice == "y" || userChoice == "Y" {
		return true
	}
	return false
}
