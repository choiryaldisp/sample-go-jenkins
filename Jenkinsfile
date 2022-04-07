pipeline{
    agent any

    environment{
        root = "/usr/local/go/bin/go"
        branch = "master"
        scmUlr = "https://github.com/choiryaldisp/sample-go-jenkins.git"
    }

    stages{
        stage('Go version'){
            steps{
                sh "${root} version"
            }
        }

        stage('Git clone'){
            steps{
                git branch:"${branch}", url:"${scmUlr}"
            }
        }

        stage('Go Test'){
            steps{
               sh "${root} test ./... -cover"
            }
        }

        stage('Go Build'){
            steps{
               sh "${root} build ./..."
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