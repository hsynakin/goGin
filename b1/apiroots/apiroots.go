package apiroots

import (
	"github.com/gin-gonic/gin"
	"github.com/hsynakin/goGin/b1/apicontrollers"
)

// Einvoiceservices ...
// das
func Einvoiceservices(api *gin.RouterGroup) {
	api.GET("/getUserFromTaxRegistrationNo/:id", apicontrollers.GETUserFromRegistrationNo)
	api.GET("/getUserFromFirstCreationTime/:date", apicontrollers.GETUserFromFirstCreationTime)
	api.POST("/uploadeinvoicexmlfile", apicontrollers.POSTUploadeInvoiceXMLFile)
}
