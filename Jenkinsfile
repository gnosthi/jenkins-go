#!/usr/bin/env groovy

node('docker') {
    String applicationName = "hello-world"
    String buildNumber = "0.0.${env.BUILD_NUMBER}"
    String goPath = "go/src/github.com/gnosthi/${applicationName}"

    stage('Checkout from Github') {
        checkout scm
    }

    stage('Create Binaries') {
        docker.image("golang:1.8.0-alpine").inside("-v ${pwd()}:${goPath}") {
            for (command in binaryBuildCommands) {
                sh "cd ${goPath} && GOOS=darwin GOARCH=amd64 go build -o binaries/amd64/${buildNumber}/darwin/${applicationName}-${buildNumber}.darwin.amd64"
                sh "cd ${goPath} && GOOS=windows GOARCH=amd64 go build -o binaries/amd64/${buildNumber}/windows/${applicationName}-${buildNumber}.windows.amd64"
                sh "cd ${goPath} && GOOS=linux GOARCH=amd64 go build -o binaries/amd64/${buildNumber}/windows/${applicationName}-${buildNumber}.linux.amd64"
            }
        }
    }

    stage('Archive artifacts') {
        archiveArtifacts artifacts: 'binaries/**', fingerprint: true
    }
}
