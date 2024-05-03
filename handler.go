package main

import (
	"github.com/labstack/echo/v4"
	"github.com/nyaruka/phonenumbers"
)

func IinVerification(c echo.Context) error {
	iin := c.Param("iin")
	if iin == "" {
		return c.JSON(400, map[string]string{"errors": "IIN parameter is required"})
	}

	if !VerifyIIN(iin) {
		return c.JSON(200, map[string]bool{"correct": false})
	}

	dob, err := GetDateOfBirth(iin)
	if err != nil {
		return c.JSON(200, map[string]bool{"correct": false})
	}

	gender, err := GetGender(iin)
	if err != nil {
		return c.JSON(200, map[string]bool{"correct": false})
	}

	return c.JSON(200, map[string]interface{}{
		"correct":       true,
		"sex":           gender,
		"date_of_birth": dob.Format("02.01.2006"),
	})
}

func AddCitizen(c echo.Context) error {
	citizen := new(Citizen)
	if err := c.Bind(citizen); err != nil {
		return c.JSON(500, map[string]interface{}{"success": false, "errors": err.Error()}) // 500 - Internal Server Error why not 400? In test problem written that we need to return 500
	}

	if !VerifyIIN(citizen.Iin) {
		return c.JSON(500, map[string]interface{}{"success": false, "errors": "IIN is not valid"})
	}

	phone := citizen.Phone
	if _, err := phonenumbers.Parse(phone, "KZ"); err != nil {
		return c.JSON(500, map[string]interface{}{"success": false, "errors": "Phone number is not valid"})
	}

	if err := DB.Create(&citizen).Error; err != nil {
		return c.JSON(500, map[string]interface{}{"success": false, "errors": err.Error()})
	}

	return c.JSON(200, map[string]interface{}{"success": true})
}

func GetCitizenByIin(c echo.Context) error {
	iin := c.Param("iin")
	if iin == "" {
		return c.JSON(400, map[string]string{"errors": "IIN parameter is required"})
	}

	if !VerifyIIN(iin) {
		return c.JSON(200, map[string]interface{}{"success": false, "errors": "IIN is not valid"})
	}

	citizen := new(Citizen)
	if err := DB.Where("iin = ?", iin).First(&citizen).Error; err != nil {
		return c.JSON(200, map[string]interface{}{"success": false, "errors": err.Error()})
	}

	return c.JSON(200, citizen)
}

// I generalize this handler because we can search by name or phone if requierments will be changed
// And this handler is more flexible for adding new search parameters
func GetCitizens(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		return c.JSON(400, map[string]string{"error": "name query parameter is required"})
	}

	var citizens []Citizen

	// I think that we need to search by full name, not by part of the name
	// So I will change the query if requirements will be changed
	// Now we search by part of the name LastName FirstName
	if err := DB.Where("name LIKE ? OR name LIKE ?", "%"+name+" %", "% "+name+" %").Find(&citizens).Error; err != nil {
		return c.JSON(500, map[string]interface{}{"success": false, "errors": err.Error()})
	}

	return c.JSON(200, citizens)
}
