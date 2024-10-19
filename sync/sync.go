package sync

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
	"sync-image/config"
	"sync-image/docker"
)

var wg sync.WaitGroup

type ImageSync struct {
	Config   *config.Config
	Operator *docker.ImageOperator
}

func (is *ImageSync) SyncImages(authStr string) error {
	for _, registry := range is.Config.Source.Registries {
		targetRepo := findTargetRepo(is.Config.Target.Repositories, registry.Name)
		if targetRepo == "" {
			return errors.New(fmt.Sprintf("No target repository found for source registry: %s", registry.Name))
		}

		for _, image := range registry.Images {
			sourceImageURL := image
			targetImageURL := buildTargetImageURL(is.Config.Target.Registry, targetRepo, extractImageName(image))

			log.Printf("Source Image: %s\n", sourceImageURL)
			log.Printf("Target Image: %s\n", targetImageURL)

			wg.Add(1)
			go func(src, target, auth string) {
				defer wg.Done()

				if err := is.Operator.PullImage(src); err != nil {
					log.Printf("Error pulling image %s: %v\n", src, err)
					return
				}
				log.Printf("Successfully pulled image: %s\n", src)

				if err := is.Operator.TagImage(src, target); err != nil {
					log.Printf("Error tagging image %s to %s: %v\n", src, target, err)
					return
				}
				log.Printf("Successfully tagged image: %s to %s\n", src, target)

				if err := is.Operator.PushImage(target, auth); err != nil {
					log.Printf("Error pushing image %s: %v\n", target, err)
					return
				}
				log.Printf("Successfully pushed image: %s\n", target)
			}(sourceImageURL, targetImageURL, authStr)
		}
	}
	wg.Wait()
	return nil
}

func findTargetRepo(repositories []config.TargetRepository, sourceRegistryName string) string {
	for _, repo := range repositories {
		if repo.Name == sourceRegistryName {
			return repo.Name
		}
	}
	return ""
}

func buildTargetImageURL(targetRegistry, targetRepo, imageName string) string {
	return fmt.Sprintf("%s/%s/%s", targetRegistry, targetRepo, imageName)
}

// 假设 imageURL 是 quay.io/prometheus/alertmanager:v0.27.0

func extractImageName(imageURL string) string {
	image := strings.Split(imageURL, "/")
	imageAndTag := image[len(image)-1]
	if strings.Contains(imageAndTag, ":") {
		return imageAndTag
	}
	return fmt.Sprintf("%s:latest", imageAndTag)
}
