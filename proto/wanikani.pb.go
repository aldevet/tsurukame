// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wanikani.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	wanikani.proto

It has these top-level messages:
	Meaning
	Reading
	Radical
	Kanji
	Vocabulary
	Subject
	Assignment
	Progress
	StudyMaterials
	User
	SubjectOverrides
	FormattedText
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Reading_Type int32

const (
	Reading_ONYOMI  Reading_Type = 1
	Reading_KUNYOMI Reading_Type = 2
	Reading_NANORI  Reading_Type = 3
)

var Reading_Type_name = map[int32]string{
	1: "ONYOMI",
	2: "KUNYOMI",
	3: "NANORI",
}
var Reading_Type_value = map[string]int32{
	"ONYOMI":  1,
	"KUNYOMI": 2,
	"NANORI":  3,
}

func (x Reading_Type) Enum() *Reading_Type {
	p := new(Reading_Type)
	*p = x
	return p
}
func (x Reading_Type) String() string {
	return proto1.EnumName(Reading_Type_name, int32(x))
}
func (x *Reading_Type) UnmarshalJSON(data []byte) error {
	value, err := proto1.UnmarshalJSONEnum(Reading_Type_value, data, "Reading_Type")
	if err != nil {
		return err
	}
	*x = Reading_Type(value)
	return nil
}
func (Reading_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type Vocabulary_PartOfSpeech int32

const (
	Vocabulary_NOUN              Vocabulary_PartOfSpeech = 1
	Vocabulary_NUMERAL           Vocabulary_PartOfSpeech = 2
	Vocabulary_INTRANSITIVE_VERB Vocabulary_PartOfSpeech = 3
	Vocabulary_ICHIDAN_VERB      Vocabulary_PartOfSpeech = 4
	Vocabulary_TRANSITIVE_VERB   Vocabulary_PartOfSpeech = 5
	Vocabulary_NO_ADJECTIVE      Vocabulary_PartOfSpeech = 6
	Vocabulary_GODAN_VERB        Vocabulary_PartOfSpeech = 7
	Vocabulary_NA_ADJECTIVE      Vocabulary_PartOfSpeech = 8
	Vocabulary_I_ADJECTIVE       Vocabulary_PartOfSpeech = 9
	Vocabulary_SUFFIX            Vocabulary_PartOfSpeech = 10
	Vocabulary_ADVERB            Vocabulary_PartOfSpeech = 11
	Vocabulary_SURU_VERB         Vocabulary_PartOfSpeech = 12
	Vocabulary_PREFIX            Vocabulary_PartOfSpeech = 13
	Vocabulary_PROPER_NOUN       Vocabulary_PartOfSpeech = 14
	Vocabulary_EXPRESSION        Vocabulary_PartOfSpeech = 15
	Vocabulary_ADJECTIVE         Vocabulary_PartOfSpeech = 16
	Vocabulary_INTERJECTION      Vocabulary_PartOfSpeech = 17
	Vocabulary_COUNTER           Vocabulary_PartOfSpeech = 18
	Vocabulary_PRONOUN           Vocabulary_PartOfSpeech = 19
	Vocabulary_CONJUNCTION       Vocabulary_PartOfSpeech = 20
)

var Vocabulary_PartOfSpeech_name = map[int32]string{
	1:  "NOUN",
	2:  "NUMERAL",
	3:  "INTRANSITIVE_VERB",
	4:  "ICHIDAN_VERB",
	5:  "TRANSITIVE_VERB",
	6:  "NO_ADJECTIVE",
	7:  "GODAN_VERB",
	8:  "NA_ADJECTIVE",
	9:  "I_ADJECTIVE",
	10: "SUFFIX",
	11: "ADVERB",
	12: "SURU_VERB",
	13: "PREFIX",
	14: "PROPER_NOUN",
	15: "EXPRESSION",
	16: "ADJECTIVE",
	17: "INTERJECTION",
	18: "COUNTER",
	19: "PRONOUN",
	20: "CONJUNCTION",
}
var Vocabulary_PartOfSpeech_value = map[string]int32{
	"NOUN":              1,
	"NUMERAL":           2,
	"INTRANSITIVE_VERB": 3,
	"ICHIDAN_VERB":      4,
	"TRANSITIVE_VERB":   5,
	"NO_ADJECTIVE":      6,
	"GODAN_VERB":        7,
	"NA_ADJECTIVE":      8,
	"I_ADJECTIVE":       9,
	"SUFFIX":            10,
	"ADVERB":            11,
	"SURU_VERB":         12,
	"PREFIX":            13,
	"PROPER_NOUN":       14,
	"EXPRESSION":        15,
	"ADJECTIVE":         16,
	"INTERJECTION":      17,
	"COUNTER":           18,
	"PRONOUN":           19,
	"CONJUNCTION":       20,
}

func (x Vocabulary_PartOfSpeech) Enum() *Vocabulary_PartOfSpeech {
	p := new(Vocabulary_PartOfSpeech)
	*p = x
	return p
}
func (x Vocabulary_PartOfSpeech) String() string {
	return proto1.EnumName(Vocabulary_PartOfSpeech_name, int32(x))
}
func (x *Vocabulary_PartOfSpeech) UnmarshalJSON(data []byte) error {
	value, err := proto1.UnmarshalJSONEnum(Vocabulary_PartOfSpeech_value, data, "Vocabulary_PartOfSpeech")
	if err != nil {
		return err
	}
	*x = Vocabulary_PartOfSpeech(value)
	return nil
}
func (Vocabulary_PartOfSpeech) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

type Subject_Type int32

const (
	Subject_RADICAL    Subject_Type = 1
	Subject_KANJI      Subject_Type = 2
	Subject_VOCABULARY Subject_Type = 3
)

var Subject_Type_name = map[int32]string{
	1: "RADICAL",
	2: "KANJI",
	3: "VOCABULARY",
}
var Subject_Type_value = map[string]int32{
	"RADICAL":    1,
	"KANJI":      2,
	"VOCABULARY": 3,
}

func (x Subject_Type) Enum() *Subject_Type {
	p := new(Subject_Type)
	*p = x
	return p
}
func (x Subject_Type) String() string {
	return proto1.EnumName(Subject_Type_name, int32(x))
}
func (x *Subject_Type) UnmarshalJSON(data []byte) error {
	value, err := proto1.UnmarshalJSONEnum(Subject_Type_value, data, "Subject_Type")
	if err != nil {
		return err
	}
	*x = Subject_Type(value)
	return nil
}
func (Subject_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

type FormattedText_Format int32

const (
	FormattedText_RADICAL    FormattedText_Format = 1
	FormattedText_KANJI      FormattedText_Format = 2
	FormattedText_JAPANESE   FormattedText_Format = 3
	FormattedText_READING    FormattedText_Format = 4
	FormattedText_VOCABULARY FormattedText_Format = 5
)

var FormattedText_Format_name = map[int32]string{
	1: "RADICAL",
	2: "KANJI",
	3: "JAPANESE",
	4: "READING",
	5: "VOCABULARY",
}
var FormattedText_Format_value = map[string]int32{
	"RADICAL":    1,
	"KANJI":      2,
	"JAPANESE":   3,
	"READING":    4,
	"VOCABULARY": 5,
}

func (x FormattedText_Format) Enum() *FormattedText_Format {
	p := new(FormattedText_Format)
	*p = x
	return p
}
func (x FormattedText_Format) String() string {
	return proto1.EnumName(FormattedText_Format_name, int32(x))
}
func (x *FormattedText_Format) UnmarshalJSON(data []byte) error {
	value, err := proto1.UnmarshalJSONEnum(FormattedText_Format_value, data, "FormattedText_Format")
	if err != nil {
		return err
	}
	*x = FormattedText_Format(value)
	return nil
}
func (FormattedText_Format) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{11, 0} }

type Meaning struct {
	Meaning          *string `protobuf:"bytes,1,opt,name=meaning" json:"meaning,omitempty"`
	IsPrimary        *bool   `protobuf:"varint,2,opt,name=is_primary,json=isPrimary" json:"is_primary,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Meaning) Reset()                    { *m = Meaning{} }
func (m *Meaning) String() string            { return proto1.CompactTextString(m) }
func (*Meaning) ProtoMessage()               {}
func (*Meaning) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Meaning) GetMeaning() string {
	if m != nil && m.Meaning != nil {
		return *m.Meaning
	}
	return ""
}

func (m *Meaning) GetIsPrimary() bool {
	if m != nil && m.IsPrimary != nil {
		return *m.IsPrimary
	}
	return false
}

type Reading struct {
	Reading          *string       `protobuf:"bytes,1,opt,name=reading" json:"reading,omitempty"`
	IsPrimary        *bool         `protobuf:"varint,2,opt,name=is_primary,json=isPrimary" json:"is_primary,omitempty"`
	Type             *Reading_Type `protobuf:"varint,3,opt,name=type,enum=proto.Reading_Type" json:"type,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Reading) Reset()                    { *m = Reading{} }
func (m *Reading) String() string            { return proto1.CompactTextString(m) }
func (*Reading) ProtoMessage()               {}
func (*Reading) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Reading) GetReading() string {
	if m != nil && m.Reading != nil {
		return *m.Reading
	}
	return ""
}

func (m *Reading) GetIsPrimary() bool {
	if m != nil && m.IsPrimary != nil {
		return *m.IsPrimary
	}
	return false
}

func (m *Reading) GetType() Reading_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Reading_ONYOMI
}

