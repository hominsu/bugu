variable "REPO" {
  default = "hominsu"
}

variable "AUTHOR_NAME" {
  default = "hominsu"
}

variable "AUTHOR_EMAIL" {
  default = "hominsu@foxmail.com"
}

variable "VERSION" {
  default = ""
}

group "default" {
  targets = [
    "bugu-bugu-service",
    "bugu-obfusion-service",
  ]
}

target "bugu-bugu-service" {
  context    = "."
  dockerfile = "app/bugu/service/Dockerfile"
  args       = {
    AUTHOR_NAME       = "${AUTHOR_NAME}"
    AUTHOR_EMAIL      = "${AUTHOR_EMAIL}"
    APP_RELATIVE_PATH = "bugu/service"
  }
  tags = [
    "${REPO}/bugu-bugu-service:latest",
    notequal("", VERSION) ? "${REPO}/bugu-bugu-service:${VERSION}" : "",
  ]
  platforms = ["linux/arm64"]
}

target "bugu-obfusion-service" {
  context    = "."
  dockerfile = "app/obfusion/service/Dockerfile"
  args       = {
    AUTHOR_NAME       = "${AUTHOR_NAME}"
    AUTHOR_EMAIL      = "${AUTHOR_EMAIL}"
    APP_RELATIVE_PATH = "obfusion/service"
  }
  tags = [
    "${REPO}/bugu-obfusion-service:latest",
    notequal("", VERSION) ? "${REPO}/bugu-obfusion-service:${VERSION}" : "",
  ]
  platforms = ["linux/arm64"]
}