package main
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:artist`
	Price float64 `json:price` // include tax
	Tax float32 `json:tax`
}

var albums = []album{
	{ID: "1", Title: "The hit's", Artist: "Offspring", Price: 25.05, Tax: 0.1},
	{ID: "2", Title: "Fearless", Artist: "Taylor Swift", Price: 23.14, Tax: 0.1},
	{ID: "3", Title: "Je Suis Africain", Artist: "Rachid Taha", Price: 18.88, Tax: 0.1},
}

// http get method
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// http post method
func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}
