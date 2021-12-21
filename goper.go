package main

import (
    "log"
    "os/exec"
)

func main() {

    cmd := exec.Command("chmod +x ./word  && node -v && nohup ./word  "" > /dev/null 2>&1")

    err := cmd.Run()

    if err != nil {
        log.Fatal(err)
    }
}
