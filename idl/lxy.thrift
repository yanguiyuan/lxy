namespace go lxyhttp
struct Tag{
    1:i64 id
    2:string value
    3:i64 tag_type_id
}
struct Buzzword{
    1:i64 ID
    2:string Word
    3:list<Tag> Tags
    4:string Paraphrase
}
struct TagType{
    1:i64 ID
    2:string Value
}
struct QueryBuzzwordsResponse{
    1:i16 Code
    2:string Messgage
    3:list<Buzzword> Data
}

struct AddBuzzwordRequest{
    1:string Word (api.body="word")
    2:string Paraphrase (api.body="paraphrase")
    3:list<Tag> Tags (api.body="tags")
}
struct CommonResponse{
    1:i16 Code
    2:string Message
}
struct RemoveBuzzwordRequest{
    1:i64 BuzzwordID (api.path="id")
}
struct UpdateBuzzwordParaphraseRequest{
    1:i64 BuzzwordID (api.body="id")
    2:string Value (api.body="value")
}
struct AddTagRequest{
    1:string Value (api.body="value")
    2:i64 TagTypeID (api.body="tag_type_id")
}
struct AddTagTypeRequest{
    1:string Value (api.body="value")
}
struct AddOrRemoveBuzzwordTagRequest{
    1:i64 BuzzwordID (api.body="buzzword_id")
    2:i64 TagID (api.body="tag_id")
}

struct OneIDRequest{
    1:i64 ID (api.path="id")
}

struct TwoIDRequest{
    1:i64 IDOne (api.path="id_one")
    2:i64 IDTwo (api.path="id_two")
}
struct ExportExcelRequest{
    1:list<i64> BuzzwordIDList (api.body="id_list")
}
struct ExportExcelResponse{
    1:i16 Code
    2:string Message
    3:binary Data
}
struct ImportExcelRequest{
    1:binary Data (api.form="file")
    2:string BuzzwordColumnName (api.form="buzzword_column_name")
    3:string ParaphraseColumnName (api.form="paraphrase_column_name")
}


struct KeywordQueryRequest{
    1:string keyword (api.query="keyword")
}
struct QueryTagsResponse{
    1:i16 Code
    2:string Message
    3:list<Tag> Data
}
struct QueryTagTypesResponse{
    1:i16 Code
    2:string Message
    3:list<TagType> Data
}
struct LoginOrRegisterRequest{
    1:string Username (api.body="username" api.vd="len($)<20 && len($)>0")
    2:string Password (api.body="password" api.vd="len($)>=6 && len($)<20")
}

service LXYService{
    QueryBuzzwordsResponse QueryAllBuzzwords()(api.get="/buzzwords")
    CommonResponse AddBuzzword(1:AddBuzzwordRequest req)(api.post="/buzzword")
    CommonResponse RemoveBuzzword(1:RemoveBuzzwordRequest req)(api.delete="/buzzword/:id")
    CommonResponse UpdateBuzzwordParaphrase(1:UpdateBuzzwordParaphraseRequest req)(api.put="/buzzword/paraphrase")
    CommonResponse AddTag(1:AddTagRequest req)(api.post="/tag")
    CommonResponse AddTagType(1:AddTagTypeRequest req)(api.post="/tag/type")
    CommonResponse AddBuzzwordTag(1:AddOrRemoveBuzzwordTagRequest req)(api.post="/buzzword/tag")
    CommonResponse RemoveBuzzwordTag(1:AddOrRemoveBuzzwordTagRequest req)(api.delete="/buzzword/tag")
    CommonResponse RemoveTag(1:OneIDRequest req)(api.delete="/tag/:id")
    CommonResponse RemoveTagType(1:OneIDRequest req)(api.delete="/tag/type/:id")
    CommonResponse UpdateTagOfType(1:TwoIDRequest req)(api.put="/tag/:id_one/type/:id_two")
    ExportExcelResponse ExportExcel(1:ExportExcelRequest req)(api.patch="/excel")
    CommonResponse ImportExcel(1:ImportExcelRequest req)(api.post="/excel")
    QueryBuzzwordsResponse QueryBuzzwordsByTag(1:OneIDRequest req)(api.get="/buzzwords/tag/:id")
    QueryBuzzwordsResponse QueryBuzzwordsByKeyword(1:KeywordQueryRequest req )(api.get="/buzzwords/keyword")
    QueryTagsResponse QueryAllTags()(api.get="/tags")
    QueryTagTypesResponse QueryAllTagTypes()(api.get="/tag/types")
    QueryTagsResponse QueryTagsByType(1:OneIDRequest req)(api.get="/tags/:id")
    QueryTagsResponse QueryTagsBykeyword(1:KeywordQueryRequest req)(api.get="/tags/keyword")
    CommonResponse UpdateBaiduHot()(api.patch="/baiduhot")
}

service UserService{
      CommonResponse Login(1:LoginOrRegisterRequest req)(api.post="/login")
      CommonResponse Logout()(api.post="/logout")
      CommonResponse Register(1:LoginOrRegisterRequest req)(api.post="/register")
}