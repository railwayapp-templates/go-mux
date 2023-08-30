# [Gorilla Mux](https://github.com/gorilla/mux) Template

This template deploys a base ready-to-use Gorilla Mux app.

[![Deploy on Railway](https://railway.app/button.svg)](https://railway.app/template/_eGomU?referralCode=ySCnWl)

## Features

 - Two examples of custom middleware:
    
    - **TrustProxy:** Checks if the request IP matches one of the provided ranges/IPs, then inspects common reverse proxy headers and sets the corresponding, fields in the HTTP request struct for use by middleware or handlers that are next.

        Useful for trusting the `X-Forwarded-For` and `X-Forwarded-Proto` headers that Railway's proxy sets.

    - **Logger:** The Logger middleware gathers metrics from the upstream handlers (status code, duration, bytes written) and logs them to stdout.

- Comes with some helpful internal packages:

    - **logger:** The internal logger package is centered around Go's [slog package](https://pkg.go.dev/golang.org/x/exp/slog) but has some pre-configured loggers for added ease of use.

    - **router:** This is where the Gorilla Mux router is initialized and global middlewares are registered like the `TrustProxy` and `Logger`middlewares. Along with registering some handlers and path prefixes.

    - **responder:** Comes with a few utility functions to send json (formatted and unformatted) or plaintext responses to the http client.

    - **tools:** Simple tools package, currently only has a `EnvPortOr` function that reads the `PORT` variable from the environment or falls back to the provided port.

    - **server:** This is where the `http.Server` settings are configured and the server is started from.

    - **routes:** Where all the handlers and paths live, this template comes with the following prefixes and handlers:

        `/`: Returns a greeting message.
        
        `/health`: Returns just a 200 status code.

        `/api/`: The api prefix, with a subrouter.

        `/api/v1/`: The v1 prefix with a subrouter, registered on the `/api/` subrouter.

        `/api/v1/temperature`: Simple mock api endpoint.

        `/api/v1/forecast/3day`: Just another mock api endpoint, but with a path parameter.

## How to use

- Download the Go modules: `go mod download`

- Start the server: `go run main.go`

- Open the browser or any api test program to `http://127.0.0.1:3000`

- Start coding!

## Notes

- The http server listens on port `3000` if no environment variable `PORT` is found, and once on Railway a `PORT` variable is automatically generated, and `EnvPortOr` will use that `PORT` instead of `3000`.