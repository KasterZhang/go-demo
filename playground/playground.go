package playground

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

// printDatasetInfo demonstrates fetching dataset metadata and printing some of it to an io.Writer.
func printDatasetInfo(projectID, datasetID string) error {
	fmt.Println(projectID)
	fmt.Println(datasetID)
	// projectID := "my-project-id"
	// datasetID := "mydataset"
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("bigquery.NewClient: %v", err)
	}
	defer client.Close()

	meta, err := client.Dataset(datasetID).Metadata(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("Dataset ID: %s\n", datasetID)
	fmt.Printf("Description: %s\n", meta.Description)
	fmt.Printf("Labels:")
	for k, v := range meta.Labels {
		fmt.Printf("\t%s: %s", k, v)
	}
	fmt.Printf("Tables:")
	it := client.Dataset(datasetID).Tables(ctx)

	cnt := 0
	for {
		t, err := it.Next()
		if err == iterator.Done {
			break
		}
		cnt++
		fmt.Printf("\t%s\n", t.TableID)
	}
	if cnt == 0 {
		fmt.Printf("\tThis dataset does not contain any tables.")
	}
	return nil
}
