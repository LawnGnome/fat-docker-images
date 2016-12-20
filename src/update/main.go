package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"text/template"

	"fat"
)

var manifest string
var reposFile string

func init() {
	flag.StringVar(&manifest, "manifest", "./manifest", "the path to the manifest tool")
	flag.StringVar(&reposFile, "repos", "", "the repositories metadata")
}

func main() {
	flag.Parse()

	repos, err := fat.ParseRepositoriesFile(reposFile)
	if err != nil {
		log.Fatalln(err)
	}

	t, err := template.New("manifest").Parse(`
image: {{.Target}}:{{.Tag}}
manifests:
  {{$tag := .Tag}}
  {{range .Architectures}}
  -
    image: {{.Source}}:{{$tag}}
    platform:
      architecture: {{.Name}}
      os: linux
  {{end}}
`)

	for _, repo := range repos {
		func(repo fat.Repository) {
			tags, err := repo.TagsInAllArchitectures()
			if err != nil {
				log.Fatalln(err)
			}

			for _, tag := range tags {
				func(repo fat.Repository, tag string) {
					// Construct a manifest file.
					file, err := ioutil.TempFile("", "fat")
					if err != nil {
						log.Println(err)
						return
					}
					defer os.Remove(file.Name())

					t.Execute(file, struct {
						Architectures []fat.Architecture
						Tag           string
						Target        string
					}{
						Architectures: repo.Architectures,
						Tag:           tag,
						Target:        repo.Target,
					})
					file.Sync()

					cmd := exec.Command(manifest, "pushml", file.Name())
					output, err := cmd.CombinedOutput()
					if err != nil {
						log.Printf("Error pushing %s:%s: %v; output: %s\n", repo.Target, tag, err, output)
					} else {
						log.Printf("Pushed %s:%s\n", repo.Target, tag)
					}
				}(repo, tag)
			}
		}(repo)
	}
}
