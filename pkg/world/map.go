package world

type MapTile struct {
	tileType int
}

type LocalMap struct {
	Tile [][][]MapTile
}

func (lMap *LocalMap) GetTile(x, y, z int) MapTile {
	return lMap.Tile[x][y][z]
}

func (lMap *LocalMap) GenerateLocalMap(xx, yy, zz int) {
	for x := 0; x < xx; x++ {
		lMap.Tile = append(lMap.Tile, [][]MapTile{})
		for y := 0; y < yy; y++ {
			lMap.Tile[x] = append(lMap.Tile[x], []MapTile{})
			for z := 0; z < zz; z++ {
				lMap.Tile[x][y] = append(lMap.Tile[x][y], MapTile{tileType: 0})
			}
		}
	}
}