type Radical struct {
	CharacterImage        *string          `protobuf:"bytes,1,opt,name=character_image,json=characterImage" json:"character_image,omitempty"`
	Mnemonic              *string          `protobuf:"bytes,2,opt,name=mnemonic" json:"mnemonic,omitempty"`
	HasCharacterImageFile *bool            `protobuf:"varint,3,opt,name=has_character_image_file,json=hasCharacterImageFile" json:"has_character_image_file,omitempty"`
	FormattedMnemonic     []*FormattedText `protobuf:"bytes,4,rep,name=formatted_mnemonic,json=formattedMnemonic" json:"formatted_mnemonic,omitempty"`
	XXX_unrecognized      []byte           `json:"-"`
}

func (m *Radical) Reset()                    { *m = Radical{} }
func (m *Radical) String() string            { return proto1.CompactTextString(m) }
func (*Radical) ProtoMessage()               {}
func (*Radical) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Radical) GetCharacterImage() string {
	if m != nil && m.CharacterImage != nil {
		return *m.CharacterImage
	}
	return ""
}

func (m *Radical) GetMnemonic() string {
	if m != nil && m.Mnemonic != nil {
		return *m.Mnemonic
	}
	return ""
}

func (m *Radical) GetHasCharacterImageFile() bool {
	if m != nil && m.HasCharacterImageFile != nil {
		return *m.HasCharacterImageFile
	}
	return false
}

func (m *Radical) GetFormattedMnemonic() []*FormattedText {
	if m != nil {
		return m.FormattedMnemonic
	}
	return nil
}

type Kanji struct {
	MeaningMnemonic          *string          `protobuf:"bytes,1,opt,name=meaning_mnemonic,json=meaningMnemonic" json:"meaning_mnemonic,omitempty"`
	MeaningHint              *string          `protobuf:"bytes,2,opt,name=meaning_hint,json=meaningHint" json:"meaning_hint,omitempty"`
	ReadingMnemonic          *string          `protobuf:"bytes,3,opt,name=reading_mnemonic,json=readingMnemonic" json:"reading_mnemonic,omitempty"`
	ReadingHint              *string          `protobuf:"bytes,4,opt,name=reading_hint,json=readingHint" json:"reading_hint,omitempty"`
	FormattedMeaningMnemonic []*FormattedText `protobuf:"bytes,5,rep,name=formatted_meaning_mnemonic,json=formattedMeaningMnemonic" json:"formatted_meaning_mnemonic,omitempty"`
	FormattedMeaningHint     []*FormattedText `protobuf:"bytes,6,rep,name=formatted_meaning_hint,json=formattedMeaningHint" json:"formatted_meaning_hint,omitempty"`
	FormattedReadingMnemonic []*FormattedText `protobuf:"bytes,7,rep,name=formatted_reading_mnemonic,json=formattedReadingMnemonic" json:"formatted_reading_mnemonic,omitempty"`
	FormattedReadingHint     []*FormattedText `protobuf:"bytes,8,rep,name=formatted_reading_hint,json=formattedReadingHint" json:"formatted_reading_hint,omitempty"`
	XXX_unrecognized         []byte           `json:"-"`
}

