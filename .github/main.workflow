workflow "Build" {
  on = "release"
  resolves = [
    "release linux/amd64",
  ]
}

action "release linux/amd64" {
  uses = "ngs/go-release.action@v1.0.1"
  env = {
    GOOS = "linux"
    GOARCH = "amd64"
    GO111MODULE = "on"
  }
  secrets = ["GITHUB_TOKEN"]
}
