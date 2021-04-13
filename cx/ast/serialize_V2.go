package ast

import "github.com/skycoin/skycoin/src/cipher/encoder"

// SerializeCXProgramV2 translates cx program to slice of bytes that we can save.
// These slice of bytes can then be deserialize in the future and
// be translated back to cx program.
func SerializeCXProgramV2(prgrm *CXProgram, includeMemory bool) (b []byte) {
	s := SerializedCXProgram{}
	initSerialization(prgrm, &s, includeMemory)

	// serialize cx program's packages,
	// structs, functions, etc.
	serializeCXProgramElements(prgrm, &s)

	// serialize cx program's program
	serializeProgram(prgrm, &s)

	s.Index = serializedCXProgramIndex{}
	sIdx := &s.Index

	// assigning relative offset
	idxSize := encoder.Size(s.Index)
	prgrmSize := encoder.Size(s.Program)
	callSize := encoder.Size(s.Calls)
	pkgSize := encoder.Size(s.Packages)
	strctSize := encoder.Size(s.Structs)
	fnSize := encoder.Size(s.Functions)
	exprSize := encoder.Size(s.Expressions)
	argSize := encoder.Size(s.Arguments)
	intSize := encoder.Size(s.Integers)

	// assigning absolute offset
	sIdx.ProgramOffset += int64(idxSize)
	sIdx.CallsOffset += sIdx.ProgramOffset + int64(prgrmSize)
	sIdx.PackagesOffset += sIdx.CallsOffset + int64(callSize)
	sIdx.StructsOffset += sIdx.PackagesOffset + int64(pkgSize)
	sIdx.FunctionsOffset += sIdx.StructsOffset + int64(strctSize)
	sIdx.ExpressionsOffset += sIdx.FunctionsOffset + int64(fnSize)
	sIdx.ArgumentsOffset += sIdx.ExpressionsOffset + int64(exprSize)
	sIdx.IntegersOffset += sIdx.ArgumentsOffset + int64(argSize)
	sIdx.StringsOffset += sIdx.IntegersOffset + int64(intSize)
	sIdx.MemoryOffset += sIdx.StringsOffset + int64(len(s.Strings))

	// serializing everything
	b = encoder.Serialize(s)

	return b
}
