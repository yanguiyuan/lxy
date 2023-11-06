# https://baike.baidu.com/item/666/17192634?fromModule=lemma_search-box
import requests
from bs4 import BeautifulSoup
url='https://baike.baidu.com/item/666/17192634?fromModule=lemma_search-box'
# 发送GET请求
response = requests.get(url)
print(response.content)
# 解析HTML文档
soup = BeautifulSoup(response.content, 'html.parser')
print(soup)
# 获取头部信息
meta_tags = soup.head.find_all('meta', attrs={'property': "og:description"})
print(meta_tags)
# 打印meta标签的内容
for meta_tag in meta_tags:
    # content = meta_tag.get('content')
    print(meta_tag['content'])
