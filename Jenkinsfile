pipeline {
  agent any
  stages {
    stage('Checkout Code') {
      steps {
        git(url: 'https://github.com/jjimgo/go_gin', branch: 'main')
      }
    }

    stage('check ls-al') {
      steps {
        sh 'ls -al'
      }
    }

    stage('Download dependencies') {
      steps {
        sh '''#!/bin/bash

# Download Go package
wget https://golang.org/dl/go1.17.5.linux-amd64.tar.gz

go version'''
      }
    }

    stage('error') {
      steps {
        sh 'go mod download'
      }
    }

  }
}