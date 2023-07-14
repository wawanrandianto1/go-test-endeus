package repository_test

import (
	"endeus/wawan/internal/mocks"
	"endeus/wawan/pkg/endeus"
	"endeus/wawan/pkg/model"
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Resep", func() {
	var (
		mockCtrl *gomock.Controller
		repo     *mocks.MockResepRepository
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		repo = mocks.NewMockResepRepository(mockCtrl)
	})

	Context("GetAll", func() {
		Context("without search", func() {
			It("get all records", func() {
				var dt []model.Resep
				dt = append(dt, model.Resep{ID: 22, Judul: "resep ayam goreng", Deskripsi: "ayam goreng", CategoryID: 1})
				dt = append(dt, model.Resep{ID: 23, Judul: "resep nasi goreng", Deskripsi: "nasi goreng", CategoryID: 2})
				repo.EXPECT().FindAll("").Return(dt, nil)

				data, err := repo.FindAll("")
				Expect(err).NotTo(HaveOccurred())
				Expect(len(data)).To(Equal(2))
			})
		})

		Context("with search", func() {
			It("get all records", func() {
				var dt []model.Resep
				dt = append(dt, model.Resep{ID: 22, Judul: "resep ayam goreng", Deskripsi: "ayam goreng", CategoryID: 1})
				repo.EXPECT().FindAll("ayam").Return(dt, nil)

				data, err := repo.FindAll("ayam")
				Expect(err).NotTo(HaveOccurred())
				Expect(len(data)).To(Equal(1))
			})
		})
	})

	Context("GetAllByCategoryID", func() {
		It("get all records by category_id", func() {
			var dt []model.Resep
			dt = append(dt, model.Resep{ID: 22, Judul: "resep ayam goreng", Deskripsi: "ayam goreng", CategoryID: 1})
			dt = append(dt, model.Resep{ID: 23, Judul: "resep nasi goreng", Deskripsi: "nasi goreng", CategoryID: 1})
			repo.EXPECT().FindAllByCategoryID(uint(1)).Return(dt, nil)

			data, err := repo.FindAllByCategoryID(uint(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(len(data)).To(Equal(2))
		})
	})

	Context("GetByID", func() {
		It("existing id, get single records by id", func() {
			dt := model.Resep{ID: 1, Judul: "resep ayam goreng", Deskripsi: "ayam goreng", CategoryID: 5}
			repo.EXPECT().FindByID(uint(1)).Return(dt, nil)

			data, err := repo.FindByID(uint(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(data.Judul).NotTo(BeEmpty())
			Expect(data.Deskripsi).NotTo(BeEmpty())
		})

		It("unknown id, throw error", func() {
			dt := model.Resep{}
			repo.EXPECT().FindByID(uint(2)).Return(dt, errors.New("id not found"))

			_, err := repo.FindByID(uint(2))
			Expect(err).To(HaveOccurred())
		})
	})

	Context("Create", func() {
		It("success", func() {
			dt := endeus.ResepParam{
				Judul:              "resep nasi goreng",
				Deskripsi:          "nasi goreng",
				CategoryID:         1,
				VideoUrl:           "http://localhost",
				Porsi:              3,
				DeskripsiBahan:     "bawang, nasi, minyak",
				LamaWaktu:          30,
				Tips:               "",
				DeskripsiPembuatan: "minyak dituang lalu dimasukkan semua",
			}
			repo.EXPECT().Create(dt).Return(nil)
			err := repo.Create(dt)
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("Unpublish", func() {
		It("will unpublish data", func() {
			repo.EXPECT().Unpublish(uint(1)).Return(nil)
			err := repo.Unpublish(uint(1))
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("Delete", func() {
		It("will delete data", func() {
			repo.EXPECT().Delete(uint(1)).Return(nil)
			err := repo.Delete(uint(1))
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("Update", func() {
		It("success", func() {
			dt := endeus.ResepParam{
				Judul:              "resep ayam goreng",
				Deskripsi:          "ayam goreng",
				CategoryID:         1,
				VideoUrl:           "http://localhost",
				Porsi:              2,
				DeskripsiBahan:     "bawang, ayam, minyak",
				LamaWaktu:          10,
				Tips:               "",
				DeskripsiPembuatan: "minyak dituang lalu dimasukkan semua",
			}
			repo.EXPECT().Update(uint(1), dt).Return(nil)
			err := repo.Update(uint(1), dt)
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
