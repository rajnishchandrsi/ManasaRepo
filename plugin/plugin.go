package plugin

import (
	"fmt"
	"os"
	"strings"

	"github.com/gogo/protobuf/gogoproto"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	validator "github.com/maanasasubrahmanyam-sd/test/secvalidator"
)

const alphaPattern = "^[a-zA-Z()\\[\\]\\-_`'\" \\n\\r\\t&]+$"
const defaultPattern = "^[a-z']+$"
const betaPattern = "^[a-zA-Z0-9\\[\\]]+$"

type plugin struct {
	*generator.Generator
	generator.PluginImports
	regexPkg      generator.Single
	fmtPkg        generator.Single
	validatorPkg  generator.Single
	osPkg         generator.Single
	jsonPkg       generator.Single
	zapPkg        generator.Single
	flagPkg       generator.Single
	stringsPkg    generator.Single
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
	p.osPkg = p.NewImport("os")
	p.flagPkg = p.NewImport("flag")
	p.jsonPkg = p.NewImport("encoding/json")
	p.zapPkg = p.NewImport("go.uber.org/zap")
	p.stringsPkg = p.NewImport("strings")
	p.validatorPkg = p.NewImport("github.com/maanasasubrahmanyam-sd/test/secvalidator")
	p.P(`type ErrorList []error`)
	//p.logger()

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

func (p *plugin) logger() {

	p.P(`var logger *`, p.zapPkg.Use(), `.Logger`)
	p.P(`func init() {`)
	p.In()
	p.P(`var fileName *string`)
	p.P(`var debugLevel *string`)

	p.P(`if `, p.flagPkg.Use(), `.Lookup("fileName") == nil {`)
	p.In()
	p.P(`fileName = `, p.flagPkg.Use(), `.String("fileName", "TestLogs" , "default message")`)
	p.Out()
	p.P(`}`)

	p.P(`if `, p.flagPkg.Use(), `.Lookup("debugLevel") == nil {`)
	p.In()
	p.P(`debugLevel = `, p.flagPkg.Use(), `.String("debugLevel", "info" , "default message")`)
	p.Out()
	p.P(`}`)

	p.P(p.flagPkg.Use(), `.Parse()`)

	p.P("rawJSON := []byte(`{\"level\": \"`+ " + p.stringsPkg.Use() + ".ToLower(*debugLevel)+`\",\"encoding\": \"json\",\"outputPaths\": [\"stdout\", \"`+*fileName+`\"],\"errorOutputPaths\": [\"stderr\"],\"encoderConfig\": {\"messageKey\": \"message\",\"levelKey\": \"level\", \"levelEncoder\": \"lowercase\"}}`)")
	p.P(`var cfg `, p.zapPkg.Use(), `.Config`)
	p.P(`var err error`)
	p.P(`if err := `, p.jsonPkg.Use(), `.Unmarshal(rawJSON, &cfg); err != nil {`)
	p.P(p.fmtPkg.Use(), `.Println("Only debug, info, error level can be set.")`)
	p.P(`panic(err)`)
	p.P(`}`)
	p.P(`logger, err = cfg.Build()`)
	p.P(`if err != nil {`)
	p.In()
	p.P(`panic(err)`)
	p.Out()
	p.P(`}`)
	p.P(`defer logger.Sync()`)
	p.Out()
	p.P(`}`)

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
				betaPatternStr := strings.Replace(betaPattern, `\`, `\\`, -1)
				p.P(`var `, p.regexName(ccTypeName, fieldName), ` = `, p.regexPkg.Use(), `.MustCompile(`, "\"", betaPatternStr, "\"", `)`)
			} else {
				// no validation
			}
		} else {
			if field.IsString() && field.Options == nil {
				fieldName := p.GetOneOfFieldName(message, field)
				p.P(`var `, p.regexName(ccTypeName, fieldName), ` = `, p.regexPkg.Use(), `.MustCompile(`, "\"", defaultPattern, "\"", `)`)
			}
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
	p.P(`func (this *`, ccTypeName, `) Secvalidator() ErrorList {`)
	p.In()
	p.P(`logger.Info("info logger construction succeeded")`)
	p.P(`logger.Debug("debug logger construction succeeded")`)
	p.P(`var errorsList ErrorList`)
	for _, field := range message.Field {
		fieldValidator := getFieldValidatorIfAny(field)
		fieldName := p.getOneOfFieldNameHelper(message, field)
		variableName := "this." + fieldName
		fmt.Fprintln(os.Stderr, "all fileds :", field)

		repeated := field.IsRepeated()
		if repeated {
			//field this.Name ValidatorMessage3 Name beta:true message ValidatorMessage3 Name
			fmt.Fprintln(os.Stderr, "Maanasa :", variableName, ccTypeName, fieldName, fieldValidator)
			p.generateValidatorRepetaed(variableName, ccTypeName, fieldName, fieldValidator)
		}

		if fieldValidator == nil && !field.IsMessage() {
			if field.IsString() && field.Options == nil {
				p.generateDefaultValidator(variableName, ccTypeName, fieldName)
			}
			continue
		}
		if field.IsString() {
			p.generateSecValidator(variableName, ccTypeName, fieldName, fieldValidator)
		}

	}
	p.P(`return errorsList`)
	p.Out()
	p.P(`}`)
}

//code this.Name ValidatorMessage3 Name beta:true message ValidatorMessage3 Name

func (p *plugin) generateValidatorRepetaed(variableName string, ccTypeName string, fieldName string, fv *validator.FieldValidator) {
	p.P(`for _, val := range `, variableName, `{`)
	if (fv.Alpha != nil && *fv.Alpha) || (fv.Beta != nil && *fv.Beta) {
		p.P(`if !`, p.regexName(ccTypeName, fieldName), `.MatchString(val) {`)
		p.In()
		errorStr := ""

		fmt.Fprintln(os.Stderr, "message", ccTypeName, fieldName)
		if fv.Alpha != nil && *fv.Alpha {
			p.P(`alPattern := "[a-zA-Z]"`)
			p.P(`reg := regexp.MustCompile(alPattern)`)
			p.P(`res := reg.ReplaceAllString(val,"")`)
			errorStr = " \" allowed " + alphaPattern
		} else if fv.Beta != nil && *fv.Beta {
			p.P(`btPattern := "[a-zA-Z0-9]"`)
			p.P(`reg := regexp.MustCompile(btPattern)`)
			p.P(`res := reg.ReplaceAllString(val,"")`)
			errorStr = " \\ allowed beta " + betaPattern
		}
		errorStr = strings.Replace(errorStr, `\`, `\\`, -1)
		errorStr = strings.Replace(errorStr, `"`, `\"`, -1)
		p.P(`errorsList = append(errorsList,`, p.validatorPkg.Use(), `.FieldError("`, fieldName, `",`, p.fmtPkg.Use(), `.Errorf("%v"," `, ccTypeName+"."+fieldName+": "+errorStr, `" +res)))`)
		p.Out()
		p.P(`}`)
		p.P(`}`)

	}
}

func (p *plugin) generateDefaultValidatorForRepeated(variableName string, ccTypeName string, fieldName string) {
	p.P(`for _, val := range `, variableName, `{`)
	p.P(`if !`, p.regexName(ccTypeName, fieldName), `.MatchString(`, variableName, `) {`)
	p.In()
	errorStr := "be a string conforming to default regex " + defaultPattern
	errorStr = strings.Replace(errorStr, `"`, `\"`, -1)
	p.P(`errorsList = append(errorsList,`, p.validatorPkg.Use(), `.FieldError("`, fieldName, `",`, p.fmtPkg.Use(), `.Errorf("%v"," `, ccTypeName+"."+fieldName+": "+errorStr, `")))`)
	p.Out()
	p.P(`}`)
	p.P(`}`)

}

func (p *plugin) generateSecValidator(variableName string, ccTypeName string, fieldName string, fv *validator.FieldValidator) {

	if (fv.Alpha != nil && *fv.Alpha) || (fv.Beta != nil && *fv.Beta) {
		p.P(`if !`, p.regexName(ccTypeName, fieldName), `.MatchString(`, variableName, `) {`)
		p.In()
		errorStr := ""

		fmt.Fprintln(os.Stderr, "message", ccTypeName, fieldName)
		if fv.Alpha != nil && *fv.Alpha {
			p.P(`alPattern := "[a-zA-Z]"`)
			p.P(`reg := regexp.MustCompile(alPattern)`)
			p.P(`res := reg.ReplaceAllString(`, variableName, `,"")`)
			errorStr = " \" allowed " + alphaPattern
		} else if fv.Beta != nil && *fv.Beta {
			p.P(`btPattern := "[a-zA-Z0-9]"`)
			p.P(`reg := regexp.MustCompile(btPattern)`)
			p.P(`res := reg.ReplaceAllString(`, variableName, `,"")`)
			errorStr = " \\ allowed beta " + betaPattern
		}
		errorStr = strings.Replace(errorStr, `\`, `\\`, -1)
		errorStr = strings.Replace(errorStr, `"`, `\"`, -1)
		p.P(`errorsList = append(errorsList,`, p.validatorPkg.Use(), `.FieldError("`, fieldName, `",`, p.fmtPkg.Use(), `.Errorf("%v"," `, ccTypeName+"."+fieldName+": "+errorStr, `" +res)))`)
		p.Out()
		p.P(`}`)
	}
}

func (p *plugin) generateDefaultValidator(variableName string, ccTypeName string, fieldName string) {
	p.P(`if !`, p.regexName(ccTypeName, fieldName), `.MatchString(`, variableName, `) {`)
	p.In()
	errorStr := "be a string conforming to default regex " + defaultPattern
	errorStr = strings.Replace(errorStr, `"`, `\"`, -1)
	p.P(`errorsList = append(errorsList,`, p.validatorPkg.Use(), `.FieldError("`, fieldName, `",`, p.fmtPkg.Use(), `.Errorf("%v"," `, ccTypeName+"."+fieldName+": "+errorStr, `")))`)
	p.Out()
	p.P(`}`)
}

func (p *plugin) regexName(ccTypeName string, fieldName string) string {
	return "_regex_" + ccTypeName + "_" + fieldName
}
