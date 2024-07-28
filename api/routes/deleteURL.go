package routes

import (
  "main/api/database" 
  "net/http"
  "github.com/gin-gonic/gin"
)


func DeleteURL(c *gin.Context){
  shortID := c.Param("shortID")

  r:= database.CreateClient(0)
  defer r.Close()

  r.Del(database.Ctx, shortID).Err()
  if err!=nil{
    c.JSON(http.StatusInternalServerError, gin.H{
      "error" : "Unable to delete the Shortened link",
    })
    return
  }
  c.JSON(http.StatusOK, gin.H{
    "message" : "Shortened link deleted",
  })
}
