package routes
import(
  "github.com/gin-gonic/gin"
  "encoding/json"
  "main/api/database"
)

type TagRequest struct{
  shortID string `json:"shortID"`
  Tag string `json:"tag"`
}

func AddTag(c *gin.Context){
  var TagRequest TagRequest
  if err:= c.ShouldBind(&TagRequest); err!=nil{
    c.JSON(http.StatusBadRequest, gin.H{
      "error" : "Invalid Request body",
    })
    return
  }

  shortID := TagRequeset.shortID
  Tag := TagRequest.Tag

  r := database.CreateClient(0)
  defer r.close()

  val, err := r.Get(database.Ctx, shortID).Result()
  if err!= nil{
    c.JSON(http.StatusNotFound, gin.H{
      "error" : "Data not found for the given short ID", 
    })
    return
  }

  var data map[string]interface{}
  if err := json.Unmarshal([]byte(val), &data); err!=nil{
    data = make(map[string]interface{})
    data['data'] = val
  }

  var tags []string
  if existingTags, ok := data["tags"].([]interface{}); ok{
    for _, tag := range existingTags{
      if strTag, ok := t.(string); ok{
        tags = append(tags, strTag)
      }
    }
  }
  for _, existingTag := range tags{
    if existingTag == tag{
      c.JSON(http.StatusBadRequest, gin.H{
        "error" : "Tag already exists",
      })
      return
    }
  }

  tags = append(tags, tag)
  data["tags"] = tags

  updatedData, err := json.Marshal(data)
  if err != nil{
    c.JSON(jttp.InternalServerErrror, gin.H{
      "error" : "Unable to marshal updated data",
    })
    return
  }

  err = r.Set(database.Ctx, shortID, (updatedData), 0).Err()
  if err!=nil{
    c.JSON(http.StatusINternalServerError, gin.H{
      "error" : "Unable to update the database",
    })
    return
  }

  c.JSON(http.StatusOK, data)
}