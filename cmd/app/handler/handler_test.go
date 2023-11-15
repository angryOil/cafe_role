package handler

import (
	"cafe_role/internal/controller"
	"cafe_role/internal/repository"
	"cafe_role/internal/repository/infla"
	"cafe_role/internal/service"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("handler", func() {
	h := NewHandler(controller.NewController(service.NewService(repository.NewRepository(infla.NewDB()))))
	Describe("getList 메소드는", func() {
		Describe("path 값 1을 넣을경우", func() {
			r := httptest.NewRequest(http.MethodGet, "/roles/1", nil)
			w := httptest.NewRecorder()
			h.ServeHTTP(w, r)

			It("200 ok 를 반환한다", func() {
				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(w.Body).To(Equal(http.StatusOK))
			})
		})
	})
})
