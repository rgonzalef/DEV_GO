package main

import (
	"fmt"
	"math/rand"
)

//CHARACTERS

type Character struct {
	Name		string
	Race 		Race
	Faction 	Faction
	Level		int
	Health		int
	Attack		int
	Shield		int
	Experience	int

}

type Race string
const(
	Human	Race = "Humano"
	Elfo	Race = "Elfo"
	Enano 	Race = "Enano"
	Goblin	Race = "Gablin"
	Orco 	Race = "Orco"
	Trol	Race = "Trol"
)


type Faction string
const (
	Alliance	Faction = "Alianza"
	Horde		Faction = "Horda"
	Neutral		Faction = "Neutral"
	Belligerent	Faction = "Combativo"

)

type Characters struct {
	characters []Character
}

func (ch *Characters) AddCharacter(name string, race Race, faction Faction, health, attack, shield, experience int ){
	newCharacter := Character{
		Name: name,
		Race: race,
		Faction: faction,
		Level: 1,
		Health: health,
		Attack: attack,
		Shield: shield,
		Experience: experience,
	}
	ch.characters = append(ch.characters, newCharacter)

}

func (ch *Characters) showFaction(personajes *Characters , choice int) []Character {
	
	switch choice {
	case 1:
		fmt.Println("Personajes de La Alianza")
		showPersonajes := personajes.findCharacters(Alliance)
		return showPersonajes

	case 2:
		fmt.Println("Personajes de La Horda")
		showPersonajes := personajes.findCharacters(Horde)
		return showPersonajes
	default:
		fmt.Println("Selección no válida")
		return nil
	}
}

func (ch *Characters) findCharacters(faction Faction) []Character {
	var foundCharacters []Character
	for _, fCharacter := range ch.characters{
		if fCharacter.Faction == faction{
			foundCharacters = append(foundCharacters, fCharacter)
		}
	}
	return foundCharacters

}

func showCharacters(personajes []Character){
	for i, char := range personajes{
		fmt.Println(i+1)
		fmt.Printf(" Nombre: %s \n Raza: %s \n Nivel: %d \n Vida: %d \n Ataque: %d \n Escudo: %d \n ", char.Name, char.Race, char.Level, char.Health, char.Attack, char.Shield )
		fmt.Println("_______________________")
	}

} 

//MISSIONS

type Mission struct {
	MissionName			string
	Place				Stage
	StageCharacter		Character
	Level				int
	Reward				int
}


type Stage string
const (
	Home	Stage = "Casa"
	Forest	Stage = "Bosque"
	Castle	Stage = "Castillo"
	Swamp	Stage = "Pantano"
)
type Missions struct {
	missions []Mission
}

func (mi *Missions) AddMission(name string, place Stage, stChacarter Character, level int, reward int) []Mission {
	newMission := Mission {
		MissionName: name,
		Place: place,
		StageCharacter: stChacarter,
		Level: level,
		Reward:	reward,
	}
	mi.missions = append(mi.missions, newMission)
	return mi.missions
}

var Ogro = Character{
	Name: "Ogro",
	Level: 1, 
	Health: 80,
	Attack: 7,
	Shield:	60,
	Faction: Horde ,
}
var Anciano = Character{
	Name: "Anciano",
	Level: 1, 
	Health: 50,
	Attack: 4,
	Shield:	30,
	Faction: Neutral,
}

var Caballero = Character{
	Name: "Caballero",
	Level: 1, 
	Health: 70,
	Attack: 6,
	Shield:	70,
	Faction: Alliance,
}

//FUNCION CLASH
func clash(characterSelected Character, missions []Mission) {
	for _, mission := range missions {
		missionChar := mission.StageCharacter
		
		if missionChar.Faction == Neutral {
			fmt.Printf("El %s habla contigo y te da una recompensa \n",missionChar.Name)
			
		} else if (characterSelected.Faction == missionChar.Faction){
			fmt.Printf("Ambos personajes conversan\n")
			

		} else {
			fmt.Printf("Enfrentamiento entre %s y %s\n", characterSelected.Name, missionChar.Name)
			for characterSelected.Health > 0 && missionChar.Health > 0 {
				// Ataque del personaje seleccionado
				damage := missionChar.Shield - characterSelected.Attack*rand.Intn(10)
				if damage > 0 {
					missionChar.Health -= damage
					fmt.Printf(" Has generado %d de daño a %s, su salud esta en %d \n", damage, missionChar.Name, missionChar.Health)
				}
				if missionChar.Health <= 0 {
					fmt.Printf("%s ha ganado!\n", characterSelected.Name)
					break
				}
				// Ataque del personaje de la misión
				damage = characterSelected.Shield -missionChar.Attack*rand.Intn(10)
				if damage > 0 {
					characterSelected.Health -= damage
					fmt.Printf(" %s te Ha generado %d de daño, tu salud esta en %d \n", missionChar.Name, damage, characterSelected.Health )
				}
				if characterSelected.Health <= 0 {
					fmt.Printf("%s ha ganado!\n", missionChar.Name)
				}
			}
		}
		
	}
	
}


func main(){
	
	personajes := &Characters{}
	missions := &Missions{}

	personajes.AddCharacter("Arthas", Human, Alliance, 100, 8, 80, 1)
	personajes.AddCharacter("Sylvanas", Elfo, Alliance, 90, 7, 60, 1)
	personajes.AddCharacter("Thrall", Orco, Horde, 110, 9, 80, 1)
	personajes.AddCharacter("Garona", Orco, Horde, 90, 6, 60, 1 )

	//fmt.Println(personajes)
	var choiceFaction int
	var choiceCharacter int
	fmt.Println()
	fmt.Println("Bienvenido al juego cli-WOW")
	fmt.Println("===========================")
	fmt.Println()
	fmt.Println("Escoge tu bando: \n 1. La Alianza \n 2. La Horda ")
	fmt.Scan(&choiceFaction)
	fmt.Println()
	
	//funcion imprime personajes segun selección
	chosenCharacters:= personajes.showFaction(personajes, choiceFaction)
	showCharacters(chosenCharacters)
	
	fmt.Println()
	fmt.Println("Escoge tu personaje: ")
	fmt.Scan(&choiceCharacter)
	fmt.Println()
	fmt.Printf("El personaje seleccionado es %s \n", chosenCharacters[choiceCharacter-1].Name)

	//carga en una variable el personaje seleccionado
	characterSelected :=  chosenCharacters[choiceCharacter-1]
	//fmt.Println(characterSelected)
	
 	//mission1 := missions.AddMission("Mision 1", Home, Anciano, 1, 5)
	mission2 := missions.AddMission("Mision 2", Forest, Ogro, 1, 5)
	mission3 := missions.AddMission("Mision 3", Castle, Caballero, 1, 5)
	//fmt.Println(missions)
	
	clash(characterSelected, mission2)
	clash(characterSelected, mission3)


}