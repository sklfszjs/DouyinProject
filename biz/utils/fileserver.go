package utils

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func RunFileServer() {
	p, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	http.Handle("/", http.FileServer(http.Dir(p)))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println(err)
	}
}

// func RunFileServer() {
// 	p, _ := filepath.Abs(filepath.Dir(os.Args[0]))
// 	// http.Handle("/statistic", http.FileServer(http.Dir(p)+"/statistic"))
// 	// fmt.Println(http.Dir(p) + "/statistic")
// 	// err := http.ListenAndServe(":8888", nil)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }
// 	mux := http.NewServeMux()
// 	mux.Handle("/statistic/", http.StripPrefix("/statistic/", http.FileServer(http.Dir(p)+"/statistic/")))
// 	err := http.ListenAndServe(":8888", mux)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }
