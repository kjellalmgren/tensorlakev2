package main

/*
Services: tensorlake
	Author: Kjell Osse Almgren, Tetracon AB
	Date: 2020-04-12
	Description: Service to feed Selma-SME UX, just for test purpose
	Compiled: go 1.14.2 linux/amd64
	Architecture:
	win32: GOOS=windows GOARCH=386 go build -v
	win64: GOOS=windows GOARCH=amd64 go build -v
	arm64: GOOS=linux GOARCH=arm64 go build -v
	arm: GOOS=linux GOARCH=arm go build -v
	exprimental: GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags '-w -s' -a -installsuffix cgo -o selmasme
	expriemntal: CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -tags selmasme -ldflags '-w'
	exprimental: GOOS=linux GOARCH=arm64 go build -a --ldflags 'extldflags "-static"' -tags selmasme -installsuffix selma-sme .
*/

//

import (
	"crypto/x509"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"tensorlake/segments"
	"tensorlake/version"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

//
//	Processes
//
// Banner text
const (
	// TETRACON banner
	TETRACON = `
_________    __
|__    __|   | |
   |  |  ___ | |_   ____  ___   ___ ___  _ __ 
   |  | / _ \|  _| /  __|/ _ \ / __/ _ \| '_ \
   |  | \ __/| |_  | |  | (_| | (_| (_) | | | | 
   |__| \___| \__| |_|   \__,_|\___\___/|_| |_| 
Server-version: %s Model-version: %s Model-date: %s
`
)

//
var (
	srv  bool
	vrsn bool
	date bool
)

var (
	client *http.Client
	pool   *x509.CertPool
)

// init documwentation
func init() {

	// instanciate a new logger
	var log = logrus.New()
	log.Formatter = new(logrus.TextFormatter)
	log.Level = logrus.DebugLevel
	color.Set(color.FgHiGreen)
	fmt.Fprint(os.Stderr, fmt.Sprintf(TETRACON, version.ServerVersion(), version.ModelVersion(), version.ModelDate()))
	color.Unset()
}

//
// our main function
func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/segment_training_v2", segments.SegmentTrainingV2).Methods("GET", "POST", "OPTIONS")
	router.HandleFunc("/segment_training_v3", segments.SegmentTrainingV3).Methods("GET", "POST", "OPTIONS")
	router.HandleFunc("/segment_training_v31", segments.SegmentTrainingV31).Methods("GET", "POST", "OPTIONS")
	router.HandleFunc("/segment_training_v32", segments.SegmentTrainingV32).Methods("GET", "POST", "OPTIONS")
	router.HandleFunc("/segment_training_v4", segments.SegmentTrainingV4).Methods("GET", "POST", "OPTIONS")
	router.HandleFunc("/segment_evaluation_v4", segments.SegmentEvaluationV4).Methods("GET", "POST", "OPTIONS")
	router.HandleFunc("/iris_training_v1", segments.IrisTrainingV1).Methods("GET", "POST", "OPTIONS")
	router.HandleFunc("/iris_evaluation_v1", segments.IrisEvaluationV1).Methods("GET", "POST", "OPTIONS")
	//
	// Healt services local
	router.HandleFunc("/ping", HealthCheckHandler).Methods("GET")
	color.Set(color.FgHiRed)
	fmt.Fprint(os.Stdout, fmt.Sprintf(getHostname()))
	color.Set(color.FgHiGreen)
	fmt.Printf(" Listen on TLS-server ")
	color.Set(color.FgHiYellow)
	fmt.Printf("localhost")
	color.Set(color.FgHiCyan)
	fmt.Printf(":8443\r\n")
	color.Unset()
	//
	// ===================================================================================
	/*
		err := http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", router)

			if err != nil {
				fmt.Printf("ListenAndServeTLS Error: %s", err.Error())
				log.Fatal(err)
				logrus.Fatal(err)
			}
	*/
	err := http.ListenAndServe(":8443", router)
	if err != nil {
		fmt.Printf("ListenAndServeTLS Error: %s", err.Error())
		log.Fatal(err)
		logrus.Fatal(err)
	}
	// ===================================================================================
	// ===================================================================================
	// Create a CA certificate pool and add cert.pem to it
	//caCert, err := ioutil.ReadFile("cert.pem")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//caCertPool := x509.NewCertPool()
	//caCertPool.AppendCertsFromPEM(caCert)

	// Create the TLS Config with the CA pool and enable Client certificate validation
	// ##
	// ClientAuth: tls.RequireAndVerifyClientCert,
	// ##
	//tlsConfig := &tls.Config{
	//	ClientCAs:  caCertPool,
	//	ClientAuth: tls.NoClientCert,
	//}
	//tlsConfig.InsecureSkipVerify = false
	//tlsConfig.BuildNameToCertificate()

	// Create a Server instance to listen on port 8443 with the TLS config
	//server := &http.Server{
	//	Addr:      ":8443",
	//	TLSConfig: tlsConfig,
	//}

	// Listen to HTTPS connections with the server certificate and wait
	//log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))

	// ===================================================================================
}

//
//	getHostName documentation
func getHostname() string {

	hostname, err1 := os.Hostname()
	if err1 != nil {
		//log.Printf("Hostname: %s", hostname)
		fmt.Println("Error when try to resolve Hostname: ", hostname)
	}
	return hostname
}

//
// HealthCheckHandler documentation
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	//w.WriteHeader(http.StatusOK)
	//w.Header().Set("Content-Type", "application/json")

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	//io.WriteString(w, `[{"alive": "true",`)
	//io.WriteString(w, `"status": "200"}]`)

	type HealthCheck struct {
		HealthStatus int `json:"healthStatus,omitempty"`
	}
	var healtCheck HealthCheck
	healtCheck.HealthStatus = http.StatusOK
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//
	if err := json.NewEncoder(w).Encode(healtCheck); err != nil {
		panic(err)
	}
	fmt.Printf("Http-Status %d received\r\n", http.StatusOK)
}
