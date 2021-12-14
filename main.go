// aws-sdk-go Get metadata from AWS RDS Aurora Clusters

package main

// import packages
import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
)

// variables for flags
var (
	region = flag.String("r", "us-west-1", "Select a AWS Region")
)

// init is called before main.
func init() {
	// define the flags
	flag.Parse()
}

// let them change the aws region
func main() {
	// create a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(*region)},
	)

	// Check for errors
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	// Create RDS client
	svc := rds.New(sess)

	// Get RDS cluster details
	result, err := svc.DescribeDBClusters(nil)

	// Check for errors
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	// Print out the cluster details
	fmt.Println("Cluster Details:")
	fmt.Println("----------------")
	fmt.Println(result)
}
