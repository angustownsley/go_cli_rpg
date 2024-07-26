package main

import ("fmt"
"strconv")

type combatStats interface {
	power() int
	resistance() int
	speedCalc() int
}

type armour struct {
	name string
	defence int
}

type weapon struct {
	name string
	power int
}

type player struct {
	name string
	maxHp int
	hp int
	strength int
	defence int
	speed int
	currentRoom int
	armour armour
	weapon weapon
}

func (p player) power() int {
	return p.strength + p.weapon.power
}
func (p player) speedCalc() int {
	return p.speed
}

func (p player) resistance() int {
	return p.defence + p.armour.defence
}

type enemy struct {
	name string
	description string
	hp int
	strength int
	defence int
	speed int
}

func newEnemy(name string, description string,  strength int, hp int, defence int, speed int) *enemy {
	e := enemy{name, description,hp, strength, defence,speed }
	return &e
}

func (e enemy) power() int {
	return e.strength
}
func (e enemy) speedCalc() int {
	return e.speed
}

func (e enemy) resistance() int {
	return e.defence
}

func newPlayer(name string) *player {
	p := player{name: name}
	p.hp = 20
	p.maxHp = 20
	p.strength = 5
	p.defence = 0
	p.speed = 5
	return &p
}

func calculateStats(s combatStats) [3]int {
	var statsArr [3]int

	statsArr[0] = s.power()
	statsArr[1] = s.resistance()
	statsArr[2] = s.speedCalc()

	return  statsArr
}

func combatRound(p *player, e *enemy) {
	var playerStats = calculateStats(p)
	var enemyStats = calculateStats(e)

	if(playerStats[2] >= enemyStats[2]) {
		fmt.Println(p.name + " strikes first!")
		e.hp -= playerStats[0]-enemyStats[1]
		if(e.hp > 0) {
			fmt.Println("The" + e.name + " strikes back!")
			p.hp -= enemyStats[0]-playerStats[1]
		}
		return
	}
	fmt.Println(e.name + " strikes first!")
	p.hp -= enemyStats[0]-playerStats[1]
	if(p.hp > 0) {
		fmt.Println(p.name + " strikes back!")
			e.hp -= playerStats[0]-enemyStats[1]
	}
}

func combat (p *player, e *enemy) {
	fmt.Println("There is a " + e.name)
	fmt.Println(e.description)
	var yn string
	fmt.Println("Would you like to fight the " + e.name + "? Y/N")
	fmt.Scanln(&yn)
	if(yn == "Y" || yn == "y") {
		for e.hp >= 0 {
		fmt.Println(p.name + " HP:" + strconv.Itoa(p.hp))
		fmt.Println(e.name + " HP:" + strconv.Itoa(e.hp))
		fmt.Println("Attack or Run")
		var atkOrRun string
		fmt.Scanln(&atkOrRun)
		if(atkOrRun == "run") {
			fmt.Println("You run away")
			return
		}
		combatRound(p, e)
		if(p.hp <= 0) {
			fmt.Println("You Died")
			return
		}
		}
		fmt.Println("You defeated the " + e.name)
		return
	}
}

func main() {
	var player1 = newPlayer("Angus")
	var weapon = weapon{name:"Doomhammer", power:20}
	player1.weapon = weapon

	var enemy1 = newEnemy("Bat", "It's a bat, flying around the room", 20, 600, 0, 8)

	combat(player1, enemy1)
}