package sign

import (
	"strings"

	"github.com/google/go-cmp/cmp"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	kmmv1beta1 "github.com/rh-ecosystem-edge/kernel-module-management/api/v1beta1"
	v1 "k8s.io/api/core/v1"
)

var _ = Describe("GetRelevantSign", func() {

	const (
		unsignedImage = "my.registry/my/image"
		keySecret     = "securebootkey"
		certSecret    = "securebootcert"
		filesToSign   = "/modules/simple-kmod.ko:/modules/simple-procfs-kmod.ko"
		kernelVersion = "1.2.3"
	)

	var (
		h Helper
	)

	BeforeEach(func() {
		h = NewSignerHelper()
	})

	expected := &kmmv1beta1.Sign{
		UnsignedImage: unsignedImage,
		KeySecret:     &v1.LocalObjectReference{Name: keySecret},
		CertSecret:    &v1.LocalObjectReference{Name: certSecret},
		FilesToSign:   strings.Split(filesToSign, ":"),
	}

	DescribeTable("should set fields correctly", func(mod kmmv1beta1.Module, km kmmv1beta1.KernelMapping) {
		actual, err := h.GetRelevantSign(mod.Spec, km, kernelVersion)
		Expect(err).NotTo(HaveOccurred())
		Expect(
			cmp.Diff(expected, actual),
		).To(
			BeEmpty(),
		)
	},
		Entry(
			"no km.Sign",
			kmmv1beta1.Module{
				Spec: kmmv1beta1.ModuleSpec{
					ModuleLoader: kmmv1beta1.ModuleLoaderSpec{
						Container: kmmv1beta1.ModuleLoaderContainerSpec{
							Sign: &kmmv1beta1.Sign{
								UnsignedImage: unsignedImage,
								KeySecret:     &v1.LocalObjectReference{Name: keySecret},
								CertSecret:    &v1.LocalObjectReference{Name: certSecret},
								FilesToSign:   strings.Split(filesToSign, ":"),
							},
						},
					},
				},
			},
			kmmv1beta1.KernelMapping{},
		),
		Entry(
			"no container.Sign",
			kmmv1beta1.Module{
				Spec: kmmv1beta1.ModuleSpec{
					ModuleLoader: kmmv1beta1.ModuleLoaderSpec{
						Container: kmmv1beta1.ModuleLoaderContainerSpec{},
					},
				},
			},
			kmmv1beta1.KernelMapping{
				Sign: &kmmv1beta1.Sign{
					UnsignedImage: unsignedImage,
					KeySecret:     &v1.LocalObjectReference{Name: keySecret},
					CertSecret:    &v1.LocalObjectReference{Name: certSecret},
					FilesToSign:   strings.Split(filesToSign, ":"),
				},
			},
		),
		Entry(
			"default UnsignedImage",
			kmmv1beta1.Module{
				Spec: kmmv1beta1.ModuleSpec{
					ModuleLoader: kmmv1beta1.ModuleLoaderSpec{
						Container: kmmv1beta1.ModuleLoaderContainerSpec{
							Sign: &kmmv1beta1.Sign{
								UnsignedImage: unsignedImage,
							},
						},
					},
				},
			},
			kmmv1beta1.KernelMapping{
				Sign: &kmmv1beta1.Sign{
					KeySecret:   &v1.LocalObjectReference{Name: keySecret},
					CertSecret:  &v1.LocalObjectReference{Name: certSecret},
					FilesToSign: strings.Split(filesToSign, ":"),
				},
			},
		),
		Entry(
			"default UnsignedImage and KeySecret",
			kmmv1beta1.Module{
				Spec: kmmv1beta1.ModuleSpec{
					ModuleLoader: kmmv1beta1.ModuleLoaderSpec{
						Container: kmmv1beta1.ModuleLoaderContainerSpec{
							Sign: &kmmv1beta1.Sign{
								UnsignedImage: unsignedImage,
								KeySecret:     &v1.LocalObjectReference{Name: keySecret},
							},
						},
					},
				},
			},
			kmmv1beta1.KernelMapping{
				Sign: &kmmv1beta1.Sign{
					CertSecret:  &v1.LocalObjectReference{Name: certSecret},
					FilesToSign: strings.Split(filesToSign, ":"),
				},
			},
		),
		Entry(
			"default UnsignedImage, KeySecret, and CertSecret",
			kmmv1beta1.Module{
				Spec: kmmv1beta1.ModuleSpec{
					ModuleLoader: kmmv1beta1.ModuleLoaderSpec{
						Container: kmmv1beta1.ModuleLoaderContainerSpec{
							Sign: &kmmv1beta1.Sign{
								UnsignedImage: unsignedImage,
								KeySecret:     &v1.LocalObjectReference{Name: keySecret},
								CertSecret:    &v1.LocalObjectReference{Name: certSecret},
							},
						},
					},
				},
			},
			kmmv1beta1.KernelMapping{
				Sign: &kmmv1beta1.Sign{
					FilesToSign: strings.Split(filesToSign, ":"),
				},
			},
		),
		Entry(
			"default FilesToSign only",
			kmmv1beta1.Module{
				Spec: kmmv1beta1.ModuleSpec{
					ModuleLoader: kmmv1beta1.ModuleLoaderSpec{
						Container: kmmv1beta1.ModuleLoaderContainerSpec{
							Sign: &kmmv1beta1.Sign{
								FilesToSign: strings.Split(filesToSign, ":"),
							},
						},
					},
				},
			},
			kmmv1beta1.KernelMapping{
				Sign: &kmmv1beta1.Sign{
					UnsignedImage: unsignedImage,
					KeySecret:     &v1.LocalObjectReference{Name: keySecret},
					CertSecret:    &v1.LocalObjectReference{Name: certSecret},
				},
			},
		),
	)

})
var _ = Describe("GetRelevantSign", func() {

	const (
		unsignedImage = "my.registry/my/image"
		keySecret     = "securebootkey"
		certSecret    = "securebootcert"
		filesToSign   = "/modules/${KERNEL_VERSION}/simple-kmod.ko:/modules/${KERNEL_VERSION}/simple-procfs-kmod.ko"
		kernelVersion = "1.2.3"
	)

	var (
		h Helper
	)

	BeforeEach(func() {
		h = NewSignerHelper()
	})

	expected := &kmmv1beta1.Sign{
		UnsignedImage: unsignedImage + ":" + kernelVersion,
		KeySecret:     &v1.LocalObjectReference{Name: keySecret},
		CertSecret:    &v1.LocalObjectReference{Name: certSecret},
		FilesToSign:   strings.Split("/modules/"+kernelVersion+"/simple-kmod.ko:/modules/"+kernelVersion+"/simple-procfs-kmod.ko", ":"),
	}

	DescribeTable("should set fields correctly", func(mod kmmv1beta1.Module, km kmmv1beta1.KernelMapping) {
		actual, _ := h.GetRelevantSign(mod.Spec, km, kernelVersion)
		Expect(
			cmp.Diff(expected, actual),
		).To(
			BeEmpty(),
		)
	},
		Entry(
			"no km.Sign",
			kmmv1beta1.Module{
				Spec: kmmv1beta1.ModuleSpec{
					ModuleLoader: kmmv1beta1.ModuleLoaderSpec{
						Container: kmmv1beta1.ModuleLoaderContainerSpec{
							Sign: &kmmv1beta1.Sign{
								UnsignedImage: unsignedImage + ":${KERNEL_VERSION}",
								KeySecret:     &v1.LocalObjectReference{Name: keySecret},
								CertSecret:    &v1.LocalObjectReference{Name: certSecret},
								FilesToSign:   strings.Split(filesToSign, ":"),
							},
						},
					},
				},
			},
			kmmv1beta1.KernelMapping{},
		),
		Entry(
			"no container.Sign",
			kmmv1beta1.Module{
				Spec: kmmv1beta1.ModuleSpec{
					ModuleLoader: kmmv1beta1.ModuleLoaderSpec{
						Container: kmmv1beta1.ModuleLoaderContainerSpec{},
					},
				},
			},
			kmmv1beta1.KernelMapping{
				Sign: &kmmv1beta1.Sign{
					UnsignedImage: unsignedImage + ":${KERNEL_VERSION}",
					KeySecret:     &v1.LocalObjectReference{Name: keySecret},
					CertSecret:    &v1.LocalObjectReference{Name: certSecret},
					FilesToSign:   strings.Split(filesToSign, ":"),
				},
			},
		),
	)
})
