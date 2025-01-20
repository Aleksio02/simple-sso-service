package connector

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"simple-sso-service/modules/ext/service"
)

func GetToken(authService service.AuthService, code string) map[string]string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", authService.GetTokenLink()+"?code="+code, nil)
	req.Header.Set("Accept", "application/json")
	response, _ := client.Do(req)
	defer response.Body.Close()
	responseBodyAsReader, _ := ioutil.ReadAll(response.Body)
	responseBody := map[string]string{}
	json.Unmarshal(responseBodyAsReader, &responseBody)
	return responseBody
}
