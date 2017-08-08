# gosync

I want to be the fastest way to concurrently sync files and directories to/from S3.

Gosync will concurrently transfer your files to and from S3 (or across different S3 
buckets). It will validate checksyms to ensure that only new or changed files are
synced.

# Installation

Ensure you have Go 1.2 or greater installed and your GOPATH is set.

Clone the repo:

    go get github.com/ivancevich/gosync

Change into the gosync directory and run make:

    cd $GOPATH/src/github.com/ivancevich/gosync/
    make

# Setup

Set environment variables (Security Token is optional):

    AWS_ACCESS_KEY_ID=xxx
    AWS_ACCESS_KEY_SECRET=yyy
    AWS_SECURITY_TOKEN=xxx

# Usage

    gosync OPTIONS SOURCE TARGET

## Syncing from local directory to S3

    gosync /files s3://bucket/files

## Syncing from S3 to local directory

    gosync s3://bucket/files /files

## Syncing from S3 to S3

    gosync s3://source_bucket s3://target_bucket

## Syncing from S3 to another directory in S3

    gosync s3://source_bucket/dir s3://target_bucket/another_dir

## Help

For full list of options and commands:

    gosync -h

# Contributing

1. Fork it
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create new Pull Request
