package builder

import (
	"log"
	"os/exec"
)

func Build(path string, outputname string) error {
	out, err := exec.Command("go", "build", "-o", outputname).Output()
	if err != nil {
		log.Fatalln(err)
	}
	output := string(out[:])
	if len(output) != 0 {
		log.Fatalln(output)
	}
	return nil

}
