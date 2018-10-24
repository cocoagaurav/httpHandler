package main

import (
	"bytes"
	"github.com/gorilla/mux"
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
	Db = Opendatabase()
	r := mux.NewRouter()
	It("will run the register handler", func() {
		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte(`{"name":"gomega404","id":410,"age":12}`)))
		rr := httptest.NewRecorder()
		r.HandleFunc("/register", registerHandler)
		r.ServeHTTP(rr, req)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(rr.Code).To(Equal(http.StatusOK))
	})

	It("will run the register handler for existing user", func() {
		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte(`{"name":"gomega404","id":407,"age":12}`)))
		rr := httptest.NewRecorder()
		r.HandleFunc("/register", registerHandler)
		r.ServeHTTP(rr, req)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(rr.Code).To(Equal(http.StatusInternalServerError))
	})

	Describe("test the login handler", func() {

		It("will test login handler for wrong user id", func() {
			req, err := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(`{"name":"gaurav","id":2,"age":23}`)))
			rr := httptest.NewRecorder()

			r.HandleFunc("/login", loginhandler)
			r.ServeHTTP(rr, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(rr.Code).To(Equal(http.StatusNotFound))

		})

		It("will test login handler for right user", func() {
			req, err := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(`{"name":"gaurav","id":1,"age":23}`)))
			rr := httptest.NewRecorder()
			r.HandleFunc("/login", loginhandler)
			r.ServeHTTP(rr, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(rr.Code).To(Equal(http.StatusFound))

		})
	})

})
