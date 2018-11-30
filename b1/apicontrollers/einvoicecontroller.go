package apicontrollers

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hsynakin/goGin/b1/models"
)

var xmlUsers models.Users

// GETUserFromRegistrationNo ...
// Bu fonksiyon gönderilen vergi numarasına göre ilgili kaydı bulur.
func GETUserFromRegistrationNo(c *gin.Context) {
	var user models.EInvoiceUsers
	var taxRegistrationNo = c.Params.ByName("id")

	if len(xmlUsers.Users) == 0 {
		retVal := updateUsersArray()
		if !retVal {
			c.JSON(http.StatusBadRequest, models.GetGenericStatusResponse("400", "Users.xml dosyası bulunamadı"))
			return
		}
	}

	RecordFound := false

	for i := 0; i < len(xmlUsers.Users); i++ {
		if taxRegistrationNo == xmlUsers.Users[i].Identifier {
			user = models.EInvoiceUsers{
				Identifier:        xmlUsers.Users[i].Identifier,
				Alias:             xmlUsers.Users[i].Alias,
				Title:             xmlUsers.Users[i].Title,
				Type:              xmlUsers.Users[i].Type,
				FirstCreationTime: xmlUsers.Users[i].FirstCreationTime,
			}
			RecordFound = true
			break
		}
	}
	if RecordFound {
		c.JSON(http.StatusOK, user) //return values user
	} else {
		c.JSON(http.StatusNotFound, models.GetGenericStatusResponse("404", taxRegistrationNo+" vergi numaralı kayıt bulunamadı"))
	}
}

// GETUserFromFirstCreationTime ...
// get FirstCreationTime>=:date
func GETUserFromFirstCreationTime(c *gin.Context) {
	firstCrationTimeUsers := []models.EInvoiceUsers{}
	var taxRegistrationNo = c.Params.ByName("date")

	convertTime, err := time.Parse("2006-01-02", taxRegistrationNo)
	convertTime = convertTime.Local().Add(time.Hour * -3)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.GetGenericStatusResponse("400", "Hatalı bir zaman girilmiştir. Zaman formatı Yıl-Ay-Gün şeklinde olması gerekmektedir. Ör: 2018-11-28"))
		return
	}

	if len(xmlUsers.Users) == 0 {
		retVal := updateUsersArray()
		if !retVal {
			c.JSON(http.StatusBadRequest, models.GetGenericStatusResponse("400", "Users.xml dosyası bulunamadı"))
			return
		}
	}

	RecordFound := false

	for i := 0; i < len(xmlUsers.Users); i++ {
		if xmlUsers.Users[i].FirstCreationTime.After(convertTime) {
			user := models.EInvoiceUsers{
				Identifier:        xmlUsers.Users[i].Identifier,
				Alias:             xmlUsers.Users[i].Alias,
				Title:             xmlUsers.Users[i].Title,
				Type:              xmlUsers.Users[i].Type,
				FirstCreationTime: xmlUsers.Users[i].FirstCreationTime,
			}
			firstCrationTimeUsers = append(firstCrationTimeUsers, user)
			RecordFound = true
		}
	}
	if RecordFound {
		c.JSON(http.StatusOK, firstCrationTimeUsers) //return values user
	} else {
		c.JSON(http.StatusNotFound, models.GetGenericStatusResponse("404", convertTime.String()+" tarihinde kayıt bulunamadı"))
	}
}

// POSTUploadeInvoiceXMLFile ...
// update xml file
func POSTUploadeInvoiceXMLFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, models.GetGenericStatusResponse("400", "Dosya Yüklenemedi"))
		return
	}

	filePath := "./upload"
	exist, erri := FileOrDirectoryExists(filePath)

	log.Println(erri)
	// TODO: aşağıdaki dizini sadece dizin yok ise açalım.
	if exist == false {
		log.Println("FilePath: ", filePath)
		os.Mkdir(filePath, 0700)
	}

	filePath += "/" + file.Filename

	// Upload the file to specific dst.
	c.SaveUploadedFile(file, filePath)

	/*src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	defer src.Close()

	out, err := os.Create(filePath)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	defer out.Close()

	io.Copy(out, src)*/

	InputFunction := updateUsersArray()

	if InputFunction {
		c.JSON(http.StatusOK, models.GetGenericStatusResponse("404", "Xml başarılı bir şekilde güncellenmiştir."))
	} else {
		c.JSON(http.StatusNotFound, models.GetGenericStatusResponse("404", "Dosya bulunamadı."))
	}
}

func updateUsersArray() bool {
	xmlFile, err := os.Open("Users.xml")
	if err != nil {
		return false
	}

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	xml.Unmarshal(byteValue, &xmlUsers)

	return true
}

func FileOrDirectoryExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
