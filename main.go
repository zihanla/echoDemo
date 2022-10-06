package main

import "demo_1/router"

func main() {
	//e := echo.New()
	//	//e.GET("/", func(c echo.Context) error {
	//	//	return c.String(http.StatusOK, "Hello, zihan!")
	//	//})
	//	//e.Logger.Fatal(e.Start(":8080"))

	router.Run()
}
