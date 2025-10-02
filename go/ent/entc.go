//go:build ignore

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
	/*	spec := new(ogen.Spec)

		oas, err := entoas.NewExtension(
			entoas.Spec(spec),
			entoas.DefaultPolicy(entoas.PolicyExclude),
			entoas.Mutations(func(_ *gen.Graph, spec *ogen.Spec) error {
				spec.AddPathItem("/login", ogen.NewPathItem().
					SetDescription("JWT login").
					SetPost(ogen.NewOperation().
						SetRequestBody(
							ogen.NewRequestBody().
								SetRequired(true).
								SetDescription("Request a JSON Web Token (JWT) using your login credentials").
								SetJSONContent(
									ogen.NewSchema().
										SetType("object").
										AddRequiredProperties(
											ogen.NewProperty().
												SetName("email").
												SetSchema(ogen.String()),
											ogen.NewProperty().
												SetName("password").
												SetSchema(ogen.Password()),
										))).
						SetOperationID("login").
						SetSummary("Accepts credentials and returns JWT").
						AddTags("Authentication").
						AddResponse("200", ogen.NewResponse().
							SetDescription("Login success").
							SetJSONContent(
								ogen.NewSchema().SetType("object").
									SetProperties(
										&ogen.Properties{
											*ogen.NewProperty().SetName("code").SetSchema(ogen.Int()),
											*ogen.NewProperty().SetName("expires").SetSchema(ogen.DateTime()),
											*ogen.NewProperty().SetName("token").SetSchema(ogen.String()),
										},
									),
							),
						).
						AddResponse("401", ogen.NewResponse().
							SetDescription("Unknown email or password").
							SetRef("#/components/responses/400"),
						),
					),
				)

				return nil
			}),
			entoas.Mutations(func(_ *gen.Graph, spec *ogen.Spec) error {
				spec.AddPathItem("/user/{email}/otk-requests", ogen.NewPathItem().
					SetDescription("Create a new reset password request").
					SetPost(ogen.NewOperation().
						SetOperationID("resetPassword").
						SetSummary("Creates a new OTK and fires email to user to reset their password").
						AddTags("Authentication").
						AddParameters(
							ogen.NewParameter().
								SetName("email").
								SetIn("path").
								SetRequired(true).
								SetSchema(ogen.NewSchema().
									SetType("string"))).
						AddResponse("200", ogen.NewResponse().
							SetDescription("OTK create success; check email for reset link").
							SetJSONContent(
								ogen.NewSchema().
									SetType("object").
									SetProperties(&ogen.Properties{*ogen.NewProperty().
										SetName("success").
										SetSchema(ogen.Bool())},
									))),
					),
				)
				return nil
			}),
			entoas.Mutations(func(_ *gen.Graph, spec *ogen.Spec) error {
				spec.AddPathItem("/user/me", ogen.NewPathItem().
					SetDescription("Get information about the currently logged in user").
					SetGet(ogen.NewOperation().
						SetOperationID("myInfo").
						AddTags("User").
						AddResponse("200", ogen.NewResponse().
							SetDescription("User is authenticated").
							SetJSONContent(
								ogen.NewSchema().
									SetType("object").
									AddRequiredProperties(
										ogen.NewProperty().
											SetName("my_pulid").
											SetSchema(ogen.String()),
										ogen.NewProperty().
											SetName("my_group_pulid").
											SetSchema(ogen.String()),
										ogen.NewProperty().
											SetName("my_tenant_pulid").
											SetSchema(ogen.String()),
									))),
					),
				)
				return nil
			}),
			entoas.Mutations(func(_ *gen.Graph, spec *ogen.Spec) error {
				spec.AddNamedSchemas(ogen.NewNamedSchema("SomeTYpe", ogen.NewSchema().
					SetType("object").
					SetProperties(
						&ogen.Properties{
							*ogen.NewProperty().
								SetName("success").
								SetSchema(ogen.Bool()),
						},
					),
				))
				return nil
			}),
		)
		if err != nil {
			log.Fatalf("creating entoas extension: %v", err)
		}*/

	gen.AddAcronym("GLS")
	gen.AddAcronym("DHL")

	graphql, err := entgql.NewExtension(
		entgql.WithConfigPath("./gqlgen.yml"),
		// Generate GQL schema from the Ent's schema.
		entgql.WithSchemaGenerator(),
		// Generate the where to a separate schema
		// file and load it in the gqlgen.yml config.
		entgql.WithSchemaPath("./ent.graphql"),
		entgql.WithWhereFilters(true),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	tmplate := gen.NewTemplate("matches.tmpl")
	tmplate.Funcs(template.FuncMap{
		"last": func(x int, a interface{}) bool {
			value := reflect.ValueOf(a)
			return x == reflect.Indirect(value).Len()-1
		},
	})
	_, err = tmplate.ParseFiles("./ent/schema/templates/matches.tmpl")
	if err != nil {
		log.Fatalf("generating matches template: %v", err)
	}

	opts := []entc.Option{
		entc.Dependency(
			entc.DependencyName("Bucket"),
			entc.DependencyType(&blob.Bucket{}),
		),
		entc.Extensions(graphql),
		//entc.TemplateFiles("./ent/schema/templates/matches.tmpl"),
		entc.FeatureNames("privacy", "entql", "sql/upsert", "sql/versioned-migration", "intercept"),
	}

	err = entc.Generate(
		"./ent/schema",
		&gen.Config{
			Features: []gen.Feature{gen.FeatureVersionedMigration},
			Templates: []*gen.Template{
				tmplate,
			},
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
