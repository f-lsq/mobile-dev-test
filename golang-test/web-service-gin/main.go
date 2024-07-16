package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album
type album struct {
	ID     string  `json: "id"`
	Title  string  `json: "title"`
	Artist string  `json: "artist"`
	Price  float64 `json: "price"`
}

// albums slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()                 // initialises a new Gin router
	router.GET("/albums", getAlbums)        // app.get("/albums", (req, res) => {getAlbums(req, res)})
	router.POST("/albums", postAlbums)      // app.post("/albums", (req, res) => {postAlbums(req, res)})
	router.GET("/albums/:id", getAlbumByID) // app.get("/albums/:id", (req, res) => {getAlbumByID(req, res)})

	router.Run("localhost:8080") // similar to app.listen(8080)
}

// getAlbums reponds with the list of all albums as JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
	// .IndentedJSON serialises the struct into JSON and add it to the response
	// similar to res.status(200).json(albums) in Express?
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum
	// is this same as let newAlbum = req.body?
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
	// res.status(201).json(newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id") // similar to let id = req.params.id

	// Loop over the list of albums, looking for an album whose ID value matches the parameter
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	// gin.H is a shortcut to create a map with "string" keys and "interface{}" values (map[string]interface{})
	// -> Used to send JSON responses with key-value pairs, where the keys are strings and values is of any type

	// similar to res.status(404).json({
	// "message": "album not found"
	// })
}
