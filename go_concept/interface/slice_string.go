package main

import (
	"fmt"
	"strings"
)

func main() {
	metadataSlice := make([]string, 1)
	//metadataSlice  := []string{https://gateway.ipfscdn.io/ipfs/QmYKEvw2hRySeoDuLudifg8D7sz4vMH8qRDasGnw8kaELh/0 https://gateway.ipfscdn.io/ipfs/QmT85Bm5pBAYy5uM6t2mC3iyvSv68hXqwFvcS3BdwUnNuV/0 https://gateway.ipfscdn.io/ipfs/QmQDBwRQMZiMoFz4u528UPuUMZsdxtsXWPfgMqwrAHNuxi/0 https://gateway.ipfscdn.io/ipfs/QmS6KQyFz6KPEtDyWGiaSM5uH3w2Qs4HQt5CWFoGQbt3Fn/0]}
	metadataSlice = []string{" ", " ", " ", " ", "https://gateway.ipfscdn.io/ipfs/QmYKEvw2hRySeoDuLudifg8D7sz4vMH8qRDasGnw8kaELh/0", "https://gateway.ipfscdn.io/ipfs/QmT85Bm5pBAYy5uM6t2mC3iyvSv68hXqwFvcS3BdwUnNuV/0", "https://gateway.ipfscdn.io/ipfs/QmQDBwRQMZiMoFz4u528UPuUMZsdxtsXWPfgMqwrAHNuxi/0", "https://gateway.ipfscdn.io/ipfs/QmS6KQyFz6KPEtDyWGiaSM5uH3w2Qs4HQt5CWFoGQbt3Fn/0"}
	metadataSlice = metadataSlice[4:]
	fmt.Println("metadataSlice", metadataSlice)
	for i, v := range metadataSlice {
		metadataSlice[i] = strings.TrimSpace(v)
		fmt.Println(metadataSlice)
	}

	// metadataString ,,,,https://gateway.ipfscdn.io/ipfs/QmYKEvw2hRySeoDuLudifg8D7sz4vMH8qRDasGnw8kaELh/0,https://gateway.ipfscdn.io/ipfs/QmT85Bm5pBAYy5uM6t2mC3iyvSv68hXqwFvcS3BdwUnNuV/0,https://gateway.ipfscdn.io/ipfs/QmQDBwRQMZiMoFz4u528UPuUMZsdxtsXWPfgMqwrAHNuxi/0,https://gateway.ipfscdn.io/ipfs/QmS6KQyFz6KPEtDyWGiaSM5uH3w2Qs4HQt5CWFoGQbt3Fn/0
	// metadataString2 ,,,,https://gateway.ipfscdn.io/ipfs/QmYKEvw2hRySeoDuLudifg8D7sz4vMH8qRDasGnw8kaELh/0,https://gateway.ipfscdn.io/ipfs/QmT85Bm5pBAYy5uM6t2mC3iyvSv68hXqwFvcS3BdwUnNuV/0,https://gateway.ipfscdn.io/ipfs/QmQDBwRQMZiMoFz4u528UPuUMZsdxtsXWPfgMqwrAHNuxi/0,https://gateway.ipfscdn.io/ipfs/QmS6KQyFz6KPEtDyWGiaSM5uH3w2Qs4HQt5CWFoGQbt3Fn/0
	// metadataelements [,,,,https://gateway.ipfscdn.io/ipfs/QmYKEvw2hRySeoDuLudifg8D7sz4vMH8qRDasGnw8kaELh/0,https://gateway.ipfscdn.io/ipfs/QmT85Bm5pBAYy5uM6t2mC3iyvSv68hXqwFvcS3BdwUnNuV/0,https://gateway.ipfscdn.io/ipfs/QmQDBwRQMZiMoFz4u528UPuUMZsdxtsXWPfgMqwrAHNuxi/0,https://gateway.ipfscdn.io/ipfs/QmS6KQyFz6KPEtDyWGiaSM5uH3w2Qs4HQt5CWFoGQbt3Fn/0]
	// metadataelements2 [",,,,https://gateway.ipfscdn.io/ipfs/QmYKEvw2hRySeoDuLudifg8D7sz4vMH8qRDasGnw8kaELh/0,https://gateway.ipfscdn.io/ipfs/QmT85Bm5pBAYy5uM6t2mC3iyvSv68hXqwFvcS3BdwUnNuV/0,https://gateway.ipfscdn.io/ipfs/QmQDBwRQMZiMoFz4u528UPuUMZsdxtsXWPfgMqwrAHNuxi/0,https://gateway.ipfscdn.io/ipfs/QmS6KQyFz6KPEtDyWGiaSM5uH3w2Qs4HQt5CWFoGQbt3Fn/0"]

}
