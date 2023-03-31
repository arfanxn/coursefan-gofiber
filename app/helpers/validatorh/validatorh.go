package validatorh

import (
	"fmt"
	"mime/multipart"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/arfanxn/coursefan-gofiber/app/exceptions"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/numh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/synch"
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
	syncronizer := synch.NewSyncronizer()
	structValue := reflect.ValueOf(structure)
	structType := structValue.Type()
	for i := 0; i < structValue.NumField(); i++ {
		syncronizer.WG().Add(1)
		go func(structField reflect.StructField, fieldValue reflect.Value) {
			defer syncronizer.WG().Done()
			if syncronizer.Err() != nil {
				return
			}
			jsonFieldName := strcase.ToSnake(structField.Name)
			jsonTags := strings.Split(structField.Tag.Get("json"), ",")
			if len(jsonTags) >= 1 {
				jsonFieldName = jsonTags[0]
			}
			rules := strings.Split(structField.Tag.Get("fhlidate"), ",")
			fileHeader, ok := fieldValue.Interface().(*multipart.FileHeader)
			if !ok || (len(rules) == 0) {
				return // return if no rules
			}
			// Get max media upload size from application environment variable
			maxUploadSize, err := strconv.ParseInt(os.Getenv("MEDIA_MAX_SIZE"), 10, 64)
			if err != nil {
				syncronizer.Err(err)
				return
			}
			maxUploadSize = numh.MegabyteToByte(maxUploadSize)
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
					mb, err := strconv.ParseInt(strings.SplitAfter(rule, "min=")[1], 10, 64)
					if err != nil {
						syncronizer.Err(err)
						return
					}
					min = numh.MegabyteToByte(mb)
					break
				case strings.Contains(rule, "max="):
					mb, err := strconv.ParseInt(strings.SplitAfter(rule, "max=")[1], 10, 64)
					if err != nil {
						syncronizer.Err(err)
						return
					}
					max = numh.MegabyteToByte(mb)
					break
				case strings.Contains(rule, "mimes="):
					mimeTypes = strings.Split(strings.SplitAfter(rule, "mimes=")[1], " ")
					break
				}
			}
			// if file is not required in rule and the file is not provided then immediately return nil error
			if !required && (fileHeader == nil) {
				return
			}
			// check if file is required but not provided
			if (required) && ((fileHeader == nil) || (fileHeader.Size == 0)) {
				syncronizer.M().Lock()
				validationErrs = append(validationErrs,
					exceptions.NewValidationError(jsonFieldName, jsonFieldName+" is required"))
				syncronizer.M().Unlock()
				return
			}
			// check if file size is not between the specified min and max size
			if fileHeader.Size < min || fileHeader.Size > max {
				syncronizer.M().Lock()
				validationErrs = append(
					validationErrs,
					exceptions.NewValidationError(
						jsonFieldName,
						fmt.Sprintf("%s must be between %d and %d size", jsonFieldName, min, max),
					),
				)
				syncronizer.M().Unlock()
				return
			}
			// check if mimetype are available in rules
			if len(mimeTypes) > 0 {
				file, err := fileHeader.Open()
				if err != nil {
					syncronizer.Err(err)
					return
				}
				defer file.Close()
				fileHeaderMime, err := mimetype.DetectReader(file)
				matchedMimeTypes := sliceh.Filter(mimeTypes, func(mimeType string) bool {
					return strings.ToLower(fileHeaderMime.String()) == strings.ToLower(mimeType)
				})

				// if no matching mime types found append an error
				if len(matchedMimeTypes) == 0 {
					syncronizer.M().Lock()
					validationErrs = append(
						validationErrs,
						exceptions.NewValidationError(jsonFieldName,
							fmt.Sprintf(
								"%s must be a file of types %s",
								jsonFieldName,
								strings.Join(mimeTypes, " or "),
							),
						),
					)
					syncronizer.M().Unlock()
					return
				}
			}
		}(structType.Field(i), structValue.Field(i))
	}
	syncronizer.WG().Wait()

	if err := syncronizer.Err(); err != nil {
		panic(err)
	}

	syncronizer.Close()

	return validationErrs
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