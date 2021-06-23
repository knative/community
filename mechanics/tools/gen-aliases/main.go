package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"k8s.io/apimachinery/pkg/util/sets"
	peribolos "k8s.io/test-infra/prow/config/org"
	"sigs.k8s.io/yaml"
)

func main() {
	if len(os.Args) < 2 {
		os.Args = append(os.Args, "knative")
	}
	org := os.Args[1]
	if len(os.Args) < 3 {
		os.Args = append(os.Args, filepath.Join("peribolos", os.Args[1]+".yaml"))
	}
	if len(os.Args) < 4 {
		os.Args = append(os.Args, "OWNERS_ALIASES")
	}
	infile, outfile := os.Args[2], os.Args[3]

	f, err := ioutil.ReadFile(infile)
	if err != nil {
		log.Print("Unable to open ", err)
		os.Exit(1)
	}
	var cfg peribolos.FullConfig
	if err := yaml.Unmarshal(f, &cfg); err != nil {
		log.Printf("Unable to parse ", err)
		os.Exit(1)
	}
	out := AliasConfig{
		Aliases: map[string][]string{},
	}
	dashes := func(r rune) rune {
		if r >= 'a' && r <= 'z' {
			return r
		}
		return '-'
	}
	for _, team := range expandTeams(cfg.Orgs[org].Teams) {
		name := strings.Map(dashes, strings.ToLower(team.Name))
		out.Aliases[name] = team.Members.List()
	}

	output, err := yaml.Marshal(out)
	if err != nil {
		log.Printf("Could not serialize: ", err)
		os.Exit(1)
	}
	preamble := fmt.Sprintf(`# This file is auto-generated from peribolos.
# Do not modify this file, instead modify %s

`, infile)
	output = append([]byte(preamble), output...)
	ioutil.WriteFile(outfile, output, 0644)
	log.Print("Wrote ", outfile)
}

func expandTeams(seed map[string]peribolos.Team) []expandedTeam {
	var retval []expandedTeam
	for name, team := range seed {
		children := expandTeams(team.Children)
		this := expandedTeam{
			Name:    name,
			Members: sets.NewString(team.Members...),
		}
		this.Members.Insert(team.Maintainers...)
		for _, child := range children {
			this.Members = this.Members.Union(child.Members)
		}
		retval = append(retval, this)
		retval = append(retval, children...)
	}

	return retval
}

type AliasConfig struct {
	Aliases map[string][]string `json:"aliases"`
}

type expandedTeam struct {
	Name    string
	Members sets.String
}
