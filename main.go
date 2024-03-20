package main

import (
	"fmt"
	"modul1/model"
)

func main() {
	s := model.NewStudent("pera", "peric", "100")
	s.PredstaviSe()
	s.NoviBrojIndeksa("101010")
	fmt.Println(s)
	s.PredstaviSe()

	r := model.NewRadnik("milan", "tucakov", "radnik")
	r.PredstaviSe()

	Ispis(s)
	Ispis(r)
}

func Ispis(o model.Osoba) {

	o.PredstaviSe()

}
