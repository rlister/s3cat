package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"os"
	"regexp"
)

// parsing s3://bucket/key
var s3Regex = regexp.MustCompile("s3://([^/]*)(.*)")

// take s3://bucket/key and return bucket, /key
func parseS3Uri(uri string) (string, string) {
	values := s3Regex.FindStringSubmatch(uri)
	if len(values) == 0 {
		panic("Objects should have form s3://bucket/key")
	}
	return values[1], values[2]
}

// error handler
func check(e error) {
	if e != nil {
		panic(e.Error())
	}
}

func main() {
	sess := session.Must(session.NewSession())
	client := s3.New(sess)

	// loop cmdline args
	for _, arg := range os.Args[1:] {
		bucket, key := parseS3Uri(arg)

		// get response from S3
		resp, err := client.GetObject(&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		})
		check(err)

		// write response to stdout
		_, err = io.Copy(os.Stdout, resp.Body)
		check(err)
		resp.Body.Close()
	}
}
