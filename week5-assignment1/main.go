package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

// Movie struct
type Movie struct {
    ID       string  `json:"id"`
    Title    string  `json:"title"`
    Director string  `json:"director"`
    Year     int     `json:"year"`
    Rating   float64 `json:"rating"` // คะแนนรีวิว
}

// In-memory movie database
var movies = []Movie{
    {ID: "1", Title: "Inception", Director: "Christopher Nolan", Year: 2010, Rating: 8.8},
    {ID: "2", Title: "The Matrix", Director: "The Wachowskis", Year: 1999, Rating: 8.7},
    {ID: "3", Title: "Interstellar", Director: "Christopher Nolan", Year: 2014, Rating: 8.6},
    {ID: "4", Title: "Transformers", Director: "Michael Bay", Year: 2007, Rating: 7.1},
}

// Handler to get movies, with optional rating filter
func getMovies(c *gin.Context) {
    ratingQuery := c.Query("rating")

    if ratingQuery != "" {
        filter := []Movie{}
        for _, movie := range movies {
            // ใช้ fmt.Sprintf เพื่อเปรียบเทียบ float เป็น string
            if fmt.Sprintf("%.1f", movie.Rating) == ratingQuery {
                filter = append(filter, movie)
            }
        }
        c.JSON(http.StatusOK, filter)
        return
    }

    c.JSON(http.StatusOK, movies)
}

func main() {
    r := gin.Default()

    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "healthy"})
    })

    api := r.Group("/api/v1")
    {
        api.GET("/movies", getMovies)
    }

    r.Run(":8080")
}
