package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    const Username = "USERNAME"
    const Password = "PASSWORD"

    payload := map[string]string{
      	"source": "universal",
	"url": "https://mls.foreclosure.com/listing/search.html?ci=bay%20shore&st=ny&utm_source=internal&utm_medium=link&utm_campaign=MLS_top_links",
          }

    jsonValue, _ := json.Marshal(payload)

    client := &http.Client{}
    request, _ := http.NewRequest("POST",
        "https://realtime.oxylabs.io/v1/queries",
        bytes.NewBuffer(jsonValue),
    )

    request.SetBasicAuth(Username, Password)
    response, _ := client.Do(request)

    responseText, _ := ioutil.ReadAll(response.Body)
    fmt.Println(string(responseText))
}
