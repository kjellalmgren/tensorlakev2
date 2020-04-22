package segments

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// SegmentTrainingV2 documentation, this holds the end-point to API definitions
func SegmentTrainingV2(w http.ResponseWriter, r *http.Request) {

	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me")
	//
	file, err := ioutil.ReadFile("json/segment_training_v2.csv")
	if err != nil {
		fmt.Fprintf(w, "Error reading segment_training_v2.csv - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write(file)
	w.WriteHeader(http.StatusOK)
}

// SegmentTrainingV3 documentation, this holds the end-point to API definitions
func SegmentTrainingV3(w http.ResponseWriter, r *http.Request) {

	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me")
	//
	file, err := ioutil.ReadFile("json/segment_training_v3.csv")
	if err != nil {
		fmt.Fprintf(w, "Error reading segment_training_v3.csv - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write(file)
	w.WriteHeader(http.StatusOK)
}

// SegmentTrainingV31 documentation, this holds the end-point to API definitions
func SegmentTrainingV31(w http.ResponseWriter, r *http.Request) {

	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me")
	//
	file, err := ioutil.ReadFile("json/segment_training_v31.csv")
	if err != nil {
		fmt.Fprintf(w, "Error reading segment_training_v31.csv - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write(file)
	w.WriteHeader(http.StatusOK)
}

// SegmentTrainingV32 documentation, this holds the end-point to API definitions
func SegmentTrainingV32(w http.ResponseWriter, r *http.Request) {

	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me")
	//
	file, err := ioutil.ReadFile("json/segment_training_v32.csv")
	if err != nil {
		fmt.Fprintf(w, "Error reading segment_training_v32.csv - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write(file)
	w.WriteHeader(http.StatusOK)
}

// SegmentTrainingV4 documentation, this holds the end-point to API definitions
func SegmentTrainingV4(w http.ResponseWriter, r *http.Request) {

	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me")
	//

	file, err := ioutil.ReadFile("json/segment_training_v4.csv")
	if err != nil {
		fmt.Fprintf(w, "Error reading segment_training_v4.csv - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write(file)
	w.WriteHeader(http.StatusOK)
}

// SegmentEvaluationV4 documentation, this holds the end-point to API definitions
func SegmentEvaluationV4(w http.ResponseWriter, r *http.Request) {

	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me")
	//

	file, err := ioutil.ReadFile("json/segment_evaluation_v4.csv")
	if err != nil {
		fmt.Fprintf(w, "Error reading segment_evaluation_v4.csv - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write(file)
	w.WriteHeader(http.StatusOK)
}

// IrisTrainingV1 documentation, this holds the end-point to API definitions
func IrisTrainingV1(w http.ResponseWriter, r *http.Request) {

	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me")
	//

	file, err := ioutil.ReadFile("json/iris_training_v1.csv")
	if err != nil {
		fmt.Fprintf(w, "Error reading iris_training_v1.csv - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write(file)
	w.WriteHeader(http.StatusOK)
}

// IrisEvaluationV1 documentation, this holds the end-point to API definitions
func IrisEvaluationV1(w http.ResponseWriter, r *http.Request) {

	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me")
	//

	file, err := ioutil.ReadFile("json/iris_evaluation_v1.csv")
	if err != nil {
		fmt.Fprintf(w, "Error reading iris_evaluation_v1.csv - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write(file)
	w.WriteHeader(http.StatusOK)
}
