package assetman

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"image"
	"image/draw"
	_ "image/png"
	"io"
	"log"
	"net/http"

	"github.com/laykku/tower/engine"
)

func LoadShader(path string) (code string) {
	data := loadFile(path)
	return string(data)
}

func LoadTexture(path string) ([]uint8, int32, int32) {
	data := loadFile(path)
	reader := bytes.NewReader(data)
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatalln(err)
	}

	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)

	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	return rgba.Pix, int32(bounds.Size().X), int32(bounds.Size().Y)
}

func LoadMesh(path string) *engine.MeshData {
	data := loadFile(path)
	reader := bytes.NewReader(data)

	name, _ := readString(reader)
	fmt.Println("Loaded mesh:", name)

	var vertexCount int32
	if err := binary.Read(reader, binary.LittleEndian, &vertexCount); err != nil {
		log.Fatal("Error loading mesh:", err)
	}

	positions := make([]float32, vertexCount*3)
	if err := binary.Read(reader, binary.LittleEndian, &positions); err != nil {
		log.Fatalf("error reading vertex data")
	}

	var uv0Count int32
	if err := binary.Read(reader, binary.LittleEndian, &uv0Count); err != nil {
		log.Fatal("Error loading mesh:", err)
	}

	uv0 := make([]float32, uv0Count*2)
	if err := binary.Read(reader, binary.LittleEndian, &uv0); err != nil {
		log.Fatalf("error reading vertex data")
	}

	var indexCount int32
	if err := binary.Read(reader, binary.LittleEndian, &indexCount); err != nil {
		log.Fatal("Error loading mesh:", err)
	}

	indices := make([]uint32, indexCount)
	if err := binary.Read(reader, binary.LittleEndian, &indices); err != nil {
		log.Fatal("Error loading mesh:", err)
	}

	meshData := &engine.MeshData{
		Positions: positions,
		Uv0:       uv0,
		Triangles: indices,
	}

	return meshData
}

func readString(file io.Reader) (string, error) {
	var strLen int32
	if err := binary.Read(file, binary.LittleEndian, &strLen); err != nil {
		return "", err
	}

	strBytes := make([]byte, strLen)
	if _, err := file.Read(strBytes); err != nil {
		return "", err
	}
	return string(strBytes), nil
}

func loadFile(path string) []byte {

	if bytes := GetEmbeddedResource(path); bytes != nil {
		return bytes
	}

	resp, err := http.Get(fmt.Sprintf("assets/%s", path))
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return data
}
