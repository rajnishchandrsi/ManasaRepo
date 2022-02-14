// Copyright 2016 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

package secvalidator
import (
	"fmt"
	"os")
// Validator is a general interface that allows a message to be validated.
type Validator interface {
	Validate() error
}

func CallValidatorIfExists(candidate interface{}) error {
	if validator, ok := candidate.(Validator); ok {
		return validator.Validate()
	}
	return nil
}

type fieldError struct {
	fieldStack []string
	nestedErr  error
}

// Generator is the type whose methods generate the output, stored in the associated response structure.
type Generator struct {
	*bytes.Buffer

	Request  *plugin.CodeGeneratorRequest  // The input.
	Response *plugin.CodeGeneratorResponse // The output.

	Param             map[string]string // Command-line parameters.
	PackageImportPath string            // Go import path of the package we're generating code for
	ImportPrefix      string            // String to prefix to imported package file names.
	ImportMap         map[string]string // Mapping from .proto file name to import path

	Pkg map[string]string // The names under which we import support packages

	outputImportPath GoImportPath                   // Package we're generating code for.
	allFiles         []*FileDescriptor              // All files in the tree
	allFilesByName   map[string]*FileDescriptor     // All files by filename.
	genFiles         []*FileDescriptor              // Those files we will generate output for.
	file             *FileDescriptor                // The file we are compiling now.
	packageNames     map[GoImportPath]GoPackageName // Imported package names in the current file.
	usedPackages     map[GoImportPath]bool          // Packages used in current file.
	usedPackageNames map[GoPackageName]bool         // Package names used in the current file.
	addedImports     map[GoImportPath]bool          // Additional imports to emit.`
	typeNameToObject map[string]Object              // Key is a fully-qualified name in input syntax.
	init             []string                       // Lines to emit in the init function.
	indent           string
	pathType         pathType // How to generate output filenames.
	writeOutput      bool
	annotateCode     bool                                       // whether to store annotations
	annotations      []*descriptor.GeneratedCodeInfo_Annotation // annotations to store

	customImports  []string
	writtenImports map[string]bool // For de-duplicating written imports
}

func (f *fieldError) Error() string {
	return "invalid" + f.nestedErr.Error()
}

// FieldError wraps a given Validator error providing a message call stack.
func FieldError(fieldName string, err error) error {
	if fErr, ok := err.(*fieldError); ok {
		fErr.fieldStack = append([]string{fieldName}, fErr.fieldStack...)
		return err
	}
	return &fieldError{
		fieldStack: []string{fieldName},
		nestedErr:  err,
	}
}

func (g *Generator) getOneOfFieldNameHelper(message *Descriptor, field *descriptor.FieldDescriptorProto) string {
	goTyp, _ := g.GoType(message, field)
	fmt.Fprintln(os.Stderr, "gotype is: ", goTyp)
	fieldname := CamelCase(*field.Name)
	if gogoproto.IsCustomName(field) {
		fieldname = gogoproto.GetCustomName(field)
	}
	if gogoproto.IsEmbed(field) {
		fieldname = EmbedFieldName(goTyp)
	}
	for _, f := range methodNames {
		if f == fieldname {
			return fieldname + "_"
		}
	}
	if !gogoproto.IsProtoSizer(message.file.FileDescriptorProto, message.DescriptorProto) {
		if fieldname == "Size" {
			return fieldname + "_"
		}
	}
	return fieldname
}