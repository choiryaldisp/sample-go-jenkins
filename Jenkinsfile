// Run on an agent where we want to use Go
node {
    // Ensure the desired Go version is installed
    def root = "/usr/local/go/bin/go"

    stage 'Checkout'
    git url: 'https://github.com/choiryaldisp/sample-go-jenkins.git'

    stage 'preTest'
    sh "${root} go version"

    stage 'Test'
    sh "${root} go test ./... -cover"

    stage 'Build'
    sh "${root} go build ./..."
}