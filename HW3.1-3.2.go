package main

import "fmt"
	


type  passengerCar struct{
	Mark string
	ReleaseYear int
	BodyVolume float64
	run bool
	window bool
	workload int
}

type  truck struct{
	Mark string
	ReleaseYear int
	BodyVolume float64
	run bool
	window bool
	workload int
}

func main()  {
	batcar := passengerCar{Mark : "batmobile", ReleaseYear : 2004, BodyVolume : 3.86, run : true, window : false, workload : 0}

	batcar.ReleaseYear = 2008 //изменение года создания

	kamaz := truck{Mark : "kamaz", ReleaseYear : 2000, BodyVolume : 10.85, run : true, window : true, workload : 0}
	
	if batcar.Mark == kamaz.Mark{ //сравнение машин
	fmt.Println("Это машины одной марки")
	} else { 
	fmt.Println("Это разные машины")
	}

	fmt.Println("Марка:" ,batcar.Mark, "Год выпуска:" ,batcar.ReleaseYear, "Объем кузова:" ,batcar.BodyVolume)
	fmt.Println("Марка:" ,kamaz.Mark, "Год выпуска:" ,kamaz.ReleaseYear, "Объем кузова:" ,kamaz.BodyVolume)
}