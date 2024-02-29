// main.go

package main

import "fmt"

// Block represents a simple data structure
type Block struct {
	ID    int
	Data  string
	Color string
}

// DisplayAllBlocks prints information about all the blocks
func DisplayAllBlocks(blocks []Block) {
	fmt.Println("All Blocks:")
	for _, b := range blocks {
		fmt.Printf("ID: %d, Data: %s, Color: %s\n", b.ID, b.Data, b.Color)
	}
}

// NewBlock creates a new Block with given ID, Data, and Color
func NewBlock(id int, data, color string) Block {
	return Block{
		ID:    id,
		Data:  data,
		Color: color,
	}
}

// ModifyBlock modifies the given block's data and color
func ModifyBlock(b *Block, newData, newColor string) {
	b.Data = newData
	b.Color = newColor
	fmt.Println("\n===Modified===\n")
}

func main() {
	// Sample usage
	blocks := []Block{
		NewBlock(1, "Sabeen", "Black"),
		NewBlock(2, "Ali", "Yellow"),
	}

	DisplayAllBlocks(blocks)

	// Modify the second block
	ModifyBlock(&blocks[1], "Hamza", "Blue")

	DisplayAllBlocks(blocks)
}
