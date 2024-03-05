# ASCII Art Generator Web Server 🎯

This is a simple web server that generates ASCII art based on user input and displays it in a web page. Users can also download the generated ASCII art as a text file.

## How to Run 🏁

1. Make sure you have Go installed on your system.
2. Clone this repository to your local machine.
3. Navigate to the project directory in your terminal.

```bash
cd path/to/project/directory
```

1. Start the server by running the following command:

```bash
go run main.go
```

The server will start on port 8080. Access it in your web browser by visiting http://localhost:8080/.

## Usage 📟

1. Generating ASCII Art:

   To generate ASCII art, enter your text in the input box provided on the main page.
   Choose the style of ASCII art from the available options (thinkertoy, shadow, or standard) using the radio buttons.
   Click the "Generate" button to see the ASCII art on the web page.

2. Downloading ASCII Art:

   After generating the ASCII art, you can download it as a text file.
   Click the "Download" button, and it will prompt you to save the ASCII art as ascii_art.txt.
   Error Handling
   The server is equipped with basic error handling for certain scenarios, and it will display appropriate error messages and corresponding GIFs for the following cases:

400 Bad Request: Occurs when the input for generating ASCII art is invalid.
404 Page Not Found: Occurs when accessing a page or resource that doesn't exist.
405 Method Not Allowed: Occurs when an unsupported HTTP method is used for certain routes.

## Project Structure ⛩️

main.go: Contains the main code for the web server.
index.html: HTML template for the web page.
error.html: HTML template for displaying error messages with GIFs.
images/: Directory containing GIF files for error messages and other images.
banners/: Directory containing text files representing ASCII art styles.

## Credits ⭐️

This project is inspired by the love for ASCII art and is developed for educational purposes. The ASCII art styles used in this project are sourced from various open-source ASCII art repositories.

- tmachado
- rinrino

## Disclaimer 📸

Please note that the ASCII art styles provided are for educational and non-commercial use only. All ASCII art styles used in this project are sourced from open-source repositories and are not owned or copyrighted by the author of this project. If you are the creator of any of the ASCII art styles used here and wish to have it removed or attributed differently, kindly reach out to the project maintainer.
