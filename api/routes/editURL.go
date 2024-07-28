package routes

import (
	"main/api/models"
	"main/api/database"
	"net/http"
  "time"
	"github.com/gin-gonic/gin"
)
func EditURL(c *gin.Context){
  shortID := c.Param("shortID")
  var body models.Request

  if err := c.ShouldBind(&body); err!= nil{
    c.JSON(http.StatusBadRequest, gin.H{
      "error" : "Cannot parse JSON",
    })
  }

  r := database.CreateClient(0)
  defer r.Close()

  val,err := r.Get(database.Ctx, shortID).Result()
  if err!=nil || val == ""{
    c.JSON(http.StatusNotFound,gin.H{
      "error" : "ShortID does not exist",
    })
  }

  err = r.Set(database.Ctx, shortID, body.URL, body.Expiry*3600*time.Second).Err()
  if err!=nil{
    c.JSON(http.StatusInternalServerError, gin.H{
      "error" : "Unable to update the Shortened link",
    })
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "message" : "Shortened link updated",
  })
}