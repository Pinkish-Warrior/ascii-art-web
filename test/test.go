package stylize

// import (
// 	"fmt"
// 	"html/template"
// 	"log"
// 	"net/http"
// 	"os"
// 	"strings"
// )

// type PageData struct {
// 	AsciiArt string
// }

// func errorHandler(w http.ResponseWriter, statusCode int, message string, gifPath string) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	w.WriteHeader(statusCode)

// 	tmpl, err := template.ParseFiles("images/error400.html")
// 	if err != nil {
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	err = tmpl.Execute(w, struct {
// 		Message string
// 		GIFPath string
// 	}{Message: message, GIFPath: gifPath})

// 	if err != nil {
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// }

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		if r.URL.Path != "/" {
// 			errorHandler(w, http.StatusNotFound, "404 Page Not Found", "/images/error404.gif")
// 			return
// 		}
// 		http.ServeFile(w, r, "index.html")
// 	})

// 	// Serve static files
// 	fs := http.FileServer(http.Dir("."))
// 	http.Handle("/static/", http.StripPrefix("/static/", fs))

// 	// Serve the image file
// 	http.HandleFunc("/images/goworkout.gif", func(w http.ResponseWriter, r *http.Request) {
// 		http.ServeFile(w, r, "images/goworkout.gif")
// 	})

// 	// Serve the gif for 400 ERROR
// 	http.HandleFunc("/images/error400.gif", func(w http.ResponseWriter, r *http.Request) {
// 		http.ServeFile(w, r, "images/error400.gif")
// 	})

// 	// Serve the gif for 405 ERROR
// 	http.HandleFunc("/images/error405.webp", func(w http.ResponseWriter, r *http.Request) {
// 		http.ServeFile(w, r, "images/error405.webp")
// 	})

// 	// Serve the gif for 404 ERROR
// 	http.HandleFunc("/images/error404.gif", func(w http.ResponseWriter, r *http.Request) {
// 		http.ServeFile(w, r, "images/error404.gif")
// 	})

// 	// Serve the Art
// 	http.HandleFunc("/ascii-art", generateHandler)

// 	// Serve the Download
// 	http.HandleFunc("/download", downloadHandler)

// 	fmt.Println("Starting server at port 8080\n", " 👉 http://localhost:8080/\n", "Use Control+C to STOP server")

// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatal(err)
// 	}
// }

// func generateASCII(str string, filename string) string {

// 	// Split the string into individual words using whitespace as the separator.
// 	words := strings.Split(str, "\n")

// 	bannerPath := fmt.Sprintf("banners/%s", filename)
// 	bannerBytes, err := os.ReadFile(bannerPath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	bannerLines := strings.Split(strings.ReplaceAll(string(bannerBytes), "\r\n", "\n"), "\n")

// 	// Create a strings.Builder to store the resulting ASCII art.
// 	var result strings.Builder
// 	for i, word := range words {
// 		if word == "" {
// 			if i < len(words)-1 {
// 				result.WriteString("\n")
// 			}
// 			continue
// 		}
// 		for row := 1; row < 9; row++ {
// 			for _, char := range word {
// 				if int(char) >= 32 && int(char) <= 126 {
// 					for lineIndex, line := range bannerLines {
// 						// Determine which line corresponds to the current character and row.
// 						if lineIndex == (int(char)-32)*9+row {
// 							result.WriteString(line)
// 						}
// 					}
// 				}
// 			}
// 			result.WriteString("\n")
// 		}
// 	}

// 	return result.String()
// }

// func generateHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		errorHandler(w, http.StatusMethodNotAllowed, "405 Method Not Allowed", "/images/error405.webp")
// 		return
// 	}

// 	input := r.FormValue("input")

// 	if input < "" || input > "~" {
// 		errorHandler(w, http.StatusBadRequest, "400 Bad Request", "/images/error400.gif")
// 		return
// 	}

// 	// Get the filename from the query parameter
// 	filename := r.FormValue("filename")
// 	switch filename {
// 	case "thinkertoy":
// 		filename = "thinkertoy.txt"
// 	case "shadow":
// 		filename = "shadow.txt"
// 	default:
// 		filename = "standard.txt"
// 	}

// 	asciiArt := generateASCII(input, filename)

// 	if len(asciiArt) > 0 && asciiArt[len(asciiArt)-1] == '\n' {
// 		asciiArt = asciiArt[:len(asciiArt)-1]
// 	}

// 	// Define the HTML template
// 	tmpl, err := template.ParseFiles("index.html")
// 	if err != nil {
// 		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	tmplData := struct {
// 		ASCIIArt template.HTML
// 	}{
// 		ASCIIArt: template.HTML(asciiArt),
// 	}

// 	err = tmpl.Execute(w, tmplData)
// 	if err != nil {
// 		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// }

// func downloadHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	asciiArt := r.FormValue("ascii")
// 	if asciiArt == "" {
// 		http.Error(w, "400 Bad Request", http.StatusBadRequest)
// 		return
// 	}

// 	// Set the appropriate headers for the file download
// 	w.Header().Set("Content-Disposition", "attachment; filename=ascii_art.txt")
// 	w.Header().Set("Content-Type", "text/plain")

// 	// Write the ASCII art to the response writer
// 	_, err := w.Write([]byte(asciiArt))
// 	if err != nil {
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// }
