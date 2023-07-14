package repository_test

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
		repo     *mocks.MockCategoryRepository
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		repo = mocks.NewMockCategoryRepository(mockCtrl)
	})

	Context("GetAll", func() {
		It("get all records", func() {
			var dt []model.Category
			dt = append(dt, model.Category{ID: 1, Name: "kategori1"})
			dt = append(dt, model.Category{ID: 2, Name: "kategori2"})
			repo.EXPECT().FindAll().Return(dt, nil)

			data, err := repo.FindAll()
			Expect(err).NotTo(HaveOccurred())
			Expect(data).To(HaveLen(2))
		})
	})
})
