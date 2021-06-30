package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Bike struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

func GetBikes(c *gin.Context) {
	res := []Bike{{
		Id:          1,
		Title:       "Estilo de vida",
		Url:         "https://images.pexels.com/photos/7512060/pexels-photo-7512060.jpeg?auto=compress&cs=tinysrgb&dpr=2&h=650&w=940",
		Description: "Sabemos que tu moto es tu estilo de vida, por eso cuidamos de ti dantote el mejor servicio y la mejor calidad.",
	},
		{
			Id:          2,
			Title:       "Experiencias",
			Url:         "https://images.pexels.com/photos/5542036/pexels-photo-5542036.jpeg?auto=compress&cs=tinysrgb&dpr=2&h=650&w=940",
			Description: "Porque una moto te lleva a donde quieras, a disfrutar tu tiempo en las cosas que mas quieres y a vivir experiencias que te haran vibrar.",
		},
		{
			Id:          3,
			Title:       "Disfruta",
			Url:         "https://images.pexels.com/photos/7564495/pexels-photo-7564495.jpeg?auto=compress&cs=tinysrgb&dpr=2&h=650&w=940",
			Description: "El viaje no debe ser incomodo. Por más largo que sea el viaje, nuestras motos te haran disfrutar del camino, ya que para nosotros la comodidad no es un lujo es una necesidad",
		},
		{
			Id:          4,
			Title:       "Pasíon",
			Url:         "https://images.pexels.com/photos/5078778/pexels-photo-5078778.jpeg?auto=compress&cs=tinysrgb&dpr=2&h=650&w=940",
			Description: "Más que un hobby es una pasión, es algo que nos lleva a exigirnos más, a creer que todo es posible cuando estamos abordo de ella.",
		},
	}
	c.JSON(http.StatusOK, res)
}
