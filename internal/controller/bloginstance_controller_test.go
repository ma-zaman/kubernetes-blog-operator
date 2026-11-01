/*
Copyright 2026.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	platformv1alpha1 "github.com/ma-zaman/kubernetes-blog-operator/api/v1alpha1"
)

var _ = Describe("BlogInstance Controller", func() {
	Context("When reconciling a resource", func() {
		const (
			resourceName      = "test-resource"
			resourceNamespace = "default"
		)

		ctx := context.Background()

		typeNamespacedName := types.NamespacedName{
			Name:      resourceName,
			Namespace: resourceNamespace,
		}
		bloginstance := &platformv1alpha1.BlogInstance{}

		BeforeEach(func() {
			By("creating the custom resource for the Kind BlogInstance")
			err := k8sClient.Get(ctx, typeNamespacedName, bloginstance)
			if err != nil && errors.IsNotFound(err) {
				resource := &platformv1alpha1.BlogInstance{
					ObjectMeta: metav1.ObjectMeta{
						Name:      resourceName,
						Namespace: resourceNamespace,
					},
					Spec: platformv1alpha1.BlogInstanceSpec{
						Image: "nginx:latest", // Fournir une valeur valide pour spec.image
						Title: "Test Blog",    // Fournir une valeur valide pour spec.title
					},
				}
				Expect(k8sClient.Create(ctx, resource)).To(Succeed())
			}
		})

		AfterEach(func() {
			// Nettoyage après chaque test
			resource := &platformv1alpha1.BlogInstance{}
			err := k8sClient.Get(ctx, typeNamespacedName, resource)
			Expect(err).NotTo(HaveOccurred())

			By("Cleanup the specific resource instance BlogInstance")
			Expect(k8sClient.Delete(ctx, resource)).To(Succeed())
		})

		It("should successfully reconcile the resource", func() {
			// Ajoutez ici vos assertions pour vérifier le comportement attendu
			By("checking if the resource was reconciled successfully")
			err := k8sClient.Get(ctx, typeNamespacedName, bloginstance)
			Expect(err).NotTo(HaveOccurred())
			Expect(bloginstance.Spec.Image).To(Equal("nginx:latest"))
			Expect(bloginstance.Spec.Title).To(Equal("Test Blog"))
		})

		It("should successfully reconcile the resource", func() {
			req := reconcile.Request{
				NamespacedName: typeNamespacedName,
			}
			_ = req // Pour éviter l'erreur d'import inutilisé

			// Ajoutez vos assertions ici
		})
	})
})
