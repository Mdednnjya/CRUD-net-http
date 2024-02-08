package handler

import (
	"encoding/JSON"
	"net/http"
	"strconv"

	"CRUD-nethttp/model"
)

type HTTPHandler struct {
	Candidates []*model.Candidate
	CandidateNum int
}

func (h *HTTPHandler) HandleListAndCreate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.read(r, w)
	case http.MethodPost:
		h.create(r, w)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *HTTPHandler) HandleDetailAndModify(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.readByID(r, w)
	case http.MethodPut:
		h.update(r, w)
	case http.MethodDelete:
		h.delete(r, w)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}


func (h *HTTPHandler) create(r *http.Request,w http.ResponseWriter) {
	var newCandidate model.Candidate

	if err := json.NewDecoder(r.Body).Decode(&newCandidate); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}

	h.CandidateNum++
	h.Candidates = append(h.Candidates, &newCandidate)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(h.Candidates)
}

func (h *HTTPHandler) read(r *http.Request,w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(h.Candidates)
}

func (h *HTTPHandler) readByID(r *http.Request,w http.ResponseWriter) {
	param := r.URL.Path[len("/Candidates/"):]
	var candidate *model.Candidate
	candidateID, err := strconv.Atoi(param)
	if err != nil {
		http.Error(w, "Invalid candidate ID", http.StatusBadRequest)
		return
	}
	for _, c := range h.Candidates {
		if candidateID != c.CandidateNumber {
			continue
		}
		candidate = c
		break
	}
	if candidate == nil {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(candidate)
}

func (h *HTTPHandler) update(r *http.Request,w http.ResponseWriter) {
	param := r.URL.Path[len("/Candidates/"):]
	var Candidate model.Candidate
	if err := json.NewDecoder(r.Body).Decode(&Candidate); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}
	candidateID, err := strconv.Atoi(param)
	if err != nil {
		http.Error(w, "Invalid candidate ID", http.StatusBadRequest)
		return
	}
	for _, c := range h.Candidates {
		if candidateID == c.CandidateNumber {
            c.Name = Candidate.Name
            c.Vision = Candidate.Vision
            c.Mission = Candidate.Mission
			break
		}
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(h.Candidates)
}
	
func (h *HTTPHandler) delete(r *http.Request,w http.ResponseWriter) {
	param := r.URL.Path[len("/Candidates/"):]
	var updatedCandidates []*model.Candidate
	candidateID, err := strconv.Atoi(param)
	if err != nil {
		http.Error(w, "Invalid candidate ID", http.StatusBadRequest)
		return
	}
	for _, c := range h.Candidates {
		if candidateID == c.CandidateNumber {
				continue
		}
		updatedCandidates = append(updatedCandidates, c)
	}
	h.Candidates = updatedCandidates
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(h.Candidates)
}