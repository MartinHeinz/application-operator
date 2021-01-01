package v1alpha1

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// Run `go test ./...` in `api/v1alpha1/` directory
var _ = Describe("Default webhook", func() {

	// Define utility constants for object names and testing.
	const (
		ApplicationName      = "test-app"
		ApplicationNamespace = "default"

		timeout  = time.Second * 10
		interval = time.Millisecond * 250
	)

	Context("When creating Application", func() {
		It("Should set replicas count", func() {
			By("By updating spec.Replicas")
			ctx := context.Background()
			application := &Application{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "applications.apps.martinheinz.dev/v1alpha1",
					Kind:       "Application",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      ApplicationName,
					Namespace: ApplicationNamespace,
				},
				Spec: ApplicationSpec{
					Replicas: 0,
				},
			}
			Expect(k8sClient.Create(ctx, application)).Should(Succeed())

			applicationLookupKey := types.NamespacedName{Name: ApplicationName, Namespace: ApplicationNamespace}
			createdApplication := &Application{}

			Eventually(func() bool {
				err := k8sClient.Get(ctx, applicationLookupKey, createdApplication)
				if err != nil {
					return false
				}
				return true
			}, timeout, interval).Should(BeTrue())
			Expect(createdApplication.Spec.Replicas).Should(Equal(int32(1)))

		})
	})
})
