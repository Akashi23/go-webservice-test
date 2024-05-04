package main

import (
	"github.com/labstack/echo/v4"
	"github.com/nyaruka/phonenumbers"
)

// IinVerification godoc
// @Summary Get sex and date of birth by IIN
// @Description Verify IIN
// @Tags iin
// @Produce json
// @Param iin path string true "IIN"
// @Success 200 {object} map[string]interface{}
// @Router /iin_check/:iin [get]
func IinVerification(c echo.Context) error {
	iin := c.Param("iin")
	if iin == "" {
		return c.JSON(400, map[string]string{"errors": "IIN parameter is required"})
	}

	if !VerifyIIN(iin) {
		return c.JSON(200, map[string]bool{"correct": false})
	}

	gender, err := GetGender(iin)
	if err != nil {
		return c.JSON(200, map[string]bool{"correct": false})
	}

	dob, err := GetDateOfBirth(iin)
	if err != nil {
		return c.JSON(200, map[string]bool{"correct": false})
	}

	return c.JSON(200, map[string]interface{}{
		"correct":       true,
		"sex":           gender,
		"date_of_birth": dob.Format("02.01.2006"),
	})
}

// AddCitizen godoc
// @Summary Add citizen
// @Description Add citizen
// @Tags people
// @Accept json
// @Produce json
// @Param citizen body Citizen true "Citizen"
// @Success 200 {object} map[string]interface{} "success"
// @Router /people/info [post]
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

// GetCitizenByIin godoc
// @Summary Get citizen by IIN
// @Description Get citizen by IIN
// @Tags people
// @Produce json
// @Param iin path string true "IIN"
// @Success 200 {object} Citizen
// @Router /people/info/iin/:iin [get]
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

// GetCitizens godoc
// @Summary Get citizens by name
// @Description Get citizens by name
// @Tags people
// @Produce json
// @Param name query string true "Name"
// @Success 200 {object} []Citizen
// @Router /people/info [get]
func GetCitizens(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" { // if name is empty we might return all citizens if requirements will be changed
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

// Healthcheck godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /health [get]
func Healthcheck(c echo.Context) error {
	return c.JSON(200, map[string]string{"status": "ok"})
}
