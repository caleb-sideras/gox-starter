package gox

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/caleb-sideras/goxstack/gox/utils"
	"github.com/gorilla/mux"
)

const (
	DIR      = "/"
	GO_EXT   = ".go"
	HTML_EXT = ".html"
	TXT_EXT  = ".txt"
	PAGE     = "page"
	INDEX    = "index"
	METADATA = "metadata"
	DATA     = "data"
	RENDER   = "render"
	HANDLE   = "handle"
	ETAG     = "etag_file"
	BODY     = "-body"

	PAGE_BODY      = PAGE + BODY
	PAGE_FILE      = PAGE + HTML_EXT
	INDEX_FILE     = INDEX + HTML_EXT
	METADATA_FILE  = METADATA + HTML_EXT
	DATA_FILE      = DATA + GO_EXT
	RENDER_FILE    = RENDER + GO_EXT
	HANDLE_FILE    = HANDLE + GO_EXT
	PAGE_BODY_FILE = PAGE_BODY + HTML_EXT
	ETAG_FILE      = ETAG + TXT_EXT
)

var FILE_CHECK_LIST = map[string]bool{
	DATA_FILE:     true,
	RENDER_FILE:   true,
	HANDLE_FILE:   true,
	INDEX_FILE:    true,
	PAGE_FILE:     true,
	METADATA_FILE: true,
}

type Gox struct {
	OutputDir string
}

type GoxDir struct {
	FileType string
	FilePath string
}

func NewGox(outputDir string) *Gox {
	return &Gox{
		OutputDir: outputDir,
	}
}

