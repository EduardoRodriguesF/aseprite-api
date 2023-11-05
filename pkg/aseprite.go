package aseprite

import (
	"os"
	"os/exec"
)

type ExportOptions struct {
	OutputFile string
}

type Aseprite struct {
	bin string
}

type AsepriteExporter struct {
	aseprite *Aseprite
	File     *os.File
	Options  ExportOptions
}

func NewAseprite(path string) *Aseprite {
	ase := &Aseprite{bin: path}

	return ase
}

func (ase *Aseprite) command(args ...string) ([]byte, error) {
	defaultArgs := []string{"-b"}
	args = append(defaultArgs, args...)

	return exec.Command(ase.bin, args...).Output()
}

func (ase *Aseprite) Export(file *os.File, options ExportOptions) *AsepriteExporter {
	export := &AsepriteExporter{
		File:     file,
		Options:  options,
		aseprite: ase,
	}

	return export
}

func (exporter *AsepriteExporter) Sheet() ([]byte, error) {
	args := []string{exporter.File.Name(), "--sheet", exporter.Options.OutputFile}

	return exporter.aseprite.command(args...)
}
