package rbac

import (
	"context"
	"errors"

	"github.com/golang/mock/gomock"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/utils/pointer"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	kmmv1beta1 "github.com/rh-ecosystem-edge/kernel-module-management/api/v1beta1"
	"github.com/rh-ecosystem-edge/kernel-module-management/internal/client"
)

var (
	ctrl *gomock.Controller
	clnt *client.MockClient
)

var _ = Describe("GenerateModuleLoaderServiceAccountName", func() {
	mod := kmmv1beta1.Module{
		ObjectMeta: metav1.ObjectMeta{Name: "test-module", Namespace: "namespace"},
	}

	It("should return the default module loader ServiceAccount name", func() {
		Expect(GenerateModuleLoaderServiceAccountName(mod)).To(Equal("test-module-module-loader"))
	})
})

var _ = Describe("GenerateDevicePluginServiceAccountName", func() {
	mod := kmmv1beta1.Module{
		ObjectMeta: metav1.ObjectMeta{Name: "test-module", Namespace: "namespace"},
	}

	It("should return the default device plugin ServiceAccount name", func() {
		Expect(GenerateDevicePluginServiceAccountName(mod)).To(Equal("test-module-device-plugin"))
	})
})

var _ = Describe("CreateModuleLoaderRBAC", func() {
	const (
		moduleName = "test-module"
		namespace  = "namespace"
	)

	var (
		rc RBACCreator

		mod kmmv1beta1.Module

		requestedServiceAccount *corev1.ServiceAccount

		requestedRoleBinding *rbacv1.RoleBinding

		expectedServiceAccount *corev1.ServiceAccount

		expectedRoleBinding *rbacv1.RoleBinding
	)

	ctx := context.Background()

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		clnt = client.NewMockClient(ctrl)
		rc = NewCreator(clnt, scheme)

		mod = kmmv1beta1.Module{
			TypeMeta: metav1.TypeMeta{
				APIVersion: kmmv1beta1.GroupVersion.String(),
				Kind:       "Module",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      moduleName,
				Namespace: namespace,
			},
		}
		requestedServiceAccount = &corev1.ServiceAccount{
			ObjectMeta: metav1.ObjectMeta{
				Name:      moduleName + "-module-loader",
				Namespace: namespace,
			},
		}
		requestedRoleBinding = &rbacv1.RoleBinding{
			ObjectMeta: metav1.ObjectMeta{
				Name:      moduleName + "-module-loader",
				Namespace: namespace,
			},
		}
		expectedServiceAccount = &corev1.ServiceAccount{
			ObjectMeta: metav1.ObjectMeta{
				Name:      moduleName + "-module-loader",
				Namespace: namespace,
				OwnerReferences: []metav1.OwnerReference{
					{
						APIVersion:         mod.APIVersion,
						BlockOwnerDeletion: pointer.Bool(true),
						Controller:         pointer.Bool(true),
						Kind:               mod.Kind,
						Name:               moduleName,
						UID:                mod.UID,
					},
				},
			},
		}
		expectedRoleBinding = &rbacv1.RoleBinding{
			ObjectMeta: metav1.ObjectMeta{
				Name:      moduleName + "-module-loader",
				Namespace: namespace,
				OwnerReferences: []metav1.OwnerReference{
					{
						APIVersion:         mod.APIVersion,
						BlockOwnerDeletion: pointer.Bool(true),
						Controller:         pointer.Bool(true),
						Kind:               mod.Kind,
						Name:               moduleName,
						UID:                mod.UID,
					},
				},
			},
			Subjects: []rbacv1.Subject{
				{
					Kind:      rbacv1.ServiceAccountKind,
					Name:      expectedServiceAccount.Name,
					Namespace: mod.Namespace,
				},
			},
			RoleRef: rbacv1.RoleRef{
				APIGroup: rbacv1.GroupName,
				Kind:     "ClusterRole",
				Name:     "kmm-operator-module-loader",
			},
		}
	})

	It("should add the default module loader ServiceAccount and RoleBinding", func() {
		gomock.InOrder(
			clnt.EXPECT().Get(ctx, gomock.Any(), requestedServiceAccount).Return(apierrors.NewNotFound(schema.GroupResource{}, "whatever")),
			clnt.EXPECT().Create(ctx, expectedServiceAccount).Return(nil),
			clnt.EXPECT().Get(ctx, gomock.Any(), requestedRoleBinding).Return(apierrors.NewNotFound(schema.GroupResource{}, "whatever")),
			clnt.EXPECT().Create(ctx, expectedRoleBinding).Return(nil),
		)

		err := rc.CreateModuleLoaderRBAC(context.Background(), mod)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should return an error when the ServiceAccount fetch fails", func() {
		gomock.InOrder(
			clnt.EXPECT().Get(ctx, gomock.Any(), requestedServiceAccount).Return(errors.New("some-error")),
		)

		err := rc.CreateModuleLoaderRBAC(context.Background(), mod)
		Expect(err).To(HaveOccurred())
	})

	It("should return an error when the ServiceAccount creation fails", func() {
		gomock.InOrder(
			clnt.EXPECT().Get(ctx, gomock.Any(), requestedServiceAccount).Return(apierrors.NewNotFound(schema.GroupResource{}, "whatever")),
			clnt.EXPECT().Create(ctx, expectedServiceAccount).Return(errors.New("some-error")),
		)

		err := rc.CreateModuleLoaderRBAC(context.Background(), mod)
		Expect(err).To(HaveOccurred())
	})

	It("should return an error when the RoleBinding fetch fails", func() {
		gomock.InOrder(
			clnt.EXPECT().Get(ctx, gomock.Any(), requestedServiceAccount).Return(apierrors.NewNotFound(schema.GroupResource{}, "whatever")),
			clnt.EXPECT().Create(ctx, expectedServiceAccount).Return(nil),
			clnt.EXPECT().Get(ctx, gomock.Any(), requestedRoleBinding).Return(errors.New("some-error")),
		)

		err := rc.CreateModuleLoaderRBAC(context.Background(), mod)
		Expect(err).To(HaveOccurred())
	})

	It("should return an error when the RoleBinding creation fails", func() {
		gomock.InOrder(
			clnt.EXPECT().Get(ctx, gomock.Any(), requestedServiceAccount).Return(apierrors.NewNotFound(schema.GroupResource{}, "whatever")),
			clnt.EXPECT().Create(ctx, expectedServiceAccount).Return(nil),
			clnt.EXPECT().Get(ctx, gomock.Any(), requestedRoleBinding).Return(apierrors.NewNotFound(schema.GroupResource{}, "whatever")),
			clnt.EXPECT().Create(ctx, expectedRoleBinding).Return(errors.New("some-error")),
		)

		err := rc.CreateModuleLoaderRBAC(context.Background(), mod)
		Expect(err).To(HaveOccurred())
	})
})

