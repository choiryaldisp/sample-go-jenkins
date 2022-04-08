pipeline{
    agent any

    environment{
        root = "/usr/local/go/bin/go"
        branch = "master"
        scmUlr = "https://github.com/choiryaldisp/sample-go-jenkins.git"
    }

    stages{
        stage('Git clone'){
            steps{
                git branch:"${branch}", url:"${scmUlr}"
            }
        }

        stage('Go Build Docker'){
            steps{
               sh "docker build -t sample-go-jenkins ."
            }
        }

        stage('Go Deploy'){
            steps{
               echo "Deploy success"
            }
        }
    }
}


// // Run on an agent where we want to use Go
// node {
//     // Ensure the desired Go version is installed
//     def root = "/usr/local/go/bin/go"

//     stage 'Checkout'
//     git url: 'https://github.com/choiryaldisp/sample-go-jenkins.git'

//     stage 'preTest'
//     sh "${root} version"

//     stage 'Test'
//     sh "${root} test ./... -cover"

//     stage 'Build'
//     sh "${root} build ./..."
// }