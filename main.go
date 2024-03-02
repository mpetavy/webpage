package main

import (
	"context"
	"embed"
	"fmt"
	"github.com/mpetavy/common"
	"github.com/mpetavy/webpage/components"
	"net/http"
	"time"
)

var (
	httpServer      *http.Server
	ctxServer       context.Context
	ctxServerCancel context.CancelFunc
	tlsInfo         string
)

//go:embed go.mod
var resources embed.FS

func init() {
	common.Init("", "", "", "", "webpage", "", "", "", &resources, start, stop, nil, 0)
}

func badRequest(w http.ResponseWriter, msg string) {
	http.Error(w, msg, http.StatusBadRequest)
}

func interalError(w http.ResponseWriter, msg string) {
	http.Error(w, msg, http.StatusInternalServerError)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := func() error {
		page, err := components.NewPage()
		if common.Error(err) {
			return err
		}

		html, err := page.HTML()
		if common.Error(err) {
			return err
		}

		_, err = w.Write([]byte(html))
		if common.Error(err) {
			return err
		}

		return nil
	}()

	if common.Error(err) {
		badRequest(w, err.Error())

	}
}

func methodHandler(method string, next http.HandlerFunc) http.HandlerFunc {
	common.DebugFunc()

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method {
			next.ServeHTTP(w, r)

			return
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	}
}

func start() error {
	common.DebugFunc()

	mux := http.NewServeMux()
	mux.HandleFunc("/", methodHandler(http.MethodGet, homeHandler))

	tlsConfig, err := common.NewTlsConfigFromFlags()
	if common.Error(err) {
		return err
	}

	if tlsConfig != nil {
		tlsInfo = " [TLS]"
	}

	timeout := time.Second * 10
	port := 8100

	httpServer = &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           mux,
		TLSConfig:         tlsConfig,
		ReadTimeout:       timeout,
		ReadHeaderTimeout: timeout,
		WriteTimeout:      timeout,
		ErrorLog:          common.LogError,
	}
	httpServer.SetKeepAlivesEnabled(false)

	common.Info(fmt.Sprintf("Server%s started on port: %d", tlsInfo, port))

	ctxServer, ctxServerCancel = context.WithCancel(context.Background())

	go func() {
		if tlsConfig != nil {
			err = httpServer.ListenAndServeTLS("", "")
		} else {
			err = httpServer.ListenAndServe()
		}
	}()

	time.Sleep(common.MillisecondToDuration(*common.FlagServiceTimeout))

	if err != nil && err == http.ErrServerClosed {
		<-ctxServer.Done()

		err = nil
	}

	if common.Error(err) {
		return err
	}

	return nil
}

func stop() error {
	common.DebugFunc()

	if httpServer == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := httpServer.Shutdown(ctx)
	if common.Error(err) {
		return err
	}

	common.Info(fmt.Sprintf("Server%s closed", tlsInfo))

	httpServer = nil

	ctxServerCancel()

	return nil
}

func main() {
	common.Run(nil)
}
