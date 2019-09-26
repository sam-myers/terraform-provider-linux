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

data "local_file" "ssh_private_key" {
  filename = "{{.PrivateKeyPath}}"
}

data "linux_ssh_connection" "docker" {
  user        = "terraform"
  host        = "localhost"
  port        = 2222
  private_key = data.local_file.ssh_private_key.content
}

resource "linux_file" "test_txt" {
  depends_on    = ["docker_container.container"]
  content       = "foo bar biz baz"
  destination   = "/home/terraform/test.txt"
  connection_id = data.linux_ssh_connection.docker.id
}
