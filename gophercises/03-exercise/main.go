package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)
const sourceFile = "data.json"
func main()  {
	fmt.Println("wtf")
	mux := setupServer()
	filePath := flag.String("source",sourceFile,"A json file that holds a story")
	flag.Parse()
	stories, err := getStories(*filePath)
	if err != nil {
		panic(err)
	}
	storyHandle := StoryHandler(stories, mux)
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", storyHandle)
}
func storyNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)

	fmt.Fprintf(w, "Provide valid story name!")
}
func StoryHandler(stories map[string]Story, fallback http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		pathUrl := request.URL.Path
		if pathUrl != "/story" {
			fallback.ServeHTTP(writer, request)
			return
		}
		story := request.URL.Query()["name"]
		//fmt.Println("storyName is ",story)

		if len(story) == 0 {
			storyNotFound(writer, request)
			return
		}
		name := story[0]
		currStory, ok := stories[name]
		if !ok {
			storyNotFound(writer, request)
			return
		}
		jsonStory, _ := json.Marshal(currStory)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		writer.Write(jsonStory)
	}
}

func setupServer() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}
func parseJson(data []byte) (map[string]Story,error) {
	mp := make(map[string]Story)
	err := json.Unmarshal(data, &mp)
	return mp, err
}
func getStories(filePath string) (map[string]Story,error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err

	}
	return parseJson(data)
}

