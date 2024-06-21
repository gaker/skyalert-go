package skyalert_test

import (
	"time"

	"github.com/gaker/skyalert-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("SkyalertGo", func() {
	var file = []byte("2024-06-11 14:09:57.00 F M 79.6   92.8  93      0      42  66.3   000 1 1 00019 045454.59025 3 1 1 1 1 1")

	It("should fail to parse a bad timestamp", func() {
		a, err := skyalert.New([]byte("2024-06-54 14:09:57.00 F M 79.6   92.8  93      0      42  66.3   000 0 0 00019 045454.59025 3 1 1 1 1 1")).Parse()
		Expect(a).To(BeNil())
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(Equal("parsing time \"2024-06-54 14:09:57.00\": day out of range"))
	})

	It("should parse a timestamp", func() {
		loc, _ := time.LoadLocation("America/Chicago")
		a, err := skyalert.New(file).WithLocation(loc).Parse()
		Expect(err).To(BeNil())
		Expect(a.Timestamp.UTC().UnixMicro()).To(Equal(int64(1718132997000000)))
	})

	It("should parse a the temp scale", func() {
		a, err := skyalert.New(file).Parse()
		Expect(err).To(BeNil())
		Expect(a.TempScale).To(Equal("F"))
	})

	It("should parse a the Wind Scale", func() {
		a, err := skyalert.New(file).Parse()
		Expect(err).To(BeNil())
		Expect(a.WindScale).To(Equal("M"))
	})

	It("should parse a the SkyTemp", func() {
		a, err := skyalert.New(file).Parse()
		Expect(err).To(BeNil())
		Expect(a.SkyTemp).To(Equal(79.6))
	})

	It("should parse a the Ambient Temp", func() {
		a, err := skyalert.New(file).Parse()
		Expect(err).To(BeNil())
		Expect(a.AmbientTemp).To(Equal(92.8))
	})

	It("should parse a the SensorTemp", func() {
		a, err := skyalert.New(file).Parse()
		Expect(err).To(BeNil())
		Expect(a.SensorTemp).To(Equal(93.0))
	})

	It("should parse a the WindSpeed", func() {
		a, err := skyalert.New(file).Parse()
		Expect(err).To(BeNil())
		Expect(a.WindSpeed).To(Equal(0.0))
	})

	It("should parse a the Humidity", func() {
		a, err := skyalert.New(file).Parse()
		Expect(err).To(BeNil())
		Expect(a.Humidity).To(Equal(42))
	})

	It("should parse a the DewPoint", func() {
		a, err := skyalert.New(file).Parse()
		Expect(err).To(BeNil())
		Expect(a.DewPoint).To(Equal(66.3))
	})

	It("should parse a the DewHeaterPercentage", func() {
		a, err := skyalert.New(file).Parse()
		Expect(err).To(BeNil())
		Expect(a.DewHeaterPercentage).To(Equal(0))
	})

	It("should parse a the RainFlag", func() {
		a, err := skyalert.New(file).Parse()
		Expect(err).To(BeNil())
		Expect(a.RainFlag).To(Equal(1))
	})

	It("should parse a the WetFlag", func() {
		a, err := skyalert.New(file).Parse()
		Expect(err).To(BeNil())
		Expect(a.WetFlag).To(Equal(1))
	})

	It("should parse a the SinceGoodData", func() {
		a, err := skyalert.New(file).Parse()
		Expect(err).To(BeNil())
		Expect(a.SinceGoodData).To(Equal(19))
	})

	It("should parse a the CloudCondition", func() {
		a, err := skyalert.New(file).Parse()
		Expect(err).To(BeNil())
		Expect(int(a.CloudCondition)).To(Equal(3))
	})

	It("should parse a the WindCondition", func() {
		a, err := skyalert.New(file).Parse()
		Expect(err).To(BeNil())
		Expect(int(a.WindCondition)).To(Equal(1))
	})

	It("should parse a the RainCondition", func() {
		a, err := skyalert.New(file).Parse()
		Expect(err).To(BeNil())
		Expect(int(a.RainCondition)).To(Equal(1))
	})

	It("should parse a the DarknessCondition", func() {
		a, err := skyalert.New(file).Parse()
		Expect(err).To(BeNil())
		Expect(int(a.DarknessCondition)).To(Equal(1))
	})

	It("should request a roof close", func() {
		a, err := skyalert.New(file).Parse()
		Expect(err).To(BeNil())
		Expect(a.RoofCloseRequested).To(Equal(1))
	})

	It("should parse a the AlertCondition", func() {
		a, err := skyalert.New(file).Parse()
		Expect(err).To(BeNil())
		Expect(int(a.AlertCondition)).To(Equal(1))
	})
})
