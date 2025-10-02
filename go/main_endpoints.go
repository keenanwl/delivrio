package main

import (
	delivrio "delivrio.io"
	"delivrio.io/go/appconfig"
	"delivrio.io/go/endpoints"
	"delivrio.io/go/endpoints/delivrioroutes"
	"delivrio.io/go/gengql"
	"delivrio.io/go/restcustomer"
	"delivrio.io/go/shopify/ratelookup"
	"delivrio.io/go/utils"
	"delivrio.io/go/viewer"
	"embed"
	"entgo.io/contrib/entgql"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"github.com/go-chi/jwtauth/v5"
	"io/fs"
	"net/http"
	"os"
	"path"
	"strings"
)

//go:embed images/*
var staticReturnImages embed.FS

//go:embed restcustomer/swagger-ui/4-18-1/*
var swaggerFiles embed.FS

func configureRouter(port string, conf appconfig.DelivrioConfig, logger *httplog.Logger) (*chi.Mux, error) {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(httplog.RequestLogger(logger))
	router.Use(serverNameMiddleware(conf.ServerID))

	router.Route("/rest/v1", func(r chi.Router) {
		r.Use(endpoints.AddSpan(conf.ServerID, "customerREST"))
		r.Use(endpoints.AddClient(client))
		r.Use(endpoints.CheckRESTAPICredentials())

		r.Get(delivrioroutes.Label, endpoints.GetLabels)
		r.Post(delivrioroutes.Label, endpoints.GetLabels)
		r.Get(delivrioroutes.Order, restcustomer.HandleOrderGet)
		r.Post(delivrioroutes.Orders, restcustomer.HandleOrderCreate)
		r.Post(delivrioroutes.Products, restcustomer.HandleProductsCreate)
		r.Post(delivrioroutes.Shipments, restcustomer.HandleShipmentsCreate)
		r.Get(delivrioroutes.Return, restcustomer.HandleReturnGet)
		r.Get(delivrioroutes.ReturnReasons, restcustomer.HandleReturnReasonsGet)
	})

	router.Route(delivrioroutes.API, func(r chi.Router) {
		r.Route("/", func(r chi.Router) {
			r.Use(endpoints.AddSpan(conf.ServerID, "internalGraphQLMutation"))
			r.Use(endpoints.AddTX(client))
			r.Use(viewer.AddAnonymousViewer())

			r.Get(delivrioroutes.AddressLookup, endpoints.AddressAutocompleteHandler)
			r.Post(delivrioroutes.RequestEmail, endpoints.RequestEmailHandler)
			r.Post(delivrioroutes.RequestPassword, endpoints.ResetPasswordHandler)
			r.Post(delivrioroutes.Register, endpoints.RegistrationHandler)
			r.Post(delivrioroutes.Login, authHandler(conf))
		})

		InternalRESTMiddleware := []func(http.Handler) http.Handler{
			endpoints.AddSpan(conf.ServerID, "internalREST"),
			endpoints.AddClient(client),
			// Viewer may be added by endpoints based on specific authentication
			viewer.AddAnonymousViewer(),
		}
		r.With(InternalRESTMiddleware...).Get(delivrioroutes.NodeHealth, endpoints.HandlerHealthCheck)

		r.With(InternalRESTMiddleware...).Post(delivrioroutes.ShopifyLookup, ratelookup.ShopifyLookupHandler)
		r.With(InternalRESTMiddleware...).Post(delivrioroutes.ShopifyLookupPickupPoints, ratelookup.PickupPointsHandler)
		// Refactor to Get
		// No tx due to simple mutations
		r.With(InternalRESTMiddleware...).Post(delivrioroutes.PrintClientPing, endpoints.PrintClientPing)
		r.With(InternalRESTMiddleware...).Post(delivrioroutes.PrintClientRequestLabel, endpoints.PrintClientRequestLabel)
		r.With(InternalRESTMiddleware...).Post(delivrioroutes.RegisterScan, endpoints.ScannedLabelRegister)
		r.With(InternalRESTMiddleware...).Post(delivrioroutes.ScanList, endpoints.ScannedLabelList)
		r.With(InternalRESTMiddleware...).Post(delivrioroutes.UpdateStatus, endpoints.StatusUpdate)
		r.With(InternalRESTMiddleware...).Post(delivrioroutes.SignatureUpload, endpoints.SignatureUpload)

		returnMiddleware := []func(http.Handler) http.Handler{
			endpoints.AddSpan(conf.ServerID, "internalREST"),
			endpoints.AddClient(client),
			// Viewer may be added by endpoints based on specific authentication
			viewer.AddAnonymousViewer(),
			endpoints.AllowCors,
		}
		r.With(returnMiddleware...).Get(delivrioroutes.ReturnView, endpoints.ReturnOrderView)
		r.With(returnMiddleware...).Post(delivrioroutes.CreateReturnOrder, endpoints.CreateReturnOrder)
		r.With(returnMiddleware...).Options(delivrioroutes.CreateReturnOrder, endpoints.AllowCorsHandler)
		r.With(returnMiddleware...).Post(delivrioroutes.ReturnDeliveryOptions, endpoints.AddReturnOrderDeliveryOptionsAndRequestLabel)
		r.With(returnMiddleware...).Options(delivrioroutes.ReturnDeliveryOptions, endpoints.AllowCorsHandler)
		r.With(returnMiddleware...).Get(
			delivrioroutes.ReturnLabelDownload,
			endpoints.ReturnColliFileDownloadHandler(endpoints.ReturnLabelDownloadWriter),
		)
		r.With(returnMiddleware...).Get(
			delivrioroutes.ReturnLabelPNG,
			endpoints.ReturnColliFileViewerHandler(endpoints.ReturnLabelPNGViewerWriter),
		)
		r.With(returnMiddleware...).Get(
			delivrioroutes.ReturnQRCodeDownload,
			endpoints.ReturnColliFileDownloadHandler(endpoints.ReturnQRCodeDownloadWriter),
		)
		r.With(returnMiddleware...).Get(
			delivrioroutes.ReturnLabel,
			endpoints.ReturnColliFileViewerHandler(endpoints.ReturnLabelViewerWriter),
		)
		r.With(returnMiddleware...).Get(
			delivrioroutes.ReturnQRCode,
			endpoints.ReturnColliFileViewerHandler(endpoints.ReturnQRCodeViewerWriter),
		)

		r.With(InternalRESTMiddleware...).Post(delivrioroutes.QueryGoland, graphqlHandler())

		InternalGraphQLMiddleware := []func(http.Handler) http.Handler{
			jwtauth.Verifier(tokenAuth),
			endpoints.AddSpan(conf.ServerID, "internalGraphQL"),
			endpoints.AddClient(client),
			viewer.AddJWTViewer(),
		}
		r.With(InternalGraphQLMiddleware...).Post(delivrioroutes.Query, graphqlHandler())
	})

	router.Route(delivrioroutes.Static, func(r chi.Router) {
		authMiddleware := []func(http.Handler) http.Handler{
			jwtauth.Verifier(tokenAuth),
		}
		r.With(authMiddleware...).Handle(
			delivrioroutes.Uploads,
			http.FileServer(http.Dir("./uploads")),
		)
		r.With(authMiddleware...).Get(delivrioroutes.Restv1APIDocs, func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, fmt.Sprintf("%s%s/", delivrioroutes.Static, delivrioroutes.Restv1APIDocs), http.StatusFound)
		})
		r.With(authMiddleware...).Handle(
			fmt.Sprintf("%s/*", delivrioroutes.Restv1APIDocs),
			fsHandler(utils.RemoveFsPrefix(
				swaggerFiles,
				"restcustomer/swagger-ui/4-18-1"),
				fmt.Sprintf("%s%s", delivrioroutes.Static, delivrioroutes.Restv1APIDocs)),
		)

		r.With(endpoints.AllowCors).Get(delivrioroutes.ReturnPolyfills, func(w http.ResponseWriter, r *http.Request) {
			// Set cache control headers to prevent caching
			// Since we aren't hashing the output
			w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
			w.Header().Set("Pragma", "no-cache")
			w.Header().Set("Expires", "0")

			// Set the Content-Type header
			w.Header().Set("Content-Type", "application/javascript")

			// Serve the content of the embedded file
			_, err := w.Write([]byte(delivrio.FrontendReturnPolyfills))
			if err != nil {
				// Handle the error, e.g., log or return an error response
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		})

		r.With(endpoints.AllowCors).Get(delivrioroutes.ReturnMain, func(w http.ResponseWriter, r *http.Request) {
			// Set cache control headers to prevent caching
			// Since we aren't hashing the output
			w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
			w.Header().Set("Pragma", "no-cache")
			w.Header().Set("Expires", "0")

			// Set the Content-Type header
			w.Header().Set("Content-Type", "application/javascript")

			// Serve the content of the embedded file
			_, err := w.Write([]byte(delivrio.FrontendReturnMain))
			if err != nil {
				// Handle the error, e.g., log or return an error response
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		})

		r.Handle(
			fmt.Sprintf("%s/*", delivrioroutes.Images),
			fsHandler(
				utils.RemoveFsPrefix(staticReturnImages, "images"),
				fmt.Sprintf("%s%s", delivrioroutes.Static, delivrioroutes.Images),
			),
		)
		router.Get("/*", fsHandler(utils.RemoveFsPrefix(delivrio.FrontendContent, "ng/dist/delivery/browser"), "/"))
	})

	router.Get(delivrioroutes.Graph, playgroundHandler())

	return router, nil
}

func serverNameMiddleware(serverName string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("x-delivrio-origin", serverName)
			next.ServeHTTP(w, r)
		})
	}
}

// urlPrefix removes the URL.Path so Path does not have to
// match the embedded FS path
func fsHandler(embedded fs.FS, urlPrefix string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		baseTrim := strings.TrimPrefix(path.Clean(r.URL.Path), urlPrefix)
		f, err := embedded.Open(baseTrim)
		if err == nil {
			defer f.Close()
		}

		if os.IsNotExist(err) {
			r.URL.Path = "/"
		} else {
			r.URL.Path = baseTrim
		}
		http.FileServer(http.FS(embedded)).ServeHTTP(w, r)
	}
}

// Defining the Graphql handler
func graphqlHandler() http.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(gengql.NewSchema(client))
	h.Use(entgql.Transactioner{TxOpener: client})

	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}

// Defining the Playground handler
func playgroundHandler() http.HandlerFunc {
	h := playground.Handler("GraphQL", "/api/query")

	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}
