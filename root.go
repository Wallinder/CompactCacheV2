package compact

type RootFile struct {
	CurrentVersion     float64       `json:"currentVersion"`
	Name               string        `json:"name"`
	CopyrightText      string        `json:"copyrightText"`
	Capabilities       string        `json:"capabilities"`
	Type               string        `json:"type"`
	TileMap            string        `json:"tileMap"`
	DefaultStyles      string        `json:"defaultStyles"`
	Tiles              []string      `json:"tiles"`
	ExportTilesAllowed bool          `json:"exportTilesAllowed"`
	InitialExtent      InitialExtent `json:"initialExtent"`
	FullExtent         FullExtent    `json:"fullExtent"`
	MinScale           int           `json:"minScale"`
	MaxScale           int           `json:"maxScale"`
	TileInfo           TileInfo      `json:"tileInfo"`
	Maxzoom            int           `json:"maxzoom"`
	MinLOD             int           `json:"minLOD"`
	MaxLOD             int           `json:"maxLOD"`
	ResourceInfo       ResourceInfo  `json:"resourceInfo"`
}

type InitialExtent struct {
	Xmin             int             `json:"xmin"`
	Ymin             int             `json:"ymin"`
	Xmax             int             `json:"xmax"`
	Ymax             int             `json:"ymax"`
	SpatialReference InitialExtentSR `json:"spatialReference"`
}
type InitialExtentSR struct {
	Wkid       int `json:"wkid"`
	LatestWkid int `json:"latestWkid"`
}

type FullExtent struct {
	Xmin             int          `json:"xmin"`
	Ymin             int          `json:"ymin"`
	Xmax             int          `json:"xmax"`
	Ymax             int          `json:"ymax"`
	SpatialReference FullExtentSR `json:"spatialReference"`
}
type FullExtentSR struct {
	Wkid       int `json:"wkid"`
	LatestWkid int `json:"latestWkid"`
}

type TileInfo struct {
	Rows             int            `json:"rows"`
	Cols             int            `json:"cols"`
	Dpi              int            `json:"dpi"`
	Format           string         `json:"format"`
	Origin           TileInfoOrigin `json:"origin"`
	SpatialReference TileInfoSR     `json:"spatialReference"`
	Lods             []TileInfoLods `json:"lods"`
}

type TileInfoOrigin struct {
	X int `json:"x"`
	Y int `json:"y"`
}
type TileInfoSR struct {
	Wkid       int `json:"wkid"`
	LatestWkid int `json:"latestWkid"`
}

type TileInfoLods struct {
	Level      int     `json:"level"`
	Resolution float64 `json:"resolution"`
	Scale      int     `json:"scale"`
}

type ResourceInfo struct {
	StyleVersion    int       `json:"styleVersion"`
	TileCompression string    `json:"tileCompression"`
	CacheInfo       CacheInfo `json:"cacheInfo"`
}

type CacheInfo struct {
	StorageInfo StorageInfo `json:"storageInfo"`
}
type StorageInfo struct {
	PacketSize    int    `json:"packetSize"`
	StorageFormat string `json:"storageFormat"`
}
