package main

/*
//1.Football-players test

type Player struct {
	Name		string
	Identifier 	string
}

func main(){
    players := []string{"John Smith", "David Johnson", "Michael Garcia", "Sarah Williams", "Daniel Martinez", "Emily Brown", "James Rodriguez", "Sophia Lee", "Lucas Oliveira", "Olivia Taylor", "Mateo Hernandez", "Emma Wilson", "Gabriel Silva"}

	fmt.Println("Lista jugadores 2024")
	imprimirLista(0,players)

	//registro de nuevos jugadores
	newPlayer := []string{"New Player 1", "New Player 2", "New Player 3"}

	//ELiminar los 7 primeros jugadores
	var playersRetirados int  =  7
	players = players[playersRetirados:]
	fmt.Println("=======================")
	fmt.Println("Lista jugadores Final 2024")
	imprimirLista(playersRetirados,players)

	//agregar
	for _, name := range newPlayer {
		players = append(players, name)
	}

	//mostrar la lista 2025
	fmt.Println("=======================")
	fmt.Println("Lista jugadores 2025: ")
	imprimirLista(playersRetirados,players)


}
//imprimir lista

func imprimirLista(num int,players []string){
	for i, player := range players{
		fmt.Printf("%d. %s\n", i+num+1001, sanitizeName(player))
	}
}


//funcionEliminarCaracteresEspeciales

func sanitizeName(name string) string {
	//eliminar caracteres especiales
	return strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r == ' ' {
			return r
		}
		return -1

	}, name)
}
*/

/*
//===========================================
//2. Fizz-buzz test
// func main(){
//     var n int
//     fmt.Println("Ingrese el valor de n:")
//     fmt.Scan(&n)
//     fizzBuzz(n)

// }
// func fizzBuzz(n int){
//     for i := 1; i <= n ; i++ {

//         if(i%3 == 0 && i%5 == 0){
//             fmt.Println("FizzBuzz")
//         } else if(i%3 == 0){
//             fmt.Println("Fizz")
//         } else if(i%5 == 0){
//             fmt.Println("Buzz")

//         }else{
//             fmt.Println(i)
//         }

// 	}
// }

//=============================================================
*/

//===========================================
//3.Menu-packaging test
func main(){

}