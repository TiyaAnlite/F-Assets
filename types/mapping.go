package types

import (
	"github.com/TiyaAnlite/F-Assests/pb"
	"github.com/duke-git/lancet/v2/maputil"
)

var Languages2OrmMapping = map[pb.Languages]LanguageType{
	pb.Languages_UnknownLanguage:    UnknownLanguage,
	pb.Languages_SimplifiedChinese:  SimplifiedChineseLanguage,
	pb.Languages_TraditionalChinese: TraditionalChineseLanguage,
	pb.Languages_Japanese:           JapaneseLanguage,
}

var Orm2LanguagesMapping = map[LanguageType]pb.Languages{
	UnknownLanguage:            pb.Languages_UnknownLanguage,
	SimplifiedChineseLanguage:  pb.Languages_SimplifiedChinese,
	TraditionalChineseLanguage: pb.Languages_TraditionalChinese,
	JapaneseLanguage:           pb.Languages_Japanese,
}

func Languages2Orm(lang pb.Languages) LanguageType {
	return maputil.GetOrDefault(Languages2OrmMapping, lang, UnknownLanguage)
}

func Orm2Languages(lang LanguageType) pb.Languages {
	return maputil.GetOrDefault(Orm2LanguagesMapping, lang, pb.Languages_UnknownLanguage)
}

func PriceUnit2Orm(priceUnit pb.PriceUnit) string {
	return maputil.GetOrDefault(pb.PriceUnit_name, int32(priceUnit), "UnknownPriceUnit")
}

func Orm2PriceUnit(priceUnit string) pb.PriceUnit {
	return pb.PriceUnit(maputil.GetOrDefault(pb.PriceUnit_value, priceUnit, 0))
}
