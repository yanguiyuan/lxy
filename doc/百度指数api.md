### 百度指数API



##### 资讯指数

```
https://index.baidu.com/api/FeedSearchApi/getFeedIndex?word=[[%7B%22name%22:%22%E5%8D%A1%E5%A1%94%E5%B0%94%22,%22wordType%22:1%7D]]&area=0&days=30
```

##### 搜索指数

```

https://index.baidu.com/api/SearchApi/index?area=0&word=[[%7B%22name%22:%22%E5%8D%A1%E5%A1%94%E5%B0%94%22,%22wordType%22:1%7D]]&days=30
```

响应：

```
{
    "status": 0,
    "data": {
        "userIndexes": [
            {
                "word": [
                    {
                        "name": "\u5361\u5854\u5c14",
                        "wordType": 1
                    }
                ],
                "all": {
                    "startDate": "2023-06-30",
                    "endDate": "2023-07-29",
                    "data": "iA2gfilBifil2lfi1l1fi%1%fiA2qfi2%ifi%%1fil1gfil12fi%gifi%llfiABifi%2Afi%l2fiBBLfi2qlfiA%LfiAALfi1lAfiAg2fi%22fi%iifilBLfiAL%fiA1%fi%LAfiALAfB%%lfiA%i"
                },
                "pc": {
                    "startDate": "2023-06-30",
                    "endDate": "2023-07-29",
                    "data": "2A1fBlifBigf%2if%2if%LBf22Bf2lAfBABfBA%f%BAf2lAf%i%f2l1f2A1fBlifBl1f%qAf%A2f%%1f%ilf2LBfB1ifBlgf%%1f%2qf%q1f%Aqf2ABfBgi"
                },
                "wise": {
                    "startDate": "2023-06-30",
                    "endDate": "2023-07-29",
                    "data": "iiLBfiqgqfiiB2fiigAfiql2fiq%LfiqqgfiiBifiii1fiiqgfiqA%fiqg1fiiqAfiiqgfiqA1fgg1fiiAAfii%Bfiiq2fii1gfiiLifiqABfiB2qfiqLgfiiBLfiil%fiq1gfiiBAfBqgifilAq"
                },
                "type": "day"
            }
        ],
        "generalRatio": [
            {
                "word": [
                    {
                        "name": "\u5361\u5854\u5c14",
                        "wordType": 1
                    }
                ],
                "all": {
                    "avg": 1581,//整体日均值
                    "yoy": -51,//整体同比
                    "qoq": -24//整体环比
                },
                "pc": {
                    "avg": 429,
                    "yoy": -43,//
                    "qoq": -13
                },
                "wise": {
                    "avg": 1152,//移动日均值
                    "yoy": -53,//移动同比
                    "qoq": -28//移动环比
                }
            }
        ],
        "uniqid": "5d33bc95ba05114d62471efdeb4dd6de"
    },
    "logid": 2851695081,
    "message": 0
}
```

##### 词汇相关搜索

```
POST https://index.baidu.com/insight/word/sug
payload:
{words: ["卡塔尔"], source: "pc_landpage_comp"}
```

响应：

```
{
    "status": 0,
    "data": {
        "result": [
            {
                "word": "卡塔尔",
                "type": "word",
                "recInfo": {
                    "dimension": "",
                    "id": "",
                    "industry": "",
                    "isnpr": 0,
                    "istopic": 0,
                    "name": "卡塔尔",
                    "wordType": 1
                }
            },
            {
                "word": "卡塔尔航空",
                "type": "word",
                "recInfo": {
                    "dimension": "",
                    "id": "",
                    "industry": "",
                    "isnpr": 0,
                    "istopic": 0,
                    "name": "卡塔尔航空",
                    "wordType": 1
                }
            },
            {
                "word": "卡塔尔世界杯",
                "type": "word",
                "recInfo": {
                    "dimension": "",
                    "id": "",
                    "industry": "",
                    "isnpr": 0,
                    "istopic": 0,
                    "name": "卡塔尔世界杯",
                    "wordType": 1
                }
            },
            {
                "word": "卡塔尔航空官网",
                "type": "word",
                "recInfo": {
                    "dimension": "",
                    "id": "",
                    "industry": "",
                    "isnpr": 0,
                    "istopic": 0,
                    "name": "卡塔尔航空官网",
                    "wordType": 1
                }
            },
            {
                "word": "卡塔尔埃米尔",
                "type": "word",
                "recInfo": {
                    "dimension": "",
                    "id": "",
                    "industry": "",
                    "isnpr": 0,
                    "istopic": 0,
                    "name": "卡塔尔埃米尔",
                    "wordType": 1
                }
            },
            {
                "word": "卡塔尔属于哪个国家",
                "type": "word",
                "recInfo": {
                    "dimension": "",
                    "id": "",
                    "industry": "",
                    "isnpr": 0,
                    "istopic": 0,
                    "name": "卡塔尔属于哪个国家",
                    "wordType": 1
                }
            },
            {
                "word": "卡塔尔是发达国家吗",
                "type": "word",
                "recInfo": {
                    "dimension": "",
                    "id": "",
                    "industry": "",
                    "isnpr": 0,
                    "istopic": 0,
                    "name": "卡塔尔是发达国家吗",
                    "wordType": 1
                }
            },
            {
                "word": "卡塔尔时间",
                "type": "word",
                "recInfo": {
                    "dimension": "",
                    "id": "",
                    "industry": "",
                    "isnpr": 0,
                    "istopic": 0,
                    "name": "卡塔尔时间",
                    "wordType": 1
                }
            },
            {
                "word": "卡塔尔世界杯主题曲",
                "type": "word",
                "recInfo": {
                    "dimension": "",
                    "id": "",
                    "industry": "",
                    "isnpr": 0,
                    "istopic": 0,
                    "name": "卡塔尔世界杯主题曲",
                    "wordType": 1
                }
            },
            {
                "word": "卡塔尔首都",
                "type": "word",
                "recInfo": {
                    "dimension": "",
                    "id": "",
                    "industry": "",
                    "isnpr": 0,
                    "istopic": 0,
                    "name": "卡塔尔首都",
                    "wordType": 1
                }
            },
            {
                "word": "卡塔尔属于哪个洲",
                "type": "word",
                "recInfo": {
                    "dimension": "",
                    "id": "",
                    "industry": "",
                    "isnpr": 0,
                    "istopic": 0,
                    "name": "卡塔尔属于哪个洲",
                    "wordType": 1
                }
            }
        ]
    }
}
```

