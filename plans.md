**Player**

```
Name: string
Health: int 
Strength: int 
CurrentRoom: int? 
Armour: Armour struct
Weapon: Weapon struct
Defence: int
Speed: int
```

**Enemy**

```
Name: string
Description: string
Health: int
Strength: int
Defence: int
Speed: int
```

**Room**

```
Doors: struct {
    N:bool,
    S:bool,
    W:bool,
    E:bool
}
Destinations: struct {
    N:Layout[index],
    S:Layout[index],
    W:Layout[index],
    E:Layout[index]
}
Enemies: slice
Treasure: (Weapon struct  || Armour struct )
```

**Layout**

```
Rooms: slice[]Rooms
```