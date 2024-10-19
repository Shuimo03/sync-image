package sync

import (
	"log"
	"sync-image/config"
	"testing"
)

func TestImageSync_SyncImages(t *testing.T) {
	cnf, err := config.LoadConfig("D:\\code\\sync-image\\cnf\\image.yaml")

	if err != nil {
		t.Fatal("Load config error:", err)
	}

	for _, registry := range cnf.Source.Registries {

		targetRepo := findTargetRepo(cnf.Target.Repositories, registry.Name)
		if targetRepo == "" {
			t.Fatal("Target repository not found:", registry.Name)
		}

		for _, image := range registry.Images {
			sourceImageURL := image
			targetImageURL := buildTargetImageURL(cnf.Target.Registry, targetRepo, extractImageName(image))

			log.Printf("Source Image: %s\n", sourceImageURL)
			log.Printf("Target Image: %s\n", targetImageURL)
		}
	}
}
