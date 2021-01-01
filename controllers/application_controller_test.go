package controllers

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	appsv1alpha1 "github.com/MartinHeinz/application-operator/api/v1alpha1"
)

// Run `go test ./...` in `controllers/` directory
var _ = Describe("Application controller", func() {

	// Define utility constants for object names and testing.
	const (
		ApplicationName      = "test-app"
		ApplicationNamespace = "default"

		timeout  = time.Second * 30
		interval = time.Millisecond * 250
	)

	Context("When updating Application Status", func() {
		It("Should increase Application Status.Instances count when pods are created", func() {
			By("By creating a new Application")
			ctx := context.Background()
			application := &appsv1alpha1.Application{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "applications.apps.martinheinz.dev/v1alpha1",
					Kind:       "Application",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      ApplicationName,
					Namespace: ApplicationNamespace,
				},
				Spec: appsv1alpha1.ApplicationSpec{
					Replicas: 2,
				},
			}
			Expect(k8sClient.Create(ctx, application)).Should(Succeed())

			applicationLookupKey := types.NamespacedName{Name: ApplicationName, Namespace: ApplicationNamespace}
			createdApplication := &appsv1alpha1.Application{}

			Eventually(func() bool {
				err := k8sClient.Get(ctx, applicationLookupKey, createdApplication)
				if err != nil {
					return false
				}
				return true
			}, timeout, interval).Should(BeTrue())
			Expect(createdApplication.Spec.Replicas).Should(Equal(int32(2)))

		})
	})
})
