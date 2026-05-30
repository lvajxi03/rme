pipeline {
agent any

```
environment {
    IMAGE_NAME = "registry.lab.local/rme"
    APP_REPO   = "git@github.com/lvajxi03/rme.git"
    INFRA_REPO = "git@github.com/lvajxi03/k3-lab.git"

    VERSION = "${env.BUILD_NUMBER}"
}

stages {

    stage('Checkout') {
        steps {
            checkout scm
        }
    }

    stage('Prepare') {
        steps {
            script {
                env.VERSION = sh(
                    script: 'git rev-parse --short HEAD',
                    returnStdout: true
                ).trim()
            }
        }
    }

    stage('Build') {
        steps {
            sh '''
                CGO_ENABLED=0 go build -o rme .
            '''
        }
    }

    stage('Build Image') {
        steps {
            sh '''
                /kaniko/executor \
                  --context $WORKSPACE \
                  --dockerfile Dockerfile \
                  --destination ${IMAGE_NAME}:${VERSION} \
                  --destination ${IMAGE_NAME}:latest \
                  --insecure \
                  --skip-tls-verify
            '''
        }
    }

    stage('Update Helm Values') {
        steps {
            dir('infra') {

                git(
                    url: "${INFRA_REPO}",
                    branch: "main"
                )

                sh '''
                    sed -i \
                      "s/tag:.*/tag: \\"${VERSION}\\"/" \
                      apps/rme/values.yaml

                    git config user.name "Jenkins"
                    git config user.email "jenkins@lab.local"

                    git add apps/rme/values.yaml
                    git commit -m "Deploy rme ${VERSION}" || true
                    git push
                '''
            }
        }
    }
}

post {
    success {
        echo "Build complete. ArgoCD will deploy automatically."
    }
}
```

}
