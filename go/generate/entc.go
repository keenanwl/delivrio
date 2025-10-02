//go:build ignore
// +build ignore

package main

import (
	"log"
	"reflect"
	"text/template"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"gocloud.dev/blob"
)

func main() {

	gen.AddAcronym("DAO")
	gen.AddAcronym("DF")
	gen.AddAcronym("DSV")
	gen.AddAcronym("GLS")
	gen.AddAcronym("DHL")
	gen.AddAcronym("USPS")

	graphql, err := entgql.NewExtension(
		entgql.WithConfigPath("./gqlgen.yml"),
		// Generate GQL schema from the Ent's schema.
		entgql.WithSchemaGenerator(),
		// Generate the where to a separate schema
		// file and load it in the gqlgen.yml config.
		entgql.WithSchemaPath("./gengql/ent.graphql"),
		entgql.WithWhereFilters(true),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	templateMatches := gen.NewTemplate("matches.tmpl")
	templateMatches.Funcs(template.FuncMap{
		"last": func(x int, a interface{}) bool {
			value := reflect.ValueOf(a)
			return x == reflect.Indirect(value).Len()-1
		},
	})
	_, err = templateMatches.ParseFiles("./schema/templates/matches.tmpl")
	if err != nil {
		log.Fatalf("generating matches template: %v", err)
	}

	templateClone := gen.NewTemplate("cloneable.tmpl")
	templateClone.Funcs(template.FuncMap{})
	_, err = templateClone.ParseFiles("./schema/templates/cloneable.tmpl")
	if err != nil {
		log.Fatalf("generating clone template: %v", err)
	}

	opts := []entc.Option{
		entc.Dependency(
			entc.DependencyName("Bucket"),
			entc.DependencyType(&blob.Bucket{}),
		),
		entc.Extensions(graphql),
		entc.FeatureNames("privacy", "entql", "sql/upsert", "sql/versioned-migration", "intercept", "schema/snapshot"),
	}

	err = entc.Generate(
		"delivrio.io/go/schema",
		&gen.Config{
			Features: []gen.Feature{gen.FeatureVersionedMigration},
			Templates: []*gen.Template{
				templateMatches,
				templateClone,
			},
			Target:  "./ent",
			Package: "delivrio.io/go/ent",
		},
		opts...,
	)
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}

	/*	file, err := os.ReadFile("./ent/openapi.json")
		if err != nil {
			log.Fatalf("reading JSON OpenAPI file: %v", err)
		}

		yml, err := yaml.JSONToYAML(file)
		if err != nil {
			log.Fatalf("converting JSON to yaml: %v", err)
		}

		err = ioutil.WriteFile("./ent/openapi.yaml", yml, 0777)
		if err != nil {
			log.Fatalf("writing ent/openapi.yaml file: %v", err)
		}*/

}