func (m *Kanji) Reset()                    { *m = Kanji{} }
func (m *Kanji) String() string            { return proto1.CompactTextString(m) }
func (*Kanji) ProtoMessage()               {}
func (*Kanji) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Kanji) GetMeaningMnemonic() string {
	if m != nil && m.MeaningMnemonic != nil {
		return *m.MeaningMnemonic
	}
	return ""
}

func (m *Kanji) GetMeaningHint() string {
	if m != nil && m.MeaningHint != nil {
		return *m.MeaningHint
	}
	return ""
}

func (m *Kanji) GetReadingMnemonic() string {
	if m != nil && m.ReadingMnemonic != nil {
		return *m.ReadingMnemonic
	}
	return ""
}

func (m *Kanji) GetReadingHint() string {
	if m != nil && m.ReadingHint != nil {
		return *m.ReadingHint
	}
	return ""
}

func (m *Kanji) GetFormattedMeaningMnemonic() []*FormattedText {
	if m != nil {
		return m.FormattedMeaningMnemonic
	}
	return nil
}

func (m *Kanji) GetFormattedMeaningHint() []*FormattedText {
	if m != nil {
		return m.FormattedMeaningHint
	}
	return nil
}

func (m *Kanji) GetFormattedReadingMnemonic() []*FormattedText {
	if m != nil {
		return m.FormattedReadingMnemonic
	}
	return nil
}

func (m *Kanji) GetFormattedReadingHint() []*FormattedText {
	if m != nil {
		return m.FormattedReadingHint
	}
	return nil
}

type Vocabulary struct {
	MeaningExplanation          *string                   `protobuf:"bytes,1,opt,name=meaning_explanation,json=meaningExplanation" json:"meaning_explanation,omitempty"`
	ReadingExplanation          *string                   `protobuf:"bytes,2,opt,name=reading_explanation,json=readingExplanation" json:"reading_explanation,omitempty"`
	FormattedMeaningExplanation []*FormattedText          `protobuf:"bytes,6,rep,name=formatted_meaning_explanation,json=formattedMeaningExplanation" json:"formatted_meaning_explanation,omitempty"`
	FormattedReadingExplanation []*FormattedText          `protobuf:"bytes,7,rep,name=formatted_reading_explanation,json=formattedReadingExplanation" json:"formatted_reading_explanation,omitempty"`
	Sentences                   []*Vocabulary_Sentence    `protobuf:"bytes,3,rep,name=sentences" json:"sentences,omitempty"`
	PartsOfSpeech               []Vocabulary_PartOfSpeech `protobuf:"varint,4,rep,name=parts_of_speech,json=partsOfSpeech,enum=proto.Vocabulary_PartOfSpeech" json:"parts_of_speech,omitempty"`
	Audio                       *string                   `protobuf:"bytes,5,opt,name=audio" json:"audio,omitempty"`
	XXX_unrecognized            []byte                    `json:"-"`
}

func (m *Vocabulary) Reset()                    { *m = Vocabulary{} }
func (m *Vocabulary) String() string            { return proto1.CompactTextString(m) }
func (*Vocabulary) ProtoMessage()               {}
func (*Vocabulary) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Vocabulary) GetMeaningExplanation() string {
	if m != nil && m.MeaningExplanation != nil {
		return *m.MeaningExplanation
	}
	return ""
}

func (m *Vocabulary) GetReadingExplanation() string {
	if m != nil && m.ReadingExplanation != nil {
		return *m.ReadingExplanation
	}
	return ""
}

func (m *Vocabulary) GetFormattedMeaningExplanation() []*FormattedText {
	if m != nil {
		return m.FormattedMeaningExplanation
	}
	return nil
}

func (m *Vocabulary) GetFormattedReadingExplanation() []*FormattedText {
	if m != nil {
		return m.FormattedReadingExplanation
	}
	return nil
}

func (m *Vocabulary) GetSentences() []*Vocabulary_Sentence {
	if m != nil {
		return m.Sentences
	}
	return nil
}

func (m *Vocabulary) GetPartsOfSpeech() []Vocabulary_PartOfSpeech {
	if m != nil {
		return m.PartsOfSpeech
	}
	return nil
}

func (m *Vocabulary) GetAudio() string {
	if m != nil && m.Audio != nil {
		return *m.Audio
	}
	return ""
}

