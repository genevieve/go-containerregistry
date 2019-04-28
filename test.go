package main

import (
	"fmt"

	resource "github.com/concourse/registry-image-resource"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

func main() {
	repository := "mcr.microsoft.com/windows/servercore:ltsc2019-amd64"
	digest := "sha256:b26b4b226ea91f4ee269b4d8608febaa2b5bf2869729cd30d97e8ccbe1f65891"

	ref := repository + "@" + digest

	n, err := name.ParseReference(ref, name.WeakValidation)
	if err != nil {
		panic(err)
	}

	imageOpts := []remote.ImageOption{
		remote.WithTransport(resource.RetryTransport),
	}

	image, err := remote.Image(n, imageOpts...)
	if err != nil {
		panic(err)
	}

	layers, err := image.Layers()
	if err != nil {
		panic(err)
	}

	for _, layer := range layers {
		mediaType, err := layer.MediaType()
		fmt.Printf("Media Type: %s\n", mediaType)
		if err != nil {
			panic(err)
		}

		_, err = layer.Compressed()
		if err != nil {
			panic(err)
		}
	}
}
