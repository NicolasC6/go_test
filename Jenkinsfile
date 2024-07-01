//Jenkinsfile (Declarative Pipeline)
pipeline {
    agent any
    
    triggers {
        pollSCM('') // Enabling being build on Push
    }

    stages {
        stage('Build') {
            steps {
                echo 'Building..'
                echo 'Building..'
                cd hello
                go run .
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}