type Vocabulary_Sentence struct {
	Japanese         *string `protobuf:"bytes,1,opt,name=japanese" json:"japanese,omitempty"`
	English          *string `protobuf:"bytes,2,opt,name=english" json:"english,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Vocabulary_Sentence) Reset()                    { *m = Vocabulary_Sentence{} }
func (m *Vocabulary_Sentence) String() string            { return proto1.CompactTextString(m) }
func (*Vocabulary_Sentence) ProtoMessage()               {}
func (*Vocabulary_Sentence) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

func (m *Vocabulary_Sentence) GetJapanese() string {
	if m != nil && m.Japanese != nil {
		return *m.Japanese
	}
	return ""
}

func (m *Vocabulary_Sentence) GetEnglish() string {
	if m != nil && m.English != nil {
		return *m.English
	}
	return ""
}

type Subject struct {
	Id                     *int32      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Level                  *int32      `protobuf:"varint,2,opt,name=level" json:"level,omitempty"`
	Slug                   *string     `protobuf:"bytes,3,opt,name=slug" json:"slug,omitempty"`
	DocumentUrl            *string     `protobuf:"bytes,4,opt,name=document_url,json=documentUrl" json:"document_url,omitempty"`
	Japanese               *string     `protobuf:"bytes,5,opt,name=japanese" json:"japanese,omitempty"`
	Readings               []*Reading  `protobuf:"bytes,6,rep,name=readings" json:"readings,omitempty"`
	Meanings               []*Meaning  `protobuf:"bytes,7,rep,name=meanings" json:"meanings,omitempty"`
	ComponentSubjectIds    []int32     `protobuf:"varint,8,rep,name=component_subject_ids,json=componentSubjectIds" json:"component_subject_ids,omitempty"`
	AmalgamationSubjectIds []int32     `protobuf:"varint,12,rep,name=amalgamation_subject_ids,json=amalgamationSubjectIds" json:"amalgamation_subject_ids,omitempty"`
	Radical                *Radical    `protobuf:"bytes,9,opt,name=radical" json:"radical,omitempty"`
	Kanji                  *Kanji      `protobuf:"bytes,10,opt,name=kanji" json:"kanji,omitempty"`
	Vocabulary             *Vocabulary `protobuf:"bytes,11,opt,name=vocabulary" json:"vocabulary,omitempty"`
	XXX_unrecognized       []byte      `json:"-"`
}

func (m *Subject) Reset()                    { *m = Subject{} }
func (m *Subject) String() string            { return proto1.CompactTextString(m) }
func (*Subject) ProtoMessage()               {}
func (*Subject) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Subject) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Subject) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *Subject) GetSlug() string {
	if m != nil && m.Slug != nil {
		return *m.Slug
	}
	return ""
}

func (m *Subject) GetDocumentUrl() string {
	if m != nil && m.DocumentUrl != nil {
		return *m.DocumentUrl
	}
	return ""
}

func (m *Subject) GetJapanese() string {
	if m != nil && m.Japanese != nil {
		return *m.Japanese
	}
	return ""
}

func (m *Subject) GetReadings() []*Reading {
	if m != nil {
		return m.Readings
	}
	return nil
}

func (m *Subject) GetMeanings() []*Meaning {
	if m != nil {
		return m.Meanings
	}
	return nil
}

func (m *Subject) GetComponentSubjectIds() []int32 {
	if m != nil {
		return m.ComponentSubjectIds
	}
	return nil
}

func (m *Subject) GetAmalgamationSubjectIds() []int32 {
	if m != nil {
		return m.AmalgamationSubjectIds
	}
	return nil
}

func (m *Subject) GetRadical() *Radical {
	if m != nil {
		return m.Radical
	}
	return nil
}

func (m *Subject) GetKanji() *Kanji {
	if m != nil {
		return m.Kanji
	}
	return nil
}

func (m *Subject) GetVocabulary() *Vocabulary {
	if m != nil {
		return m.Vocabulary
	}
	return nil
}

type Assignment struct {
	Id               *int32        `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Level            *int32        `protobuf:"varint,2,opt,name=level" json:"level,omitempty"`
	SubjectId        *int32        `protobuf:"varint,3,opt,name=subject_id,json=subjectId" json:"subject_id,omitempty"`
	SubjectType      *Subject_Type `protobuf:"varint,4,opt,name=subject_type,json=subjectType,enum=proto.Subject_Type" json:"subject_type,omitempty"`
	AvailableAt      *int32        `protobuf:"varint,5,opt,name=available_at,json=availableAt" json:"available_at,omitempty"`
	StartedAt        *int32        `protobuf:"varint,6,opt,name=started_at,json=startedAt" json:"started_at,omitempty"`
	SrsStage         *int32        `protobuf:"varint,7,opt,name=srs_stage,json=srsStage" json:"srs_stage,omitempty"`
	PassedAt         *int32        `protobuf:"varint,8,opt,name=passed_at,json=passedAt" json:"passed_at,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Assignment) Reset()                    { *m = Assignment{} }
func (m *Assignment) String() string            { return proto1.CompactTextString(m) }
func (*Assignment) ProtoMessage()               {}
func (*Assignment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Assignment) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Assignment) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *Assignment) GetSubjectId() int32 {
	if m != nil && m.SubjectId != nil {
		return *m.SubjectId
	}
	return 0
}

func (m *Assignment) GetSubjectType() Subject_Type {
	if m != nil && m.SubjectType != nil {
		return *m.SubjectType
	}
	return Subject_RADICAL
}

func (m *Assignment) GetAvailableAt() int32 {
	if m != nil && m.AvailableAt != nil {
		return *m.AvailableAt
	}
	return 0
}

func (m *Assignment) GetStartedAt() int32 {
	if m != nil && m.StartedAt != nil {
		return *m.StartedAt
	}
	return 0
}

func (m *Assignment) GetSrsStage() int32 {
	if m != nil && m.SrsStage != nil {
		return *m.SrsStage
	}
	return 0
}

func (m *Assignment) GetPassedAt() int32 {
	if m != nil && m.PassedAt != nil {
		return *m.PassedAt
	}
	return 0
}

type Progress struct {
	MeaningWrong     *bool       `protobuf:"varint,3,opt,name=meaning_wrong,json=meaningWrong" json:"meaning_wrong,omitempty"`
	ReadingWrong     *bool       `protobuf:"varint,4,opt,name=reading_wrong,json=readingWrong" json:"reading_wrong,omitempty"`
	IsLesson         *bool       `protobuf:"varint,5,opt,name=is_lesson,json=isLesson" json:"is_lesson,omitempty"`
	Assignment       *Assignment `protobuf:"bytes,6,opt,name=assignment" json:"assignment,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Progress) Reset()                    { *m = Progress{} }
func (m *Progress) String() string            { return proto1.CompactTextString(m) }
func (*Progress) ProtoMessage()               {}
func (*Progress) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Progress) GetMeaningWrong() bool {
	if m != nil && m.MeaningWrong != nil {
		return *m.MeaningWrong
	}
	return false
}

func (m *Progress) GetReadingWrong() bool {
	if m != nil && m.ReadingWrong != nil {
		return *m.ReadingWrong
	}
	return false
}

func (m *Progress) GetIsLesson() bool {
	if m != nil && m.IsLesson != nil {
		return *m.IsLesson
	}
	return false
}

func (m *Progress) GetAssignment() *Assignment {
	if m != nil {
		return m.Assignment
	}
	return nil
}