func (g *Gox) Build(startDir string, packageDir string) {

	log.Println("---------------------WALKING DIRECTORY---------------------")
	dirFiles, err := walkDirectoryStructure(startDir)
	if err != nil {
		log.Fatalf("error walking the path %v: %v", startDir, err)
	}

	for k, v := range dirFiles {
		log.Println("Directory:", k)
		for ext, files := range v {
			log.Println("  ", ext)
			for _, file := range files {
				log.Println("   -", file)
			}
		}
	}

	// used for generated output
	var pages []string
	var indexGroup []string
	var dataFunctions []string
	var renderFunctions []string
	var handleFunctions []string
	imports := utils.NewStringSet()

	log.Println("---------------------EXTRACTING YOUR CODE---------------------")
	for dir, files := range dirFiles {
		if len(files) > 0 {
			log.Println("Directory:", dir)

			dataPath := filepath.Join(dir, DATA_FILE)
			// pagePath := filepath.Join(dir, PAGE_FILE)
			renderPath := filepath.Join(dir, RENDER_FILE)
			handlerPath := filepath.Join(dir, HANDLE_FILE)

			var goFiles []GoxDir
			if _, ok := files[GO_EXT]; ok {
				goFiles = files[GO_EXT]
			}

			var htmlFiles []GoxDir
			if _, ok := files[HTML_EXT]; ok {
				htmlFiles = files[HTML_EXT]
			}

			ndir := removeDirWithUnderscorePostfix(dir)
			leafNode := filepath.Base(ndir)
			leafPath := ndir[5:]

			// GO
			for _, gd := range goFiles {
				switch gd.FileType {
				case DATA_FILE:
					log.Println("   data.go")

					hasExpVar, pkName, err := hasExportedVariable(dataPath, "Data")
					if err != nil {
						panic(err)
					}

					hasExpFunc, _, err := hasExportedFuction(dataPath, "Data")
					if err != nil {
						panic(err)
					}

					if pkName == "" {
						// im well aware this is "impossible"
						errStr := "No defined package name in " + dataPath
						panic(errors.New(errStr))
					}

					if !hasExpVar && !hasExpFunc {
						// if nothing -> render page.html + add to path -> goto page.html

					} else if hasExpVar {
						log.Println("   - Extracted -> var Data")
						pages = append(pages, formatData(pkName, leafPath, htmlFiles))
						imports.Add(`"` + packageDir + filepath.Dir(dataPath) + `"`)
					} else if hasExpFunc {
						log.Println("   - Extracted -> func Data")
						dataFunctions = append(dataFunctions, formatData(pkName, leafPath, htmlFiles))
						imports.Add(`"` + packageDir + filepath.Dir(dataPath) + `"`)
					}

				case RENDER_FILE:
					log.Println("   render.go")

					expFns, pkName, err := getExportedFuctions(renderPath)
					if err != nil {
						panic(err)
					}

					if pkName == "" {
						errStr := "No defined package name in " + renderPath
						panic(errors.New(errStr))
					}

					if expFns != nil {
						// Could have edge case where there is expFns, but none meeting our criteria, so this import will prevent a build
						imports.Add(`"` + packageDir + filepath.Dir(renderPath) + `"`)
						for _, expFn := range expFns {
							if expFn == "Render" || strings.HasSuffix(expFn, "_") {
								renderFunctions = append(renderFunctions, formatDefaultFunction(pkName, expFn, strings.TrimSuffix(expFn, "_"), leafNode, "Render"))
								log.Println("   - Extracted -> func", expFn)
							} else {
								log.Println("   - Exported function --", expFn, "-- ignored")
							}
						}
					} else {
						log.Println("   - No exported functions in", renderPath)
					}

				case HANDLE_FILE:

					log.Println("   handle.go")

					expFns, pkName, err := getExportedFuctions(handlerPath)
					if err != nil {
						panic(err)
					}

					if pkName == "" {
						errStr := "No defined package name in " + handlerPath
						panic(errors.New(errStr))
					}

					if expFns != nil {
						// Could have edge case where there is expFns, but none meeting our criteria, so this import will prevent a build
						imports.Add(`"` + packageDir + filepath.Dir(handlerPath) + `"`)
						for _, expFn := range expFns {
							if expFn == "Handler" || strings.HasSuffix(expFn, "_") {
								handleFunctions = append(handleFunctions, formatDefaultFunction(pkName, expFn, strings.TrimSuffix(expFn, "_"), leafNode, "Handler"))
								log.Println("   - Extracted -> func", expFn)
							} else {
								log.Println("   - Exported function --", expFn, "-- ignored")
							}
						}
					} else {
						log.Println("   - No exported functions in", handlerPath)
					}
				}
			}

			// HTML
			for _, gd := range htmlFiles {
				switch gd.FileType {
				case INDEX_FILE:
					log.Println("   index.html")
					// FIX THIS
					if leafPath == "" {
						leafPath = "/"
					}
					indexGroup = append(indexGroup, `"`+leafPath+`"`+" : "+`"`+gd.FilePath+`",`)
					log.Println("   - Index pair:", leafPath, "->", gd.FilePath)

				case PAGE_FILE:
					// var fileDst string
					// for k := range htmlFiles {
					// 	if filepath.Base(k) == "page.html" {
					// 		fileDst = removeDirWithUnderscorePostfix(k)[5:]
					// 		break
					// 	}
					// }

					// if fileDst == "" {
					// 	log.Panicln("Please provide a data.go &/or page.html for directory:", dir)
					// }

					// err := utils.RenderFiles(fileDst, g.OutputDir, utils.MapKeysToSlice(htmlFiles), struct{}{}, "")
					// if err != nil {
					// 	panic(err)
					// }
				}
			}
		}
	}

	err = generateCode(imports, pages, indexGroup, dataFunctions, renderFunctions, handleFunctions)
	if err != nil {
		panic(err)
	}

	err = g.renderStaticFiles()
	if err != nil {
		panic(err)
	}
}

func (g *Gox) Run(r *mux.Router, port string, servePath string) {

	http.Handle("/", r)
	// http.Handle(servePath, http.StripPrefix(servePath, http.FileServer(http.Dir(g.OutputDir))))
	g.handleRoutes(r, g.getETags())
	log.Fatal(http.ListenAndServe(port, nil))
}

func (g *Gox) getETags() map[string]string {
	log.Println("GENERATING ETAGS...")
	var eTags map[string]string
	eTags = make(map[string]string)

	file, err := os.Open(filepath.Join(g.OutputDir, ETAG_FILE))
	if err != nil {
		log.Fatalf("could not create file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")
		if len(parts) == 2 {
			eTags[parts[0]] = parts[1]
		}
	}
	return eTags
}

