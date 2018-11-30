package main

import (
	"bytes"
	"github.com/cocoagaurav/httpHandler/handler"
	"github.com/go-chi/chi"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTestproc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testproc Suite")
}

var _ = Describe("test the register handler", func() {
	r := chi.NewRouter()
	It("will run the register handler", func() {
		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte(`{"emailid":"bharadwaj@api.com","password":"simple","name":"bharadwaj","age":23}`)))
		rr := httptest.NewRecorder()
		r.HandleFunc("/register", handler.RegisterHandler)
		r.ServeHTTP(rr, req)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(rr.Code).To(Equal(http.StatusCreated))
	})

	It("will run the register handler for existing user", func() {
		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte(`{"emailid":"gaurav@api.com","password":"simple","name":"gaurav","age":23}`)))
		rr := httptest.NewRecorder()
		r.HandleFunc("/register", handler.RegisterHandler)
		r.ServeHTTP(rr, req)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(rr.Code).To(Equal(http.StatusInternalServerError))
	})

	Describe("test the login handler", func() {

		It("will test login handler for wrong user id", func() {
			req, err := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(`{"emailid":"fakeUser@api.com","password":"simple"}`)))
			rr := httptest.NewRecorder()

			r.HandleFunc("/login", handler.Loginhandler)
			r.ServeHTTP(rr, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(rr.Code).To(Equal(http.StatusNotFound))

		})

		It("will test login handler for right user", func() {
			req, err := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(`{"emailid":"gaurav@api.com","password":"simple"}`)))
			rr := httptest.NewRecorder()
			r.HandleFunc("/login", handler.Loginhandler)
			r.ServeHTTP(rr, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(rr.Code).To(Equal(http.StatusOK))

		})
	})

})