type StudyMaterials struct {
	Id               *int32   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	SubjectId        *int32   `protobuf:"varint,2,opt,name=subject_id,json=subjectId" json:"subject_id,omitempty"`
	SubjectType      *string  `protobuf:"bytes,6,opt,name=subject_type,json=subjectType" json:"subject_type,omitempty"`
	MeaningNote      *string  `protobuf:"bytes,3,opt,name=meaning_note,json=meaningNote" json:"meaning_note,omitempty"`
	ReadingNote      *string  `protobuf:"bytes,4,opt,name=reading_note,json=readingNote" json:"reading_note,omitempty"`
	MeaningSynonyms  []string `protobuf:"bytes,5,rep,name=meaning_synonyms,json=meaningSynonyms" json:"meaning_synonyms,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *StudyMaterials) Reset()                    { *m = StudyMaterials{} }
func (m *StudyMaterials) String() string            { return proto1.CompactTextString(m) }
func (*StudyMaterials) ProtoMessage()               {}
func (*StudyMaterials) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *StudyMaterials) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *StudyMaterials) GetSubjectId() int32 {
	if m != nil && m.SubjectId != nil {
		return *m.SubjectId
	}
	return 0
}

func (m *StudyMaterials) GetSubjectType() string {
	if m != nil && m.SubjectType != nil {
		return *m.SubjectType
	}
	return ""
}

func (m *StudyMaterials) GetMeaningNote() string {
	if m != nil && m.MeaningNote != nil {
		return *m.MeaningNote
	}
	return ""
}

func (m *StudyMaterials) GetReadingNote() string {
	if m != nil && m.ReadingNote != nil {
		return *m.ReadingNote
	}
	return ""
}

func (m *StudyMaterials) GetMeaningSynonyms() []string {
	if m != nil {
		return m.MeaningSynonyms
	}
	return nil
}

type User struct {
	Username                      *string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Level                         *int32  `protobuf:"varint,2,opt,name=level" json:"level,omitempty"`
	MaxLevelGrantedBySubscription *int32  `protobuf:"varint,3,opt,name=max_level_granted_by_subscription,json=maxLevelGrantedBySubscription" json:"max_level_granted_by_subscription,omitempty"`
	ProfileUrl                    *string `protobuf:"bytes,4,opt,name=profile_url,json=profileUrl" json:"profile_url,omitempty"`
	StartedAt                     *int32  `protobuf:"varint,5,opt,name=started_at,json=startedAt" json:"started_at,omitempty"`
	Subscribed                    *bool   `protobuf:"varint,6,opt,name=subscribed" json:"subscribed,omitempty"`
	XXX_unrecognized              []byte  `json:"-"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto1.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *User) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *User) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *User) GetMaxLevelGrantedBySubscription() int32 {
	if m != nil && m.MaxLevelGrantedBySubscription != nil {
		return *m.MaxLevelGrantedBySubscription
	}
	return 0
}

func (m *User) GetProfileUrl() string {
	if m != nil && m.ProfileUrl != nil {
		return *m.ProfileUrl
	}
	return ""
}

func (m *User) GetStartedAt() int32 {
	if m != nil && m.StartedAt != nil {
		return *m.StartedAt
	}
	return 0
}

func (m *User) GetSubscribed() bool {
	if m != nil && m.Subscribed != nil {
		return *m.Subscribed
	}
	return false
}

type SubjectOverrides struct {
	Subject          []*Subject `protobuf:"bytes,1,rep,name=subject" json:"subject,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *SubjectOverrides) Reset()                    { *m = SubjectOverrides{} }
func (m *SubjectOverrides) String() string            { return proto1.CompactTextString(m) }
func (*SubjectOverrides) ProtoMessage()               {}
func (*SubjectOverrides) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *SubjectOverrides) GetSubject() []*Subject {
	if m != nil {
		return m.Subject
	}
	return nil
}

type FormattedText struct {
	Format           []FormattedText_Format `protobuf:"varint,1,rep,name=format,enum=proto.FormattedText_Format" json:"format,omitempty"`
	Text             *string                `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *FormattedText) Reset()                    { *m = FormattedText{} }
func (m *FormattedText) String() string            { return proto1.CompactTextString(m) }
func (*FormattedText) ProtoMessage()               {}
func (*FormattedText) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *FormattedText) GetFormat() []FormattedText_Format {
	if m != nil {
		return m.Format
	}
	return nil
}

func (m *FormattedText) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

func init() {
	proto1.RegisterType((*Meaning)(nil), "proto.Meaning")
	proto1.RegisterType((*Reading)(nil), "proto.Reading")
	proto1.RegisterType((*Radical)(nil), "proto.Radical")
	proto1.RegisterType((*Kanji)(nil), "proto.Kanji")
	proto1.RegisterType((*Vocabulary)(nil), "proto.Vocabulary")
	proto1.RegisterType((*Vocabulary_Sentence)(nil), "proto.Vocabulary.Sentence")
	proto1.RegisterType((*Subject)(nil), "proto.Subject")
	proto1.RegisterType((*Assignment)(nil), "proto.Assignment")
	proto1.RegisterType((*Progress)(nil), "proto.Progress")
	proto1.RegisterType((*StudyMaterials)(nil), "proto.StudyMaterials")
	proto1.RegisterType((*User)(nil), "proto.User")
	proto1.RegisterType((*SubjectOverrides)(nil), "proto.SubjectOverrides")
	proto1.RegisterType((*FormattedText)(nil), "proto.FormattedText")
	proto1.RegisterEnum("proto.Reading_Type", Reading_Type_name, Reading_Type_value)
	proto1.RegisterEnum("proto.Vocabulary_PartOfSpeech", Vocabulary_PartOfSpeech_name, Vocabulary_PartOfSpeech_value)
	proto1.RegisterEnum("proto.Subject_Type", Subject_Type_name, Subject_Type_value)
	proto1.RegisterEnum("proto.FormattedText_Format", FormattedText_Format_name, FormattedText_Format_value)
}