// handleRoutes() binds Mux handlers to user defined functions, and creates default handlers to serve static pages
func (g *Gox) handleRoutes(r *mux.Router, eTags map[string]string) {
	log.Println("---------------------PAGES HANDLERS-----------------------")
	for route := range PagesList {
		log.Println(route)
		r.HandleFunc(route+"{slash:/?}",
			func(w http.ResponseWriter, r *http.Request) {
				log.Println("- - - - - - - - - - - -")
				var eTagPath string
				var pagePath string
				if utils.IsHtmxRequest(r) {
					log.Println("HX-Request")
					htmxUrl, err := lastElementOfURL(utils.GetHtmxRequestURL(r))
					if err != nil {
						log.Println("Error parsing HX-Current-URL:", err)
						http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					}
					if _, ok := IndexList[htmxUrl]; ok {
						if IndexList[htmxUrl] == IndexList[r.URL.Path] {
							eTagPath = filepath.Join(r.URL.Path, PAGE_BODY_FILE)
							pagePath = filepath.Join(g.OutputDir, eTagPath)
							log.Println("Serving file: body - matching index")
						} else {
							eTagPath = filepath.Join(r.URL.Path, PAGE_FILE)
							pagePath = filepath.Join(g.OutputDir, eTagPath)
							log.Println("Serving file: page - new index")
						}
					} else {
						eTagPath = filepath.Join(r.URL.Path, PAGE_FILE)
						pagePath = filepath.Join(g.OutputDir, eTagPath)
						log.Println("Serving file: page - no index")
					}
				} else {
					eTagPath = filepath.Join(r.URL.Path, PAGE_FILE)
					pagePath = filepath.Join(g.OutputDir, eTagPath)
					log.Println("Serving file: page - no HX-Request")
				}
				log.Println("Path:", pagePath)
				log.Println("ETag:", eTags[eTagPath])
				w.Header().Set("ETag", eTags[eTagPath])
				http.ServeFile(w, r, pagePath)
			},
		)
	}

	log.Println("---------------------DATA HANDLERS-----------------------")
	for route, data := range DataList {
		log.Println(route)
		r.HandleFunc(route+"{slash:/?}",

			func(w http.ResponseWriter, r *http.Request) {

				log.Println("- - - - - - - - - - - -")
				log.Println("Fetching Data...")
				funcReturn := data.Data(w, r)

				if funcReturn.Error != nil {
					log.Println("Error Fetching Data")
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}

				tmplPaths := append(data.AdditionalTemplates, funcReturn.Templates...)
				tmpl, err := template.ParseFiles(tmplPaths...)
				// inefficient - best to parse on run to a map -> {"route":template.Template} or something
				if err != nil {
					log.Println("Error Parsing Files")
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}

				buffer := &bytes.Buffer{}

				if utils.IsHtmxRequest(r) {
					log.Println("HX-Request")

					htmxUrl, err := lastElementOfURL(utils.GetHtmxRequestURL(r))
					if err != nil {
						log.Println("Error parsing HX-Current-URL:", err)
						http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					}
					if _, ok := IndexList[htmxUrl]; ok {
						if IndexList[htmxUrl] == IndexList[r.URL.Path] {
							tmpl.ExecuteTemplate(buffer, "page", funcReturn.Content)
							log.Println("Serving file: body - matching index")
						} else {
							tmpl.Execute(buffer, funcReturn.Content)
							log.Println("Serving file: page - new index")
						}
					} else {
						tmpl.Execute(buffer, funcReturn.Content)
						log.Println("Serving file: page - no index")
					}
				} else {
					tmpl.Execute(buffer, funcReturn.Content)
					log.Println("Serving file: page - no HX-Request")
				}

				eTag := utils.GenerateETag(string(buffer.String()))
				log.Println("ETag:", eTag)

				w.Header().Set("ETag", eTag)
				w.Write(buffer.Bytes())
			},
		)
	}

	log.Println("---------------------RENDER HANDLERS-----------------------")
	for _, route := range RenderList {
		log.Println(route.Path + DIR)
		r.HandleFunc(route.Path+"{slash:/?}",
			func(w http.ResponseWriter, r *http.Request) {
				log.Println("- - - - - - - - - - - -")
				log.Println("Serving static component")
				log.Println("Path:", filepath.Join(g.OutputDir, r.URL.Path, PAGE_FILE))
				http.ServeFile(w, r, filepath.Join(g.OutputDir, r.URL.Path, PAGE_FILE))
			},
		)
	}
	log.Println("----------------------CUSTOM HANDLERS----------------------")
	for _, route := range HandleList {
		log.Println(route.Path)
		r.HandleFunc(route.Path+"{slash:/?}", route.Handler)
	}
}

