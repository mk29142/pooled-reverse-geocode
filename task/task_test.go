package task_test

import (
	"fmt"
	"github.com/mk29142/pooled-reverse-geocode/client"
	"github.com/mk29142/pooled-reverse-geocode/domain"
	"github.com/mk29142/pooled-reverse-geocode/task"
	"github.com/mk29142/pooled-reverse-geocode/task/internal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Task", func() {
	var (
		clientFake *internal.FakeClient
		t          task.Task

		coords = domain.Coordinates{
			Latitude:  50.123,
			Longitude: 0.456,
		}

		postcode = "SS16 5HE"
	)

	BeforeEach(func() {
		clientFake = new(internal.FakeClient)
		t = task.NewTask(coords, clientFake)
	})

	Describe("Process", func() {
		var (
			result task.CoordinatesWithPostcode
			err error
		)

		JustBeforeEach(func() {
			result, err = t.Process()
		})

		When("success", func() {
			BeforeEach(func() {
				resp := client.LatLongPostcode{
					Latitude:  coords.Latitude,
					Longitude: coords.Longitude,
					Postcode:  postcode,
				}

				clientFake.PostcodeReturns(resp, nil)
			})

			It("does not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("calls client", func() {
				Expect(clientFake.PostcodeCallCount()).To(Equal(1))
				got := clientFake.PostcodeArgsForCall(0)
				Expect(got).To(Equal(coords))
			})

			It("returns correct result", func() {
				Expect(result).To(Equal(task.CoordinatesWithPostcode{
					Lat:      coords.Latitude,
					Long:     coords.Longitude,
					PostCode: postcode,
				}))
			})
		})

		When("client fails to get postcode", func() {
			BeforeEach(func() {
				clientFake.PostcodeReturns(client.LatLongPostcode{}, fmt.Errorf("something went wrong"))
			})

			It("returns error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
