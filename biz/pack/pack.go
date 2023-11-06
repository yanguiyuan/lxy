package pack

import (
	"github.com/yanguiyuan/lxy/biz/model/dbmodel"
	"github.com/yanguiyuan/lxy/biz/model/lxyhttp"
)

func FromTagList(tags []*lxyhttp.Tag) []dbmodel.Tag {
	newTags := make([]dbmodel.Tag, 0, len(tags))
	for _, tag := range tags {
		newTags = append(newTags, dbmodel.Tag{
			ID:        tag.ID,
			TagTypeID: tag.TagTypeID,
			Value:     tag.Value,
		})
	}
	return newTags
}
func ToTagListOrigin(tags []dbmodel.Tag) []*lxyhttp.Tag {
	result := make([]*lxyhttp.Tag, 0, len(tags))
	for _, tag := range tags {
		result = append(result, &lxyhttp.Tag{
			ID:        tag.ID,
			Value:     tag.Value,
			TagTypeID: tag.TagTypeID,
		})
	}
	return result
}
func ToTagList(tags []*dbmodel.Tag) []*lxyhttp.Tag {
	result := make([]*lxyhttp.Tag, 0, len(tags))
	for _, tag := range tags {
		result = append(result, &lxyhttp.Tag{
			ID:        tag.ID,
			Value:     tag.Value,
			TagTypeID: tag.TagTypeID,
		})
	}
	return result
}
func ToWordList(words []*dbmodel.Word) []*lxyhttp.Buzzword {
	result := make([]*lxyhttp.Buzzword, 0, len(words))
	for _, word := range words {
		result = append(result, &lxyhttp.Buzzword{
			ID:         word.ID,
			Word:       word.Value,
			Paraphrase: word.Paraphrase,
			Tags:       ToTagListOrigin(word.Tags),
		})
	}
	return result
}
func ToTagTypeList(tagTypes []*dbmodel.TagType) []*lxyhttp.TagType {
	result := make([]*lxyhttp.TagType, 0, len(tagTypes))
	for _, tagType := range tagTypes {
		result = append(result, &lxyhttp.TagType{
			ID:    tagType.ID,
			Value: tagType.Value,
		})
	}
	return result
}
