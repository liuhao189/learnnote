package main

import "fmt"

func main() {
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)
	//
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India"] = "新德里"

	for country, captialName := range countryCapitalMap {
		fmt.Println(country, "首都是", captialName)
	}

	captial, ok := countryCapitalMap["American"]

	if ok {
		fmt.Println("American's 首都是", captial)
	} else {
		fmt.Println("American's 首都不存在")
	}
	countryCapitalMap := map[string]string{"France": "Pairs", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
	fmt.Println("原始地图")
	for country, name := range countryCapitalMap {
		fmt.Println(country, "首都是", name)
	}
	delete(countryCapitalMap, "France")
	fmt.Println("Delete france")
	fmt.Println("Map is ")
	for country, name := range countryCapitalMap {
		fmt.Println(country, "首都是", name)
	}
}
