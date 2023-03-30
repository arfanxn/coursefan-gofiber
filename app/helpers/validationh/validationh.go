package validationh

import (
	"fmt"
	"mime/multipart"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/arfanxn/coursefan-gofiber/app/exceptions"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	validator_provider "github.com/arfanxn/coursefan-gofiber/app/providers/validators"
	"github.com/gabriel-vasile/mimetype"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
)

// ValidateStruct validates the given struct
func ValidateStruct[T any](structure T, lang string) []*exceptions.ValidationError {
	validate := validator.New()
	err := validate.Struct(structure)
	lang = strings.ToLower(lang)

	translators := map[string]func(*validator.Validate) ut.Translator{
		"en": validator_provider.EnglishTranslator,
	}
	translator := translators[lang]
	utTrans := translator(validate)
	translatedErrs := TranslateErrs(err, utTrans)

	translatedErrs = append(translatedErrs, ValidateStructFileHeader(structure, lang)...)

	return translatedErrs
}

// ValidateStructFileHeader validates the struct fields on *multipart.FileHeader type
func ValidateStructFileHeader[T any](structure T, lang string) (
	validationErrs []*exceptions.ValidationError) {
	structValue := reflect.ValueOf(structure)
	structType := structValue.Type()
	for i := 0; i < structValue.NumField(); i++ {
		field := structType.Field(i)
		jsonFieldName := strcase.ToSnake(field.Name)
		jsonTags := strings.Split(field.Tag.Get("json"), ",")
		if len(jsonTags) >= 1 {
			jsonFieldName = jsonTags[0]
		}
		rules := strings.Split(field.Tag.Get("fhlidate"), ",")
		fieldValue := structValue.Field(i).Interface()
		fileHeader, ok := fieldValue.(*multipart.FileHeader)
		if !ok || (len(rules) == 0) {
			continue
		}
		// Get max upload size from application environment variable
		maxUploadSize, err := strconv.ParseInt(os.Getenv("MAX_UPLOAD_SIZE"), 10, 64)
		if err != nil {
			panic(err)
		}
		var (
			required  bool  = false
			min       int64 = 0
			max       int64 = maxUploadSize
			mimeTypes []string
		)
		for _, rule := range rules {
			switch true {
			case rule == "required":
				required = true
				break
			case strings.Contains(rule, "min="):
				num, err := strconv.ParseInt(strings.SplitAfter(rule, "min=")[1], 10, 64)
				if err != nil {
					panic(err)
				}
				min = num
				break
			case strings.Contains(rule, "max="):
				num, err := strconv.ParseInt(strings.SplitAfter(rule, "max=")[1], 10, 64)
				if err != nil {
					panic(err)
				}
				max = num
				break
			case strings.Contains(rule, "mimes="):
				mimeTypes = strings.Split(strings.SplitAfter(rule, "mimes=")[1], " ")
				break
			}
		}
		// if file is not required in rule and the file is not provided then immediately return nil error
		if !required && (fileHeader == nil) {
			continue
		}
		// check if file is required but not provided
		if (required) && ((fileHeader == nil) || (fileHeader.Size == 0)) {
			validationErrs = append(validationErrs,
				exceptions.NewValidationError(jsonFieldName, jsonFieldName+" is required"))
			continue
		}
		// check if file size is not between the specified min and max size
		if fileHeader.Size < min || fileHeader.Size > max {
			validationErrs = append(
				validationErrs,
				exceptions.NewValidationError(
					jsonFieldName,
					fmt.Sprintf("%s must be between %d and %d size", jsonFieldName, min, max),
				),
			)
			continue
		}
		// check if mimetype are available in rules
		if len(mimeTypes) > 0 {
			file, err := fileHeader.Open()
			if err != nil {
				panic(err)
			}
			defer file.Close()
			fileHeaderMime, err := mimetype.DetectReader(file)
			matchedMimeTypes := sliceh.Filter(mimeTypes, func(mimeType string) bool {
				return strings.ToLower(fileHeaderMime.String()) == strings.ToLower(mimeType)
			})

			// if no matching mime types found append an error
			if len(matchedMimeTypes) == 0 {
				validationErrs = append(
					validationErrs,
					exceptions.NewValidationError(jsonFieldName,
						fmt.Sprintf(
							"%s must be a file of types %s",
							jsonFieldName,
							strings.Join(mimeTypes, ", "),
						),
					),
				)
				continue
			}
		}
	}
	return
}

// TranslateErrs translates errors from validation errors
func TranslateErrs(errs error, trans ut.Translator) (translatedErrs []*exceptions.ValidationError) {
	if errs == nil {
		return nil
	}
	validationErrs := errs.(validator.ValidationErrors)
	for _, validationErr := range validationErrs {
		fieldNamespace := validationErr.StructNamespace()
		fieldName := strings.Join(strings.SplitAfter(fieldNamespace, ".")[1:], ".")
		jsonFieldName := strcase.ToSnake(fieldName)
		if name := validationErr.Field(); name != "" {
			jsonFieldName = strcase.ToSnake(name)
		}
		message := validationErr.Translate(trans)
		translatedErr := exceptions.NewValidationError(jsonFieldName, message)
		translatedErrs = append(translatedErrs, translatedErr)
	}
	return
}
