package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const (
	defaultTemplate = `<!DOCTYPE html>
<html>
<head>
<meta http-equiv="content-type" content="text/html; charset=utf-8">
<title>{{ .Title}}</title>
</head>
<body>
{{ .Body}}
</body>
</html>
`
)

type content struct {
	Title string
	Body  template.HTML
}

func main() {
	//Parse flags
	filename := flag.String("file", "", "Markdown file to preview")
	skipPreview := flag.Bool("s", false, "Skip auto-preview")
	tFname := flag.String("t", "", "Alternate template name")
	flag.Parse()

	// If user did not provide a file,show usage
	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}
	if err := run(*filename, *tFname, os.Stdout, *skipPreview); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filename, tFname string, out io.Writer, skipPreview bool) error {
	// Read all the data from the input file and check for errors
	input, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	htmlData, err := parseContent(input, tFname)
	if err != nil {
		return err
	}
	//create temporary file and check for errors
	temp, err := os.CreateTemp("", "mdp*.html")
	if err != nil {
		return err
	}
	if err := temp.Close(); err != nil {
		return err
	}

	// outName := fmt.Sprintf("%s.html", filepath.Base(filename))
	outName := temp.Name()
	// fmt.Println(outName)
	fmt.Fprintln(out, outName)

	if err := saveHTML(outName, htmlData); err != nil {
		return err
	}
	if skipPreview {
		return nil
	}

	defer os.Remove(outName)
	return preview(outName)

}

func saveHTML(outName string, htmlData []byte) error {

	return os.WriteFile(outName, htmlData, 0644)

}

func parseContent(input []byte, tFname string) ([]byte, error) {

	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)

	//Parse the contents of the defaultTemplate const into a new Template
	t, err := template.New("mdp").Parse(defaultTemplate)
	if err != nil {
		return nil, err
	}
	//If user provided alternative template file,replace template
	if tFname != "" {
		t, err = template.ParseFiles(tFname)
		if err != nil {
			return nil, err
		}
	}
	//Instantiate the content type,adding the title and body
	c := content{
		Title: "Markdown Preview Tool",
		Body:  template.HTML(body),
	}

	var buffer bytes.Buffer
	//Execute the template with the content type
	if err := t.Execute(&buffer, c); err != nil {
		return nil, err
	}

	// buffer.WriteString(header)
	// buffer.Write(body)
	// buffer.WriteString(footer)

	return buffer.Bytes(), nil
}

func preview(fname string) error {
	cName := ""
	cParams := []string{}

	//Define executable based on OS
	switch runtime.GOOS {
	case "linux":
		cName = "xdg-open"
	case "windows":
		cName = "cmd.exe"
		cParams = []string{"/c", "start"}
	case "darwin":
		cName = "open"
	default:
		return fmt.Errorf("OS not supported")
	}

	//Append filename to parameters slice
	cParams = append(cParams, fname)

	//Locate executable in PATH
	cPath, err := exec.LookPath(cName)
	if err != nil {
		return err
	}

	//Open the file using default program
	err = exec.Command(cPath, cParams...).Run()

	//Give the browser some time to open the file before deleting it
	time.Sleep(2 * time.Second)
	return err
}

//1. Go back to the example in Chapter 1, Your First Command-Line Program
///in Go, on page 1, and update the wc tool to read data from files in
//addition to STDIN.
//2. Update the wc tool to process multiple files.
//3. Update the mdp tool template by adding another field that shows the
//name of the file being previewed.
//4. Update the mdp tool allowing the user to specify a default template using
// an environment variable.
//5. Update the mdp tool allowing the user to provide the input Markdown via
//STDIN.
