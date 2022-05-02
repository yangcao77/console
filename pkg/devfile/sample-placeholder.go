package devfile

const SamplePlaceholderJSON = `[
    {
        "name": "nodejs-basic",
        "displayName": "Basic NodeJS",
        "description": "A simple Hello world NodeJS application",
        "icon": "https://nodejs.org/static/images/logos/nodejs-new-pantone-black.svg",
		"type": "sample",
        "tags": [
            "NodeJS",
            "Express"
        ],
        "projectType": "nodejs",
        "language": "nodejs",
        "git": {
            "remotes": {
                "origin": "https://github.com/redhat-developer/devfile-sample.git"
            }
        }
    },
    {
        "name": "code-with-quarkus",
        "displayName": "Basic Quarkus",
        "description": "A simple Hello World Java application using Quarkus",
        "icon": "https://design.jboss.org/quarkus/logo/final/SVG/quarkus_icon_rgb_default.svg",
		"type": "sample",
        "tags": [
            "Java",
            "Quarkus"
        ],
        "projectType": "quarkus",
        "language": "java",
        "git": {
            "remotes": {
                "origin": "https://github.com/elsony/devfile-sample-code-with-quarkus.git"
            }
        }
    },
    {
        "name": "java-springboot-basic",
        "displayName": "Basic Spring Boot",
        "description": "A simple Hello World Java Spring Boot application using Maven",
        "icon": "https://spring.io/images/projects/spring-edf462fec682b9d48cf628eaf9e19521.svg",
		"type": "sample",
        "tags": [
            "Java",
            "Spring"
        ],
        "projectType": "springboot",
        "language": "java",
        "git": {
            "remotes": {
                "origin": "https://github.com/elsony/devfile-sample-java-springboot-basic.git"
            }
        }
    },
    {
        "name": "python-basic",
        "displayName": "Basic Python",
        "description": "A simple Hello World application using Python",
        "icon": "https://raw.githubusercontent.com/devfile-samples/devfile-stack-icons/main/python.svg",
		"type": "sample",
        "tags": [
            "Python"
        ],
        "projectType": "python",
        "language": "python",
        "git": {
            "remotes": {
                "origin": "https://github.com/yangcao77/devfile-sample-python-basic.git"
            }
        }
    },
    {
        "name": "go-basic",
        "displayName": "Basic Go",
        "description": "A simple Hello World application using Go",
        "icon": "https://raw.githubusercontent.com/devfile-samples/devfile-stack-icons/main/golang.svg",
		"type": "sample",
        "tags": [
            "Go"
        ],
        "projectType": "go",
        "language": "go",
        "git": {
            "remotes": {
                "origin": "https://github.com/yangcao77/go-sample.git"
            }
        }
    }
]
`
