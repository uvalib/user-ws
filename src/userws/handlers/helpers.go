package handlers

import (
   "encoding/json"
   "log"
   "net/http"
   "strings"
   "userws/api"
)

func encodeStandardResponse(w http.ResponseWriter, status int, user *api.User) {
   jsonAttributes(w)
   w.WriteHeader(status)
   if err := json.NewEncoder(w).Encode(api.StandardResponse{Status: status, Message: http.StatusText(status), User: user}); err != nil {
      log.Fatal(err)
   }
}

func encodeHealthCheckResponse(w http.ResponseWriter, healthy bool, message string) {
   status := http.StatusOK
   if healthy == false {
      status = http.StatusInternalServerError
   }
   jsonAttributes(w)
   w.WriteHeader(status)
   if err := json.NewEncoder(w).Encode(api.HealthCheckResponse{CheckType: api.HealthCheckResult{Healthy: healthy, Message: message}}); err != nil {
      log.Fatal(err)
   }
}

func encodeVersionResponse(w http.ResponseWriter, status int, version string) {
   jsonAttributes(w)
   w.WriteHeader(status)
   if err := json.NewEncoder(w).Encode(api.VersionResponse{Version: version}); err != nil {
      log.Fatal(err)
   }
}

func encodeRuntimeResponse(w http.ResponseWriter, status int, version string, cpus int, goroutines int, heapcount uint64, alloc uint64) {
   jsonAttributes(w)
   w.WriteHeader(status)
   if err := json.NewEncoder(w).Encode(api.RuntimeResponse{Version: version, CPUCount: cpus, GoRoutineCount: goroutines, ObjectCount: heapcount, AllocatedMemory: alloc}); err != nil {
      log.Fatal(err)
   }
}

func jsonAttributes(w http.ResponseWriter) {
   w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

func parametersOk(userID string, token string) bool {
   // validate inbound parameters
   return len(strings.TrimSpace(userID)) != 0 &&
      len(strings.TrimSpace(token)) != 0

}

//
// end of file
//

