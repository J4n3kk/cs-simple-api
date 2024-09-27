pipeline {
    agent { docker { image 'golang:1.23.1-alpine3.20' } }
    stages {
        stage('Check Docker') {
      steps {
        sh 'which docker'
        sh 'docker --version'
      }
    }
        stage('build') {
            steps {
                sh 'go version'
            }
        }
    }
}