var _ = Describe("CreateDevicePluginServiceRBAC", func() {
	const (
		moduleName = "test-module"
		namespace  = "namespace"
	)

	var (
		rc RBACCreator

		mod kmmv1beta1.Module

		requestedServiceAccount *corev1.ServiceAccount

		requestedRoleBinding *rbacv1.RoleBinding

		expectedServiceAccount *corev1.ServiceAccount

		expectedRoleBinding *rbacv1.RoleBinding
	)

	ctx := context.Background()

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		clnt = client.NewMockClient(ctrl)
		rc = NewCreator(clnt, scheme)

		mod = kmmv1beta1.Module{
			TypeMeta: metav1.TypeMeta{
				APIVersion: kmmv1beta1.GroupVersion.String(),
				Kind:       "Module",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      moduleName,
				Namespace: namespace,
			},
		}
		requestedServiceAccount = &corev1.ServiceAccount{
			ObjectMeta: metav1.ObjectMeta{
				Name:      moduleName + "-device-plugin",
				Namespace: namespace,
			},
		}
		requestedRoleBinding = &rbacv1.RoleBinding{
			ObjectMeta: metav1.ObjectMeta{
				Name:      moduleName + "-device-plugin",
				Namespace: namespace,
			},
		}
		expectedServiceAccount = &corev1.ServiceAccount{
			ObjectMeta: metav1.ObjectMeta{
				Name:      moduleName + "-device-plugin",
				Namespace: namespace,
				OwnerReferences: []metav1.OwnerReference{
					{
						APIVersion:         mod.APIVersion,
						BlockOwnerDeletion: pointer.Bool(true),
						Controller:         pointer.Bool(true),
						Kind:               mod.Kind,
						Name:               moduleName,
						UID:                mod.UID,
					},
				},
			},
		}
		expectedRoleBinding = &rbacv1.RoleBinding{
			ObjectMeta: metav1.ObjectMeta{
				Name:      moduleName + "-device-plugin",
				Namespace: namespace,
				OwnerReferences: []metav1.OwnerReference{
					{
						APIVersion:         mod.APIVersion,
						BlockOwnerDeletion: pointer.Bool(true),
						Controller:         pointer.Bool(true),
						Kind:               mod.Kind,
						Name:               moduleName,
						UID:                mod.UID,
					},
				},
			},
			Subjects: []rbacv1.Subject{
				{
					Kind:      rbacv1.ServiceAccountKind,
					Name:      expectedServiceAccount.Name,
					Namespace: mod.Namespace,
				},
			},
			RoleRef: rbacv1.RoleRef{
				APIGroup: rbacv1.GroupName,
				Kind:     "ClusterRole",
				Name:     "kmm-operator-device-plugin",
			},
		}
	})

	It("should add the default device plugin ServiceAccount and RoleBinding", func() {
		gomock.InOrder(
			clnt.EXPECT().Get(ctx, gomock.Any(), requestedServiceAccount).Return(apierrors.NewNotFound(schema.GroupResource{}, "whatever")),
			clnt.EXPECT().Create(ctx, expectedServiceAccount).Return(nil),
			clnt.EXPECT().Get(ctx, gomock.Any(), requestedRoleBinding).Return(apierrors.NewNotFound(schema.GroupResource{}, "whatever")),
			clnt.EXPECT().Create(ctx, expectedRoleBinding).Return(nil),
		)

		err := rc.CreateDevicePluginRBAC(context.Background(), mod)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should return an error when the ServiceAccount fetch fails", func() {
		gomock.InOrder(
			clnt.EXPECT().Get(ctx, gomock.Any(), requestedServiceAccount).Return(errors.New("some-error")),
		)

		err := rc.CreateDevicePluginRBAC(context.Background(), mod)
		Expect(err).To(HaveOccurred())
	})

	It("should return an error when the ServiceAccount creation fails", func() {
		gomock.InOrder(
			clnt.EXPECT().Get(ctx, gomock.Any(), requestedServiceAccount).Return(apierrors.NewNotFound(schema.GroupResource{}, "whatever")),
			clnt.EXPECT().Create(ctx, expectedServiceAccount).Return(errors.New("some-error")),
		)

		err := rc.CreateDevicePluginRBAC(context.Background(), mod)
		Expect(err).To(HaveOccurred())
	})

	It("should return an error when the RoleBinding fetch fails", func() {
		gomock.InOrder(
			clnt.EXPECT().Get(ctx, gomock.Any(), requestedServiceAccount).Return(apierrors.NewNotFound(schema.GroupResource{}, "whatever")),
			clnt.EXPECT().Create(ctx, expectedServiceAccount).Return(nil),
			clnt.EXPECT().Get(ctx, gomock.Any(), requestedRoleBinding).Return(errors.New("some-error")),
		)

		err := rc.CreateDevicePluginRBAC(context.Background(), mod)
		Expect(err).To(HaveOccurred())
	})

	It("should return an error when the RoleBinding creation fails", func() {
		gomock.InOrder(
			clnt.EXPECT().Get(ctx, gomock.Any(), requestedServiceAccount).Return(apierrors.NewNotFound(schema.GroupResource{}, "whatever")),
			clnt.EXPECT().Create(ctx, expectedServiceAccount).Return(nil),
			clnt.EXPECT().Get(ctx, gomock.Any(), requestedRoleBinding).Return(apierrors.NewNotFound(schema.GroupResource{}, "whatever")),
			clnt.EXPECT().Create(ctx, expectedRoleBinding).Return(errors.New("some-error")),
		)

		err := rc.CreateDevicePluginRBAC(context.Background(), mod)
		Expect(err).To(HaveOccurred())
	})
})
