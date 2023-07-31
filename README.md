# Learn golang with love

# 12. ****Gin****

## 1. Routes

### Basic route with handler Function
``` 
    var r = gin.Default() // set gin routes
    r.GET("/", homeHandler)
    
    func homeHandler(context *gin.Context) {
        //context.String(http.StatusOK, "Hello") // return String
        //context.JSON(http.StatusOK, "Json return") // return Json
        name := context.DefaultQuery("name", "Guest") // get Query Params
        age := context.DefaultQuery("age", "18")
        context.JSON(http.StatusOK, age+name)
    }
```

### Group Routes
``` 
	// Using Group in API
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/", homeHandler)
			v1.POST("/", postHandler)
			v1.GET("detail/:id", detailHandler)
			v1.POST("/upload", handlerUploadImage)
		}
	}
```


## 2. MiddleWare

### What is a middleWare function 
```
func testGlobalMiddleWare(context *gin.Context) {
    log.Println("in mid")
	context.Next()
}
```

### Use MiddleWare 
``` 
	var r = gin.Default() // set gin routes
	r.Use(testGlobalMiddleWare) // use middleware
```

## 3. Upload Single file & Multiple files

## 4. Logging

## 5. Data Binding & Validation

## 6. 


