package service_test

import (
	"endeus/wawan/internal/mocks"
	"endeus/wawan/pkg/endeus"
	"endeus/wawan/pkg/model"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Resep", func() {
	var (
		mockCtrl *gomock.Controller
		svc      *mocks.MockResepService
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		svc = mocks.NewMockResepService(mockCtrl)
	})

	Context("Get all resep", func() {
		It("return with custom response", func() {
			var ctData []endeus.ResepCustomResponse
			ctData = append(ctData, endeus.ResepCustomResponse{
				ID:         1,
				CategoryID: 2,
				Judul:      "title",
				Deskripsi:  "desc",
				VideoUrl:   "url",
				Publish:    true,
			})
			ctData = append(ctData, endeus.ResepCustomResponse{
				ID:         2,
				CategoryID: 3,
				Judul:      "title",
				Deskripsi:  "desc",
				VideoUrl:   "url",
				Publish:    true,
			})
			svc.EXPECT().GetAll("").Return(ctData, nil)

			data, err := svc.GetAll("")
			Expect(err).NotTo(HaveOccurred())
			Expect(len(data)).To(Equal(2))
		})

		It("with param return with custom response", func() {
			var ctData []endeus.ResepCustomResponse
			ctData = append(ctData, endeus.ResepCustomResponse{
				ID:         1,
				CategoryID: 2,
				Judul:      "title",
				Deskripsi:  "desc",
				VideoUrl:   "url",
				Publish:    true,
			})
			ctData = append(ctData, endeus.ResepCustomResponse{
				ID:         2,
				CategoryID: 3,
				Judul:      "title",
				Deskripsi:  "desc",
				VideoUrl:   "url",
				Publish:    true,
			})
			svc.EXPECT().GetAll("title").Return(ctData, nil)

			data, err := svc.GetAll("title")
			Expect(err).NotTo(HaveOccurred())
			Expect(len(data)).To(Equal(2))
		})
	})

	Context("Get all resep by category_id", func() {
		It("return with custom response", func() {
			var ctData []endeus.ResepCustomResponse
			ctData = append(ctData, endeus.ResepCustomResponse{
				ID:         1,
				CategoryID: 5,
				Judul:      "title",
				Deskripsi:  "desc",
				VideoUrl:   "url",
				Publish:    true,
			})
			ctData = append(ctData, endeus.ResepCustomResponse{
				ID:         2,
				CategoryID: 5,
				Judul:      "title",
				Deskripsi:  "desc",
				VideoUrl:   "url",
				Publish:    true,
			})
			svc.EXPECT().GetAllByCategoryID(uint(5)).Return(ctData, nil)

			data, err := svc.GetAllByCategoryID(uint(5))
			Expect(err).NotTo(HaveOccurred())
			Expect(len(data)).To(Equal(2))
		})
	})

	Context("Get by ID", func() {
		It("return single data", func() {
			var ctData = model.Resep{
				ID:         1,
				CategoryID: 5,
				Judul:      "title",
				Deskripsi:  "desc",
				VideoUrl:   "url",
				Publish:    true,
				Bahan: model.Bahan{
					ID:        2,
					ResepID:   1,
					Porsi:     5,
					Deskripsi: "keterangan",
				},
				CaraBuat: model.CaraBuat{
					ID:        2,
					ResepID:   1,
					LamaWaktu: 10,
					Tips:      "",
					Deskripsi: "dimasak",
				},
			}
			svc.EXPECT().GetByID(uint(1)).Return(ctData, nil)

			data, err := svc.GetByID(uint(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(data.Judul).NotTo(BeEmpty())
		})
	})

	Context("Create New", func() {
		It("record new data", func() {
			var param = endeus.ResepParam{
				CategoryID:         5,
				Judul:              "title",
				Deskripsi:          "desc",
				VideoUrl:           "url",
				Porsi:              5,
				DeskripsiBahan:     "keterangan",
				LamaWaktu:          10,
				Tips:               "",
				DeskripsiPembuatan: "dimasak",
			}
			svc.EXPECT().CreateNew(param).Return(nil)
			err := svc.CreateNew(param)
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("Update", func() {
		It("update data", func() {
			var param = endeus.ResepParam{
				CategoryID:         6,
				Judul:              "title1",
				Deskripsi:          "desc1",
				VideoUrl:           "url1",
				Porsi:              2,
				DeskripsiBahan:     "keterangan",
				LamaWaktu:          10,
				Tips:               "",
				DeskripsiPembuatan: "dimasak",
			}
			svc.EXPECT().Update(uint(1), param).Return(nil)
			err := svc.Update(uint(1), param)
			Expect(err).NotTo(HaveOccurred())
		})
	})

})
