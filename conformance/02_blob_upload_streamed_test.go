package conformance

import (
	"net/http"

	"github.com/bloodorangeio/reggie"
	g "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var test02BlobUploadStreamed = func() {
	g.Context("Blob Upload Streamed", func() {
		g.Specify("PATCH request with blob in body should yield 202 response", func() {
			req := client.NewRequest(reggie.POST, "/v2/<name>/blobs/uploads/")
			resp, err := client.Do(req)
			Expect(err).To(BeNil())
			lastResponse = resp

			req = client.NewRequest(reggie.PATCH, lastResponse.GetRelativeLocation()).
				SetHeader("Content-Type", "application/octet-stream").
				SetBody(blobA)
			resp, err = client.Do(req)
			Expect(err).To(BeNil())
			Expect(resp.StatusCode()).To(Equal(http.StatusAccepted))
			lastResponse = resp
		})

		g.Specify("PUT request to session URL with digest should yield 201 response", func() {
			req := client.NewRequest(reggie.PUT, lastResponse.GetRelativeLocation()).
				SetQueryParam("digest", blobADigest).
				SetHeader("Content-Type", "application/octet-stream").
				SetHeader("Content-Length", blobALength)
			resp, err := client.Do(req)
			Expect(err).To(BeNil())
			Expect(resp.StatusCode()).To(Equal(http.StatusCreated))
		})
	})
}
