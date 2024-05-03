package main

import "github.com/labstack/echo/v4"

func IinVerification(c echo.Context) error {
	iin := c.Param("iin")
	if iin == "" {
		return c.JSON(400, map[string]string{"error": "iin parameter is required"})
	}

	if !VerifyIIN(iin) {
		return c.JSON(500, map[string]bool{"correct": false}) // 500 - Internal Server Error Why 500?
	}

	dob, err := GetDateOfBirth(iin)
	if err != nil {
		return c.JSON(500, map[string]bool{"correct": false})
	}

	gender, err := GetGender(iin)
	if err != nil {
		return c.JSON(500, map[string]bool{"correct": false})
	}

	return c.JSON(200, map[string]interface{}{
		"correct":       true,
		"sex":           gender,
		"date_of_birth": dob.Format("01.02.2006"),
	})
}
