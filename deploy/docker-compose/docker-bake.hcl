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
  ]
}

target "bugu-bugu-service" {
  context    = "."
  dockerfile = "app/bugu/service/Dockerfile"
  args       = {
    AUTHOR_NAME       = "${AUTHOR_NAME}"
    AUTHOR_EMAIL      = "${AUTHOR_EMAIL}"
    APP_RELATIVE_PATH = "wx/service"
  }
  tags = [
    "${REPO}/bugu-bugu-service:latest",
    notequal("", VERSION) ? "${REPO}/bugu-bugu-service:${VERSION}" : "",
  ]
  platforms = ["linux/arm64", "linux/amd64"]
}