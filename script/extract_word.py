import pandas as pd
path="data/流行语分类终.xlsx"
df=pd.read_excel(path)
words=df['流行语名称']
words.to_excel("data/word.xlsx")