func lastElementOfURL(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	if u.Path == "" || u.Path == "/" {
		return "/", nil
	}

	return "/" + path.Base(u.Path), nil
}

// RenderStaticFiles() renders all static files defined by the user
// Returns a map of all rendered paths
func (g *Gox) renderStaticFiles() error {
	output := ""

	// Rendering routes defined with page.html
	for path, data := range PagesList {
		err := utils.RenderFiles[interface{}](filepath.Join(path, PAGE_FILE), g.OutputDir, append(data.AdditionalTemplates, data.Data.Templates...), data.Data.Content, "")
		if err != nil {
			return err
		}
		content, err := os.ReadFile(filepath.Join(g.OutputDir, path, PAGE_FILE))
		if err != nil {
			return err
		}
		output += fmt.Sprintf("%s:%s\n", filepath.Join(path, PAGE_FILE), utils.GenerateETag(string(content)))

		err = utils.RenderFiles[interface{}](filepath.Join(path, PAGE_BODY_FILE), g.OutputDir, append(data.AdditionalTemplates, data.Data.Templates...), data.Data.Content, "page")
		if err != nil {
			return err
		}
		content, err = os.ReadFile(filepath.Join(g.OutputDir, path, PAGE_BODY_FILE))
		if err != nil {
			return err
		}
		output += fmt.Sprintf("%s:%s\n", filepath.Join(path, PAGE_BODY_FILE), utils.GenerateETag(string(content)))
	}

	file, err := utils.CreateFile(ETAG_FILE, g.OutputDir)
	if err != nil {
		return err
	}
	_, err = file.Write([]byte(output))
	if err != nil {
		return err
	}

	// Rendering .html files returned from functions defined in render.go
	for _, renderDefault := range RenderList {
		var err error
		switch fn := renderDefault.Handler.(type) {
		case func() utils.RenderFilesType:
			renderType := renderDefault.Handler.(func() utils.RenderFilesType)()
			err = utils.RenderFiles(filepath.Join(renderDefault.Path, PAGE_FILE), g.OutputDir, renderType.StrArr, renderType.Value, renderType.Str)
		case func() utils.RenderTemplateType:
			renderType := renderDefault.Handler.(func() utils.RenderTemplateType)()
			err = utils.RenderTemplate(filepath.Join(renderDefault.Path, PAGE_FILE), g.OutputDir, renderType.Tmpl, renderType.Value, renderType.Str)
		default:
			log.Printf("Unknown function type: %T\n", fn)
		}

		if err != nil {
			return err
		}
	}
	return nil
}

func walkDirectoryStructure(startDir string) (map[string]map[string][]GoxDir, error) {

	result := make(map[string]map[string][]GoxDir)

	err := filepath.Walk(startDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && strings.HasPrefix(info.Name(), "_") {
			return filepath.SkipDir
		}

		if info.IsDir() && path != startDir {
			files := make(map[string][]GoxDir)

			filepath.Walk(path, func(innerPath string, innerInfo os.FileInfo, innerErr error) error {

				if innerInfo.IsDir() && strings.HasPrefix(innerInfo.Name(), "_") {
					return filepath.SkipDir
				}

				ext := filepath.Ext(innerPath)
				if !innerInfo.IsDir() && filepath.Dir(innerPath) == path && FILE_CHECK_LIST[filepath.Base(innerPath)] {
					if _, exists := files[ext]; !exists {
						files[ext] = []GoxDir{}
					}
					files[ext] = append(files[ext], GoxDir{filepath.Base(innerPath), innerPath})
				}
				return nil
			})

			currDir := path
			urlToIndex := make(map[string]string)
			for {
				indexFile := filepath.Join(currDir, INDEX_FILE)
				if _, err := os.Stat(indexFile); !os.IsNotExist(err) {
					if _, ok := files[filepath.Ext(indexFile)]; !ok {
						files[filepath.Ext(indexFile)] = []GoxDir{}
					}
					files[filepath.Ext(indexFile)] = append(files[filepath.Ext(indexFile)], GoxDir{filepath.Base(indexFile), indexFile})
					urlToIndex[removeDirWithUnderscorePostfix(path)[5:]] = indexFile
					break
				}
				currDir = filepath.Dir(currDir)
				if currDir == "." || currDir == "/" {
					return errors.New("MISSING: " + INDEX_FILE)
				}
			}

			result[path] = files
		}
		return nil
	})

	return result, err
}

