package repository_test

import (
	"endeus/wawan/internal/mocks"
	"endeus/wawan/pkg/model"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Caramasak", func() {
	var (
		mockCtrl *gomock.Controller
		repo     *mocks.MockCaraMasakRepository
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		repo = mocks.NewMockCaraMasakRepository(mockCtrl)
	})

	Context("GetByResepID", func() {
		It("get single records by resepID", func() {
			var dt = model.CaraMasak{ID: 1, VideoUrl: "http://localhost", LamaMasak: 60, ResepID: 20}
			repo.EXPECT().GetByResepID(uint(20)).Return(dt, nil)

			data, err := repo.GetByResepID(uint(20))
			Expect(err).NotTo(HaveOccurred())
			Expect(data.ResepID).To(Equal(uint(20)))
			Expect(data.VideoUrl).NotTo(BeEmpty())
		})
	})
})
