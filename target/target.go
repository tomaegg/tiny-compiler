package target

import (
	"bufio"
	"io"

	log "github.com/sirupsen/logrus"

	"tinygo.org/x/go-llvm"
)

type BinGenerator struct {
	mod llvm.Module
}

func NewBinGenerator(mod llvm.Module) *BinGenerator {
	return &BinGenerator{mod: mod}
}

func (bg *BinGenerator) Generate(fout io.Writer, codeGen llvm.CodeGenFileType) {
	err := llvm.InitializeNativeTarget()
	if err != nil {
		log.Fatal(err)
	}

	err = llvm.InitializeNativeAsmPrinter()
	if err != nil {
		log.Fatal(err)
	}

	triple := llvm.DefaultTargetTriple()
	target, err := llvm.GetTargetFromTriple(triple)
	if err != nil {
		log.Panic(err)
	}

	const cpu = "generic"
	tm := target.CreateTargetMachine(
		triple,
		cpu, // CPU
		"",  // Features
		llvm.CodeGenLevelDefault,
		llvm.RelocDefault,
		llvm.CodeModelDefault,
	)
	layout := tm.CreateTargetData().String()
	log.WithFields(log.Fields{
		"triple":      triple,
		"cpu":         cpu,
		"data_layout": layout,
	}).Info("generating code")

	bg.mod.SetTarget(triple)
	bg.mod.SetDataLayout(layout)

	if err := llvm.VerifyModule(bg.mod, llvm.PrintMessageAction); err != nil {
		log.Fatal("LLVM IR verification failed: ", err)
	}

	buf, err := tm.EmitToMemoryBuffer(bg.mod, codeGen)
	if err != nil {
		log.Fatal("failed to emit assemblyfile file: ", err)
	}

	writer := bufio.NewWriter(fout)
	n, err := writer.Write(buf.Bytes())
	if err != nil {
		log.Fatal("failed to write object file: ", err)
	}
	if err := writer.Flush(); err != nil {
		log.Fatal("failed to flush writer: ", err)
	}
	log.Infof("code generated with total %d bytes", n)
}
