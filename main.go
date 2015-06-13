package main

import (
	"flag"
	log "github.com/Sirupsen/logrus"
	"github.com/lawrencecraft/terrainmodel/drawer"
	"github.com/lawrencecraft/terrainmodel/generator"
	"os"
)

func main() {
	log.SetLevel(log.DebugLevel)
	roughness := flag.Float64("roughness", 1.0, "Between 0 and 1. Roughness of the generated image")
	sizeX := flag.Int("x", 1025, "X value of the generated image")
	sizeY := flag.Int("y", 1025, "Y value of the generated image")
	path := flag.String("path", "map.png", "Output path of the png")
	flag.Parse()

	file, err := os.Create(*path)
	if err != nil {
		log.Fatal("Cannot open ", *path, ", error: ", err)
	}

	gen := generator.NewDiamondSquareGenerator(float32(*roughness), *sizeX, *sizeY)
	draw := drawer.NewPngDrawer(file)

	t, err := gen.Generate()
	if err != nil {
		log.Fatal("Error during generation: ", err)
	}

	draw.Draw(t)
}
