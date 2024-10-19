package docker

import (
	"context"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"io"
	"os"
)

type ImageOperator struct {
	Client *client.Client
}

func NewImageOperator() (*ImageOperator, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return &ImageOperator{Client: cli}, nil
}

func (op *ImageOperator) PullImage(imageURL string) error {
	ctx := context.Background()
	reader, err := op.Client.ImagePull(ctx, imageURL, image.PullOptions{})
	if err != nil {
		return err
	}
	defer reader.Close()
	_, err = io.Copy(os.Stdout, reader)
	return err
}

func (op *ImageOperator) TagImage(source, target string) error {
	ctx := context.Background()
	if err := op.Client.ImageTag(ctx, source, target); err != nil {
		return err
	}
	return nil
}

func (op *ImageOperator) PushImage(imageURL, auth string) error {
	ctx := context.Background()
	reader, err := op.Client.ImagePush(ctx, imageURL, image.PushOptions{
		RegistryAuth: auth,
	})
	if err != nil {
		return err
	}
	defer reader.Close()
	_, err = io.Copy(os.Stdout, reader)
	return err
}
