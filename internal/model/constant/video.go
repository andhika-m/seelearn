package constant

type VideoCategory string

const (
	VideoCategoryMatematika VideoCategory = "matematika"
	VideoCategoryIPA        VideoCategory = "ipa"
	VideoCategoryIPS        VideoCategory = "ips"
	VideoCategoryPPKN       VideoCategory = "ppkn"
	VideoCategorySejarah    VideoCategory = "sejarah"
)

type ClassCategory string

const (
	ClassCategorySatu  ClassCategory = "kelas-1"
	ClassCategoryDua   ClassCategory = "kelas-2"
	ClassCategoryTiga  ClassCategory = "kelas-3"
	ClassCategoryEmpat ClassCategory = "kelas-4"
	ClassCategoryLima  ClassCategory = "kelas-5"
	ClassCategoryEnam  ClassCategory = "kelas-6"
)
