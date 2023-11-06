package main

import (
	"github.com/yanguiyuan/yuan/pkg/config"
	"gorm.io/gen/field"

	//"github.com/yanguiyuan/yuan/pkg/code-gen/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	c := config.NewConfig()
	db, err := gorm.Open(mysql.Open(c.GetString("mysql.dsn")), &gorm.Config{})
	if err != nil {
		return
	}
	g := gen.NewGenerator(gen.Config{
		OutPath:      c.GetString("gorm-gen.dal"),
		ModelPkgPath: c.GetString("gorm-gen.model"),
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	g.UseDB(db)
	tagType := g.GenerateModel("tag_types")
	tag := g.GenerateModel("tags", gen.FieldRelate(field.HasOne, "TagType", tagType, &field.RelateConfig{
		GORMTag: map[string][]string{
			"foreignKey": []string{"tag_type_id"},
		},
	}))
	word := g.GenerateModel("words", gen.FieldRelate(field.Many2Many, "Tags", tag, &field.RelateConfig{
		GORMTag: map[string][]string{
			"many2many": []string{"word_tags"},
		},
	}))
	wordTag := g.GenerateModel("word_tags")
	g.ApplyBasic(tagType, tag, word, wordTag)
	g.Execute()
}
