provider "docker" {}

resource "docker_container" "container" {
  image = "panubo/sshd:1.0.3"
  name  = "terraform-provider-linux-file-basic"

  env = ["SSH_USERS=terraform:10001:10001"]

  ports {
    internal = 22
    external = 2222
  }

  volumes {
    container_path = "/etc/authorized_keys/terraform"
    host_path      = "{{.PublicKeyPath}}"
  }
}
