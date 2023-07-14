package service_test

import (
	"endeus/wawan/internal/mocks"
	"endeus/wawan/pkg/model"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Category", func() {
	var (
		mockCtrl *gomock.Controller
		svc      *mocks.MockCategoryService
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		svc = mocks.NewMockCategoryService(mockCtrl)
	})

	Context("get all category", func() {
		It("success get data", func() {
			var dummyModel []model.Category
			dummyModel = append(dummyModel, model.Category{
				ID:   1,
				Name: "kategori1",
			})
			dummyModel = append(dummyModel, model.Category{
				ID:   2,
				Name: "kategori2",
			})
			svc.EXPECT().GetAll().Return(dummyModel, nil)
			data, err := svc.GetAll()
			Expect(err).NotTo(HaveOccurred())
			Expect(len(data)).To(Equal(2))
		})
	})

})
