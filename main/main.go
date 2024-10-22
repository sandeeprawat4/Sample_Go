package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id string `json:"id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
	fmt.Println("Endpoint Hit: homePage")
}
func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	//myRouter.HandleFunc("/articles/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func main() {
	/*Articles = []Article{
		Article{Id: "101", Title: "Fun", Desc: "Latest movie", Content: "Comedy"},
		Article{Id: "102", Title: "Political", Desc: "Drama", Content: "Mix"},
	}*/
	var arr = []int{3,3}
	twoSum(arr, 6)
}

func twoSum(nums []int, target int) []int {
	var m map[int]int
	m = make (map[int]int)
    for i, j:= range nums{
		var k = m[j]
		fmt.Println("K value", k, "and current arr is ", i, "and curr val is ", j)
		if k!=0 {
			k++;
			m[j]=k;
		}else{
			m[j]=1;
		}
	}
	fmt.Println(len(m))
	var first = -1;
	var second = -1;
	for i, j:= range nums{
		var temp = target - j;
		if(m[temp]!=0){
			if(temp == j && m[temp]==1){
				continue
			}
			if(first==-1){
				first=i;
			}else{
				second=i;
				break;
			}
		}
	}
	return []int{first, second}
}
