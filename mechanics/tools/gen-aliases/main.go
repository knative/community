package main

import (
	"fmt"
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

	f, err := os.ReadFile(infile)
	if err != nil {
		log.Fatal("Unable to open:", err)
	}
	var cfg peribolos.FullConfig
	if err := yaml.Unmarshal(f, &cfg); err != nil {
		log.Fatal("Unable to parse:", err)
	}
	out := AliasConfig{
		Aliases: map[string][]string{},
	}
	dashes := func(r rune) rune {
		if r >= 'a' && r <= 'z' {
			return r
		}
		if r >= '0' && r <= '9' {
			return r
		}
		return '-'
	}
	for _, team := range expandTeams(cfg.Orgs[org].Teams) {
		name := strings.Map(dashes, strings.ToLower(team.Name))
		out.Aliases[name] = sets.List(team.Members)
	}

	output, err := yaml.Marshal(out)
	if err != nil {
		log.Fatal("Could not serialize:", err)
	}
	preamble := fmt.Sprintf(`# This file is auto-generated from peribolos.
# Do not modify this file, instead modify %s

`, infile)
	output = append([]byte(preamble), output...)
	prevOut, err := os.ReadFile(outfile)

	if err == nil && string(prevOut) == string(output) {
		log.Print("No changes needed for ", outfile)
		return
	} else if err != nil {
		log.Fatal("Failed to read file:", err)
	}

	if err := os.WriteFile(outfile, output, 0o644); err != nil {
		log.Fatal("Failed to write file:", err)
	}

	log.Print("Wrote ", outfile)
}

func expandTeams(seed map[string]peribolos.Team) []expandedTeam {
	var retval []expandedTeam
	for name, team := range seed {
		children := expandTeams(team.Children)
		this := expandedTeam{
			Name:    name,
			Members: sets.New(team.Members...),
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
	Members sets.Set[string]
}
