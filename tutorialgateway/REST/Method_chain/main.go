package main

import "fmt"

//
// 🏠 STRUCT DEFINITIONS
//

// House is the main object (like the whole house)
type House struct {
	Door Door
}

// Door is part of the house
type Door struct {
	house *House // can go back to house if needed
}

// Open the door and return a Room
func (d Door) Open() Room {
	fmt.Println("🚪 Opening the front door...")
	return Room{door: &d}
}

// Room is inside the house
type Room struct {
	door *Door
}

// Enter the room and return a Desk
func (r Room) Enter() Desk {
	fmt.Println("🚶 Entering the living room...")
	return Desk{room: &r}
}

// Desk is in the room
type Desk struct {
	room *Room
}

// Open the drawer and return the Drawer
func (d Desk) Drawer() Drawer {
	fmt.Println("🗄️  Opening the drawer on the desk...")
	return Drawer{desk: &d}
}

// Drawer holds the note
type Drawer struct {
	desk *Desk
}

// Open drawer and return Note
func (dr Drawer) Open() Note {
	fmt.Println("📂 Drawer is now open...")
	return Note{drawer: &dr, name: "Alice"} // the note has "Alice" written
}

// Note has the friend's name
type Note struct {
	drawer *Drawer
	name   string
}

// Read the note and return the name
func (n Note) Read() string {
	fmt.Println("📖 Reading the note...")
	return n.name
}

// 🚀 MAIN FUNCTION
func main() {
	// 1️⃣ Create a house with a door
	house := House{
		Door: Door{house: &House{}},
	}

	// 2️⃣ The full method chain (compact version)
	friendName := house.Door.Open().Enter().Drawer().Open().Read()

	// 3️⃣ Print result
	fmt.Println("👋 Friend's name:", friendName)
}

/*
| Step | Code         | Returns  | Meaning                          |
| ---- | ------------ | -------- | -------------------------------- |
| 1    | `house.Door` | `Door`   | Access the door of the house     |
| 2    | `.Open()`    | `Room`   | Open the door and step inside    |
| 3    | `.Enter()`   | `Desk`   | Enter the living room            |
| 4    | `.Drawer()`  | `Drawer` | Open the desk drawer             |
| 5    | `.Open()`    | `Note`   | Open the drawer revealing a note |
| 6    | `.Read()`    | `string` | Read the note and get the name   |

*/
