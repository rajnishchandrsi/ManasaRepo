package plugin

import (
	"fmt"
	"os"
	"strings"

	"github.com/gogo/protobuf/gogoproto"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	validator "github.com/maanasasubrahmanyam-sd/test"
)

const alphaPattern = "^[a-zA-Z()\\[\\]\\-_`'\" \\n\\r\\t&]+$"
const defaultPattern = "^[a-z']+$"
const betaPattern = "^[a-zA-Z0-9]+$"

type plugin struct {
	*generator.Generator
	generator.PluginImports
	regexPkg      generator.Single
	fmtPkg        generator.Single
	validatorPkg  generator.Single
	useGogoImport bool
}

func NewPlugin(useGogoImport bool) generator.Plugin {
	return &plugin{useGogoImport: useGogoImport}
}

func (p *plugin) Name() string {
	return "validator"
}

func (p *plugin) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *plugin) Generate(file *generator.FileDescriptor) {
	p.PluginImports = generator.NewPluginImports(p.Generator)
	p.regexPkg = p.NewImport("regexp")
	p.fmtPkg = p.NewImport("fmt")
	p.validatorPkg = p.NewImport("github.com/maanasasubrahmanyam-sd/test")
	for _, msg := range file.Messages() {
		if msg.DescriptorProto.GetOptions().GetMapEntry() {
			continue
		}
		p.generateRegexVars(file, msg)
		if gogoproto.IsProto3(file.FileDescriptorProto) {
			p.generateProto3Message(file, msg)
		}
	}
}

func getFieldValidatorIfAny(field *descriptor.FieldDescriptorProto) *validator.FieldValidator {
	if field.Options != nil {
		v, err := proto.GetExtension(field.Options, validator.E_Field)
		if err == nil && v.(*validator.FieldValidator) != nil {
			return (v.(*validator.FieldValidator))
		}
	}
	return nil
}

func (p *plugin) generateRegexVars(file *generator.FileDescriptor, message *generator.Descriptor) {
	ccTypeName := generator.CamelCaseSlice(message.TypeName())
	for _, field := range message.Field {
		validator := getFieldValidatorIfAny(field)
		if validator != nil {
			fieldName := p.GetOneOfFieldName(message, field)
			if validator.Alpha != nil && *validator.Alpha {
				alphaPatternStr := strings.Replace(alphaPattern, `\`, `\\`, -1) 
				alphaPatternStr = strings.Replace(alphaPatternStr, `"`, `\"`, -1)

				p.P(`var `, p.regexName(ccTypeName, fieldName), ` = `, p.regexPkg.Use(), `.MustCompile(`, "\"", alphaPatternStr, "\"", `)`)
			} else if validator.Beta != nil && *validator.Beta {
				betaPatternStr := strings.Replace(betaPattern, `"`, `\"`, -1)
				p.P(`var `, p.regexName(ccTypeName, fieldName), ` = `, p.regexPkg.Use(), `.MustCompile(`, "\"", betaPatternStr, "\"", `)`)
			} else {
				// no validation
			}
		} else {
			fieldName := p.GetOneOfFieldName(message, field)
			defaultPatternStr := strings.Replace(defaultPattern, `"`, `\"`, -1)
			p.P(`var `, p.regexName(ccTypeName, fieldName), ` = `, p.regexPkg.Use(), `.MustCompile(`, "\"", defaultPatternStr, "\"", `)`)
		}
	}
}

func (p *plugin) GetFieldName(message *generator.Descriptor, field *descriptor.FieldDescriptorProto) string {
	fieldName := p.Generator.GetFieldName(message, field)
	if p.useGogoImport {
		return fieldName
	}
	if gogoproto.IsEmbed(field) {
		fieldName = generator.CamelCase(*field.Name)
	}
	return fieldName
}

func (p *plugin) generateProto3Message(file *generator.FileDescriptor, message *generator.Descriptor) {
	ccTypeName := generator.CamelCaseSlice(message.TypeName())
	p.P(`func (this *`, ccTypeName, `) Secvalidator() []error {`)
	p.In()
	p.P(`errorsList := []error{}`)
	for _, field := range message.Field {
		fieldValidator := getFieldValidatorIfAny(field)
		fieldName := p.GetOneOfFieldName(message, field)
		variableName := "this." + fieldName

		if fieldValidator == nil && !field.IsMessage() {
			p.generateDefaultValidator(variableName, ccTypeName, fieldName)
			continue
		}

		if field.IsString() {
			fmt.Fprintln(os.Stderr, variableName, ccTypeName, fieldName, fieldValidator)
			p.generateSecValidator(variableName, ccTypeName, fieldName, fieldValidator)
		}

	}
	p.P(`return errorsList`)
	p.Out()
	p.P(`}`)
}

//code
func (p *plugin) generateSecValidator(variableName string, ccTypeName string, fieldName string, fv *validator.FieldValidator) {
	if (fv.Alpha != nil && *fv.Alpha) || (fv.Beta != nil && *fv.Beta) {
		p.P(`if !`, p.regexName(ccTypeName, fieldName), `.MatchString(`, variableName, `) {`)
		p.In()
		errorStr := ""

		if fv.Alpha != nil && *fv.Alpha {
			errorStr = "be a string conforming to alpha regex " + alphaPattern
		} else if fv.Beta != nil && *fv.Beta {
			errorStr = "be a string conforming to beta regex " + betaPattern
		}
		errorStr = strings.Replace(errorStr, `\`, `\\`, -1)
		errorStr = strings.Replace(errorStr, `"`, `\"`, -1)
		p.P(`errorsList = append(errorsList,`, p.validatorPkg.Use(), `.FieldError("`, fieldName, `",`, p.fmtPkg.Use(), `.Errorf(this.`+fieldName+`+" `, errorStr, `")))`)
		p.Out()
		p.P(`}`)
	}
}

func (p *plugin) generateDefaultValidator(variableName string, ccTypeName string, fieldName string) {
	p.P(`if !`, p.regexName(ccTypeName, fieldName), `.MatchString(`, variableName, `) {`)
	p.In()
	errorStr := "be a string conforming to default regex " + defaultPattern
	errorStr = strings.Replace(errorStr, `"`, `\"`, -1)
	p.P(`errorsList = append(errorsList,`, p.validatorPkg.Use(), `.FieldError("`, fieldName, `",`, p.fmtPkg.Use(), `.Errorf(this.`+fieldName+`+" `, errorStr, `")))`)
	p.Out()
	p.P(`}`)
}

func (p *plugin) regexName(ccTypeName string, fieldName string) string {
	return "_regex_" + ccTypeName + "_" + fieldName
}