func init() { proto1.RegisterFile("wanikani.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0xcb, 0x6e, 0xdb, 0x46,
	0x17, 0x86, 0x2e, 0xb4, 0xa8, 0x23, 0x59, 0xa6, 0xc7, 0x4e, 0x20, 0xd8, 0x70, 0x7e, 0x9b, 0xff,
	0x22, 0xfe, 0xff, 0x02, 0x2e, 0xea, 0x02, 0x6d, 0x16, 0x5d, 0x94, 0x96, 0xe5, 0x84, 0xb2, 0x4d,
	0x0a, 0x23, 0xcb, 0x49, 0x56, 0xc4, 0x58, 0x1a, 0xcb, 0x4c, 0x78, 0x11, 0x38, 0x94, 0x63, 0x3d,
	0x40, 0xdf, 0xa1, 0xeb, 0x02, 0xdd, 0xf7, 0x2d, 0xda, 0x45, 0x77, 0x7d, 0x83, 0xf6, 0x29, 0xba,
	0x2b, 0xe6, 0x42, 0x89, 0x92, 0x92, 0xd4, 0x2b, 0xe9, 0x9c, 0xf3, 0x9d, 0x8f, 0xe7, 0x32, 0x73,
	0xe6, 0x40, 0xe3, 0x03, 0x89, 0xfc, 0xf7, 0x24, 0xf2, 0x8f, 0xc6, 0x49, 0x9c, 0xc6, 0x48, 0x13,
	0x3f, 0xe6, 0x09, 0x54, 0x2e, 0x29, 0x89, 0xfc, 0x68, 0x84, 0x9a, 0x50, 0x09, 0xe5, 0xdf, 0x66,
	0x61, 0xbf, 0x70, 0x58, 0xc5, 0x99, 0x88, 0xf6, 0x00, 0x7c, 0xe6, 0x8d, 0x13, 0x3f, 0x24, 0xc9,
	0xb4, 0x59, 0xdc, 0x2f, 0x1c, 0xea, 0xb8, 0xea, 0xb3, 0xae, 0x54, 0x98, 0x3f, 0x16, 0xa0, 0x82,
	0x29, 0x19, 0x2a, 0x92, 0x44, 0xfe, 0xcd, 0x48, 0x94, 0xf8, 0x2f, 0x24, 0xe8, 0x39, 0x94, 0xd3,
	0xe9, 0x98, 0x36, 0x4b, 0xfb, 0x85, 0xc3, 0xc6, 0xf1, 0x96, 0x8c, 0xf2, 0x48, 0xd1, 0x1e, 0x5d,
	0x4d, 0xc7, 0x14, 0x0b, 0x80, 0xf9, 0x05, 0x94, 0xb9, 0x84, 0x00, 0xd6, 0x5c, 0xe7, 0xad, 0x7b,
	0x69, 0x1b, 0x05, 0x54, 0x83, 0xca, 0x79, 0x5f, 0x0a, 0x45, 0x6e, 0x70, 0x2c, 0xc7, 0xc5, 0xb6,
	0x51, 0x32, 0x7f, 0xe7, 0xa1, 0x91, 0xa1, 0x3f, 0x20, 0x01, 0x7a, 0x0e, 0x1b, 0x83, 0x3b, 0x92,
	0x90, 0x41, 0x4a, 0x13, 0xcf, 0x0f, 0xc9, 0x88, 0xaa, 0x10, 0x1b, 0x33, 0xb5, 0xcd, 0xb5, 0x68,
	0x07, 0xf4, 0x30, 0xa2, 0x61, 0x1c, 0xf9, 0x03, 0x11, 0x67, 0x15, 0xcf, 0x64, 0xf4, 0x2d, 0x34,
	0xef, 0x08, 0xf3, 0x96, 0x88, 0xbc, 0x5b, 0x3f, 0x90, 0xa1, 0xeb, 0xf8, 0xc9, 0x1d, 0x61, 0xad,
	0x05, 0xc2, 0x33, 0x3f, 0xa0, 0xa8, 0x05, 0xe8, 0x36, 0x4e, 0x42, 0x92, 0xa6, 0x74, 0xe8, 0xcd,
	0xe8, 0xcb, 0xfb, 0xa5, 0xc3, 0xda, 0xf1, 0xb6, 0xca, 0xf6, 0x2c, 0x03, 0x5c, 0xd1, 0x87, 0x14,
	0x6f, 0xce, 0xf0, 0x97, 0x0a, 0x6e, 0xfe, 0x55, 0x02, 0xed, 0x9c, 0x44, 0xef, 0x7c, 0xf4, 0x3f,
	0x30, 0x54, 0x77, 0xe6, 0x64, 0x32, 0x9b, 0x0d, 0xa5, 0xcf, 0x9c, 0xd0, 0x01, 0xd4, 0x33, 0xe8,
	0x9d, 0x1f, 0xa5, 0x2a, 0xa5, 0x9a, 0xd2, 0xbd, 0xf2, 0xa3, 0x94, 0xb3, 0xa9, 0x36, 0xcd, 0xd9,
	0x4a, 0x92, 0x4d, 0xe9, 0xf3, 0x6c, 0x19, 0x54, 0xb0, 0x95, 0x25, 0x9b, 0xd2, 0x09, 0x36, 0x0c,
	0x3b, 0xb9, 0x54, 0x97, 0xa3, 0xd4, 0x3e, 0x93, 0x72, 0x73, 0x9e, 0xf2, 0x52, 0x12, 0x1d, 0x78,
	0xba, 0xca, 0x29, 0x02, 0x58, 0xfb, 0x0c, 0xdf, 0xf6, 0x32, 0xdf, 0x6a, 0x7c, 0x2b, 0x79, 0x57,
	0x1e, 0x15, 0x1f, 0x5e, 0x2a, 0xcb, 0x42, 0x7c, 0x0b, 0x05, 0xd2, 0x1f, 0x15, 0x1f, 0x9e, 0xd7,
	0xcf, 0xfc, 0x75, 0x0d, 0xe0, 0x3a, 0x1e, 0x90, 0x9b, 0x49, 0xc0, 0x6f, 0xc6, 0x97, 0xb0, 0x95,
	0x25, 0x4c, 0x1f, 0xc6, 0x01, 0x89, 0x48, 0xea, 0xc7, 0x91, 0xea, 0x36, 0x52, 0xa6, 0xf6, 0xdc,
	0xc2, 0x1d, 0xb2, 0x08, 0xf2, 0x0e, 0xb2, 0xef, 0x48, 0x99, 0xf2, 0x0e, 0x6f, 0x60, 0x6f, 0xb5,
	0xb8, 0x79, 0xd7, 0xcf, 0xd5, 0x78, 0x77, 0xb9, 0xc6, 0x9f, 0x64, 0xfe, 0x58, 0x50, 0x95, 0x47,
	0x31, 0xe3, 0xd5, 0x98, 0x5f, 0x40, 0x95, 0xd1, 0x28, 0xa5, 0xd1, 0x80, 0xb2, 0x66, 0x49, 0xb0,
	0xec, 0x28, 0x96, 0x79, 0xed, 0x8e, 0x7a, 0x0a, 0x82, 0xe7, 0x60, 0x74, 0x06, 0x1b, 0x63, 0x92,
	0xa4, 0xcc, 0x8b, 0x6f, 0x3d, 0x36, 0xa6, 0x74, 0x70, 0x27, 0xae, 0x61, 0xe3, 0xf8, 0xd9, 0xaa,
	0x7f, 0x97, 0x24, 0xa9, 0x7b, 0xdb, 0x13, 0x28, 0xbc, 0x2e, 0xdc, 0x32, 0x11, 0x6d, 0x83, 0x46,
	0x26, 0x43, 0x3f, 0x6e, 0x6a, 0xa2, 0xb0, 0x52, 0xd8, 0xf9, 0x1e, 0xf4, 0xec, 0xa3, 0x7c, 0x90,
	0xbc, 0x23, 0x63, 0x12, 0x51, 0x96, 0x8d, 0x9a, 0x99, 0xcc, 0x07, 0x25, 0x8d, 0x46, 0x81, 0xcf,
	0xee, 0x54, 0x63, 0x32, 0xd1, 0xfc, 0xad, 0x08, 0xf5, 0xfc, 0x77, 0x91, 0x0e, 0x65, 0xc7, 0xed,
	0x3b, 0x72, 0xce, 0x39, 0xfd, 0xcb, 0x36, 0xb6, 0x2e, 0x8c, 0x22, 0x7a, 0x02, 0x9b, 0xb6, 0x73,
	0x85, 0x2d, 0xa7, 0x67, 0x5f, 0xd9, 0xd7, 0x6d, 0xef, 0xba, 0x8d, 0x4f, 0x8c, 0x12, 0x32, 0xa0,
	0x6e, 0xb7, 0x5e, 0xd9, 0xa7, 0x96, 0x23, 0x35, 0x65, 0xb4, 0x05, 0x1b, 0xcb, 0x30, 0x8d, 0xc3,
	0x1c, 0xd7, 0xb3, 0x4e, 0x3b, 0xed, 0x16, 0x57, 0x1b, 0x6b, 0xa8, 0x01, 0xf0, 0xd2, 0x9d, 0xb9,
	0x55, 0x04, 0xc2, 0xca, 0x21, 0x74, 0xb4, 0x01, 0x35, 0x3b, 0xa7, 0xa8, 0xf2, 0x51, 0xdb, 0xeb,
	0x9f, 0x9d, 0xd9, 0x6f, 0x0c, 0xe0, 0xff, 0xad, 0x53, 0xe1, 0x5a, 0x43, 0xeb, 0x50, 0xed, 0xf5,
	0x71, 0x5f, 0x32, 0xd5, 0xb9, 0xa9, 0x8b, 0xdb, 0x1c, 0xb6, 0xce, 0x39, 0xba, 0xd8, 0xed, 0xb6,
	0xb1, 0x27, 0x72, 0x6a, 0xf0, 0xcf, 0xb6, 0xdf, 0x74, 0x71, 0xbb, 0xd7, 0xb3, 0x5d, 0xc7, 0xd8,
	0xe0, 0xbe, 0xf3, 0x4f, 0x18, 0x22, 0x1d, 0xe7, 0xaa, 0x8d, 0x85, 0xc6, 0x75, 0x8c, 0x4d, 0x5e,
	0x84, 0x96, 0xdb, 0xe7, 0x3a, 0x03, 0x71, 0xa1, 0x8b, 0x5d, 0x41, 0xb5, 0xc5, 0xb9, 0x5b, 0xae,
	0xd3, 0xe9, 0x3b, 0x12, 0xba, 0x6d, 0xfe, 0x5d, 0x82, 0x4a, 0x6f, 0x72, 0xf3, 0x8e, 0x0e, 0x52,
	0xd4, 0x80, 0xa2, 0x3f, 0x14, 0x6d, 0xd0, 0x70, 0xd1, 0x1f, 0xf2, 0xf6, 0x05, 0xf4, 0x9e, 0x06,
	0xa2, 0xfc, 0x1a, 0x96, 0x02, 0x42, 0x50, 0x66, 0xc1, 0x64, 0xa4, 0xa6, 0x9f, 0xf8, 0xcf, 0x47,
	0xde, 0x30, 0x1e, 0x4c, 0x42, 0x1a, 0xa5, 0xde, 0x24, 0x09, 0xb2, 0x91, 0x97, 0xe9, 0xfa, 0x49,
	0xb0, 0xd0, 0x69, 0x6d, 0xa9, 0xd3, 0xff, 0x07, 0x5d, 0x9d, 0x7c, 0xa6, 0x2e, 0x52, 0x63, 0xf1,
	0x75, 0xc3, 0x33, 0x3b, 0xc7, 0xaa, 0xfb, 0xc7, 0xd4, 0xd5, 0xc8, 0xb0, 0xea, 0x72, 0xe1, 0x99,
	0x1d, 0x1d, 0xc3, 0x93, 0x41, 0x1c, 0x8e, 0xe3, 0x88, 0xc7, 0xc5, 0x64, 0x96, 0x9e, 0x3f, 0x64,
	0x62, 0xe2, 0x68, 0x78, 0x6b, 0x66, 0x54, 0x15, 0xb0, 0x87, 0x0c, 0xbd, 0x80, 0x26, 0x09, 0x49,
	0x30, 0x22, 0xa1, 0xb8, 0x45, 0x0b, 0x6e, 0x75, 0xe1, 0xf6, 0x34, 0x6f, 0xcf, 0x79, 0x1e, 0x42,
	0x25, 0x91, 0x0f, 0x69, 0xb3, 0xba, 0x5f, 0xc8, 0x27, 0x21, 0xb5, 0x38, 0x33, 0x23, 0x13, 0xb4,
	0xf7, 0xfc, 0x8d, 0x6a, 0x82, 0xc0, 0xd5, 0x15, 0x4e, 0xbc, 0x5b, 0x58, 0x9a, 0xd0, 0x57, 0x00,
	0xf7, 0xb3, 0x5b, 0xd6, 0xac, 0x09, 0xe0, 0xe6, 0xca, 0xf5, 0xc3, 0x39, 0x90, 0x79, 0xa4, 0xde,
	0xfd, 0x1a, 0x54, 0xb0, 0x75, 0x6a, 0xb7, 0xac, 0x0b, 0xa3, 0x80, 0xaa, 0xa0, 0x9d, 0x5b, 0x4e,
	0x87, 0x3f, 0xfb, 0x0d, 0x80, 0x6b, 0xb7, 0x65, 0x9d, 0xf4, 0x2f, 0x2c, 0xfc, 0xd6, 0x28, 0x99,
	0x3f, 0x14, 0x01, 0x2c, 0xc6, 0xfc, 0x51, 0xc4, 0x9b, 0xf4, 0xc8, 0xf6, 0xef, 0x01, 0xcc, 0x4b,
	0x22, 0x0e, 0x81, 0x86, 0xab, 0x2c, 0xab, 0x02, 0xfa, 0x06, 0xea, 0x99, 0x59, 0x2c, 0x2b, 0xe5,
	0x85, 0x65, 0x45, 0x55, 0x4b, 0x2e, 0x2b, 0x35, 0x05, 0x14, 0x31, 0x1f, 0x40, 0x9d, 0xdc, 0x13,
	0x3f, 0x20, 0x37, 0x01, 0xf5, 0x48, 0x2a, 0x8e, 0x88, 0x86, 0x6b, 0x33, 0x9d, 0x95, 0x8a, 0x2f,
	0xa7, 0x24, 0xe1, 0x73, 0x92, 0xf0, 0x47, 0x4d, 0x7e, 0x59, 0x6a, 0xac, 0x14, 0xed, 0x42, 0x95,
	0x25, 0xcc, 0x63, 0x29, 0x5f, 0x5b, 0x2a, 0xc2, 0xaa, 0xb3, 0x84, 0xf5, 0xb8, 0xcc, 0x8d, 0x63,
	0xc2, 0x98, 0x74, 0xd5, 0xa5, 0x51, 0x2a, 0xac, 0xd4, 0xfc, 0xa5, 0x00, 0x7a, 0x37, 0x89, 0x47,
	0x09, 0x65, 0x0c, 0xfd, 0x17, 0xd6, 0xb3, 0xf9, 0xfe, 0x21, 0x89, 0xa3, 0x91, 0xda, 0x59, 0xb2,
	0x05, 0xe1, 0x35, 0xd7, 0x71, 0x50, 0x36, 0xaa, 0x25, 0xa8, 0x2c, 0x41, 0x4a, 0x29, 0x41, 0xbb,
	0x50, 0xf5, 0x99, 0x17, 0x50, 0xc6, 0xe2, 0x48, 0xe4, 0xa3, 0x63, 0xdd, 0x67, 0x17, 0x42, 0xe6,
	0xed, 0x25, 0xb3, 0xd2, 0x8b, 0x64, 0xe6, 0xed, 0x9d, 0xf7, 0x04, 0xe7, 0x40, 0x9d, 0xb2, 0x5e,
	0x30, 0x8a, 0x9d, 0xb2, 0x5e, 0x34, 0x4a, 0xe6, 0x1f, 0x05, 0x68, 0xf4, 0xd2, 0xc9, 0x70, 0x7a,
	0x49, 0x52, 0x9a, 0xf8, 0x24, 0x60, 0x2b, 0xed, 0x5b, 0x6c, 0x54, 0x71, 0xb9, 0x51, 0x07, 0x4b,
	0x8d, 0x5a, 0x93, 0x57, 0x76, 0xa9, 0x27, 0x59, 0x29, 0xa2, 0x38, 0xa5, 0xea, 0xc6, 0x67, 0x6b,
	0x91, 0x13, 0xa7, 0x34, 0xbf, 0xeb, 0x08, 0xc8, 0xe2, 0xae, 0x23, 0x20, 0xb9, 0x3d, 0x8c, 0x4d,
	0xa3, 0x38, 0x9a, 0x86, 0x4c, 0x6c, 0x38, 0xf3, 0x3d, 0xac, 0xa7, 0xd4, 0xe6, 0x9f, 0x05, 0x28,
	0xf7, 0x19, 0x4d, 0xf8, 0xb0, 0x98, 0x30, 0x9a, 0x44, 0x24, 0x9c, 0x3d, 0x0b, 0x99, 0xfc, 0x89,
	0x63, 0xf9, 0x0a, 0x0e, 0x42, 0xf2, 0xe0, 0x09, 0xc1, 0x1b, 0x25, 0x24, 0xe2, 0xc7, 0xe4, 0x66,
	0xca, 0xaf, 0x2f, 0x1b, 0x24, 0xfe, 0x58, 0x3c, 0xa5, 0xf2, 0xb4, 0xee, 0x85, 0xe4, 0xe1, 0x82,
	0xe3, 0x5e, 0x4a, 0xd8, 0xc9, 0xb4, 0x97, 0x03, 0xa1, 0xff, 0x40, 0x6d, 0x9c, 0xc4, 0x7c, 0x5d,
	0xcd, 0x8d, 0x32, 0x50, 0x2a, 0x3e, 0xc9, 0x16, 0xcf, 0xa1, 0xb6, 0x7c, 0x0e, 0x9f, 0x89, 0xba,
	0x73, 0xbe, 0x1b, 0x3a, 0x14, 0x65, 0xd5, 0x71, 0x4e, 0x63, 0x7e, 0x07, 0x86, 0xba, 0x06, 0xee,
	0x3d, 0x4d, 0x12, 0x7f, 0x48, 0xc5, 0xe8, 0x50, 0x85, 0x6f, 0x16, 0x16, 0x66, 0x9a, 0x42, 0xe2,
	0xcc, 0x6c, 0xfe, 0x5c, 0x80, 0xf5, 0x85, 0x1d, 0x00, 0x7d, 0x0d, 0x6b, 0x72, 0x0b, 0x10, 0xae,
	0x8d, 0xe3, 0xdd, 0x8f, 0x6d, 0x0a, 0x4a, 0xc2, 0x0a, 0xca, 0x87, 0x78, 0x4a, 0x1f, 0xb2, 0x4d,
	0x57, 0xfc, 0x37, 0xcf, 0x61, 0x4d, 0xa2, 0x3e, 0x39, 0x40, 0xea, 0xa0, 0x77, 0xac, 0xae, 0xe5,
	0xb4, 0x7b, 0x6d, 0xa3, 0x24, 0x50, 0x6d, 0xeb, 0xd4, 0x76, 0x5e, 0x1a, 0xe5, 0xa5, 0xd9, 0xa2,
	0x9d, 0x68, 0x3f, 0x15, 0x8b, 0xaf, 0xcf, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xf5, 0xcd, 0xfc,
	0xcd, 0x54, 0x0d, 0x00, 0x00,
}