func formatDefaultFunction(pkName string, fnName string, pathName string, leafPath string, rootPath string) string {
	if fnName == rootPath {
		return `{"/` + leafPath + `", ` + pkName + `.` + fnName + `},`
	} else {
		return `{"/` + leafPath + `/` + strings.ToLower(pathName) + `", ` + pkName + `.` + fnName + `},`
	}
}

func formatCustomFunction(pkName string, fnName string) string {
	return `{` + pkName + `.` + fnName + `},`
}

func formatData(pkName string, leafPath string, dirHtmlFiles []GoxDir) string {
	var additionalTemplates []string
	// root case
	if leafPath == "" {
		leafPath = "/"
	}
	for _, file := range dirHtmlFiles {
		if file.FileType == INDEX_FILE {
			additionalTemplates = append([]string{`"` + file.FilePath + `",`}, additionalTemplates...)
		} else {
			additionalTemplates = append(additionalTemplates, `"`+file.FilePath+`",`)
		}
	}
	return `"` + leafPath + `": {Data:` + pkName + `.` + "Data" + `, AdditionalTemplates: []string{` + strings.Join(additionalTemplates, " ") + `},},`
}

func getAstVals(path string) (*ast.File, error) {
	_, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func getExportedFuctions(path string) ([]string, string, error) {

	node, err := getAstVals(path)
	if err != nil {
		return nil, "", err
	}

	var pkName string
	var expFns []string
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.File:
			pkName = x.Name.Name
		case *ast.FuncDecl:
			if x.Name.IsExported() {
				expFns = append(expFns, x.Name.Name)
			}
		}
		return true
	})
	return expFns, pkName, nil
}

func hasExportedFuction(path string, funcName string) (bool, string, error) {

	node, err := getAstVals(path)
	if err != nil {
		return false, "", err
	}

	var pkName string
	var expFn bool
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.File:
			pkName = x.Name.Name
		case *ast.FuncDecl:
			if x.Name.IsExported() && x.Name.Name == funcName {
				expFn = true
			}
		}
		return true
	})
	return expFn, pkName, nil
}

func hasExportedVariable(path string, varName string) (bool, string, error) {

	node, err := getAstVals(path)
	if err != nil {
		return false, "", err
	}

	hasVar := false
	var pkName string
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.File:
			pkName = x.Name.Name
		case *ast.GenDecl:
			if x.Tok == token.VAR {
				for _, spec := range x.Specs {
					vspec := spec.(*ast.ValueSpec)
					if vspec.Names[0].Name == varName {
						hasVar = true
						break
					}
				}
			}
		}
		return true
	})
	return hasVar, pkName, nil
}

func generateCode(imports utils.StringSet, pages []string, indexGroup []string, dataFunctions []string, renderFunctions []string, handleFunctions []string) error {

	code := `
// Code generated by gox; DO NOT EDIT.
package gox
import (
	` + imports.Join("\n\t") + `
)

var IndexList = map[string]string{
	` + strings.Join(indexGroup, "\n\t") + `
}

var PagesList = map[string]PageData{
	` + strings.Join(pages, "\n\t") + `
}

var DataList = map[string]RenderData{
	` + strings.Join(dataFunctions, "\n\t") + `
}

var RenderList = []RenderDefault{
	` + strings.Join(renderFunctions, "\n\t") + `
}

var HandleList = []HandlerDefault{
	` + strings.Join(handleFunctions, "\n\t") + `
}
`
	log.Println(code)

	err := ioutil.WriteFile("../gox/generated.go", []byte(code), 0644)
	if err != nil {
		return err
	}
	return nil
}

func removeDirWithUnderscorePostfix(path string) string {
	segments := strings.Split(path, "/")
	var output []string
	if len(segments) == 0 {
		return path
	}
	for _, segment := range segments {
		if !strings.HasSuffix(segment, "_") {
			output = append(output, segment)
		}
	}

	return filepath.Join(output...)
}
